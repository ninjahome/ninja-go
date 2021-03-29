package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
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
		msgFromClientQueue: make(chan *pbs.WsMsg, _wsConfig.WsMsgQueueSize),
		threads:            make(map[string]*thread.Thread),
		dataBase:           db,
	}

	ws.RegisterService(CPUserOnline, ws.online)
	return ws
}

func (ws *Service) RegisterService(path string, handler ChatHandler) {
	ws.apis.HandleFunc(path, handler)
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

	webSocket.SetReadLimit(int64(_wsConfig.WsBufferSize))
	webSocket.SetReadDeadline(time.Now().Add(_wsConfig.PongWait))
	webSocket.SetPongHandler(func(string) error { webSocket.SetReadDeadline(time.Now().Add(_wsConfig.PongWait)); return nil })

	if err := ws.newOnlineUser(webSocket); err != nil {
		utils.LogInst().Err(err).Send()
		return
	}
}

func (ws *Service) StartService(nodeID string, omq chan *pbs.WsMsg) {

	ws.msgToOtherPeerQueue = omq
	ws.id = nodeID
	t := thread.NewThreadWithName(DispatchThreadName, func(stop chan struct{}) {
		ws.wsCliMsgDispatch(stop)
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

func (ws *Service) ShutDown() {
	for _, t := range ws.threads {
		t.Stop()
	}
	ws.threads = nil
	_ = ws.dataBase.Close()
	_ = ws.server.Close()
}

func (ws *Service) OnlineFromOtherPeer(msg *pbs.WsMsg) error {
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

func (ws *Service) OfflineFromOtherPeer(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Online)
	if !ok {
		return fmt.Errorf("this is not a valid offline p2p message")
	}
	//TODO:: verify peer's authorization
	ws.onlineSet.del(body.Online.UID)
	ws.userTable.del(body.Online.UID)
	return nil
}

func (ws *Service) PeerImmediateCryptoMsg(msg *pbs.WsMsg) error {
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

func (ws *Service) PeerUnreadMsg(msg *pbs.WsMsg) error {

	unBody, ok := msg.Payload.(*pbs.WsMsg_Unread)
	if !ok {
		return fmt.Errorf("cast to unread message body failed")
	}
	unread := unBody.Unread

LoadMore:
	unreadMsg, hasMore := ws.loadDbUnread(unread)
	if len(unreadMsg) == 0 {
		return nil
	}

	result := &pbs.WsMsg{
		Typ: pbs.WsMsgType_UnreadAck,
		Payload: &pbs.WsMsg_UnreadAck{UnreadAck: &pbs.WSUnreadAck{
			NodeID:   ws.id,
			Receiver: unread.Receiver,
			Payload:  unreadMsg,
		}},
	}

	ws.msgToOtherPeerQueue <- result

	user, ok := ws.userTable.get(unread.Receiver)
	if ok {
		if err := user.writeToCli(result); err != nil {
			return err
		}
	}

	if hasMore {
		goto LoadMore
	}
	return nil
}

func (ws *Service) PeerUnreadAckMsg(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_UnreadAck)
	if !ok {
		return fmt.Errorf("cast to unread ack message body failed")
	}
	receiver := body.UnreadAck.Receiver
	user, ok := ws.userTable.get(receiver)
	if !ok {
		return nil
	}
	return user.writeToCli(msg)
}
