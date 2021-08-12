package proxy

import (
	"fmt"
	"strconv"
)

const (
	proxyListenPort = 8088
)

var proxyAddr = []string{
	"39.99.198.143:9099",
	"47.113.87.58:9099",
}

type Config struct {
	ListenAddr string   `json:"listen_addr"`
	ProxyAddr  []string `json:"proxy_addr"`
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
	s += fmt.Sprintf("\r\n-------------------------------------------------------\r\n")
	return s
}

func InitConfig(c *Config) {
	_proxyConfig = c
}

func DefaultConfig() *Config {

	l := ":" + strconv.Itoa(proxyListenPort)

	return &Config{
		ListenAddr: l,
		ProxyAddr:  proxyAddr,
	}
}

func (c *Config) NewProxyWebServer() *WebProxyServer {
	return NewWebServer(c.ListenAddr, c.ProxyAddr)
}
