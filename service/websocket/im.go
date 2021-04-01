package websocket

import (
	"fmt"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/utils"
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

	return ws.IMP2pWorker.BroadCast(msg.Data())
}

func (ws *Service) ImmediateMsgForP2pNetwork(w *worker.TopicWorker) {
	ws.IMP2pWorker = w

	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Msgf("immediate message listening thread exit:=>%s", err)

			return
		}
		p2pMsg := &pbs.WsMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Msg("failed parse p2p message")
			continue
		}
		if p2pMsg.Typ != pbs.WsMsgType_ImmediateMsg {
			utils.LogInst().Warn().Msg("unknown msg typ in p2p immediate message channel")
			continue
		}
		if err := ws.peerImmediateMsg(p2pMsg); err != nil {
			utils.LogInst().Warn().Err(err).Send()
		}
	}
}

func (ws *Service) peerImmediateMsg(msg *pbs.WsMsg) error {
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
