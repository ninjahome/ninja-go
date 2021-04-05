package iosLib

import "github.com/ninjahome/ninja-go/wallet"

func NewWallet(auth string) string {
	key := wallet.NewKey()
	return key.StoreString(auth)
}

type IosApp struct {
	key *wallet.Key
}

var _inst = &IosApp{}

func InitApp(cipherTxt, auth string) error {
	parsedKey, err := wallet.LoadKeyFromJsonStr(cipherTxt, auth)
	if err != nil {
		return err
	}
	_inst.key = parsedKey
	return nil
}
