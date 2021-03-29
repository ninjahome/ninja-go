package contact

import "sync"

var (
	_instance *Service
	once      sync.Once
)

func Inst() *Service {
	once.Do(func() {
		_instance = newContactServer()
	})
	return _instance
}

func newContactServer() *Service {
	s := &Service{}
	return s
}
