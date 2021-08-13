package config

import (
	"encoding/json"
	"github.com/ninjahome/ninja-go/tools"
	"os"
	"path"
	"strconv"
)

const (
	homeDir        = ".extra_node"
	configFileName = "config.json"
	ListenPort     = 9099
)

type Config struct {
	ListenAddr string `json:"listen_addr"`
	WalletFile string `json:"wallet_file"`
}

var extra_config *Config

func DefaultConfig() *Config {
	return &Config{
		ListenAddr: ":" + strconv.Itoa(ListenPort),
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

func getConfigFileName() (string, error) {
	h, err := tools.Home()
	if err != nil {
		return "", err
	}

	filename := path.Join(h, homeDir, configFileName)

	return filename, nil
}

func (c *Config) Save() error {
	filename, _ := getConfigFileName()

	j, _ := json.MarshalIndent(*c, "\t", " ")
	return tools.Save2File(j, filename)
}

func InitConfig() (*Config, error) {
	filename, err := getConfigFileName()
	if err != nil {
		return nil, err
	}

	var c *Config

	if !tools.FileExists(filename) {
		c = DefaultConfig()
		c.Save()
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

func InitSystem() error {
	h, _ := GetExtraHome()
	if tools.FileExists(h) {
		return nil
	}

	return os.Mkdir(h, 0755)
}
