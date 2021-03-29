package contact

import (
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
	"sync"
)

const (
	PathOperateContact = "/contact/operate"
	PathQueryContact   = "/contact/query"
	ServiceThreadName  = "contact http service thread"
	DBPatternHead      = "ContactMap_%s"
)

var (
	_instance      *Service
	once           sync.Once
	ErBodyCast     = fmt.Errorf("invalid add contact message body")
	ErVerifyFailed = fmt.Errorf("verify signature failed")
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
	id          string
	p2pChan     chan *pbs.ContactMsg
	apis        *http.ServeMux
	server      *http.Server
	threads     map[string]*thread.Thread
	dataBase    *leveldb.DB
	contactLock sync.RWMutex
}

func (s *Service) StartService(id string, queue chan *pbs.ContactMsg) {
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

func (s *Service) operateContact(w http.ResponseWriter, r *http.Request) {
	msg := &pbs.ContactMsg{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write(pbs.ErrAck(err.Error()))
		return
	}

	if err = proto.Unmarshal(data, msg); err != nil {
		w.Write(pbs.ErrAck(err.Error()))
		return
	}

	switch msg.Typ {
	case pbs.ContactMsgType_MTAddContact, pbs.ContactMsgType_MTUpdateContact:
		body, ok := msg.PayLoad.(*pbs.ContactMsg_AddOrUpdate)
		if !ok {
			w.Write(pbs.ErrAck(ErBodyCast.Error()))
			return
		}

		item := body.AddOrUpdate
		if !msg.Verify(item.Data()) {
			w.Write(pbs.ErrAck(ErVerifyFailed.Error()))
			return
		}

		s.p2pChan <- msg
		if err := s.save(msg.From, item); err != nil {
			w.Write(pbs.ErrAck(err.Error()))
			return
		}

		w.Write(pbs.OkAck())

	case pbs.ContactMsgType_MTDeleteContact:
		body, ok := msg.PayLoad.(*pbs.ContactMsg_DelC)
		if !ok {
			w.Write(pbs.ErrAck(ErBodyCast.Error()))
			return
		}

		cid := body.DelC
		if !msg.Verify([]byte(cid)) {
			w.Write(pbs.ErrAck(ErVerifyFailed.Error()))
			return
		}
		s.p2pChan <- msg
		if err := s.del(msg.From, cid); err != nil {
			w.Write(pbs.ErrAck(err.Error()))
			return
		}

		w.Write(pbs.OkAck())
	}
}
