package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"time"
)

type wsUser struct {
	UID     string
	onTime  time.Time
	conn    *websocket.Conn
	inChan  chan *pbs.WSCryptoMsg
	outChan chan *pbs.WSCryptoMsg
}

func (u *wsUser) reader(stop chan struct{}) {
	defer u.conn.Close()
	for {
		mt, message, err := u.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				utils.LogInst().Err(err)
			}
			break
		}

		if mt != int(pbs.SrvMsgType_CryptoMsg) {
			utils.LogInst().Warn().Int("web socket read invalid msg type", mt).Send()
			continue
		}

		msg := &pbs.WSCryptoMsg{}

		if err := proto.Unmarshal(message, msg); err != nil {
			utils.LogInst().Warn().Msgf("web socket read invalid:%x", message)
			continue
		}
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("web socket reader thread exit")
			return
		default:
			u.inChan <- msg
		}
	}
}

func (u *wsUser) writer(stop chan struct{}) {

	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		u.conn.Close()
	}()

	for {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("web socket writer thread exit")
			return
		case message, ok := <-u.outChan:
			u.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				u.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := u.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			data, _ := proto.Marshal(message)
			w.Write(data)
			if err := w.Close(); err != nil {
				return
			}
		case <-pingTicker.C:
			if err := u.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return
			}
			if err := u.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}

}

func (u *wsUser) write(msg *pbs.WSCryptoMsg) error {
	u.outChan <- msg
	return nil
}

func (ws *WebSocketService) newOnlineUser(conn *websocket.Conn) error {

	online := &pbs.WSOnline{}
	if err := online.ReadOnlineFromCli(conn); err != nil {
		conn.Close()
		return err
	}

	wu := &wsUser{
		conn:   conn,
		UID:    online.UID,
		onTime: time.Unix(online.UnixTime, 0),
		inChan: ws.msgFromClientQueue,
	}

	ws.msgToOtherPeerQueue <- &pbs.P2PMsg{
		MsgTyp:  pbs.P2PMsgType_P2pOnline,
		Payload: &pbs.P2PMsg_Online{Online: online},
	}

	ws.onlineSet.add(wu.UID)
	ws.userTable.add(wu)

	tid := fmt.Sprintf("chat read:%s", wu.UID)
	t := thread.NewThreadWithName(tid, func(stop chan struct{}) {
		wu.reader(stop)
	})
	ws.threads[tid] = t
	t.Run()

	tid = fmt.Sprintf("chat writer:%s", wu.UID)
	t = thread.NewThreadWithName(tid, func(stop chan struct{}) {
		wu.writer(stop)
	})
	ws.threads[tid] = t
	t.Run()

	return nil
}
