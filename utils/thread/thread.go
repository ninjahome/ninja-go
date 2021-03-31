package thread

import (
	"fmt"
	"time"
)

type Runner func(stop chan struct{})
type BeforeExit func()
type Thread struct {
	ID        int
	name      string
	stop      chan struct{}
	runFunc   Runner
	beFunc    BeforeExit
	startTime time.Time
}

func (t *Thread) WillExit(be BeforeExit) {
	t.beFunc = be
}

func (t *Thread) Run() {
	go func() {
		if r := recover(); r != nil {
			t.Stop()
			fmt.Printf("thread panice by:%s", r)
		}
		t.runFunc(t.stop)
	}()
}

func (t *Thread) IsAlive() bool {
	_inst.locker.RLock()
	defer _inst.locker.RUnlock()
	_, ok := _inst.queue[t.ID]
	return ok
}

func (t *Thread) Stop() {
	_inst.locker.Lock()
	defer _inst.locker.Unlock()
	if t.stop == nil {
		return
	}

	if t.beFunc != nil {
		t.beFunc()
	}

	t.stop <- struct{}{}

	close(t.stop)
	t.stop = nil

	delete(_inst.nameID, t.name)
	delete(_inst.queue, t.ID)
}

func (t *Thread) String() string {
	return fmt.Sprintf("id:%d, name:%s, time:%s", t.ID, t.name, t.startTime)
}
