package node

import (
	"encoding/json"
	"github.com/ninjahome/ninja-go/utils"
)

//---service debug
func (nt *NinjaStation) DebugTopicMsg(topic, msg string) string {
	if err := nt.pubSub.SendMsg(MessageChannel(topic), []byte(msg)); err != nil {
		return err.Error()
	}
	return "publish success!"
}

func (nt *NinjaStation) DebugTopicPeers(topic string) string {
	utils.LogInst().Debug().Msgf("p2p cmd service query for topic[%s]", topic)
	peers := nt.pubSub.PeersOfTopic(topic)
	bts, _ := json.Marshal(peers)
	return string(bts)
}
