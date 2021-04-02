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
	"net/url"
	"time"
)

type WSClient struct {
	isOnline        bool
	endpoint        string
	wsConn          *websocket.Conn
	key             *wallet.Key
	reader, writer  *thread.Thread
	callback        CliCallBack
	msgChanToServer chan *pbs.WsMsg
	peerKeys        map[string][]byte
}

type CliCallBack interface {
	InputMsg(*pbs.WSCryptoMsg) error
	WebSocketClosed()
	UnreadMsg([]*pbs.WSCryptoMsg) error
}

func NewWSClient(addr string, key *wallet.Key, cb CliCallBack) (*WSClient, error) {
	if key == nil || !key.IsOpen() {
		return nil, fmt.Errorf("ivnalid key")
	}

	cc := &WSClient{
		endpoint:        addr,
		key:             key,
		isOnline:        false,
		msgChanToServer: make(chan *pbs.WsMsg, 1024),
		peerKeys:        make(map[string][]byte),
		callback:        cb,
	}

	cc.reader = thread.NewThread(cc.reading)
	cc.writer = thread.NewThread(cc.writing)
	return cc, nil
}

func (cc *WSClient) Online() error {
	u := url.URL{Scheme: "ws", Host: cc.endpoint, Path: websocket2.CPUserOnline}
	wsConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	onlineMsg := &pbs.WsMsg{}
	if err := onlineMsg.Online(wsConn, cc.key); err != nil {
		return err
	}
	cc.wsConn = wsConn
	wsConn.SetPingHandler(func(appData string) error {
		fmt.Println("ping pong time......")
		return wsConn.WriteMessage(websocket.PongMessage, []byte{})
	})
	cc.reader.Run()
	cc.writer.Run()
	return nil
}

func (cc *WSClient) PullMsg(startSeq int64) error {
	request := &pbs.WSPullUnread{
		Receiver:     cc.key.Address.String(),
		FromUnixTime: startSeq,
	}

	msgWrap := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_PullUnread,
		Payload: &pbs.WsMsg_Unread{Unread: request},
	}
	data, err := proto.Marshal(msgWrap)
	if err != nil {
		return err
	}
	return cc.wsConn.WriteMessage(websocket.TextMessage, data)
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
	if !cc.isOnline {
		return fmt.Errorf("please online yourself first")
	}

	msg := &pbs.WSCryptoMsg{
		From:     cc.key.Address.String(),
		To:       to,
		PayLoad:  body,
		UnixTime: time.Now().Unix(),
	}

	key, err := cc.getAesKey(msg.To)
	if err != nil {
		return err
	}
	dst, _ := openssl.AesECBEncrypt(msg.PayLoad, key, openssl.PKCS7_PADDING)
	msg.PayLoad = dst
	msgWrap := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_ImmediateMsg,
		Payload: &pbs.WsMsg_Message{Message: msg},
	}
	cc.msgChanToServer <- msgWrap
	return nil
}

func (cc *WSClient) procMsgFromServer() error {
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
			return fmt.Errorf("unknown websocket message:%s", err)
		}
		if !ack.OlAck.Success {
			return fmt.Errorf("online failed")
		}
		cc.isOnline = true

	case pbs.WsMsgType_ImmediateMsg:
		if cc.callback == nil {
			fmt.Println("no input message callback")
			return nil
		}

		msgWrap, ok := wsMsg.Payload.(*pbs.WsMsg_Message)
		if !ok {
			return fmt.Errorf("unknown websocket message:%s", err)
		}
		msg := msgWrap.Message
		key, err := cc.getAesKey(msg.From)
		if err != nil {
			return err
		}
		dst, _ := openssl.AesECBDecrypt(msg.PayLoad, key, openssl.PKCS7_PADDING)
		msg.PayLoad = dst
		if err := cc.callback.InputMsg(msg); err != nil {
			return fmt.Errorf("process input message err:%s", err)
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

func (cc *WSClient) writing(stop chan struct{}) {
	defer cc.ShutDown()
	for {
		select {
		case outMsg := <-cc.msgChanToServer:
			data, err := proto.Marshal(outMsg)
			if err != nil {
				fmt.Println("invalid crypto message", err)
				continue
			}
			if err := cc.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
				fmt.Println("write crypto message", err)
				return
			}

		case <-stop:
			fmt.Println("write thread exit")
			_ = cc.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		}
	}
}

func (cc *WSClient) ShutDown() {

	if cc.msgChanToServer != nil {
		return
	}
	cc.callback.WebSocketClosed()
	cc.reader.Stop()
	cc.writer.Stop()
	cc.wsConn.Close()
	cc.isOnline = false
	close(cc.msgChanToServer)
	cc.msgChanToServer = nil
}
