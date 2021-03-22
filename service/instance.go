package service

import (
	"sync"
	"time"
)

type HTTPTimeouts struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type WebSocketService struct {
}

var (
	_instance *WebSocketService
	once      sync.Once
)

func Inst() *WebSocketService {
	once.Do(func() {
		_instance = newWebSocket()
	})
	return _instance
}

func newWebSocket() *WebSocketService {

	if _srvConfig == nil {
		panic("init service config first")
	}

	ws := &WebSocketService{}
	return ws
}

func (ws *WebSocketService) StartService() error {
	return nil
}
