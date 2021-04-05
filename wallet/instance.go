package wallet

import (
	"fmt"

	"github.com/ninjahome/bls-wallet/bls"
	"github.com/ninjahome/ninja-go/common"
	"sync"
)

func init() {
	if err := bls.Init(bls.BLS12_381); err != nil {
		panic(err)
	}
	if err := bls.SetETHmode(bls.EthModeDraft07); err != nil {
		panic(err)
	}
}

type Wallet interface {
	Active(password, selectAddr string) error
	KeyInUsed() *Key
	CreateNewKey(auth string) error
}

var _instance Wallet
var once sync.Once

func Inst() Wallet {
	once.Do(func() {
		_instance = newWallet()
	})
	return _instance
}

func newWallet() Wallet {
	cw := &NinjaWallet{
		keys: make(map[common.Address]*Key),
	}
	return cw
}

type NinjaWallet struct {
	activeKey *Key
	sync.RWMutex
	keys map[common.Address]*Key
}

func (c *NinjaWallet) Active(password, addr string) error {
	if config == nil {
		return fmt.Errorf("please init wallet instance config first")
	}
	ks := NewKeyStore(config.Dir)
	validKeyFiles, err := ks.ValidKeyFiles()
	if err != nil {
		return err
	}
	var selFile string
	if addr == "" && len(validKeyFiles) == 1 {
		for a, path := range validKeyFiles {
			selFile = path
			addr = a
		}
	} else {
		path, ok := validKeyFiles[addr]
		if !ok {
			return fmt.Errorf("no valid key file for addr:[%s]", addr)
		}
		selFile = path
	}

	address, _ := common.HexToAddress(addr)
	key, err := ks.GetKey(address, selFile, password)
	if err != nil {
		return err
	}

	c.activeKey = key
	c.Lock()
	c.keys[address] = key
	c.Unlock()
	return nil
}

func (c *NinjaWallet) KeyInUsed() *Key {
	return c.activeKey
}

func (c *NinjaWallet) CreateNewKey(auth string) error {
	ks := NewKeyStore(config.Dir)
	key := NewKey()
	if err := ks.StoreKey(key, auth); err != nil {
		return err
	}
	return nil
}
