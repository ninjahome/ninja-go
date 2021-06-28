package client

import (
	"fmt"
	"github.com/forgoer/openssl"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	websocket2 "github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/ninjahome/ninja-go/wallet"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"net/url"
	"time"
)

var (
	DefaultBootWsService = []string{
		"202.182.101.145:6666",
		"167.179.78.33:6666",
		"198.13.44.159:6666",
	}

	ErrUnknownMsg    = fmt.Errorf("unknown websocket message")
	ErrNoMsgCallback = fmt.Errorf("no message reciver")
)

const (
	DevType_IOS = 1
	DevType_Android = 2
)

type WSClient struct {
	DevTyp int
	DeviceToken string
	IsOnline bool
	endpoint string
	wsConn   *websocket.Conn
	key      *wallet.Key
	reader   *thread.Thread
	callback CliCallBack
	peerKeys map[string][]byte
}

func RandomBootNode() string {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(DefaultBootWsService))
	return DefaultBootWsService[idx]
}

type CliCallBack interface {
	ImmediateMessage(*pbs.WSCryptoMsg) error
	WebSocketClosed()
	UnreadMsg(*pbs.WSUnreadAck) error
	OnlineSuccess()
}

//devType DevType_IOS/DevType_Android

func NewWSClient(deviceToken, addr string, devType int, key *wallet.Key, cb CliCallBack) (*WSClient, error) {
	if key == nil || !key.IsOpen() {
		return nil, fmt.Errorf("ivnalid key")
	}

	cc := &WSClient{
		DevTyp: devType,
		DeviceToken: deviceToken,
		endpoint: addr,
		key:      key,
		IsOnline: false,
		peerKeys: make(map[string][]byte),
		callback: cb,
	}

	cc.reader = thread.NewThread(cc.reading)
	return cc, nil
}

func (cc *WSClient) Online() error {
	u := url.URL{Scheme: "ws", Host: cc.endpoint, Path: websocket2.CPUserOnline}

	dialer:=websocket.DefaultDialer

	dialer.ReadBufferSize = websocket2.DefaultWsBuffer
	dialer.WriteBufferSize = websocket2.DefaultWsBuffer

	wsConn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	onlineMsg := &pbs.WsMsg{}
	if err := onlineMsg.Online(wsConn, cc.key,cc.DeviceToken,cc.DevTyp); err != nil {
		return err
	}
	cc.wsConn = wsConn
	wsConn.SetPingHandler(func(appData string) error {
		fmt.Println("ping pong time......")
		return wsConn.WriteMessage(websocket.PongMessage, []byte{})
	})
	cc.reader.Run()
	return nil
}

func (cc *WSClient) PullUnreadMsg(startSeq int64) error {
	request := &pbs.WSPullUnread{
		Receiver:     cc.key.Address.String(),
		FromUnixTime: startSeq,
	}

	msgWrap := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_PullUnread,
		Payload: &pbs.WsMsg_Unread{Unread: request},
	}
	return cc.wsConn.WriteMessage(websocket.TextMessage, msgWrap.Data())
}

func (cc *WSClient) getAesKey(to string) ([]byte, error) {
	worker, ok := cc.peerKeys[to]
	if ok {
		return worker, nil
	}

	key, err := cc.key.SharedKey(to)
	if err != nil {
		return nil, err
	}
	cc.peerKeys[to] = key
	return key, nil
}

func (cc *WSClient) Write(to string, body []byte) error {
	if !cc.IsOnline {
		return fmt.Errorf("please online yourself first")
	}
	key, err := cc.getAesKey(to)
	if err != nil {
		return err
	}

	from := cc.key.Address.String()
	msgWrap := &pbs.WsMsg{}

	if err := cc.wsConn.WriteMessage(websocket.TextMessage,
		msgWrap.AesCryptData(from, to, body, key)); err != nil {
		return err
	}
	return nil
}

func (cc *WSClient) procMsgFromServer() error {
	if cc.callback == nil {
		return ErrNoMsgCallback
	}

	_, message, err := cc.wsConn.ReadMessage()
	if err != nil {
		fmt.Println("read:", err)
		return err
	}
	wsMsg := &pbs.WsMsg{}
	if err := proto.Unmarshal(message, wsMsg); err != nil {
		return err
	}

	switch wsMsg.Typ {
	case pbs.WsMsgType_OnlineACK:
		ack, ok := wsMsg.Payload.(*pbs.WsMsg_OlAck)
		if !ok {
			return ErrUnknownMsg
		}
		if !ack.OlAck.Success {
			return fmt.Errorf("online failed")
		}
		cc.IsOnline = true
		go cc.callback.OnlineSuccess()
	case pbs.WsMsgType_ImmediateMsg:
		msgWrap, ok := wsMsg.Payload.(*pbs.WsMsg_Message)
		if !ok {
			return ErrUnknownMsg
		}
		msg := msgWrap.Message
		key, err := cc.getAesKey(msg.From)
		if err != nil {
			return err
		}
		dst, _ := openssl.AesECBDecrypt(msg.PayLoad, key, openssl.PKCS7_PADDING)
		msg.PayLoad = dst
		if err := cc.callback.ImmediateMessage(msg); err != nil {
			return err
		}
	case pbs.WsMsgType_UnreadAck:
		ack, ok := wsMsg.Payload.(*pbs.WsMsg_UnreadAck)
		if !ok {
			return ErrUnknownMsg
		}

		for _, msg := range ack.UnreadAck.Payload {
			key, err := cc.getAesKey(msg.From)
			if err != nil {
				return err
			}
			dst, _ := openssl.AesECBDecrypt(msg.PayLoad, key, openssl.PKCS7_PADDING)
			msg.PayLoad = dst
		}
		if err := cc.callback.UnreadMsg(ack.UnreadAck); err != nil {
			return err
		}
	}
	return nil
}

func (cc *WSClient) reading(stop chan struct{}) {

	defer cc.ShutDown()
	for {
		select {
		case <-stop:
			fmt.Println("reading thread exit")
			return
		default:
			if err := cc.procMsgFromServer(); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func (cc *WSClient) ShutDown() {
	if !cc.IsOnline {
		return
	}
	cc.callback.WebSocketClosed()
	cc.reader.Stop()
	cc.wsConn.Close()
	cc.IsOnline = false
	fmt.Println("websocket client shutdown......")
}
