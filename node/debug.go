package node

import (
	"encoding/json"
	"github.com/ninjahome/ninja-go/node/worker"
	"github.com/ninjahome/ninja-go/utils"
)

func (nt *NinjaNode) DebugTopicMsg(topic, msg string) string {
	worker, ok := nt.workers[topic]
	if !ok {
		return "no such topic"
	}
	if err := worker.WriteData([]byte(msg)); err != nil {
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
	for true {
		select {
		case <-w.Stop:
			utils.LogInst().Warn().Msg("debug peer message thread exit")
			return
		default:
			msg, err := w.Sub.Next(nt.ctx)
			if err != nil {
				utils.LogInst().Err(err).Send()
				return
			}
			utils.LogInst().Debug().Msg(msg.String())
		}
	}
}
