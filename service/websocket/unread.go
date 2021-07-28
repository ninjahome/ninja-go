package websocket

import (
	"fmt"

	"github.com/ninjahome/ninja-go/node/worker"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/syndtr/goleveldb/leveldb/util"
	"google.golang.org/protobuf/proto"
)

func (ws *Service)_loadDbunreadIm(unread *pbs.WSPullUnread)  ([]*pbs.WsUnreadAckMsg, bool){
	buf := make([]*pbs.WsUnreadAckMsg, 0)
	var counter = 0
	sKey := IMDBKey(unread.Receiver, unread.FromUnixTime)
	eKey := IMDBEnd(unread.Receiver)
	iter := ws.dataBase.NewIterator(&util.Range{Start: sKey, Limit: eKey}, nil)
	hasMore := false
	for iter.Next() {
		v := iter.Value()
		msgV := &pbs.WSCryptoMsg{}
		if err := proto.Unmarshal(v, msgV); err != nil {
			utils.LogInst().Warn().Err(err).Send()
			continue
		}
		counter++

		m:=&pbs.WsUnreadAckMsg{
			CryptoMsg: &pbs.WsUnreadAckMsg_Payload{Payload: msgV},
		}

		buf = append(buf, m)
		k := iter.Key()
		_ = ws.dataBase.Delete(k, nil)
		if counter >= _wsConfig.MaxUnreadMsgNoPerQuery {
			hasMore = true
			break
		}
	}

	iter.Release()


	return buf,hasMore
}

func (ws *Service)_loadDbunreadGIm(unread *pbs.WSPullUnread, leftCnt int)  ([]*pbs.WsUnreadAckMsg, bool)  {

	buf := make([]*pbs.WsUnreadAckMsg, 0)
	counter := 0
	startKey:=StartKey(unread.Receiver)
	endKey:=EndKey(unread.Receiver)
	iter := ws.dataBase.NewIterator(&util.Range{Start: []byte(startKey), Limit: []byte(endKey)}, nil)
	hasMore := false
	for iter.Next() {
		groupkey := iter.Value()

		gmsg,err:=GetGroupMsg(ws.dataBase, groupkey)
		if err!=nil{
			continue
		}
		counter++

		m:=&pbs.WsUnreadAckMsg{
			CryptoMsg: &pbs.WsUnreadAckMsg_GPayload{GPayload: gmsg},
		}

		buf = append(buf, m)
		k := iter.Key()
		_ = ws.dataBase.Delete(k, nil)
		if counter >= leftCnt {
			hasMore = true
			break
		}

	}

	iter.Release()

	return buf,hasMore

}


func (ws *Service) loadDbUnread(unread *pbs.WSPullUnread) ([]*pbs.WsUnreadAckMsg, bool) {
	acks, hasMore:=ws._loadDbunreadIm(unread)
	if hasMore{
		return acks,hasMore
	}

	left:= _wsConfig.MaxUnreadMsgNoPerQuery - len(acks)
	if left < 0{
		left = 0
	}

	acks2,hashMoreG := ws._loadDbunreadGIm(unread,left)

	acks = append(acks,acks2...)

	return acks,hashMoreG

}

func (ws *Service) findLocalUnread(request *pbs.WsMsg) error {

	unBody, ok := request.Payload.(*pbs.WsMsg_Unread)
	if !ok {
		return fmt.Errorf("cast to unread message body failed")
	}
	unread := unBody.Unread
	user, ok := ws.userTable.get(unread.Receiver)
	if !ok {
		return nil
	}

LoadMore:
	unreadMsg, hasMore := ws.loadDbUnread(unread)
	if len(unreadMsg) == 0 {
		return nil
	}
	utils.LogInst().Info().
		Int("local unread", len(unreadMsg)).
		Str("receiver", unread.Receiver).
		Send()

	result := &pbs.WsMsg{
		Typ: pbs.WsMsgType_UnreadAck,
		Payload: &pbs.WsMsg_UnreadAck{UnreadAck: &pbs.WSUnreadAck{
			NodeID:   ws.id,
			Receiver: unread.Receiver,
			Payload:  unreadMsg,
		}},
	}
	if err := user.writeToCli(result); err != nil {
		return err
	}
	if hasMore {
		goto LoadMore
	}
	return nil
}

func (ws *Service) procUnreadMsgQueryFromP2pNetwork(msg *pbs.WsMsg) error {

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

	utils.LogInst().Debug().
		Int("Peer unread", len(unreadMsg)).
		Str("receiver", unread.Receiver).
		Send()

	if err := ws.unreadP2pQueryWorker.BroadCast(result.Data()); err != nil {
		return err
	}

	user, ok := ws.userTable.get(unread.Receiver)
	if ok {
		utils.LogInst().Info().Str("receiver", unread.Receiver).Msg("multi online client")
		if err := user.writeToCli(result); err != nil {
			return err
		}
	}

	if hasMore {
		goto LoadMore
	}
	return nil
}

func (ws *Service) unreadMsgResultFromP2pNetwork(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_UnreadAck)
	if !ok {
		return fmt.Errorf("cast to unread ack message body failed")
	}
	receiver := body.UnreadAck.Receiver
	user, ok := ws.userTable.get(receiver)
	if !ok {
		return nil
	}

	utils.LogInst().Info().
		Str("receiver", receiver).
		Int("peers unread ack", len(body.UnreadAck.Payload)).
		Send()

	return user.writeToCli(msg)
}

func (ws *Service) UnreadMsgFromP2pNetwork(w *worker.TopicWorker) {
	ws.unreadP2pQueryWorker = w
	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Str("Peer unread message", err.Error()).Send()
			return
		}

		if msg.ReceivedFrom.String() == ws.id {
			continue
		}

		p2pMsg := &pbs.WsMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Str("Peer Unread message", err.Error()).Send()
			continue
		}

		switch p2pMsg.Typ {
		case pbs.WsMsgType_PullUnread:
			if err := ws.procUnreadMsgQueryFromP2pNetwork(p2pMsg); err != nil {
				utils.LogInst().Warn().Str("local unread message read", err.Error()).Send()
				continue
			}
		case pbs.WsMsgType_UnreadAck:
			if err := ws.unreadMsgResultFromP2pNetwork(p2pMsg); err != nil {
				utils.LogInst().Warn().Str("send unread", err.Error()).Send()
				continue
			}
		default:
			utils.LogInst().Warn().Str("invalid unread message Type", p2pMsg.Typ.String()).Send()
		}
	}
}
