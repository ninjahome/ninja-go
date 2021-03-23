package service

import (
	"fmt"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"net/http"
	"sync"
	"time"
)

type WebSocketService struct {
	apis     *http.ServeMux
	upGrader *websocket.Upgrader
	server   *http.Server

	userTable           *UserTable
	onlineSet           *OnlineMap
	msgFromClientQueue  chan *pbs.WSCryptoMsg
	threads             map[string]*thread.Thread
	msgToOtherPeerQueue chan *pbs.P2PMsg
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

	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
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

	ws := &WebSocketService{
		upGrader:           _srvConfig.newUpGrader(),
		apis:               apis,
		server:             server,
		userTable:          newUserTable(),
		onlineSet:          newOnlineSet(),
		msgFromClientQueue: make(chan *pbs.WSCryptoMsg, 1024),
		threads:            make(map[string]*thread.Thread),
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
			utils.LogInst().Warn().Msgf("websocket service panic by one server =>", r)
		}
	}()

	webSocket, err := ws.upGrader.Upgrade(w, r, nil)
	if err != nil {
		utils.LogInst().Err(err).Send()
		return
	}

	webSocket.SetReadLimit(512) //TODO::config
	webSocket.SetReadDeadline(time.Now().Add(pongWait))
	webSocket.SetPongHandler(func(string) error { webSocket.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	if err := ws.newOnlineUser(webSocket); err != nil {
		utils.LogInst().Err(err).Send()
		return
	}
}

func (ws *WebSocketService) StartService(omq chan *pbs.P2PMsg) {

	ws.msgToOtherPeerQueue = omq

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
	_ = ws.server.Close()
}

func (ws *WebSocketService) OnlineFromOtherPeer(online *pbs.WSOnline) error {
	ws.onlineSet.add(online.UID)
	return nil
}

func (ws *WebSocketService) PeerImmediateCryptoMsg(msg *pbs.WSCryptoMsg) error {
	u, ok := ws.userTable.get(msg.To)
	if !ok {
		return fmt.Errorf("there is no such user in my table")
	}
	return u.write(msg)
}
