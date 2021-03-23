package service

import "sync"

type OnlineMap struct {
	sync.RWMutex
	cache map[string]bool
}

func (m *OnlineMap) NotifyPeers(ws *wsUser) {
	m.Lock()
	defer m.Unlock()
	m.cache[ws.UID] = true
	//TODO::
}

func newOnlineSet() *OnlineMap {
	return &OnlineMap{
		cache: make(map[string]bool),
	}
}
