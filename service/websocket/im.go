package websocket

import (
	"errors"
	"fmt"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/push"
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

func (ws *Service)_procUnicastIM(msg *pbs.WsMsg) error  {
	body := msg.Payload.(*pbs.WsMsg_Message)

	im := body.Message

	utils.LogInst().Debug().
		Str("From", im.From).
		Str("TO", im.To).Int64("time", im.UnixTime).
		Msgf("IM Received time:%d", im.UnixTime)

	if !ws.onlineSet.contains(im.To) {
		utils.LogInst().Debug().Str("Receiver", im.To).
			Str("Status", "offline").Send()
		key := IMDBKey(im.To, im.UnixTime)

		if dt, typ, err := ws.GetToken(im.To); err != nil {
			utils.LogInst().Debug().Str("procIM", im.To).
				Str("device type", "not found in db").Send()
		} else {
			if typ == DevTypeIOS {
				ui := fmt.Sprintf("ios: uid: %s , token: %s, typ: %d", im.To, dt, typ)
				utils.LogInst().Debug().Str("procIM ", im.To).
					Str("Status", "not found in db").Send()
				utils.LogInst().Debug().Str("procIM begin to push", ui).Send()
				ws.iosPush.IOSPushMessage("you have a message", dt)
			} else if typ == DevTypeAndroid {
				ui := fmt.Sprintf("android: uid: %s , token: %s, typ: %d", im.To, dt, typ)
				utils.LogInst().Debug().Str("procIM begin to push", ui).Send()
				utils.LogInst().Debug().Str("procIM ", im.To).
					Str("Status", "not found in db").Send()
				ext := make(map[string]string)
				push.AndroidMessagePush("you have a message", dt, ext)
			}
		}

		return ws.dataBase.Put(key, im.MustData(), nil)
	}

	if user, ok := ws.userTable.get(im.To); ok {
		utils.LogInst().Debug().Str("Receiver", im.To).
			Str("Status", "Same node").Send()
		return user.writeToCli(msg)
	}

	utils.LogInst().Debug().Str("Receiver", im.To).
		Str("Status", "on other node").Send()

	return ws.IMP2pWorker.BroadCast(msg.Data())
}

func (ws *Service)_procMulticastIM(msg *pbs.WsMsg) error {
	body := msg.Payload.(*pbs.WsMsg_GroupMessage)

	gim:=body.GroupMessage

	utils.LogInst().Debug().Str("From:",gim.From).Str("To","Group").Int64("time",gim.UnixTime)

	allonline := true
	otherNode := false
	var (
		groupKey string
		err error
	)

	for i:=0;i<len(gim.To);i++{
		if !ws.onlineSet.contains(gim.To[i].MemberId){
			if allonline{
				allonline = false
				groupKey,err = SaveGroupMsg(ws.dataBase,gim)
				if err!=nil{
					return err
				}
			}

			err = SaveReceiverGroupMsg(ws.dataBase,gim.To[i].MemberId,[]byte(groupKey),gim.UnixTime)
			if err!=nil{
				return err
			}
		}else if user,ok:=ws.userTable.get(gim.To[i].MemberId);ok{
			return user.writeToCli(msg)
		}else{
			otherNode = true
		}
	}

	if otherNode{
		return ws.IMP2pWorker.BroadCast(msg.Data())
	}

	return nil
}

func (ws *Service) procIM(msg *pbs.WsMsg) error {

	switch msg.Payload.(type) {
	case *pbs.WsMsg_Message:
		return ws._procUnicastIM(msg)
	case *pbs.WsMsg_GroupMessage:
		return ws._procMulticastIM(msg)
	default:
		return errors.New("not a correct message type")
	}
}

func (ws *Service) ImmediateMsgForP2pNetwork(w *worker.TopicWorker) {
	ws.IMP2pWorker = w

	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Str("Peer IM read", err.Error())
			return
		}

		if msg.ReceivedFrom.String() == ws.id {
			continue
		}

		p2pMsg := &pbs.WsMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Str("IM Unmarshal", err.Error()).Send()
			continue
		}
		if p2pMsg.Typ != pbs.WsMsgType_ImmediateMsg {
			utils.LogInst().Warn().Str("Invalid Peer IM", p2pMsg.Typ.String()).Send()
			continue
		}
		if err := ws.peerImmediateMsg(p2pMsg); err != nil {
			utils.LogInst().Warn().Err(err).Send()
		}
	}
}


func (ws *Service) _peerImmediateMsg(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_Message)
	if !ok {
		return fmt.Errorf("this is not a valid p2p crypto message")
	}

	u, ok := ws.userTable.get(body.Message.To)
	if !ok {
		return nil
	}
	utils.LogInst().Debug().Str("Peer IM TO", body.Message.To).Send()
	return u.writeToCli(msg)
}


func (ws *Service) _peerImmediateGroupMsg(msg *pbs.WsMsg) error {
	body, ok := msg.Payload.(*pbs.WsMsg_GroupMessage)
	if !ok {
		return fmt.Errorf("this is not a valid p2p crypto message")
	}

	gim:=body.GroupMessage

	for i:=0;i<len(gim.To);i++{
		u, ok := ws.userTable.get(gim.To[i].MemberId)
		if !ok {
			continue
		}
		utils.LogInst().Debug().Str("Peer Group IM TO", gim.To[i].MemberId).Send()
		u.writeToCli(msg)
	}

	return nil
}

func (ws *Service) peerImmediateMsg(msg *pbs.WsMsg) error {

	switch msg.Payload.(type) {
	case *pbs.WsMsg_Message:
		return ws._peerImmediateMsg(msg)
	case *pbs.WsMsg_GroupMessage:
		return ws._peerImmediateGroupMsg(msg)
	default:
		return errors.New("not a correct message type from peer")
	}

}
