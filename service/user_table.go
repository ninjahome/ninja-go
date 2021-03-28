package service

import "sync"

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

func (ut *UserTable) del(user *wsUser) {
	ut.Lock()
	defer ut.Unlock()
	delete(ut.cache, user.UID)
}

func newUserTable() *UserTable {
	return &UserTable{
		cache: make(map[string]*wsUser),
	}
}
