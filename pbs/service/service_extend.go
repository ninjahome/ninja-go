package service

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/ninjahome/ninja-go/wallet"
	"time"
)

func (x *WSOnline) Verify(sig []byte) bool {
	s := &bls.Sign{}
	if err := s.Deserialize(sig); err != nil {
		fmt.Println(err)
		return false
	}

	p := &bls.PublicKey{}
	if err := p.DeserializeHexStr(x.UID); err != nil {
		fmt.Println(err)
		return false
	}

	data, err := proto.Marshal(x)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return wallet.VerifyByte(s, p, data)
}

func (x *ClientChatMsg) ReadOnlineFromCli(conn *websocket.Conn) (*WSOnline, error) {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	if mt != int(SrvMsgType_Online) {
		return nil, fmt.Errorf("first msg must be online noti")
	}
	if err := proto.UnmarshalMerge(message, x); err != nil {
		return nil, err
	}

	online, ok := x.Payload.(*ClientChatMsg_Online)
	if !ok {
		return nil, fmt.Errorf("convert to online msg failed")
	}

	if success := online.Online.Verify(x.Sig); !success {
		return nil, fmt.Errorf("verfiy signature failed")
	}

	return online.Online, nil
}

func (x *ClientChatMsg) Online(conn *websocket.Conn, key *wallet.Key) error {
	online := &WSOnline{
		UID:      key.Address.String(),
		UnixTime: time.Now().Unix(),
	}
	data, err := proto.Marshal(online)
	if err != nil {
		return err
	}
	x.Hash = nil
	x.Sig = key.SignData(data)
	x.Payload = &ClientChatMsg_Online{
		Online: online,
	}

	xData, err := proto.Marshal(x)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(int(SrvMsgType_Online), xData)
	if err != nil {
		return err
	}

	return nil
}
