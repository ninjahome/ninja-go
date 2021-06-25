package websocket

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs2 "github.com/ninjahome/ninja-go/pbs/stream"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

type wsUser struct {
	lock           sync.RWMutex
	UID            string
	OnLineTime     time.Time
	cliWsConn      *websocket.Conn
	msgFromCliChan chan *pbs.WsMsg
	msgToCliChan   chan *pbs.WsMsg
	devToken	   string
	devTyp         int
	kaTimer        *time.Ticker
}

func (u *wsUser) offLine() {
	u.lock.Lock()
	defer u.lock.Unlock()

	if u.msgToCliChan == nil {
		return
	}

	u.cliWsConn.Close()
	close(u.msgToCliChan)
	u.msgToCliChan = nil
	u.kaTimer.Stop()
	utils.LogInst().Debug().Str("WS user offline", u.UID).Send()
}

func (u *wsUser) reading(_ chan struct{}) {
	utils.LogInst().Debug().Str("WS reading thread start", u.UID).Send()
	defer utils.LogInst().Debug().Str("reading thread exit!", u.UID).Send()
	defer u.offLine()
	for {
		_, message, err := u.cliWsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure) {
				utils.LogInst().Err(err).Send()
			}
			utils.LogInst().Info().Str("WS read client", err.Error()).Send()
			return
		}

		msg := &pbs.WsMsg{}
		if err := proto.Unmarshal(message, msg); err != nil {
			utils.LogInst().Warn().Str("WS invalid client message", err.Error()).Send()
			continue
		}
		u.msgFromCliChan <- msg
	}
}

func (u *wsUser) writing(stop chan struct{}) {
	utils.LogInst().Debug().Str("WS writing thread start", u.UID).Send()
	defer utils.LogInst().Debug().Str("WS writer thread exit", u.UID).Send()

	defer u.offLine()
	for {
		select {
		case <-stop:
			return
		case message, ok := <-u.msgToCliChan:
			if !ok {
				utils.LogInst().Info().Str("WS client message chan", " closed").Send()
				u.cliWsConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Warn().Str("WS set write timeout ", err.Error()).Send()
				return
			}

			w, err := u.cliWsConn.NextWriter(websocket.TextMessage)
			if err != nil {
				utils.LogInst().Warn().Str("WS get next writer ", err.Error()).Send()
				return
			}

			_, err = w.Write(message.Data())
			if err := w.Close(); err != nil {
				utils.LogInst().Warn().Str("WS write ", err.Error()).Send()
				return
			}

		case <-u.kaTimer.C:
			utils.LogInst().Debug().Str("WS ping pong", "sent").Send()
			if err := u.cliWsConn.SetWriteDeadline(time.Now().Add(_wsConfig.WriteWait)); err != nil {
				utils.LogInst().Warn().Str("WS write deadline", err.Error()).Send()
				return
			}
			if err := u.cliWsConn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				utils.LogInst().Warn().Str("WS write ping", err.Error()).Send()
				return
			}
		}
	}
}

func (u *wsUser) writeToCli(msg *pbs.WsMsg) error {
	u.msgToCliChan <- msg
	return nil
}

func (u *wsUser) String() string {
	return fmt.Sprintf("uid:%s, online:%s, from:%s", u.UID, u.OnLineTime, u.cliWsConn.RemoteAddr())
}

type DevInfo struct {
	DevToken string   `json:"t"`
	DevTyp   int      `json:"typ"`
}

const(
	DevInfoDBKeyHead = "Deviceinfodbkey_0"
	DevInfoDBKeyEnd = "Deviceinfodbkey_1"
	SyncBuflength = 2048
	SyncDevInfoCount = 20
	DevTypeIOS = 1
	DevTypeAndroid = 2

)


func (ws *Service)SaveToken(uid,devToken string, devTyp int) error  {
	o:=&opt.WriteOptions{
		Sync: true,
	}

	di:=&DevInfo{
		DevTyp: devTyp,
		DevToken: devToken,
	}

	j,_:=json.Marshal(*di)

	return ws.dataBase.Put([]byte(DevInfoDBKeyHead+uid),j,o)
}

func (ws *Service)GetToken(uid string) (string, int, error)  {
	dibytes,err:=ws.dataBase.Get([]byte(DevInfoDBKeyHead+uid),nil)
	if err!=nil{
		return "",0,err
	}
	di:=&DevInfo{}

	err=json.Unmarshal(dibytes,di)
	if err!=nil{
		return "",0,err
	}

	return di.DevToken,di.DevTyp,nil
}


func (ws *Service) newOnlineUser(conn *websocket.Conn) error {

	msg := &pbs.WsMsg{}
	online, rawData, err := msg.ReadOnlineFromCli(conn)
	if err != nil {
		conn.Close()
		return err
	}

	wu := &wsUser{
		devToken: online.DevToken,
		devTyp: int(online.DevTyp),
		cliWsConn:      conn,
		UID:            online.UID,
		OnLineTime:     time.Now(),
		msgFromCliChan: ws.msgFromClientQueue,
		kaTimer:        time.NewTicker(_wsConfig.PingPeriod),
		msgToCliChan:   make(chan *pbs.WsMsg, _wsConfig.MaxUnreadMsgNoPerQuery),
	}
	ws.onlineSet.add(wu.UID)
	ws.SaveToken(wu.UID,wu.devToken,wu.devTyp)
	ws.userTable.add(wu)

	tid := fmt.Sprintf("chat read:%s", wu.UID)
	readTh := thread.NewThreadWithName(tid, wu.reading)
	ws.threads[tid] = readTh
	readTh.Run()

	tid = fmt.Sprintf("chat writer:%s", wu.UID)
	writeTh := thread.NewThreadWithName(tid, wu.writing)
	writeTh.DidExit(func() {
		ws.offlineUser(tid, wu.UID)
	})
	ws.threads[tid] = writeTh
	writeTh.Run()

	if err := ws.onOffLineP2pWorker.BroadCast(rawData); err != nil {
		return err
	}

	utils.LogInst().Debug().Str("WS New User", wu.UID).Send()

	return nil
}

func (ws *Service) offlineUser(threadId string, uid string) {
	delete(ws.threads, threadId)
	ws.onlineSet.del(uid)
	ws.userTable.del(uid)

	//TODO:: add signature for offline message
	msg := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_Offline,
		Payload: &pbs.WsMsg_Online{Online: &pbs.WSOnline{UID: uid}},
	}

	if err := ws.onOffLineP2pWorker.BroadCast(msg.Data()); err != nil {
		utils.LogInst().Warn().Str("offline broadcast", err.Error()).Send()
	}
	utils.LogInst().Info().Str("WS user offline", uid).Send()
}

func (ws *Service) OnOffLineForP2pNetwork(w *worker.TopicWorker) {
	ws.onOffLineP2pWorker = w

	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Str("on-off line ", err.Error()).Send()
			return
		}
		if msg.ReceivedFrom.String() == ws.id {
			continue
		}

		p2pMsg := &pbs.WsMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Str("unmarshal", err.Error()).Send()
			continue
		}

		switch p2pMsg.Typ {
		case pbs.WsMsgType_Online:
			err = ws.onlineFromOtherPeer(p2pMsg)
		case pbs.WsMsgType_Offline:
			err = ws.offlineFromOtherPeer(p2pMsg)
		default:
			err = fmt.Errorf("unknown msg typ in p2p on-off line channel")
		}

		if err != nil {
			utils.LogInst().Err(err).Send()
		}
	}
}

func (ws *Service) onlineFromOtherPeer(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Online)
	if !ok {
		return fmt.Errorf("this is not a valid online p2p message")
	}

	if !body.Online.Verify(msg.Sig) {
		return fmt.Errorf("this is an attack")
	}
	ws.onlineSet.add(body.Online.UID)
	ws.SaveToken(body.Online.UID,body.Online.DevToken,int(body.Online.DevTyp))
	utils.LogInst().Debug().Str("online", body.Online.UID).Send()
	return nil
}

func (ws *Service) offlineFromOtherPeer(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Online)
	if !ok {
		return fmt.Errorf("this is not a valid offline p2p message")
	}
	//TODO:: verify peer's authorization
	ws.onlineSet.del(body.Online.UID)
	ws.userTable.del(body.Online.UID)
	utils.LogInst().Debug().Str("offline", body.Online.UID).Send()
	return nil
}

func (ws *Service) SyncOnlineSetFromPeerNodes(stream network.Stream) error {
	defer stream.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))

	streamMsg := &pbs2.StreamMsg{}
	data := streamMsg.SyncOnline("TODO::wallet key and sig") //TODO::
	data = append(data, OnlineStreamDelim)

	_, err := rw.Write(data)
	if err != nil {
		utils.LogInst().Warn().Str("stream: write online", err.Error()).Send()
		return err
	}
	if err := rw.Flush(); err != nil {
		utils.LogInst().Warn().Str("stream: flush online", err.Error()).Send()
	}

	bts, err := rw.ReadBytes(OnlineStreamDelim)
	if err != nil {
		utils.LogInst().Warn().Str("stream: read online", err.Error()).Send()
		return err
	}

	resp := &pbs2.StreamMsg{}
	bts = bts[:len(bts)-1]
	if err := proto.Unmarshal(bts, resp); err != nil {
		utils.LogInst().Warn().Str("stream: parse data online", err.Error()).Send()
		return err
	}

	body, ok := resp.Payload.(*pbs2.StreamMsg_OnlineAck)
	if !ok {
		utils.LogInst().Warn().Str("stream: cast data online", "failed").Send()
		return fmt.Errorf("invalid onlime map data")
	}

	uidBatch := body.OnlineAck.UID
	utils.LogInst().Info().Int("synced online users", len(uidBatch)).Send()
	if len(uidBatch) == 0 {
		return nil
	}
	ws.onlineSet.addBatch(uidBatch)

	return nil
}

func (ws *Service) OnlineMapQuery(stream network.Stream) {
	defer stream.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	bts, err := rw.ReadBytes(OnlineStreamDelim)
	if err != nil {
		utils.LogInst().Warn().Str("read online", err.Error()).Send()
		return
	}

	bts = bts[:len(bts)-1]
	streamMsg := &pbs2.StreamMsg{}
	if err := proto.Unmarshal(bts, streamMsg); err != nil {
		utils.LogInst().Warn().Str("parse stream", err.Error()).Send()
		return
	}

	resp := &pbs2.StreamMsg{}
	data := resp.SyncOnlineAck(ws.onlineSet.AllUid())
	data = append(data, OnlineStreamDelim)
	if _, err := rw.Write(data); err != nil {
		utils.LogInst().Warn().Str("stream:  write online response", err.Error()).Send()
		return
	}
	if err := rw.Flush(); err != nil {
		utils.LogInst().Warn().Str("stream:  flush online response", err.Error()).Send()
	}
}

func GetUUIDFromDBDevInfoKey(key []byte) (string,error)  {
	if len(key) <= len(DevInfoDBKeyHead){
		return "",errors.New("not a  device key")
	}

	uk:=key[len(DevInfoDBKeyHead):]

	return string(uk),nil

}




func (ws *Service) SyncDevInfoFromPeerNodes(stream network.Stream) error {
	defer stream.Close()

	rw:=NewLVRW(stream,SyncBuflength)

	streamMsg:=&pbs2.StreamMsg{}
	data:=streamMsg.SyncDevInfo("TODO::wallet key and sig")

	fmt.Println("send sync dev:",hex.EncodeToString(data))

	if _,err:=rw.Write(data);err!=nil{
		return err
	}

	if _,err:=rw.Commit();err!=nil{
		return err
	}

	for{
		buf:=make([]byte,SyncBuflength)
		if n,err:=rw.Read(buf);err!=nil{
			return err
		}else{
			if IsReadEnd(buf[:n]){
				utils.LogInst().Info().Str("sync dev info", "success").Send()
				return nil
			}

			resp:=&pbs2.StreamMsg{}
			if err:=proto.Unmarshal(buf[:n],resp);err!=nil{
				utils.LogInst().Warn().Str("sync dev info", err.Error()).Send()
				return fmt.Errorf("sync dev info, unmarshal stream msg failed")
			}
			diack,ok:=resp.Payload.(*pbs2.StreamMsg_DiAck)
			if !ok{
				utils.LogInst().Warn().Str("sync dev info: cast data", "failed").Send()
				return fmt.Errorf("sync dev info, cast data failed")
			}
			if diack!= nil && diack.DiAck != nil{
				for i:=0;i<len(diack.DiAck.Dis);i++{
					di:=diack.DiAck.Dis[i]

					if err:=ws.SaveToken(di.Uid,di.DevToken,int(di.DevTyp));err!=nil{
						utils.LogInst().Warn().Str("sync dev info: save to db error", di.Uid).Send()
					}

				}
			}

		}

	}

	return nil

}


func (ws *Service)DevtokensQuery(stream network.Stream)  {
	defer stream.Close()

	rw:=NewLVRW(stream,SyncBuflength)

	buf:=make([]byte,SyncBuflength)

	n,err:=rw.ReadFull(buf)
	if err!=nil{
		utils.LogInst().Warn().Str("read dev info query", err.Error()).Send()
		return
	}
	buf = buf[:n]
	fmt.Println("rcv sync dev info:",hex.EncodeToString(buf))
	streamMsg := &pbs2.StreamMsg{}
	if err := proto.Unmarshal(buf, streamMsg); err != nil {
		utils.LogInst().Warn().Str("devinfo parse stream", err.Error()).Send()
		return
	}

	if streamMsg.MTyp != pbs2.StreamMType_MTDevInfoSync{
		utils.LogInst().Warn().Str("devinfo parse stream", "not a sync dev info msg").Send()
		return
	}


	iter:=ws.dataBase.NewIterator(&util.Range{Start: []byte(DevInfoDBKeyHead), Limit: []byte(DevInfoDBKeyEnd)},nil)
	defer iter.Release()

	var uuid string

	var resp *pbs2.StreamMsg
	var diack *pbs2.DevInfoAck

	for iter.Next(){
		uuid,err = GetUUIDFromDBDevInfoKey(iter.Key())
		if err!=nil{
			continue
		}

		di:=&DevInfo{}

		err=json.Unmarshal(iter.Value(),di)
		if err!=nil{
			continue
		}

		if resp == nil{
			resp = &pbs2.StreamMsg{}
			resp.MTyp = pbs2.StreamMType_MTDevInfoAck

			diack=&pbs2.DevInfoAck{

			}

			resp.Payload = &pbs2.StreamMsg_DiAck{DiAck: diack}
		}

		pbdi:=&pbs2.DevInfo{
			Uid: uuid,
			DevTyp: int32(di.DevTyp),
			DevToken: di.DevToken,
		}

		diack.Dis = append(diack.Dis,pbdi)

		if len(diack.Dis) >= SyncDevInfoCount{
			data, _ := proto.Marshal(resp)
			_,err:=rw.Write(data)
			if err!=nil{
				utils.LogInst().Warn().Str("devinfo ack error", err.Error()).Send()
				return
			}

			resp = nil
			diack = nil
		}
	}

	if resp != nil{
		data, _ := proto.Marshal(resp)
		_,err:=rw.Write(data)
		if err!=nil{
			utils.LogInst().Warn().Str("devinfo ack error", err.Error()).Send()
			return
		}

		resp = nil
		diack = nil
	}

	if _,err:=rw.Commit();err!=nil{
		utils.LogInst().Warn().Str("devinfo ack error", err.Error()).Send()
		return
	}

	utils.LogInst().Info().Str("devinfo ack", "success").Send()
}