package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"google.golang.org/protobuf/proto"
	"time"
)

type wsUser struct {
	UID            string
	onLineTime     time.Time
	cliWsConn      *websocket.Conn
	msgFromCliChan chan *pbs.WsMsg
	msgToCliChan   chan *pbs.WsMsg
	kaTimer        *time.Ticker
}

func (u *wsUser) offLine() {

	if u.msgToCliChan == nil {
		return
	}

	u.cliWsConn.WriteMessage(websocket.CloseMessage, []byte{})
	u.cliWsConn.Close()
	close(u.msgToCliChan)
	u.msgToCliChan = nil
	u.kaTimer.Stop()
}

func (u *wsUser) reading(_ chan struct{}) {
	utils.LogInst().Info().Msgf("reading thread for [%s] start success!", u.UID)
	defer u.offLine()
	for {
		_, message, err := u.cliWsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				utils.LogInst().Err(err).Send()
			}
			break
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
	utils.LogInst().Info().Msgf("writing thread for [%s] start success!", u.UID)
	defer u.offLine()
	for {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("web socket writer thread exit")
			return

		case message, ok := <-u.msgToCliChan:
			u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait))
			if !ok {
				u.cliWsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.cliWsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				utils.LogInst().Err(err).Send()
				return
			}

			data, _ := proto.Marshal(message)
			w.Write(data)
			if err := w.Close(); err != nil {
				utils.LogInst().Err(err).Send()
				return
			}

		case <-u.kaTimer.C:
			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Err(err).Send()
				return
			}
			if err := u.cliWsConn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				utils.LogInst().Err(err).Send()
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
	return fmt.Sprintf("uid:%s, online:%s, from:%s", u.UID, u.onLineTime, u.cliWsConn.RemoteAddr())
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
		onLineTime:     time.Now(),
		msgFromCliChan: ws.msgFromClientQueue,
		kaTimer:        time.NewTicker(_wsConfig.PingPeriod),
		msgToCliChan:   make(chan *pbs.WsMsg, _wsConfig.MaxUnreadMsgNoPerQuery),
	}

	if err := ws.p2pOnOffLineWriter.Publish(ws.ctx, rawData); err != nil {
		return err
	}
	ws.onlineSet.add(wu.UID)
	ws.userTable.add(wu)

	tid := fmt.Sprintf("chat read:%s", wu.UID)
	t := thread.NewThreadWithName(tid, wu.reading)
	t.WillExit(func() {
		ws.offlineUser(tid, wu.UID)
	})
	ws.threads[tid] = t
	t.Run()

	tid = fmt.Sprintf("chat writer:%s", wu.UID)
	t = thread.NewThreadWithName(tid, wu.writing)
	ws.threads[tid] = t
	t.Run()

	utils.LogInst().Debug().Msgf("new user[%s] online success.....", wu.UID)

	return nil
}

func (ws *Service) offlineUser(threadId string, uid string) {
	utils.LogInst().Info().Msgf("user [%s] offline ", uid)
	delete(ws.threads, threadId)
	ws.onlineSet.del(uid)
	ws.userTable.del(uid)

	//TODO:: add signature for offline message
	msg := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_Offline,
		Payload: &pbs.WsMsg_Online{Online: &pbs.WSOnline{UID: uid}},
	}

	if err := ws.p2pOnOffLineWriter.Publish(ws.ctx, msg.Data()); err != nil {
		utils.LogInst().Warn().Err(err).Send()
	}
}

func (ws *Service) OnOffLineForP2pNetwork(w *worker.TopicWorker) {
	ws.p2pOnOffLineWriter = w.Pub

	for {
		msg, err := w.Sub.Next(ws.ctx)
		if err != nil {
			utils.LogInst().Warn().Msgf("on-off line thread exit:=>%s", err)
			return
		}

		p2pMsg := &pbs.WsMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Msg("failed parse p2p message")
			continue
		}

		if p2pMsg.Typ == pbs.WsMsgType_Online {
			if err := ws.onlineFromOtherPeer(p2pMsg); err != nil {
				utils.LogInst().Warn().Msg("online from p2p network failed")
			}
		} else if p2pMsg.Typ == pbs.WsMsgType_Offline {
			if err := ws.offlineFromOtherPeer(p2pMsg); err != nil {
				utils.LogInst().Warn().Msg("offline from p2p network failed")
			}
		} else {
			utils.LogInst().Warn().Msg("unknown msg typ in p2p on-off line channel")
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
