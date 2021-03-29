package websocket

import (
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/syndtr/goleveldb/leveldb/util"
	"google.golang.org/protobuf/proto"
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

func (ws *Service) procIM(msg *pbs.WsMsg) error {
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

func (ws *Service) loadDbUnread(unread *pbs.WSPullUnread) ([]*pbs.WSCryptoMsg, bool) {

	buf := make([]*pbs.WSCryptoMsg, 0)
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
		buf = append(buf, msgV)
		k := iter.Key()
		_ = ws.dataBase.Delete(k, nil)
		if counter > _wsConfig.WsMsgSizePerUser {
			hasMore = true
			break
		}
	}

	iter.Release()
	return buf, hasMore
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
				ws.msgToOtherPeerQueue <- msg

				if err := ws.findLocalUnread(msg); err != nil {
					utils.LogInst().Warn().Msgf("dispatch ws message failed:%s", err)
				}
			default:
				utils.LogInst().Warn().Msgf("unknown message type:%s", msg.Typ)
			}
		}
	}
}
