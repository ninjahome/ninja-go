package thread

import (
	"fmt"
	"sync"
	"time"
)

var _inst *Manager

func init() {
	_inst = &Manager{
		queue:  make(map[int]*Thread),
		nameID: make(map[string]int),
	}
}
func Inst() *Manager {
	return _inst
}

type Manager struct {
	Seq    int
	locker sync.RWMutex
	queue  map[int]*Thread
	nameID map[string]int
}

func NewThreadWithName(name string, runner Runner) *Thread {
	t := &Thread{
		ID:        _inst.Seq,
		name:      name,
		stop:      make(chan struct{}, 1),
		runFunc:   runner,
		startTime: time.Now(),
	}

	_inst.locker.Lock()
	defer _inst.locker.Unlock()

	_inst.queue[t.ID] = t
	_inst.nameID[name] = t.ID
	_inst.Seq++

	return t
}

func NewThread(runner Runner) *Thread {
	_inst.locker.RLock()
	name := fmt.Sprintf("Thread[%d]", _inst.Seq)
	_inst.locker.RUnlock()
	return NewThreadWithName(name, runner)
}

//-------------------------------debug and trace-------------------
func (m *Manager) PrintAllThread() {
	m.locker.RLock()
	defer m.locker.RUnlock()

	fmt.Println("\n----------------thread debug info-----------------")
	fmt.Printf("\nsize:\t%d", len(m.queue))
	for id, t := range m.queue {
		fmt.Printf("\nid:%d\tname:%s\tstart at:%s", id, t.name, t.startTime)
	}
	fmt.Println("\n--------------------------------------------------")
}

func (m *Manager) DebugThreadNo() int {
	m.locker.RLock()
	defer m.locker.RUnlock()
	return len(m.queue)
}

func (m *Manager) ThreadByName(name string) *Thread {
	m.locker.RLock()
	defer m.locker.RUnlock()
	id, ok := m.nameID[name]
	if !ok {
		return nil
	}
	return m.queue[id]
}

func (m *Manager) AllThread(group string) string {
	m.locker.RLock()
	defer m.locker.RUnlock()

	s := fmt.Sprintf("\n----------threads list----------")
	for name, id := range m.nameID {
		s += fmt.Sprintf("\n\t %d\t %s", id, name)
	}
	s += fmt.Sprintf("\n--------------------------------\n")
	return s
}
