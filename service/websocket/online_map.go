package websocket

import (
	"encoding/json"
	"fmt"
	"sync"
)

type OnlineMap struct {
	sync.RWMutex
	syncedTimes int
	lines       map[string]bool
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
		syncedTimes: 0,
		lines:       make(map[string]bool),
	}
}

func (m *OnlineMap) AllUid() []string {
	m.RLock()
	defer m.RUnlock()

	var i = 0
	keys := make([]string, len(m.lines))
	for k := range m.lines {
		keys[i] = k
		i++
	}
	return keys
}

func (m *OnlineMap) DumpContent() string {
	m.RLock()
	ca := m.lines
	m.RUnlock()

	bts, _ := json.Marshal(ca)
	return string(bts) + fmt.Sprintf("size=[%d]", len(m.lines))
}

func (m *OnlineMap) addBatch(uid []string) {
	m.Lock()
	defer m.Unlock()
	m.syncedTimes++
	for _, id := range uid {
		m.lines[id] = true
	}
}

func (m *OnlineMap) hasSynced() bool {
	return m.syncedTimes > 0
}
