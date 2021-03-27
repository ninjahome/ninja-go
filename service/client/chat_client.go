package client

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/service"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/ninjahome/ninja-go/wallet"
	"net/url"
)

type InputFunc func(*pbs.WSCryptoMsg) error

type ChatClient struct {
	isOnline             bool
	endpoint             string
	wsConn               *websocket.Conn
	key                  *wallet.Key
	reader, writer       *thread.Thread
	callbackForServerMsg InputFunc
	out                  chan *pbs.WSCryptoMsg
}

func NewClient(addr string, key *wallet.Key) (*ChatClient, error) {
	if key == nil || !key.IsOpen() {
		return nil, fmt.Errorf("ivnalid key")
	}

	cc := &ChatClient{
		endpoint: addr,
		key:      key,
		isOnline: false,
		out:      make(chan *pbs.WSCryptoMsg),
	}

	cc.reader = thread.NewThread(cc.reading)
	cc.writer = thread.NewThread(cc.writing)
	return cc, nil
}

func (cc *ChatClient) Register(in InputFunc) {
	cc.callbackForServerMsg = in
}

func (cc *ChatClient) Online() error {
	u := url.URL{Scheme: "ws", Host: cc.endpoint, Path: service.CPUserOnline}
	wsConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	onlineMsg := &pbs.CliOnlineMsg{}
	if err := onlineMsg.Online(wsConn, cc.key); err != nil {
		return err
	}
	cc.wsConn = wsConn
	return nil
}

func (cc *ChatClient) Write(msg *pbs.WSCryptoMsg) error {
	if !cc.isOnline {
		return fmt.Errorf("please online yourself first")
	}
	cc.out <- msg
	return nil
}
func (cc *ChatClient) procMsgFromServer() error {
	mt, message, err := cc.wsConn.ReadMessage()
	if err != nil {
		fmt.Println("read:", err)
		return err
	}
	if cc.callbackForServerMsg == nil {
		fmt.Println("no input message callback")
		return nil
	}

	switch mt {
	case websocket.PingMessage:
		return cc.wsConn.WriteMessage(websocket.PongMessage, []byte{})

	case int(pbs.SrvMsgType_OnlineACK):
		cliMsg := &pbs.CliOnlineMsg{}
		if err := proto.Unmarshal(message, cliMsg); err != nil {
			return fmt.Errorf("unknown websocket message:%s", err)
		}
		ack, ok := cliMsg.Payload.(*pbs.CliOnlineMsg_OlAck)
		if !ok {
			return fmt.Errorf("convert to online ack failed")
		}
		if !ack.OlAck.Success {
			return fmt.Errorf("online failed")
		}
		cc.isOnline = true

	case int(pbs.SrvMsgType_CryptoMsg):
		msg := &pbs.WSCryptoMsg{}
		if err := proto.Unmarshal(message, msg); err != nil {
			return fmt.Errorf("unknown websocket message:%s", err)
		}
		err := cc.callbackForServerMsg(msg)
		if err != nil {
			return fmt.Errorf("process input message err:%s", err)

		}
	}
	return nil
}
func (cc *ChatClient) reading(stop chan struct{}) {

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

func (cc *ChatClient) writing(stop chan struct{}) {
	defer cc.ShutDown()

	for {
		select {
		case outMsg := <-cc.out:
			data, err := proto.Marshal(outMsg)
			if err != nil {
				fmt.Println("invalid crypto message", err)
				continue
			}
			if err := cc.wsConn.WriteMessage(int(pbs.SrvMsgType_CryptoMsg), data); err != nil {
				fmt.Println("write crypto message", err)
				return
			}

		case <-stop:
			fmt.Println("write thread exit")
			_ = cc.wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		}
	}
}

func (cc *ChatClient) ShutDown() {
	cc.reader.Stop()
	cc.writer.Stop()
	cc.wsConn.Close()
	cc.isOnline = false
}
