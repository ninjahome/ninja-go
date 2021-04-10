package iosLib

import (
	"encoding/json"
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
}

func (i IosApp) ImmediateMessage(msg *pbs.WSCryptoMsg) error {
	return i.cb.ImmediateMessage(msg.From, msg.To, msg.PayLoad, msg.UnixTime)
}

func (i IosApp) WebSocketClosed() {
	i.cb.WebSocketClosed()
}

func (i IosApp) UnreadMsg(ack *pbs.WSUnreadAck) error {
	payload := ack.Payload
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return i.cb.UnreadMsg(data)
}

var _inst = &IosApp{}

type AppCallBack interface {
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
	fmt.Println("debug infos:", cipherTxt, auth)
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

	return _inst.websocket.Write(to, payload)
}

func IsValidNinjaAddr(addr string) bool {
	_, err := common.HexToAddress(addr)
	if err != nil {
		return false
	}
	return true
}
