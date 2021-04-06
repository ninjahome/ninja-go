package iosLib

import (
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
	key *wallet.Key
	cb  AppCallBack
}

var _inst = &IosApp{}

type AppCallBack interface {
	ImmediateMessage(from, to string, payload []byte, time int64) error
	WebSocketClosed()
	UnreadMsg(jsonData []byte) error
}

func InitApp(cipherTxt, auth string, callback AppCallBack) error {
	parsedKey, err := wallet.LoadKeyFromJsonStr(cipherTxt, auth)
	if err != nil {
		return err
	}
	_inst.key = parsedKey
	_inst.cb = callback
	return nil
}

func WriteMessage(To string, payload []byte) {

}
