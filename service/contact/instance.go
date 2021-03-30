package contact

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p-pubsub"
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
	utils.LogInst().Info().Msg("contact service instance init......")
	return s
}

type Service struct {
	id              string
	ctx             context.Context
	contactOpWriter *pubsub.Topic
	contactQuery    *pubsub.Topic
	apis            *http.ServeMux
	server          *http.Server
	threads         map[string]*thread.Thread
	dataBase        *leveldb.DB
	contactLock     sync.RWMutex
}

func (s *Service) StartService(id string, ctx context.Context) {
	s.id = id
	s.ctx = ctx

	t := thread.NewThreadWithName(ServiceThreadName, func(_ chan struct{}) {
		utils.LogInst().Info().Msg("contact service thread start......")
		err := s.server.ListenAndServe()
		utils.LogInst().Err(err).Send()
		s.ShutDown()
	})
	s.threads[ServiceThreadName] = t
	t.Run()
}

func (s *Service) ShutDown() {
	utils.LogInst().Warn().Msg("contact service thread exit......")
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

	if err = s.contactOpWriter.Publish(s.ctx, msg.Data()); err != nil {
		return
	}

	if err := s.procContactOperation(msg); err != nil {
		w.Write(pbs.ErrAck(err.Error()))
		return
	}

	w.Write(pbs.OkAck())
}

func (s *Service) procContactOperation(msg *pbs.ContactMsg) error {
	switch msg.Typ {
	case pbs.ContactMsgType_MTAddContact, pbs.ContactMsgType_MTUpdateContact:
		return s.saveContact(msg)
	case pbs.ContactMsgType_MTDeleteContact:
		return s.delContact(msg)
	default:
		return fmt.Errorf("unknown contact operation")
	}
}
