package node

import (
	"encoding/json"
	"github.com/ninjahome/ninja-go/utils"
)

func (nt *NinjaStation) DebugTopicMsg(topic, msg string) string {
	worker, ok := nt.workers[topic]
	if !ok {
		return "no such topic"
	}
	if err := worker.tWriter.Publish(nt.ctx, []byte(msg)); err != nil {
		return err.Error()
	}
	return "publish success!"
}

func (nt *NinjaStation) DebugTopicPeers(topic string) string {
	utils.LogInst().Debug().Msgf("p2p cmd service query for topic[%s]", topic)
	peers := nt.PeersOfTopic(topic)
	bts, _ := json.Marshal(peers)
	return string(bts)
}
