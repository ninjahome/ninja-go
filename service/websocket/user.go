package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/libp2p/go-libp2p-pubsub"
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
	u.cliWsConn.Close()
	close(u.msgToCliChan)
	u.msgToCliChan = nil
	u.kaTimer.Stop()
}

func (u *wsUser) reader(stop chan struct{}) {
	defer u.offLine()
	for {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("web socket reader thread exit")
			return
		default:
			_, message, err := u.cliWsConn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err,
					websocket.CloseGoingAway,
					websocket.CloseAbnormalClosure) {
					utils.LogInst().Err(err)
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
}

func (u *wsUser) writer(stop chan struct{}) {

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
				utils.LogInst().Err(err)
				return
			}

			data, _ := proto.Marshal(message)
			w.Write(data)
			if err := w.Close(); err != nil {
				utils.LogInst().Err(err)
				return
			}

		case <-u.kaTimer.C:
			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Err(err)
				return
			}
			if err := u.cliWsConn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				utils.LogInst().Err(err)
				return
			}
		}
	}
}

func (u *wsUser) writeToCli(msg *pbs.WsMsg) error {
	u.msgToCliChan <- msg
	return nil
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
		msgToCliChan:   make(chan *pbs.WsMsg, _wsConfig.WsMsgSizePerUser),
	}

	if err := ws.p2pOnOffWriter.Publish(ws.ctx, rawData); err != nil {
		return err
	}
	ws.onlineSet.add(wu.UID)
	ws.userTable.add(wu)

	tid := fmt.Sprintf("chat read:%s", wu.UID)
	t := thread.NewThreadWithName(tid, func(stop chan struct{}) {
		wu.reader(stop)
		ws.OfflineUser(tid, wu, online.UID)
	})
	ws.threads[tid] = t
	t.Run()

	tid = fmt.Sprintf("chat writer:%s", wu.UID)
	t = thread.NewThreadWithName(tid, func(stop chan struct{}) {
		wu.writer(stop)
		ws.OfflineUser(tid, wu, online.UID)
	})
	ws.threads[tid] = t
	t.Run()

	return nil
}

func (ws *Service) OfflineUser(threadId string, user *wsUser, uid string) {
	delete(ws.threads, threadId)
	ws.onlineSet.del(user.UID)
	ws.userTable.del(user.UID)

	//TODO:: add signature for offline message
	msg := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_Offline,
		Payload: &pbs.WsMsg_Online{Online: &pbs.WSOnline{UID: uid}},
	}

	if err := ws.p2pOnOffWriter.Publish(ws.ctx, msg.Data()); err != nil {
		utils.LogInst().Warn().Err(err).Send()
	}
}

func (ws *Service) OnOffLineForP2pNetwork(stop chan struct{}, r *pubsub.Subscription, w *pubsub.Topic) {
	ws.p2pOnOffWriter = w
	utils.LogInst().Debug().Msg("start on-off line message listening thread for p2p network")

	for {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("on-off line thread exit by outer controller")
			return
		default:
			msg, err := r.Next(ws.ctx)
			if err != nil {
				utils.LogInst().Warn().Err(err).Send()
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
