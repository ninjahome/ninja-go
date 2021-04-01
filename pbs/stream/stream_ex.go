package stream

import "google.golang.org/protobuf/proto"

func (x *StreamMsg) SyncOnline() []byte {
	x.MTyp = StreamMType_MTOnlineSync
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
