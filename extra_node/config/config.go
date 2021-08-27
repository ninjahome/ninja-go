package config

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/ninja-go/tools"
	"github.com/ninjahome/ninja-go/wallet"
	"math/big"
	"path"
	"strconv"
)

const (
	homeDir        = ".extra_node"
	configFileName = "config.json"
	ListenPort     = 9099

	infuraUrl   = "https://kovan.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contactAddr = "0x291F8A7353E460416095602e7BEc53a12cb5F0cC"
	tokenAddr   = "0x122938b76c071142ea6b39c34ffc38e5711cada1"
)

type Config struct {
	ListenAddr      string `json:"listen_addr"`
	WalletFile      string `json:"wallet_file"`
	EthUrl          string `json:"eth_url"`
	TokenAddr       string `json:"token_addr"`
	LicenseContract string `json:"license_contract"`
}

var extra_config *Config

func DefaultConfig() *Config {
	return &Config{
		ListenAddr:      ":" + strconv.Itoa(ListenPort),
		WalletFile:      ".wallet",
		EthUrl:          infuraUrl,
		TokenAddr:       tokenAddr,
		LicenseContract: contactAddr,
	}
}

func GetExtraHome() (string, error) {
	h, err := tools.Home()
	if err != nil {
		return "", err
	}

	dir := path.Join(h, homeDir)

	return dir, nil
}

func GetConfigFileName() (string, error) {
	h, err := tools.Home()
	if err != nil {
		return "", err
	}

	filename := path.Join(h, homeDir, configFileName)

	return filename, nil
}

func (c *Config) Save() error {
	if filename, err := GetConfigFileName(); err != nil {
		return err
	} else {
		var j []byte
		if j, err = json.MarshalIndent(*c, "\t", " "); err != nil {
			return err
		} else {
			return tools.Save2File(j, filename)
		}
	}
}

func InitConfig() (*Config, error) {
	filename, err := GetConfigFileName()
	if err != nil {
		return nil, err
	}

	var c *Config

	if !tools.FileExists(filename) {
		return nil, errors.New("please initial exnode")
	} else {
		var data []byte
		data, err = tools.OpenAndReadAll(filename)
		if err != nil {
			return nil, err
		}
		c = &Config{}

		if err = json.Unmarshal(data, c); err != nil {
			return nil, err
		}
	}

	extra_config = c

	return c, nil
}

func (c *Config) GetWalletFile() string {
	h, _ := GetExtraHome()

	return path.Join(h, c.WalletFile)

}

func GetExtraConfig() *Config {
	if extra_config == nil {
		if _, err := InitConfig(); err != nil {
			panic(err)
		}
	}
	return extra_config
}

func GetBalance(addr common.Address) float64  {
	cli,err:=ethclient.Dial(infuraUrl)
	if err!=nil{
		return 0
	}
	defer cli.Close()

	var b *big.Int
	b,err=cli.BalanceAt(context.TODO(),addr,nil)
	if err!=nil{
		return 0
	}

	return wallet.BalanceHuman(b)
}
