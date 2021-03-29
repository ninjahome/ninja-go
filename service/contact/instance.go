package contact

import (
	"github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"net/http"
	"sync"
)

const (
	PathOperateContact = "/contact/operate"
	PathQueryContact   = "/contact/query"
	ServiceThreadName  = "contact http service thread"
)

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
	apis := http.NewServeMux()
	server := _srvConfig.newContactServer(apis)

	db, err := leveldb.OpenFile(_srvConfig.DataBaseDir, &opt.Options{
		Strict:      opt.DefaultStrict,
		Compression: opt.NoCompression,
		Filter:      filter.NewBloomFilter(10),
	})
	if err != nil {
		return nil
	}

	s := &Service{
		apis:     apis,
		server:   server,
		dataBase: db,
		threads:  make(map[string]*thread.Thread),
	}
	apis.HandleFunc(PathOperateContact, s.operateContact)
	apis.HandleFunc(PathQueryContact, s.queryContact)
	return s
}

type Service struct {
	id       string
	p2pChan  chan *contact.ContactMsg
	apis     *http.ServeMux
	server   *http.Server
	threads  map[string]*thread.Thread
	dataBase *leveldb.DB
}

func (s *Service) StartService(id string, queue chan *contact.ContactMsg) {
	s.id = id
	s.p2pChan = queue

	t := thread.NewThreadWithName(ServiceThreadName, func(_ chan struct{}) {
		err := s.server.ListenAndServe()
		utils.LogInst().Err(err).Send()
		s.ShutDown()
	})
	s.threads[ServiceThreadName] = t
	t.Run()
}

func (s *Service) ShutDown() {
	for _, t := range s.threads {
		t.Stop()
	}
	if s.threads == nil {
		return
	}

	s.threads = nil
	_ = s.dataBase.Close()
	_ = s.server.Close()
}
