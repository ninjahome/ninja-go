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
	utils.LogInst().Debug().Str("WS user offline", u.UID).Send()
}

func (u *wsUser) reading(_ chan struct{}) {
	utils.LogInst().Debug().Str("WS reading thread start", u.UID).Send()
	defer utils.LogInst().Debug().Str("reading thread exit!", u.UID).Send()
	defer u.offLine()
	for {
		_, message, err := u.cliWsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				utils.LogInst().Err(err).Send()
			}
			utils.LogInst().Info().Str("WS read client", err.Error()).Send()
			return
		}

		msg := &pbs.WsMsg{}
		if err := proto.Unmarshal(message, msg); err != nil {
			utils.LogInst().Warn().Str("WS invalid client message", err.Error()).Send()
			continue
		}
		u.msgFromCliChan <- msg
	}
}

func (u *wsUser) writing(stop chan struct{}) {
	utils.LogInst().Debug().Str("WS writing thread start!", u.UID).Send()
	defer utils.LogInst().Debug().Str("WS writer thread exit", u.UID).Send()

	defer u.offLine()
	for {
		select {
		case <-stop:
			return
		case message, ok := <-u.msgToCliChan:
			if !ok {
				utils.LogInst().Info().Str("WS client message chan", " closed").Send()
				u.cliWsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Warn().Str("WS set write timeout ", err.Error()).Send()
				return
			}

			w, err := u.cliWsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				utils.LogInst().Warn().Str("WS get next writer ", err.Error()).Send()
				return
			}

			_, err = w.Write(message.Data())
			if err := w.Close(); err != nil {
				utils.LogInst().Warn().Str("WS write ", err.Error()).Send()
				return
			}

		case <-u.kaTimer.C:
			utils.LogInst().Debug().Str("WS ping pong", "sent").Send()
			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Warn().Str("WS write deadline", err.Error()).Send()
				return
			}
			if err := u.cliWsConn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				utils.LogInst().Warn().Str("WS write ping", err.Error()).Send()
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
	ws.threads[tid] = readTh
	readTh.Run()

	tid = fmt.Sprintf("chat writer:%s", wu.UID)
	writeTh := thread.NewThreadWithName(tid, wu.writing)
	writeTh.DidExit(func() {
		ws.offlineUser(tid, wu.UID)
	})
	ws.threads[tid] = writeTh
	writeTh.Run()

	if err := ws.onOffLineP2pWorker.BroadCast(rawData); err != nil {
		return err
	}

	utils.LogInst().Debug().Str("WS New User", wu.UID).Send()

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
		utils.LogInst().Warn().Str("offline broadcast", err.Error()).Send()
	}
	utils.LogInst().Info().Str("WS user offline", uid).Send()
}

func (ws *Service) OnOffLineForP2pNetwork(w *worker.TopicWorker) {
	ws.onOffLineP2pWorker = w

	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Str("on-off line ", err.Error()).Send()
			return
		}
		if msg.ReceivedFrom.String() == ws.id {
			continue
		}

		p2pMsg := &pbs.WsMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Str("unmarshal", err.Error()).Send()
			continue
		}

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
	utils.LogInst().Debug().Str("online", body.Online.UID).Send()
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
	utils.LogInst().Debug().Str("online", body.Online.UID).Send()
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
		utils.LogInst().Warn().Str("stream: write online", err.Error()).Send()
		return err
	}
	if err := rw.Flush(); err != nil {
		utils.LogInst().Warn().Str("stream: flush online", err.Error()).Send()
	}

	bts, err := rw.ReadBytes(OnlineStreamDelim)
	if err != nil {
		utils.LogInst().Warn().Str("stream: read online", err.Error()).Send()
		return err
	}

	resp := &pbs2.StreamMsg{}
	bts = bts[:len(bts)-1]
	if err := proto.Unmarshal(bts, resp); err != nil {
		utils.LogInst().Warn().Str("stream: parse data online", err.Error()).Send()
		return err
	}

	body, ok := resp.Payload.(*pbs2.StreamMsg_OnlineAck)
	if !ok {
		utils.LogInst().Warn().Str("stream: cast data online", "failed").Send()
		return fmt.Errorf("invalid onlime map data")
	}

	uidBatch := body.OnlineAck.UID
	utils.LogInst().Info().Int("synced online users", len(uidBatch)).Send()
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
		utils.LogInst().Warn().Str("read online", err.Error()).Send()
		return
	}

	bts = bts[:len(bts)-1]
	streamMsg := &pbs2.StreamMsg{}
	if err := proto.Unmarshal(bts, streamMsg); err != nil {
		utils.LogInst().Warn().Str("parse stream", err.Error()).Send()
		return
	}

	resp := &pbs2.StreamMsg{}
	data := resp.SyncOnlineAck(ws.onlineSet.AllUid())
	data = append(data, OnlineStreamDelim)
	if _, err := rw.Write(data); err != nil {
		utils.LogInst().Warn().Str("stream:  write online response", err.Error()).Send()
		return
	}
	if err := rw.Flush(); err != nil {
		utils.LogInst().Warn().Str("stream:  flush online response", err.Error()).Send()
	}
}
