package contact

import (
	"fmt"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"net/http"
	"sync"
)

const (
	PathOperateContact = "/contact/operate"
	PathQueryContact   = "/contact/query"
	PathServicesPost   = "/service/post"
	ServiceThreadName  = "contact http service thread"
	DBPatternHead      = "ContactMap_%s"

	SyncStreamDelim byte = '@'
)

var FGetSrvPort func() bool



var (
	_instance      *Service
	once           sync.Once
	ErBodyCast     = fmt.Errorf("invalid add contact message body")
	ErVerifyFailed = fmt.Errorf("verify signature failed")
)

type Service struct {
	id                   string
	contactOperateWorker *worker.TopicWorker
	contactQueryWorker   *worker.TopicWorker
	apis                 *http.ServeMux
	server               *http.Server
	threads              map[string]*thread.Thread
	dataBase             *leveldb.DB
	contactLock          sync.RWMutex
	contactPeerWorker    *worker.StreamWorker //TODO::Make contact be a small block chain
}

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
	//if FGetSrvPort(){
		apis.HandleFunc(PathServicesPost, s.servicePost)
	//}

	return s
}
func (s *Service) StartService(id string, cpw *worker.StreamWorker) {
	s.id = id
	s.contactPeerWorker = cpw
	t := thread.NewThreadWithName(ServiceThreadName, func(_ chan struct{}) {
		utils.LogInst().Info().Str("Contact Serve", "Start up").Send()
		endPoint := fmt.Sprintf("%s:%d", _srvConfig.SrvIP, _srvConfig.SrvPort)
		ln, err := net.Listen("tcp4", endPoint)
		if err != nil {
			utils.LogInst().Error().
				Str("end", endPoint).
				Str("Contact Serve", err.Error()).
				Send()
			return
		}
		err = s.server.Serve(ln)
		utils.LogInst().Err(err).Str("Contact Serve", "Exit").Send()
		s.ShutDown()
	})
	s.threads[ServiceThreadName] = t
	t.Run()
}

func (s *Service) ShutDown() {
	if s.threads == nil {
		return
	}
	utils.LogInst().Warn().Str("Contact Serve", "Shutting Down").Send()
	for _, t := range s.threads {
		t.Stop()
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
	utils.LogInst().Debug().Str("From", msg.From).Msg("contact operation")
	//if err = s.contactOperateWorker.BroadCast(msg.Data()); err != nil {
	//	w.Write(pbs.ErrAck(err.Error()))
	//	return
	//}

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
