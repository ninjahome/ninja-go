package androidlib


import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ninjahome/ninja-go/common"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
	"github.com/ninjahome/ninja-go/wallet"
)

func NewWallet(auth string) string {
	key := wallet.NewKey()
	_inst.key = key
	return key.StoreString(auth)
}

func ActiveAddress() string {
	if _inst.key == nil {
		return ""
	}

	return _inst.key.Address.String()
}

type AndroidAPP struct {
	key       *wallet.Key
	cb        AppCallBack
	wsEnd     string
	websocket *client.WSClient
	unreadSeq int64
}

func (a AndroidAPP) OnlineSuccess() {
	if err := _inst.websocket.PullUnreadMsg(a.unreadSeq); err != nil {
		fmt.Println(err.Error())
	}
}

func (a AndroidAPP) ImmediateMessage(msg *pbs.WSCryptoMsg) error {
	if msg == nil{
		return errors.New("msg is nil")
	}
	switch msg.Typ {
	case pbs.ChatMsgType_TextMessage:
		return a.cb.ImmediateMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	case pbs.ChatMsgType_MapMessage:
		return a.cb.ImmediateMapMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	case pbs.ChatMsgType_ImageMessage:
		return a.cb.ImmediateImageMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	case pbs.ChatMsgType_VoiceMessage:
		return a.cb.ImmediateVoiceMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	default:
		return errors.New("msg not recognize")
	}
}

func (a AndroidAPP) WebSocketClosed() {
	a.cb.WebSocketClosed()
}

func (a AndroidAPP) UnreadMsg(ack *pbs.WSUnreadAck) error {
	payload := ack.Payload

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return a.cb.UnreadMsg(data)
}

func UnmarshalGoByte(s string) []byte {
	b, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		return nil
	}
	return b
}

var _inst = &AndroidAPP{unreadSeq: 0}

type AppCallBack interface {
	ImmediateVoiceMessage(from, to string, payload []byte, time int64) error
	ImmediateImageMessage(from, to string, payload []byte, time int64) error
	ImmediateMapMessage(from, to string, payload []byte, time int64) error
	ImmediateMessage(from, to string, payload []byte, time int64) error
	WebSocketClosed()
	UnreadMsg(jsonData []byte) error
}

func ConfigApp(addr string, callback AppCallBack) {

	if addr == "" {
		addr = client.RandomBootNode()
	}
	fmt.Println("======>", addr)
	_inst.wsEnd = addr
	_inst.cb = callback
}

func ActiveWallet(cipherTxt, auth string) error {
	key, err := wallet.LoadKeyFromJsonStr(cipherTxt, auth)
	if err != nil {
		return err
	}
	_inst.key = key
	ws, err := client.NewWSClient(_inst.wsEnd, key, _inst) //202.182.101.145//167.179.78.33//127.0.0.1//
	if err != nil {
		return err
	}
	_inst.websocket = ws

	return ws.Online()
}

func WalletIsOpen() bool {
	return _inst.key != nil && _inst.key.IsOpen()
}

func WSIsOnline() bool {
	return _inst.websocket != nil && _inst.websocket.IsOnline
}

func WSOnline() error {
	if WSIsOnline() {
		return nil
	}

	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}

	return _inst.websocket.Online()
}

func WSOffline() {
	if _inst.websocket == nil {
		fmt.Println("nil, no need to offline")
		return
	}

	_inst.websocket.ShutDown()
}

func WriteMessage(to string, payload []byte) error {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	return _inst.websocket.Write(to, pbs.ChatMsgType_TextMessage, payload)
}

func WriteMapMessage(to string, payload []byte) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	return _inst.websocket.Write(to, pbs.ChatMsgType_MapMessage, payload)
}

func WriteImageMessage(to string, payload []byte) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	return _inst.websocket.Write(to, pbs.ChatMsgType_ImageMessage, payload)
}

func WriteVoiceMessage(to string, payload []byte) error  {
	if _inst.websocket == nil {
		return fmt.Errorf("init application first please")
	}
	if !_inst.websocket.IsOnline {
		if err := _inst.websocket.Online(); err != nil {
			return err
		}
	}

	return _inst.websocket.Write(to,pbs.ChatMsgType_VoiceMessage, payload)
}


func IsValidNinjaAddr(addr string) bool {
	_, err := common.HexToAddress(addr)
	if err != nil {
		return false
	}
	return true
}
