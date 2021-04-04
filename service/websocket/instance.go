package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"net/http"
	"sync"
	"time"
)

type Service struct {
	id       string
	apis     *http.ServeMux
	upGrader *websocket.Upgrader
	server   *http.Server
	dataBase *leveldb.DB

	userTable            *UserTable
	onlineSet            *OnlineMap
	msgFromClientQueue   chan *pbs.WsMsg
	threads              map[string]*thread.Thread
	onOffLineP2pWorker   *worker.TopicWorker
	IMP2pWorker          *worker.TopicWorker
	unreadP2pQueryWorker *worker.TopicWorker
	peerStreamWorker     *worker.StreamWorker
}
type ChatHandler func(http.ResponseWriter, *http.Request)

const (
	CPUserOnline            = "/user/online"
	WSThreadName            = "websocket main service thread"
	DispatchThreadName      = "websocket message dispatcher thread"
	OnlineStreamDelim  byte = '@'
)

var (
	_instance *Service
	once      sync.Once
)

func Inst() *Service {
	once.Do(func() {
		_instance = newWebSocket()
	})
	return _instance
}

func newWebSocket() *Service {
	if _wsConfig == nil {
		panic("init service config first")
	}

	apis := http.NewServeMux()
	server := _wsConfig.newWSServer(apis)
	db, err := leveldb.OpenFile(_wsConfig.DataBaseDir, &opt.Options{
		Strict:      opt.DefaultStrict,
		Compression: opt.NoCompression,
		Filter:      filter.NewBloomFilter(10),
	})
	if err != nil {
		return nil
	}
	ws := &Service{
		upGrader:           _wsConfig.newUpGrader(),
		apis:               apis,
		server:             server,
		userTable:          newUserTable(),
		onlineSet:          newOnlineSet(),
		msgFromClientQueue: make(chan *pbs.WsMsg, _wsConfig.WsMsgNoFromCli),
		threads:            make(map[string]*thread.Thread),
		dataBase:           db,
	}
	ws.apis.HandleFunc(CPUserOnline, ws.online)
	return ws
}

func (ws *Service) StartService(nodeID string) {
	ws.id = nodeID
	dspThread := thread.NewThreadWithName(DispatchThreadName, ws.wsCliMsgDispatch)
	ws.threads[DispatchThreadName] = dspThread

	srvThread := thread.NewThreadWithName(WSThreadName, func(_ chan struct{}) {
		utils.LogInst().Info().Str("Websocket Serve", "Start up").Send()
		err := ws.server.ListenAndServe()
		utils.LogInst().Err(err).Str("Websocket Serve", "Exit").Send()
		ws.ShutDown()
	})
	ws.threads[WSThreadName] = srvThread

	dspThread.Run()
	srvThread.Run()

	return
}

func (ws *Service) ShutDown() {
	if ws.threads == nil {
		return
	}
	utils.LogInst().Warn().Msg("websocket service thread shutting down......")
	for _, t := range ws.threads {
		t.Stop()
	}
	ws.threads = nil

	_ = ws.dataBase.Close()
	_ = ws.server.Close()
	close(ws.msgFromClientQueue)
}

func (ws *Service) online(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			utils.LogInst().Warn().Msgf("websocket service panic by one server :=>%s", r)
		}
	}()

	webSocket, err := ws.upGrader.Upgrade(w, r, nil)
	if err != nil {
		utils.LogInst().Err(err).Send()
		return
	}

	webSocket.SetReadLimit(int64(_wsConfig.WsIOBufferSize))
	webSocket.SetReadDeadline(time.Now().Add(_wsConfig.PongWait))
	webSocket.SetPongHandler(func(string) error { webSocket.SetReadDeadline(time.Now().Add(_wsConfig.PongWait)); return nil })

	if err := ws.newOnlineUser(webSocket); err != nil {
		utils.LogInst().Err(err).Send()
		return
	}
}

func (ws *Service) wsCliMsgDispatch(stop chan struct{}) {
	utils.LogInst().Info().Str("Dispatch thread", "Start up").Send()
	defer utils.LogInst().Info().Str("Dispatch thread", "Exit").Send()
	for {
		select {
		case <-stop:
			return

		case msg := <-ws.msgFromClientQueue:
			switch msg.Typ {
			case pbs.WsMsgType_ImmediateMsg:
				if err := ws.procIM(msg); err != nil {
					utils.LogInst().Warn().Msgf("dispatch ws message failed:%s", err)
				}
			case pbs.WsMsgType_PullUnread:
				if err := ws.findLocalUnread(msg); err != nil {
					utils.LogInst().Warn().Msgf("dispatch ws message failed:%s", err)
				}

				if err := ws.unreadP2pQueryWorker.BroadCast(msg.Data()); err != nil {
					utils.LogInst().Warn().Msgf("broadcast unread message request failed:%s", err)
				}

			default:
				utils.LogInst().Warn().Msgf("unknown message type:%s from websocket client", msg.Typ)
			}
		}
	}
}

func (ws *Service) DebugInfo(online, local bool, usr string) string {
	s := "\n-------------------websocket debug info---------------------\n"
	if online {
		s += ws.onlineSet.DumpContent() + "\n"
	}
	if local {
		s += ws.userTable.DumpContent() + "\n"
	}
	if usr != "" {
		u, ok := ws.userTable.get(usr)
		if !ok {
			s += fmt.Sprintf("no such user:[%s] in local ", usr)
			if ws.onlineSet.contains(usr) {
				s += "but is online\n"
			}
		} else {
			s += u.String() + "\n"
		}
	}
	s += "\n----------------------------------------------------------\n"
	return s
}
