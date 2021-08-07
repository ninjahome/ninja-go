package client

import (
	crand "crypto/rand"
	"errors"
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
	"strings"
	"sync"
	"time"
)

var (
	DefaultBootWsService = []string{
		"39.99.198.143:16666",
		"47.113.87.58:16666",
		"118.186.203.36:16666",
	}

	ErrUnknownMsg    = fmt.Errorf("unknown websocket message")
	ErrNoMsgCallback = fmt.Errorf("no message reciver")
)

const (
	DevType_IOS     = 1
	DevType_Android = 2
)

type SafeWriteConn struct {
	*websocket.Conn
	writeLock sync.Mutex
}

func (swc *SafeWriteConn) WriteMessage(messageType int, data []byte) error {
	swc.writeLock.Lock()
	defer swc.writeLock.Unlock()

	return swc.Conn.WriteMessage(messageType, data)
}

type WSClient struct {
	DevTyp      int
	DeviceToken string
	IsOnline    bool
	endpoint    string
	wsConn      *SafeWriteConn
	key         *wallet.Key
	reader      *thread.Thread
	callback    CliCallBack
	peerKeys    map[string][]byte
}

func RandomBootNode() string {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(DefaultBootWsService))
	return DefaultBootWsService[idx]

	//return DefaultBootWsService[0]
}

type CliCallBack interface {
	ImmediateMessage(*pbs.WSCryptoMsg) error
	ImmediateGMessage(msg *pbs.WSCryptoGroupMsg) error
	WebSocketClosed()
	//UnreadMsg(*pbs.WSUnreadAck) error
	OnlineSuccess()
}

//devType DevType_IOS/DevType_Android

func NewWSClient(deviceToken, addr string, devType int, key *wallet.Key, cb CliCallBack) (*WSClient, error) {
	if key == nil || !key.IsOpen() {
		return nil, fmt.Errorf("ivnalid key")
	}

	cc := &WSClient{
		DevTyp:      devType,
		DeviceToken: deviceToken,
		endpoint:    addr,
		key:         key,
		IsOnline:    false,
		peerKeys:    make(map[string][]byte),
		callback:    cb,
	}

	cc.reader = thread.NewThread(cc.reading)
	return cc, nil
}

func (cc *WSClient) Online() error {
	u := url.URL{Scheme: "ws", Host: cc.endpoint, Path: websocket2.CPUserOnline}

	dialer := websocket.DefaultDialer

	dialer.ReadBufferSize = websocket2.DefaultWsBuffer
	dialer.WriteBufferSize = websocket2.DefaultWsBuffer

	wsConn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	onlineMsg := &pbs.WsMsg{}
	if err := onlineMsg.Online(wsConn, cc.key, cc.DeviceToken, cc.DevTyp); err != nil {
		return err
	}

	swc := &SafeWriteConn{Conn: wsConn}

	cc.wsConn = swc
	wsConn.SetPingHandler(func(appData string) error {
		fmt.Println("ping pong time......")
		return swc.WriteMessage(websocket.PongMessage, []byte{})
	})

	cc.wsConn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("------>websocket is close")
		cc.callback.WebSocketClosed()
		return nil
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
	return cc.wsConn.WriteMessage(websocket.BinaryMessage, msgWrap.Data())
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

func (cc *WSClient) recoverGroupKey(from string, gekey []*pbs.GroupEncryptKey) ([]byte, error) {
	lfrom := strings.ToLower(from)
	key, err := cc.getAesKey(lfrom)
	if err != nil {
		return nil, err
	}

	var gkey []byte
	for i := 0; i < len(gekey); i++ {
		if strings.ToLower(gekey[i].MemberId) == cc.key.Address.String() {
			gkey = gekey[i].EncryptKey
			break
		}
	}

	if gkey == nil {
		return nil, errors.New("i'm not in group")
	}

	return openssl.AesECBDecrypt(gkey, key, openssl.PKCS7_PADDING)
}

func getGMsgKey() []byte {
	key := make([]byte, 32)
	for {
		if n, err := crand.Read(key); err != nil {
			continue
		} else if n != len(key) {
			continue
		}

		return key
	}
}

func (cc *WSClient) groupEncryptKey(to []string) (gekey []*pbs.GroupEncryptKey, key []byte, err error) {
	from := strings.ToLower(cc.key.Address.String())
	gkey := getGMsgKey()

	gekey = make([]*pbs.GroupEncryptKey, 0)

	for i := 0; i < len(to); i++ {
		lto := strings.ToLower(to[i])
		if from == lto {
			gekey = append(gekey,&pbs.GroupEncryptKey{
				MemberId: lto,
			})
			continue
		}
		if tokey, err := cc.getAesKey(lto); err != nil {
			return nil, nil, err
		} else {

			dst, _ := openssl.AesECBEncrypt(gkey, tokey, openssl.PKCS7_PADDING)

			ek := &pbs.GroupEncryptKey{
				MemberId:   lto,
				EncryptKey: dst,
			}

			gekey = append(gekey, ek)
		}
	}

	return gekey, gkey, nil
}

func (cc *WSClient) GWrite(to []string, body []byte) error {
	if !cc.IsOnline {
		return fmt.Errorf("please online yourself first")
	}
	if !cc.IsOnline {
		return fmt.Errorf("please online yourself first")
	}

	gekey, gkey, err := cc.groupEncryptKey(to)
	if err != nil {
		return err
	}

	from := cc.key.Address.String()

	msgWrap := &pbs.WsMsg{}

	if err := cc.wsConn.WriteMessage(websocket.BinaryMessage,
		msgWrap.AesCryptGData(from, gekey, body, gkey)); err != nil {
		return err
	}
	return nil
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

	if err := cc.wsConn.WriteMessage(websocket.BinaryMessage,
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
		switch wsMsg.Payload.(type) {
		case *pbs.WsMsg_Message:
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
		case *pbs.WsMsg_GroupMessage:
			msgWrap, ok := wsMsg.Payload.(*pbs.WsMsg_GroupMessage)
			if !ok {
				return ErrUnknownMsg
			}
			msg := msgWrap.GroupMessage

			key, err := cc.recoverGroupKey(msg.From, msg.To)
			if err != nil {
				return err
			}

			dst, _ := openssl.AesECBDecrypt(msg.PayLoad, key, openssl.PKCS7_PADDING)
			msg.PayLoad = dst
			if err := cc.callback.ImmediateGMessage(msg); err != nil {
				return err
			}
		default:
			return errors.New("not a correct message")

		}

	case pbs.WsMsgType_UnreadAck:
		ack, ok := wsMsg.Payload.(*pbs.WsMsg_UnreadAck)
		if !ok {
			return ErrUnknownMsg
		}

		for i := 0; i < len(ack.UnreadAck.Payload); i++ {
			unreadmsg := ack.UnreadAck.Payload[i]
			switch unreadmsg.CryptoMsg.(type) {
			case *pbs.WsUnreadAckMsg_Payload:
				msg := unreadmsg.CryptoMsg.(*pbs.WsUnreadAckMsg_Payload)
				key, err := cc.getAesKey(msg.Payload.From)
				if err != nil {
					continue
				}
				dst, _ := openssl.AesECBDecrypt(msg.Payload.PayLoad, key, openssl.PKCS7_PADDING)
				msg.Payload.PayLoad = dst
				if err := cc.callback.ImmediateMessage(msg.Payload); err != nil {
					continue
				}
			case *pbs.WsUnreadAckMsg_GPayload:
				gmsg := unreadmsg.CryptoMsg.(*pbs.WsUnreadAckMsg_GPayload)
				gpayload := gmsg.GPayload

				key, err := cc.recoverGroupKey(gpayload.From, gpayload.To)
				if err != nil {
					continue
				}

				dst, _ := openssl.AesECBDecrypt(gpayload.PayLoad, key, openssl.PKCS7_PADDING)
				gpayload.PayLoad = dst
				if err := cc.callback.ImmediateGMessage(gpayload); err != nil {
					continue
				}
			}
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

func (ws *WSClient) Address() string {
	return ws.key.Address.String()
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
