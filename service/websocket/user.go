package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
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
	online, err := msg.ReadOnlineFromCli(conn)
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

	ws.msgToOtherPeerQueue <- msg
	ws.onlineSet.add(wu.UID)
	ws.userTable.add(wu)

	tid := fmt.Sprintf("chat read:%s", wu.UID)
	t := thread.NewThreadWithName(tid, func(stop chan struct{}) {
		wu.reader(stop)
		ws.OfflineUser(tid, wu, msg)
	})
	ws.threads[tid] = t
	t.Run()

	tid = fmt.Sprintf("chat writer:%s", wu.UID)
	t = thread.NewThreadWithName(tid, func(stop chan struct{}) {
		wu.writer(stop)
		ws.OfflineUser(tid, wu, msg)
	})
	ws.threads[tid] = t
	t.Run()

	return nil
}

func (ws *Service) OfflineUser(threadId string, user *wsUser, msg *pbs.WsMsg) {
	delete(ws.threads, threadId)
	ws.onlineSet.del(user.UID)
	ws.userTable.del(user.UID)

	//key := wallet.Inst().KeyInUsed()
	//key.SignData(msg.Payload)
	//offlineMsg := &pbs.P2PMsg_Offline{
	//	Offline:msg,
	//}
	//TODO:: add signature for offline message
	msg.Typ = pbs.WsMsgType_Offline
	ws.msgToOtherPeerQueue <- msg
}
