package stream

import "google.golang.org/protobuf/proto"

func (x *StreamMsg) SyncOnline(nodeId string) []byte {
	x.MTyp = StreamMType_MTOnlineSync
	x.Payload = &StreamMsg_OnlineSync{OnlineSync: &OnlineSync{NodeID: nodeId}}
	data, _ := proto.Marshal(x)
	return data
}

func (x *StreamMsg) SyncOnlineAck(uid []string) []byte {
	x.MTyp = StreamMType_MTOnlineAck
	x.Payload = &StreamMsg_OnlineAck{OnlineAck: &OnlineMap{UID: uid}}

	data, _ := proto.Marshal(x)
	return data
}
func (x *StreamMsg) Data() []byte {
	data, _ := proto.Marshal(x)
	return data
}
