package service

import (
	"fmt"
	"time"
)

const (
	DefaultWsPort       = 8888
	DefaultReadTimeout  = 30 * time.Second
	DefaultWriteTimeout = 30 * time.Second
	DefaultIdleTimeout  = 120 * time.Second
	DefaultHost         = "localhost"
)

type Config struct {
	ReadTimeout  time.Duration `json:"http.r.timeout"`
	WriteTimeout time.Duration `json:"http.w.timeout"`
	IdleTimeout  time.Duration `json:"http.i.timeout"`
	WsIP         string        `json:"ws.ip"`
	WsPort       int16         `json:"ws.port"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------service Config------------")
	s += fmt.Sprintf("\nhttp read timeout:%20d", c.ReadTimeout)
	s += fmt.Sprintf("\nhttp writ timeout:%20d", c.WriteTimeout)
	s += fmt.Sprintf("\nhttp idle timeout:%20d", c.IdleTimeout)
	s += fmt.Sprintf("\nws ip:%20s", c.WsIP)
	s += fmt.Sprintf("\nws port:%20d", c.WsPort)
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

var _srvConfig *Config = nil

func InitConfig(c *Config) {
	_srvConfig = c
}

func DefaultConfig() *Config {

	return &Config{
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		IdleTimeout:  DefaultIdleTimeout,
		WsIP:         DefaultHost,
		WsPort:       DefaultWsPort,
	}
}
