package websocket

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/libp2p/go-libp2p-pubsub"
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
	ctx      context.Context
	apis     *http.ServeMux
	upGrader *websocket.Upgrader
	server   *http.Server

	userTable          *UserTable
	onlineSet          *OnlineMap
	msgFromClientQueue chan *pbs.WsMsg
	threads            map[string]*thread.Thread
	p2pOnOffWriter     *pubsub.Topic
	p2pIMWriter        *pubsub.Topic
	p2pUnreadQuery     *pubsub.Topic
	dataBase           *leveldb.DB
}
type ChatHandler func(http.ResponseWriter, *http.Request)

const (
	CPUserOnline       = "/user/online"
	WSThreadName       = "websocket thread"
	DispatchThreadName = "websocket message dispatcher thread"
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
	utils.LogInst().Info().Msg("websocket service instance init......")
	return ws
}

func (ws *Service) StartService(nodeID string, ctx context.Context) {
	ws.id = nodeID
	ws.ctx = ctx
	t := thread.NewThreadWithName(DispatchThreadName, func(stop chan struct{}) {
		utils.LogInst().Info().Msg("websocket client message dispatch thread start......")
		ws.wsCliMsgDispatch(stop)
		ws.ShutDown()
	})
	ws.threads[DispatchThreadName] = t
	t.Run()

	t = thread.NewThreadWithName(WSThreadName, func(_ chan struct{}) {
		utils.LogInst().Info().Msg("websocket service thread start......")
		err := ws.server.ListenAndServe()
		utils.LogInst().Err(err).Send()
		ws.ShutDown()
	})
	ws.threads[WSThreadName] = t
	t.Run()
}

func (ws *Service) ShutDown() {
	utils.LogInst().Warn().Msg("websocket service thread exit......")
	for _, t := range ws.threads {
		t.Stop()
	}
	if ws.threads == nil {
		return
	}
	ws.threads = nil
	_ = ws.dataBase.Close()
	_ = ws.server.Close()
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

	for {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("websocket dispatch thread exit")
			return

		case msg := <-ws.msgFromClientQueue:
			switch msg.Typ {
			case pbs.WsMsgType_ImmediateMsg:
				if err := ws.procIM(msg); err != nil {
					utils.LogInst().Warn().Msgf("dispatch ws message failed:%s", err)
				}
			case pbs.WsMsgType_PullUnread:

				if err := ws.p2pUnreadQuery.Publish(ws.ctx, msg.Data()); err != nil {
					utils.LogInst().Warn().Msgf("broadcast unread message request failed:%s", err)
					continue
				}

				if err := ws.findLocalUnread(msg); err != nil {
					utils.LogInst().Warn().Msgf("dispatch ws message failed:%s", err)
				}

			default:
				utils.LogInst().Warn().Msgf("unknown message type:%s from websocket client", msg.Typ)
			}
		}
	}
}
