package contact

import (
	"fmt"
	"github.com/ninjahome/bls-wallet/bls"
	"github.com/ninjahome/ninja-go/wallet"
	"google.golang.org/protobuf/proto"
)

func ErrAck(msg string) []byte {
	ack := &ContactMsg{
		Typ: ContactMsgType_MTAck,
		PayLoad: &ContactMsg_OpAck{OpAck: &OperateAck{
			Success: false,
			Msg:     msg,
		}},
	}

	data, _ := proto.Marshal(ack)
	return data
}

func OkAck() []byte {
	ack := &ContactMsg{
		Typ: ContactMsgType_MTAck,
		PayLoad: &ContactMsg_OpAck{OpAck: &OperateAck{
			Success: true,
		}},
	}
	data, _ := proto.Marshal(ack)
	return data
}

func (x *ContactMsg) Verify(data []byte) bool {
	s := &bls.Sign{}
	if err := s.Deserialize(x.Sig); err != nil {
		fmt.Println(err)
		return false
	}

	p := &bls.PublicKey{}
	if err := p.DeserializeHexStr(x.From); err != nil {
		fmt.Println(err)
		return false
	}
	return wallet.VerifyByte(s, p, data)
}

func (x *ContactItem) Data() []byte {
	itemData, _ := proto.Marshal(x)
	return itemData
}

func (x *ContactMsg) Data() []byte {
	data, _ := proto.Marshal(x)
	return data
}
