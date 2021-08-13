package proxy

import (
	"fmt"
	"strconv"
)

const (
	ProxyListenPort = 8088

	infuraUrl   = "https://ropsten.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contactAddr = "0x52996249f64d760ac02c6b82866d92b9e7d02f06"
	tokenAddr   = "0x122938b76c071142ea6b39c34ffc38e5711cada1"
)

var proxyAddr = []string{
	"39.99.198.143:9099",
	"47.113.87.58:9099",
}

type Config struct {
	ListenAddr   string   `json:"listen_addr"`
	ProxyAddr    []string `json:"proxy_addr"`
	EthUrl       string   `json:"eth_url"`
	ContractAddr string   `json:"contract_addr"`
	TokenAddr    string   `json:"token_addr"`
}

var _proxyConfig *Config = nil

func (c *Config) String() string {
	s := fmt.Sprintf("\r\n--------------------Proxy Config--------------------")
	s += fmt.Sprintf("\r\nListen addr: %s", c.ListenAddr)
	s += fmt.Sprintf("\r\nServer Addr: ")
	for i := 0; i < len(c.ProxyAddr); i++ {
		pa := c.ProxyAddr[i]
		s += fmt.Sprintf("\r\n        %s ", pa)
	}

	s += fmt.Sprintf("\r\nEth access node: %s", c.EthUrl)
	s += fmt.Sprintf("\r\ncontract: %s", c.ContractAddr)
	s += fmt.Sprintf("\r\ntoken: %s", c.TokenAddr)

	s += fmt.Sprintf("\r\n-------------------------------------------------------\r\n")

	return s
}

func InitConfig(c *Config) {
	_proxyConfig = c
}

func DefaultConfig() *Config {
	l := ":" + strconv.Itoa(ProxyListenPort)

	var pa []string
	for i := 0; i < len(proxyAddr); i++ {
		pa = append(pa, "http://"+proxyAddr[i])
	}

	return &Config{
		ListenAddr:   l,
		ProxyAddr:    pa,
		EthUrl:       infuraUrl,
		ContractAddr: contactAddr,
		TokenAddr:    tokenAddr,
	}
}

func (c *Config) NewProxyWebServer() *WebProxyServer {
	return NewWebServer(c.ListenAddr, c.ProxyAddr)
}
