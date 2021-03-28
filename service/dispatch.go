package service

import (
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/utils"
)

func (ws *WebSocketService) sendOffline(msg *pbs.WSCryptoMsg) error {
	utils.LogInst().Debug().Msgf("save message [%s->%s]_%x", msg.From, msg.To, msg.UnixTime)
	return ws.dataBase.Put(msg.DBKey(), msg.MustData(), nil)
}

func (ws *WebSocketService) relayMsg(msg *pbs.WSCryptoMsg) error {
	ws.msgToOtherPeerQueue <- &pbs.P2PMsg{
		MsgTyp:  pbs.P2PMsgType_P2pCryptoMsg,
		Payload: &pbs.P2PMsg_Msg{Msg: msg},
	}
	return nil
}

func (ws *WebSocketService) sendToPeer(msg *pbs.WSCryptoMsg) error {
	if !ws.onlineSet.contains(msg.To) {
		return ws.sendOffline(msg)
	}
	if user, ok := ws.userTable.get(msg.To); ok {
		return user.writeToCli(msg)
	}
	return ws.relayMsg(msg)
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
