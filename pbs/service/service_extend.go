package service

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/ninjahome/ninja-go/wallet"
	"time"
)

const (
	MSGPatternHead = "TempCachedMsg_%s_%d"
	MSGPatternEnd  = "TempCachedMsg_%s_ffffffffffffffff"
)

func (x *WSOnlineData) Verify(sig []byte) bool {
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

func (x *WSOnline) ReadOnlineFromCli(conn *websocket.Conn) (*WSOnlineData, error) {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	if mt != int(SrvMsgType_Online) {
		return nil, fmt.Errorf("first msg must be online noti")
	}
	if err := proto.Unmarshal(message, x); err != nil {
		return nil, err
	}

	online := x.Payload

	if success := online.Verify(x.Sig); !success {
		return nil, fmt.Errorf("verfiy signature failed")
	}

	ack := &WSOnlineAck{
		Success: true,
		Seq:     online.UnixTime,
	}

	ackData, err := proto.Marshal(ack)
	if err != nil {
		return nil, err
	}
	if err := conn.WriteMessage(int(SrvMsgType_OnlineACK), ackData); err != nil {
		return nil, err
	}
	return online, nil
}

func (x *WSOnline) Online(conn *websocket.Conn, key *wallet.Key) error {
	online := &WSOnlineData{
		UID:      key.Address.String(),
		UnixTime: time.Now().Unix(),
	}
	data, err := proto.Marshal(online)
	if err != nil {
		return err
	}
	x.Hash = nil
	x.Sig = key.SignData(data)
	x.Payload = online

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

func (x *WSCryptoMsg) MustData() []byte {
	data, _ := proto.Marshal(x)
	return data
}

func (x *WSCryptoMsg) DBKey() []byte {
	key := fmt.Sprintf(MSGPatternHead, x.To, x.UnixTime)
	return []byte(key)
}
