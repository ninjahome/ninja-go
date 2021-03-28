package client

import (
	"fmt"
	"github.com/forgoer/openssl"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/service"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/ninjahome/ninja-go/wallet"
	"net/url"
)

type InputFunc func(*pbs.WSCryptoMsg) error

type WSClient struct {
	isOnline             bool
	endpoint             string
	wsConn               *websocket.Conn
	key                  *wallet.Key
	reader, writer       *thread.Thread
	callbackForServerMsg InputFunc
	msgChanToServer      chan *pbs.WSCryptoMsg
	peerKeys             map[string][]byte
}

func NewWSClient(addr string, key *wallet.Key) (*WSClient, error) {
	if key == nil || !key.IsOpen() {
		return nil, fmt.Errorf("ivnalid key")
	}

	cc := &WSClient{
		endpoint:        addr,
		key:             key,
		isOnline:        false,
		msgChanToServer: make(chan *pbs.WSCryptoMsg, 1024),
		peerKeys:        make(map[string][]byte),
	}

	cc.reader = thread.NewThread(cc.reading)
	cc.writer = thread.NewThread(cc.writing)
	return cc, nil
}

func (cc *WSClient) Register(in InputFunc) {
	cc.callbackForServerMsg = in
}

func (cc *WSClient) Online() error {
	u := url.URL{Scheme: "ws", Host: cc.endpoint, Path: service.CPUserOnline}
	wsConn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	onlineMsg := &pbs.WSOnline{}
	if err := onlineMsg.Online(wsConn, cc.key); err != nil {
		return err
	}
	cc.wsConn = wsConn
	cc.reader.Run()
	cc.writer.Run()
	return nil
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
	return key, nil
}

func (cc *WSClient) Write(msg *pbs.WSCryptoMsg) error {
	if !cc.isOnline {
		return fmt.Errorf("please online yourself first")
	}
	key, err := cc.getAesKey(msg.To)
	if err != nil {
		return err
	}
	dst, _ := openssl.AesECBEncrypt(msg.PayLoad, key, openssl.PKCS7_PADDING)
	msg.PayLoad = dst
	cc.msgChanToServer <- msg
	return nil
}

func (cc *WSClient) procMsgFromServer() error {
	mt, message, err := cc.wsConn.ReadMessage()
	if err != nil {
		fmt.Println("read:", err)
		return err
	}

	switch mt {
	case websocket.PingMessage:
		return cc.wsConn.WriteMessage(websocket.PongMessage, []byte{})

	case int(pbs.SrvMsgType_OnlineACK):
		ack := &pbs.WSOnlineAck{}
		if err := proto.Unmarshal(message, ack); err != nil {
			return fmt.Errorf("unknown websocket message:%s", err)
		}
		if !ack.Success {
			return fmt.Errorf("online failed")
		}
		cc.isOnline = true

	case int(pbs.SrvMsgType_CryptoMsg):
		if cc.callbackForServerMsg == nil {
			fmt.Println("no input message callback")
			return nil
		}

		msg := &pbs.WSCryptoMsg{}
		if err := proto.Unmarshal(message, msg); err != nil {
			return fmt.Errorf("unknown websocket message:%s", err)
		}

		key, err := cc.getAesKey(msg.From)
		if err != nil {
			return err
		}
		dst, _ := openssl.AesECBDecrypt(msg.PayLoad, key, openssl.PKCS7_PADDING)
		msg.PayLoad = dst
		if err := cc.callbackForServerMsg(msg); err != nil {
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

func (cc *WSClient) ShutDown() {

	if cc.msgChanToServer != nil {
		return
	}

	cc.reader.Stop()
	cc.writer.Stop()
	cc.wsConn.Close()
	cc.isOnline = false
	close(cc.msgChanToServer)
	cc.msgChanToServer = nil
}
