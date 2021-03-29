package contact

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"
)

const (
	DefaultSrvPort      = 6667
	DefaultDataBaseDir  = "contact"
	DefaultHost         = "0.0.0.0"
	DefaultReadTimeout  = 10 * time.Second
	DefaultWriteTimeout = 20 * time.Second
	DefaultIdleTimeout  = 20 * time.Second
)

type Config struct {
	ReadTimeout  time.Duration `json:"http.r.timeout"`
	WriteTimeout time.Duration `json:"http.w.timeout"`
	IdleTimeout  time.Duration `json:"http.i.timeout"`
	DataBaseDir  string        `json:"contact.database"`
	SrvPort      int16         `json:"contact.port"`
	SrvIP        string        `json:"contact.ip"`
}

var _srvConfig *Config = nil

func (c Config) String() string {
	s := fmt.Sprintf("\n<-------------websocket Config------------")
	s += fmt.Sprintf("\nhttp read timeout:%20d", c.ReadTimeout)
	s += fmt.Sprintf("\nhttp writ timeout:%20d", c.WriteTimeout)
	s += fmt.Sprintf("\nhttp idle timeout:%20d", c.IdleTimeout)
	s += fmt.Sprintf("\nws ip:%20s", c.SrvIP)
	s += fmt.Sprintf("\nmessage database dir:%20s", c.DataBaseDir)
	s += fmt.Sprintf("\nws port:%20d", c.SrvPort)
	s += fmt.Sprintf("\n----------------------------------->\n")
	return s
}

func InitConfig(c *Config) {
	_srvConfig = c
}

func DefaultConfig(isMain bool, base string) *Config {
	var dir string
	if isMain {
		dir = filepath.Join(base, string(filepath.Separator), DefaultDataBaseDir)
	} else {
		dir = filepath.Join(base, string(filepath.Separator), DefaultDataBaseDir+"_test")
	}

	return &Config{

		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
		IdleTimeout:  DefaultIdleTimeout,
		SrvPort:      DefaultSrvPort,
		DataBaseDir:  dir,
		SrvIP:        DefaultHost,
	}
}

func (c *Config) newContactServer(handler http.Handler) *http.Server {
	endPoint := fmt.Sprintf("%s:%d", c.SrvIP, c.SrvPort)
	return &http.Server{
		Addr:         endPoint,
		Handler:      handler,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		IdleTimeout:  c.IdleTimeout,
	}
}
