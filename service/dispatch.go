package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	MSGPatternHead = "TempCachedMsg_%s_%d"
	MSGPatternEnd  = "TempCachedMsg_%s_ffffffffffffffff"
)

func IMDBKey(receiver string, seq int64) []byte {
	key := fmt.Sprintf(MSGPatternHead, receiver, seq)
	return []byte(key)
}
func IMDBEnd(receiver string) []byte {
	key := fmt.Sprintf(MSGPatternEnd, receiver)
	return []byte(key)
}

func (ws *WebSocketService) procIM(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Message)
	if !ok {
		return fmt.Errorf("cast immediate message failed")
	}
	im := body.Message
	utils.LogInst().Debug().Msgf("message [%s->%s]_%x", im.From, im.To, im.UnixTime)
	if !ws.onlineSet.contains(im.To) {
		key := IMDBKey(im.To, im.UnixTime)
		return ws.dataBase.Put(key, im.MustData(), nil)
	}

	if user, ok := ws.userTable.get(im.To); ok {
		return user.writeToCli(msg)
	}

	ws.msgToOtherPeerQueue <- msg
	return nil
}
func (ws *WebSocketService) findUnread(msg *pbs.WsMsg) error {
	unBody, ok := msg.Payload.(*pbs.WsMsg_Unread)
	if !ok {
		return fmt.Errorf("cast to unread message body failed")
	}
	unread := unBody.Unread
	//TODO::need to have a deep think, from local ,from peer , from network
	if !ws.onlineSet.contains(unread.Receiver) {
	}
	sKey := IMDBKey(unread.Receiver, unread.FromUnixTime)
	eKey := IMDBEnd(unread.Receiver)
	iter := ws.dataBase.NewIterator(&util.Range{Start: sKey, Limit: eKey}, nil)
	unreadAck := &pbs.WSUnreadAck{
		NodeID:  ws.id,
		Payload: make([]*pbs.WSCryptoMsg, 0),
	}
	var counter = 0
	for iter.Next() {
		v := iter.Value()
		msgV := &pbs.WSCryptoMsg{}
		if err := proto.Unmarshal(v, msgV); err != nil {
			utils.LogInst().Warn().Err(err).Send()
			continue
		}
		counter++
		unreadAck.Payload = append(unreadAck.Payload, msgV)
		k := iter.Key()
		_ = ws.dataBase.Delete(k, nil)
		if counter > _srvConfig.WsMsgSizePerUser {
			break
		}
	}
	iter.Release()
	result := &pbs.WsMsg{
		Typ:     pbs.WsMsgType_UnreadAck,
		Payload: &pbs.WsMsg_UnreadAck{UnreadAck: unreadAck},
	}

	ws.msgToOtherPeerQueue <- result
	return iter.Error()
}

func (ws *WebSocketService) sendToPeer(msg *pbs.WsMsg) error {
	switch msg.Typ {
	case pbs.WsMsgType_ImmediateMsg:
		return ws.procIM(msg)
	case pbs.WsMsgType_PullUnread:
		return ws.findUnread(msg)
	}

	return fmt.Errorf("unknown message type:%s", msg.Typ)
}

func (ws *WebSocketService) msgDispatch(stop chan struct{}) {

	for {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("websocket dispatch thread exit")
			return

		case msg := <-ws.msgFromClientQueue:
			if err := ws.sendToPeer(msg); err != nil {
				utils.LogInst().Warn().Msgf("send ws message failed:%s", err)
			}
		}
	}
}
