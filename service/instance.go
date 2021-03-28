package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"net/http"
	"sync"
	"time"
)

type WebSocketService struct {
	id       string
	apis     *http.ServeMux
	upGrader *websocket.Upgrader
	server   *http.Server

	userTable           *UserTable
	onlineSet           *OnlineMap
	msgFromClientQueue  chan *pbs.WsMsg
	threads             map[string]*thread.Thread
	msgToOtherPeerQueue chan *pbs.WsMsg
	dataBase            *leveldb.DB
}
type ChatHandler func(http.ResponseWriter, *http.Request)

const (
	CPUserOnline       = "/user/online"
	WSThreadName       = "websocket thread"
	DispatchThreadName = "websocket message dispatcher thread"
)

var (
	_instance *WebSocketService
	once      sync.Once
)

func Inst() *WebSocketService {
	once.Do(func() {
		_instance = newWebSocket()
	})
	return _instance
}

func newWebSocket() *WebSocketService {

	if _srvConfig == nil {
		panic("init service config first")
	}

	apis := http.NewServeMux()
	server := _srvConfig.newWSServer(apis)
	db, err := leveldb.OpenFile(_srvConfig.DataBaseDir, &opt.Options{
		Strict:      opt.DefaultStrict,
		Compression: opt.NoCompression,
		Filter:      filter.NewBloomFilter(10),
	})
	if err != nil {
		return nil
	}
	ws := &WebSocketService{
		upGrader:           _srvConfig.newUpGrader(),
		apis:               apis,
		server:             server,
		userTable:          newUserTable(),
		onlineSet:          newOnlineSet(),
		msgFromClientQueue: make(chan *pbs.WsMsg, _srvConfig.WsMsgQueueSize),
		threads:            make(map[string]*thread.Thread),
		dataBase:           db,
	}

	ws.RegisterService(CPUserOnline, ws.online)
	return ws
}

func (ws *WebSocketService) RegisterService(path string, handler ChatHandler) {
	ws.apis.HandleFunc(path, handler)
}

func (ws *WebSocketService) online(w http.ResponseWriter, r *http.Request) {

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

	webSocket.SetReadLimit(int64(_srvConfig.WsBufferSize))
	webSocket.SetReadDeadline(time.Now().Add(_srvConfig.PongWait))
	webSocket.SetPongHandler(func(string) error { webSocket.SetReadDeadline(time.Now().Add(_srvConfig.PongWait)); return nil })

	if err := ws.newOnlineUser(webSocket); err != nil {
		utils.LogInst().Err(err).Send()
		return
	}
}

func (ws *WebSocketService) StartService(nodeID string, omq chan *pbs.WsMsg) {

	ws.msgToOtherPeerQueue = omq
	ws.id = nodeID
	t := thread.NewThreadWithName(DispatchThreadName, func(stop chan struct{}) {
		ws.msgDispatch(stop)
		ws.ShutDown()
	})
	ws.threads[DispatchThreadName] = t
	t.Run()

	t = thread.NewThreadWithName(WSThreadName, func(_ chan struct{}) {
		err := ws.server.ListenAndServe()
		utils.LogInst().Err(err).Send()
		ws.ShutDown()
	})
	ws.threads[WSThreadName] = t
	t.Run()
}

func (ws *WebSocketService) ShutDown() {
	for _, t := range ws.threads {
		t.Stop()
	}
	ws.threads = nil
	_ = ws.dataBase.Close()
	_ = ws.server.Close()
}

func (ws *WebSocketService) OnlineFromOtherPeer(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Online)
	if !ok {
		return fmt.Errorf("this is not a valid online p2p message")
	}

	if !body.Online.Verify(msg.Sig) {
		return fmt.Errorf("this is an attack")
	}
	ws.onlineSet.add(body.Online.UID)
	return nil
}

func (ws *WebSocketService) OfflineFromOtherPeer(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Online)
	if !ok {
		return fmt.Errorf("this is not a valid offline p2p message")
	}
	//TODO:: verify peer's authorization
	ws.onlineSet.del(body.Online.UID)
	ws.userTable.del(body.Online.UID)
	return nil
}

func (ws *WebSocketService) PeerImmediateCryptoMsg(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Message)
	if !ok {
		return fmt.Errorf("this is not a valid p2p crypto message")
	}

	u, ok := ws.userTable.get(body.Message.To)
	if !ok {
		return nil
	}
	utils.LogInst().Debug().Msgf("found to peer[%s] in my table", body.Message.To)
	return u.writeToCli(msg)
}
