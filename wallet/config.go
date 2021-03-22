package wallet

import (
	"fmt"
	"path/filepath"
)

const KeyStoreScheme = "keystore"
const TestKeyStoreScheme = "test_keystore"

type Config struct {
	Dir string `json:"keystore"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------wallet Config------------")
	s += fmt.Sprintf("\nkey store dir:%20s", c.Dir)
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var config *Config = nil

func InitConfig(c *Config) {
	config = c
}

func DefaultConfig(isMain bool, base string) *Config {

	var dir string
	if isMain {
		dir = filepath.Join(base, string(filepath.Separator), KeyStoreScheme)
	} else {
		dir = filepath.Join(base, string(filepath.Separator), TestKeyStoreScheme)
	}
	c := &Config{Dir: dir}
	return c
}
