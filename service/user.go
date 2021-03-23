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

func newWsUser(conn *websocket.Conn) (*wsUser, error) {
	conn.SetReadLimit(512)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	mt, message, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	if mt != int(pbs.SrvMsgType_Online) {
		return nil, fmt.Errorf("first msg must be online noti")
	}
	online := &pbs.WSOnline{}
	if err := proto.UnmarshalMerge(message, online); err != nil {
		return nil, err
	}

	//TODO::verify user's account balance
	wu := &wsUser{
		conn: conn,
		UID:  online.UID,
	}
	return wu, nil
}

func (ws *WebSocketService) newOnlineUser(conn *websocket.Conn) error {
	wu, err := newWsUser(conn)
	if err != nil {
		conn.Close()
		utils.LogInst().Err(err).Send()
		return err
	}
	wu.inChan = ws.msgQueue

	ws.userTable.Add(wu)
	ws.onlineSet.NotifyPeers(wu)

	thread.NewThreadWithName(fmt.Sprintf("chat read:%s", wu.UID), func(stop chan struct{}) {
		wu.reader(stop)
	}).Run()

	thread.NewThreadWithName(fmt.Sprintf("chat writer:%s", wu.UID), func(stop chan struct{}) {
		wu.writer(stop)
	}).Run()

	return nil
}
