package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/service"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/ninjahome/ninja-go/wallet"
	"net/url"
)

type InputFunc func([]byte) error

type ChatClient struct {
	endpoint       string
	conn           *websocket.Conn
	key            *wallet.Key
	reader, writer *thread.Thread
	in             InputFunc
	out            <-chan []byte
}

func NewClient(addr string, key *wallet.Key) (*ChatClient, error) {
	if key == nil || !key.IsOpen() {
		return nil, fmt.Errorf("ivnalid key")
	}

	cc := &ChatClient{
		endpoint: addr,
		key:      key,
	}

	cc.reader = thread.NewThread(cc.reading)
	cc.writer = thread.NewThread(cc.writing)
	return cc, nil
}

func (cc *ChatClient) Register(in InputFunc, out <-chan []byte) {
	cc.in = in
	cc.out = out
}

func (cc *ChatClient) Online() error {
	url := url.URL{Scheme: "ws", Host: cc.endpoint, Path: service.CPUserOnline}
	wsConn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return err
	}

	onlineMsg := &pbs.ClientChatMsg{}
	if err := onlineMsg.Online(wsConn, cc.key); err != nil {
		return err
	}
	cc.conn = wsConn
	return nil
}

func (cc *ChatClient) Write() error {
	return nil
}

func (cc *ChatClient) reading(stop chan struct{}) {

	defer cc.ShutDown()

	for {
		_, message, err := cc.conn.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			return
		}

		select {
		case <-stop:
			fmt.Println("reading thread exit")
			return
		default:
			if cc.in == nil {
				continue
			}
			err := cc.in(message)
			if err != nil {
				fmt.Println("process input message err:", err)
				continue
			}
		}
	}
}

func (cc *ChatClient) writing(stop chan struct{}) {
	defer cc.ShutDown()

	for {
		select {
		case _ = <-cc.out:

		case <-stop:
			fmt.Println("write thread exit")
			_ = cc.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		}
	}
}

func (cc *ChatClient) ShutDown() {
	cc.reader.Stop()
	cc.writer.Stop()
	cc.conn.Close()
}
