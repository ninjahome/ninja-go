package node

import (
	"encoding/json"
	"github.com/ninjahome/ninja-go/node/worker"
	"github.com/ninjahome/ninja-go/utils"
)

func (nt *NinjaNode) DebugTopicMsg(topic, msg string) string {
	w, ok := nt.tWorkers[topic]
	if !ok {
		return "no such topic"
	}
	if err := w.WriteData([]byte(msg)); err != nil {
		return err.Error()
	}
	return "publish success!"
}

func (nt *NinjaNode) DebugTopicPeers(topic string) string {
	utils.LogInst().Debug().Msgf("p2p cmd service query for topic[%s]", topic)
	peers := nt.PeersOfTopic(topic)
	bts, _ := json.Marshal(peers)
	return string(bts)
}

func (nt *NinjaNode) DebugPeerMsg(w *worker.TopicWorker) {
	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Msgf("debug peer message thread exit:=>%s", err)
			return
		}
		utils.LogInst().Debug().Msg(msg.String())
	}
}
