package service

import (
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/utils"
	"net/http"
	"sync"
	"time"
)

type WebSocketService struct {
	apis      *http.ServeMux
	upGrader  *websocket.Upgrader
	server    *http.Server
	userTable *UserTable
	onlineSet *OnlineMap
	msgQueue  chan *pbs.WSCryptoMsg
}
type ChatHandler func(http.ResponseWriter, *http.Request)

const (
	CPUserOnline = "/user/online"
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
		upGrader:  _srvConfig.newUpGrader(),
		apis:      apis,
		server:    server,
		userTable: newUserTable(),
		onlineSet: newOnlineSet(),
		msgQueue:  make(chan *pbs.WSCryptoMsg, 1024),
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
		if _, ok := err.(websocket.HandshakeError); !ok {
			utils.LogInst().Err(err).Send()
		}
		return
	}
	ws.newOnlineUser(webSocket)
}

func (ws *WebSocketService) StartService() error {
	return ws.server.ListenAndServe()
}
