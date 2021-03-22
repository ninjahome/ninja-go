package node

import (
	"encoding/json"
	"github.com/ninjahome/ninja-go/utils"
)

type RpcPushTopic struct {
	Topics  string `json:"topic"`
	Message string `json:"msg"`
}

func (nt *NinjaStation) initRpcApis() {
	//service.HttpRpcApis["/node/PeerList"] = nt.ApiPeesList
	//service.HttpRpcApis["/node/PushMsg"] = nt.ApiPushMsg
	//service.HttpRpcApis["/node/nid"] = nt.HostID
}

//--->public service apis
//func (nt *NinjaStation) ApiPeesList(request *pbs.RpcMsgItem) *pbs.RpcResponse {
//	peerStr := nt.DebugTopicPeers(string(request.Parameter))
//	return pbs.RpcOk([]byte(peerStr))
//}
//
//func (nt *NinjaStation) HostID(_ *pbs.RpcMsgItem) *pbs.RpcResponse {
//	return pbs.RpcOk([]byte(nt.p2pHost.ID()))
//}
//
//func (nt *NinjaStation) ApiPushMsg(request *pbs.RpcMsgItem) *pbs.RpcResponse {
//	param := &RpcPushTopic{}
//	if err := json.Unmarshal(request.Parameter, param); err != nil {
//		return pbs.RpcError(err.Error())
//	}
//	res := nt.DebugTopicMsg(param.Topics, param.Message)
//	return pbs.RpcOk([]byte(res))
//}

//---service debug
func (nt *NinjaStation) DebugTopicMsg(topic, msg string) string {
	if err := nt.msgManager.SendMsg(topic, []byte(msg)); err != nil {
		return err.Error()
	}
	return "publish success!"
}

func (nt *NinjaStation) DebugTopicPeers(topic string) string {
	utils.LogInst().Debug().Msgf("p2p cmd service query for topic[%s]", topic)
	peers := nt.msgManager.PeersOfTopic(topic)
	bts, _ := json.Marshal(peers)
	return string(bts)
}
