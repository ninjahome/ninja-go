package thread

import (
	"fmt"
	"time"
)

type Runner func(stop chan struct{})
type Thread struct {
	ID        int
	name      string
	stop      chan struct{}
	runFunc   Runner
	startTime time.Time
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

	if _, ok := _inst.queue[t.ID]; !ok {
		return
	}

	t.stop <- struct{}{}

	delete(_inst.nameID, t.name)
	delete(_inst.queue, t.ID)
	close(t.stop)
	t.stop = nil
}
