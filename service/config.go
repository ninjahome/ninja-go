package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

const (
	DefaultWsPort       = 8886
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

func (c *Config) newUpGrader() *websocket.Upgrader {

	return &websocket.Upgrader{
		HandshakeTimeout: time.Second * 3,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (c *Config) newWSServer(handler http.Handler) *http.Server {
	endPoint := fmt.Sprintf("%s:%d", c.WsIP, c.WsPort)
	return &http.Server{
		Addr:         endPoint,
		Handler:      handler,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		IdleTimeout:  c.IdleTimeout,
	}
}
