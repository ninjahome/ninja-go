package service

import (
	"sync"
)

type OnlineMap struct {
	sync.RWMutex
	lines map[string]bool
}

func (m *OnlineMap) add(uid string) {
	m.Lock()
	defer m.Unlock()
	m.lines[uid] = true
}

func (m *OnlineMap) contains(to string) bool {
	m.RLock()
	defer m.RUnlock()
	return m.lines[to]
}

func (m *OnlineMap) del(uid string) {
	m.RLock()
	defer m.RUnlock()
	delete(m.lines, uid)
}

func newOnlineSet() *OnlineMap {
	return &OnlineMap{
		lines: make(map[string]bool),
	}
}
