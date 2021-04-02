package websocket

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs2 "github.com/ninjahome/ninja-go/pbs/stream"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

type wsUser struct {
	lock           sync.RWMutex
	UID            string
	OnLineTime     time.Time
	cliWsConn      *websocket.Conn
	msgFromCliChan chan *pbs.WsMsg
	msgToCliChan   chan *pbs.WsMsg
	kaTimer        *time.Ticker
}

func (u *wsUser) offLine() {
	u.lock.Lock()
	defer u.lock.Unlock()

	if u.msgToCliChan == nil {
		return
	}

	u.cliWsConn.Close()
	close(u.msgToCliChan)
	u.msgToCliChan = nil
	u.kaTimer.Stop()
	utils.LogInst().Debug().Msgf("user[%s] offline add clean data.....", u.UID)
}

func (u *wsUser) reading(_ chan struct{}) {
	utils.LogInst().Debug().Msgf("reading thread for [%s] start success!", u.UID)
	defer utils.LogInst().Debug().Msgf("reading thread for [%s] exit!", u.UID)
	defer u.offLine()
	for {
		_, message, err := u.cliWsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				utils.LogInst().Err(err).Send()
			}
			utils.LogInst().Info().Msgf("websocket read thread read message failed:%s", err)
			return
		}

		msg := &pbs.WsMsg{}
		if err := proto.Unmarshal(message, msg); err != nil {
			utils.LogInst().Warn().Msgf("web socket read invalid:%x", message)
			continue
		}
		u.msgFromCliChan <- msg
	}
}

func (u *wsUser) writing(stop chan struct{}) {
	utils.LogInst().Debug().Msgf("web socket writing thread for [%s] start success!", u.UID)
	defer utils.LogInst().Debug().Msgf("web socket writer thread for [%s] exit", u.UID)

	defer u.offLine()
	for {
		select {
		case <-stop:
			return

		case message, ok := <-u.msgToCliChan:
			if !ok {
				utils.LogInst().Info().Msgf("websocket write thread message closed")
				u.cliWsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			utils.LogInst().Debug().Msgf("websocket write thread get new client message=>%s", message.String())
			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Err(err).Msg("websocket write thread set timeout failed ")
				return
			}

			w, err := u.cliWsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				utils.LogInst().Err(err).Msg("websocket write thread get next writer failed ")
				return
			}

			data, _ := proto.Marshal(message)
			w.Write(data)
			if err := w.Close(); err != nil {
				utils.LogInst().Err(err).Msg("websocket write thread close current writer failed")
				return
			}

		case <-u.kaTimer.C:
			utils.LogInst().Debug().Msg("websocket write thread ping pong time")
			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Err(err).Msg("websocket write deadline failed")
				return
			}
			if err := u.cliWsConn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				utils.LogInst().Err(err).Msg("websocket write ping data to client failed")
				return
			}
		}
	}
}

func (u *wsUser) writeToCli(msg *pbs.WsMsg) error {
	u.msgToCliChan <- msg
	return nil
}

func (u *wsUser) String() string {
	return fmt.Sprintf("uid:%s, online:%s, from:%s", u.UID, u.OnLineTime, u.cliWsConn.RemoteAddr())
}

func (ws *Service) newOnlineUser(conn *websocket.Conn) error {

	msg := &pbs.WsMsg{}
	online, rawData, err := msg.ReadOnlineFromCli(conn)
	if err != nil {
		conn.Close()
		return err
	}

	wu := &wsUser{
		cliWsConn:      conn,
		UID:            online.UID,
		OnLineTime:     time.Now(),
		msgFromCliChan: ws.msgFromClientQueue,
		kaTimer:        time.NewTicker(_wsConfig.PingPeriod),
		msgToCliChan:   make(chan *pbs.WsMsg, _wsConfig.MaxUnreadMsgNoPerQuery),
	}
	ws.onlineSet.add(wu.UID)
	ws.userTable.add(wu)

	tid := fmt.Sprintf("chat read:%s", wu.UID)
	readTh := thread.NewThreadWithName(tid, wu.reading)
	readTh.WillExit(func() {
		ws.offlineUser(tid, wu.UID)
	})
	ws.threads[tid] = readTh
	readTh.Run()

	tid = fmt.Sprintf("chat writer:%s", wu.UID)
	writeTh := thread.NewThreadWithName(tid, wu.writing)
	ws.threads[tid] = writeTh
	writeTh.Run()

	if err := ws.onOffLineP2pWorker.BroadCast(rawData); err != nil {
		return err
	}

	utils.LogInst().Debug().Msgf("new user[%s] online success.....", wu.UID)

	return nil
}

func (ws *Service) offlineUser(threadId string, uid string) {
	delete(ws.threads, threadId)
	ws.onlineSet.del(uid)
	ws.userTable.del(uid)

	//TODO:: add signature for offline message
	msg := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_Offline,
		Payload: &pbs.WsMsg_Online{Online: &pbs.WSOnline{UID: uid}},
	}

	if err := ws.onOffLineP2pWorker.BroadCast(msg.Data()); err != nil {
		utils.LogInst().Warn().Err(err).Msg("broadcast user offline message failed")
	}
	utils.LogInst().Info().Msgf("offline [%s]", uid)
}

func (ws *Service) OnOffLineForP2pNetwork(w *worker.TopicWorker) {
	ws.onOffLineP2pWorker = w

	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Msgf("on-off line thread exit:=>%s", err)
			return
		}

		p2pMsg := &pbs.WsMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Msg("failed parse p2p message")
			continue
		}

		utils.LogInst().Debug().Str("online-offline message", p2pMsg.Typ.String()).Msg(p2pMsg.String())

		switch p2pMsg.Typ {
		case pbs.WsMsgType_Online:
			err = ws.onlineFromOtherPeer(p2pMsg)
		case pbs.WsMsgType_Offline:
			err = ws.offlineFromOtherPeer(p2pMsg)
		default:
			err = fmt.Errorf("unknown msg typ in p2p on-off line channel")
		}

		if err != nil {
			utils.LogInst().Err(err).Send()
		}
	}
}

func (ws *Service) onlineFromOtherPeer(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Online)
	if !ok {
		return fmt.Errorf("this is not a valid online p2p message")
	}

	if !body.Online.Verify(msg.Sig) {
		return fmt.Errorf("this is an attack")
	}
	ws.onlineSet.add(body.Online.UID)
	return nil
}

func (ws *Service) offlineFromOtherPeer(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Online)
	if !ok {
		return fmt.Errorf("this is not a valid offline p2p message")
	}
	//TODO:: verify peer's authorization
	ws.onlineSet.del(body.Online.UID)
	ws.userTable.del(body.Online.UID)
	return nil
}

func (ws *Service) SyncOnlineSetFromPeerNodes(stream network.Stream) error {
	defer stream.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	streamMsg := &pbs2.StreamMsg{}
	data := streamMsg.SyncOnline("TODO::wallet key and sig") //TODO::
	data = append(data, OnlineStreamDelim)

	_, err := rw.Write(data)
	if err != nil {
		utils.LogInst().Err(err).Msg("stream: write online sync request data failed")
		return err
	}
	if err := rw.Flush(); err != nil {
		utils.LogInst().Err(err).Msg("stream:  online sync request flush failed")
	}

	bts, err := rw.ReadBytes(OnlineStreamDelim)
	if err != nil {
		utils.LogInst().Err(err).Msg("stream: read online sync response data failed")
		return err
	}

	resp := &pbs2.StreamMsg{}
	bts = bts[:len(bts)-1]
	if err := proto.Unmarshal(bts, resp); err != nil {
		utils.LogInst().Err(err).Msg("failed parse stream message")
		return err
	}

	body, ok := resp.Payload.(*pbs2.StreamMsg_OnlineAck)
	if !ok {
		utils.LogInst().Err(err).Msg("failed parse stream message")
		return fmt.Errorf("invalid onlime map data")
	}

	uidBatch := body.OnlineAck.UID
	utils.LogInst().Info().Msgf("sync online users[%d]", len(uidBatch))
	if len(uidBatch) == 0 {
		return nil
	}
	ws.onlineSet.addBatch(uidBatch)

	return nil
}

func (ws *Service) OnlineMapQuery(stream network.Stream) {
	defer stream.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	bts, err := rw.ReadBytes(OnlineStreamDelim)
	if err != nil {
		utils.LogInst().Err(err).Msg("stream: read online sync request data failed")
		return
	}

	bts = bts[:len(bts)-1]
	streamMsg := &pbs2.StreamMsg{}
	if err := proto.Unmarshal(bts, streamMsg); err != nil {
		utils.LogInst().Err(err).Msg("failed parse stream message")
		return
	}

	resp := &pbs2.StreamMsg{}
	data := resp.SyncOnlineAck(ws.onlineSet.AllUid())
	data = append(data, OnlineStreamDelim)
	if _, err := rw.Write(data); err != nil {
		utils.LogInst().Err(err).Msg("stream: write online set response data failed")
		return
	}
	if err := rw.Flush(); err != nil {
		utils.LogInst().Err(err).Msg("stream:  online sync response flush failed")
	}
}
