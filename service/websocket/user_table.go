package websocket

import (
	"encoding/json"
	"sync"
)

type UserTable struct {
	sync.RWMutex
	cache map[string]*wsUser
}

func (ut *UserTable) add(wu *wsUser) {
	ut.Lock()
	defer ut.Unlock()
	ut.cache[wu.UID] = wu
}

func (ut *UserTable) get(to string) (*wsUser, bool) {
	ut.RLock()
	defer ut.RUnlock()
	us, ok := ut.cache[to]
	return us, ok
}

func (ut *UserTable) del(uid string) {
	ut.Lock()
	defer ut.Unlock()
	delete(ut.cache, uid)
}

func newUserTable() *UserTable {
	return &UserTable{
		cache: make(map[string]*wsUser),
	}
}

func (ut *UserTable) DumpContent() string {
	bts, _ := json.Marshal(ut.cache)
	return string(bts)
}

func (ut *UserTable) AllUid() []string {
	ut.RLock()
	defer ut.RUnlock()
	ids := make([]string, len(ut.cache))
	var idx = 0
	for uid, _ := range ut.cache {
		ids[idx] = uid
		idx++
	}
	return ids
}
