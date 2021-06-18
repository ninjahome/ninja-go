package iosLib

import (
	"encoding/base64"
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

type IosApp struct {
	key       *wallet.Key
	cb        AppCallBack
	wsEnd     string
	websocket *client.WSClient
	unreadSeq int64
}

func (i IosApp) OnlineSuccess() {
	if err := _inst.websocket.PullUnreadMsg(i.unreadSeq); err != nil {
		fmt.Println(err.Error())
	}
}

func (i IosApp) ImmediateMessage(msg *pbs.WSCryptoMsg) error {
	return i.callback(msg)
}

func (i IosApp) WebSocketClosed() {
	i.cb.WebSocketClosed()
}

func (i IosApp) UnreadMsg(ack *pbs.WSUnreadAck) error {
	payload := ack.Payload

	for j:=0;j<len(payload);j++{
		i.callback(payload[j])
	}

	return nil
}


func (i IosApp)callback(msg *pbs.WSCryptoMsg) error {
	switch msg.Typ {
	case pbs.ChatMsgType_TextMessage:
		return i.cb.TextMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	case pbs.ChatMsgType_MapMessage:
		return i.cb.MapMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	case pbs.ChatMsgType_ImageMessage:
		return i.cb.ImageMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	case pbs.ChatMsgType_VoiceMessage:
		return i.cb.VoiceMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
	default:
		return errors.New("msg not recognize")
	}

	return nil
}

func UnmarshalGoByte(s string) []byte {
	b, e := base64.StdEncoding.DecodeString(s)
	if e != nil {
		return nil
	}
	return b
}

var _inst = &IosApp{unreadSeq: 0}

type AppCallBack interface {
	VoiceMessage(from, to string, payload []byte, time int64) error
	ImageMessage(from, to string, payload []byte, time int64) error
	MapMessage(from, to string, payload []byte, time int64) error
	TextMessage(from, to string, payload []byte, time int64) error
	WebSocketClosed()
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
