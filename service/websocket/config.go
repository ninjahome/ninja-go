package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"path/filepath"
	"time"
)

const (
	DefaultWsPort           = 6666
	DefaultReadTimeout      = 30 * time.Second
	DefaultWriteTimeout     = 30 * time.Second
	DefaultIdleTimeout      = 120 * time.Second
	DefaultHost             = "localhost"
	DefaultPongWait         = 60 * time.Second
	DefaultPingPeriod       = (DefaultPongWait * 9) / 10
	DefaultWriteWait        = 10 * time.Second
	DefaultWsBuffer         = 1 << 21
	DefaultWsMsgQueue       = 1 << 26
	DefaultWsMsgSizePerUser = 1 << 16
	DefaultHandShakeTimeOut = time.Second * 3
	DefaultDataBaseDir      = "Msg"
)

type Config struct {
	ReadTimeout      time.Duration `json:"http.r.timeout"`
	WriteTimeout     time.Duration `json:"http.w.timeout"`
	IdleTimeout      time.Duration `json:"http.i.timeout"`
	PingPeriod       time.Duration `json:"ws.ping.timeout"`
	PongWait         time.Duration `json:"ws.pong.timeout"`
	WriteWait        time.Duration `json:"ws.w.timeout"`
	HsTimeout        time.Duration `json:"ws.hs.timeout"`
	WsBufferSize     int           `json:"ws.buffer.size"`
	WsMsgQueueSize   int           `json:"ws.msg.size"`
	WsMsgSizePerUser int           `json:"ws.user_msg.size"`
	WsIP             string        `json:"ws.ip"`
	WsPort           int16         `json:"ws.port"`
	DataBaseDir      string        `json:"ws.msg.database"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------------Websocket Config-------------------")
	s += fmt.Sprintf("\nhttp read timeout:\t%d", c.ReadTimeout)
	s += fmt.Sprintf("\nhttp writ timeout:\t%d", c.WriteTimeout)
	s += fmt.Sprintf("\nhttp idle timeout:\t%d", c.IdleTimeout)
	s += fmt.Sprintf("\nws ping timeout:\t%d", c.PingPeriod)
	s += fmt.Sprintf("\nws pong timeout:\t%d", c.PongWait)
	s += fmt.Sprintf("\nws wait timeout:\t%d", c.WriteWait)
	s += fmt.Sprintf("\nws handshake timeout:\t%d", c.HsTimeout)
	s += fmt.Sprintf("\nws buffer size:\t\t%d", c.WsBufferSize)
	s += fmt.Sprintf("\nws msg queue size:\t%d", c.WsMsgQueueSize)
	s += fmt.Sprintf("\nws msg size/user:\t%d", c.WsMsgSizePerUser)
	s += fmt.Sprintf("\nws ip:\t\t\t%s", c.WsIP)
	s += fmt.Sprintf("\nmessage database dir:\t%s", c.DataBaseDir)
	s += fmt.Sprintf("\nws port:\t\t%d", c.WsPort)
	s += fmt.Sprintf("\n-------------------------------------------------------\n")
	return s
}

var _wsConfig *Config = nil

func InitConfig(c *Config) {
	_wsConfig = c
}

func DefaultConfig(isMain bool, base string) *Config {

	var dir string
	if isMain {
		dir = filepath.Join(base, string(filepath.Separator), DefaultDataBaseDir)
	} else {
		dir = filepath.Join(base, string(filepath.Separator), DefaultDataBaseDir+"_test")
	}

	return &Config{
		ReadTimeout:      DefaultReadTimeout,
		WriteTimeout:     DefaultWriteTimeout,
		IdleTimeout:      DefaultIdleTimeout,
		PingPeriod:       DefaultPingPeriod,
		PongWait:         DefaultPongWait,
		WriteWait:        DefaultWriteWait,
		WsBufferSize:     DefaultWsBuffer,
		WsMsgQueueSize:   DefaultWsMsgQueue,
		WsMsgSizePerUser: DefaultWsMsgSizePerUser,
		HsTimeout:        DefaultHandShakeTimeOut,
		WsIP:             DefaultHost,
		WsPort:           DefaultWsPort,
		DataBaseDir:      dir,
	}
}

func (c *Config) newUpGrader() *websocket.Upgrader {

	return &websocket.Upgrader{
		HandshakeTimeout: _wsConfig.HsTimeout,
		ReadBufferSize:   _wsConfig.WsBufferSize,
		WriteBufferSize:  _wsConfig.WsBufferSize,
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
