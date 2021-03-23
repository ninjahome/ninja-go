package service

import "sync"

type UserTable struct {
	sync.RWMutex
	cache map[string]*wsUser
}

func (u *UserTable) Add(wu *wsUser) {
	u.Lock()
	defer u.Unlock()
	u.cache[wu.UID] = wu
}

func newUserTable() *UserTable {
	return &UserTable{
		cache: make(map[string]*wsUser),
	}
}
