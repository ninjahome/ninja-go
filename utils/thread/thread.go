package thread

import "time"

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
		t.runFunc(t.stop)
	}()
}

func (t *Thread) Stop() {

	t.stop <- struct{}{}

	_inst.locker.Lock()
	defer _inst.locker.Unlock()

	delete(_inst.nameID, t.name)
	delete(_inst.queue, t.ID)
}
