package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"net/http"
	"path/filepath"
	"time"
)

const (
	DefaultWsPort              = 6666
	DefaultHost                = "localhost"
	DefaultPongWait            = 60 * time.Second
	DefaultPingPeriod          = (DefaultPongWait * 9) / 10
	DefaultWriteWait           = 10 * time.Second
	DefaultWsBuffer            = 1 << 21
	DefaultWsMsgNoFromCli      = 1 << 8
	DefaultUnreadMsgNoPerQuery = 1 << 13
	DefaultHandShakeTimeOut    = time.Second * 3
	DefaultDataBaseDir         = "Msg"
)

type Config struct {
	PingPeriod             time.Duration  `json:"ping.timeout"`
	PongWait               time.Duration  `json:"pong.timeout"`
	WriteWait              time.Duration  `json:"write.timeout"`
	HandShakeTimeout       time.Duration  `json:"handshake.timeout"`
	WsIOBufferSize         utils.ByteSize `json:"io.buffer.size"`
	MaxUnreadMsgNoPerQuery int            `json:"max.unread.size"`
	WsMsgNoFromCli         int            `json:"max.client.msg.no"`
	WsIP                   string         `json:"ws.ip"`
	WsPort                 int16          `json:"ws.port"`
	DataBaseDir            string         `json:"ws.msg.database"`
}

func (c Config) String() string {
	s := fmt.Sprintf("\n--------------------Websocket Config-------------------")
	s += fmt.Sprintf("\nws ping timeout:\t%s", c.PingPeriod)
	s += fmt.Sprintf("\nws pong timeout:\t%s", c.PongWait)
	s += fmt.Sprintf("\nwrite wait timeout:\t%s", c.WriteWait)
	s += fmt.Sprintf("\nws handshake timeout:\t%s", c.HandShakeTimeout)
	s += fmt.Sprintf("\nws io buffer size:\t%s", c.WsIOBufferSize)
	s += fmt.Sprintf("\nmax message for cli:\t%d", c.WsMsgNoFromCli)
	s += fmt.Sprintf("\nmax query unread msg:\t%d", c.MaxUnreadMsgNoPerQuery)
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
		PingPeriod:             DefaultPingPeriod,
		PongWait:               DefaultPongWait,
		WriteWait:              DefaultWriteWait,
		WsIOBufferSize:         DefaultWsBuffer,
		WsMsgNoFromCli:         DefaultWsMsgNoFromCli,
		MaxUnreadMsgNoPerQuery: DefaultUnreadMsgNoPerQuery,
		HandShakeTimeout:       DefaultHandShakeTimeOut,
		WsIP:                   DefaultHost,
		WsPort:                 DefaultWsPort,
		DataBaseDir:            dir,
	}
}

func (c *Config) newUpGrader() *websocket.Upgrader {

	return &websocket.Upgrader{
		HandshakeTimeout: _wsConfig.HandShakeTimeout,
		ReadBufferSize:   int(_wsConfig.WsIOBufferSize),
		WriteBufferSize:  int(_wsConfig.WsIOBufferSize),
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
}

func (c *Config) newWSServer(handler http.Handler) *http.Server {
	endPoint := fmt.Sprintf("%s:%d", c.WsIP, c.WsPort)
	return &http.Server{
		Addr:    endPoint,
		Handler: handler,
	}
}
