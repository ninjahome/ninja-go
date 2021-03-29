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

func (x *WsMsg) ReadOnlineFromCli(conn *websocket.Conn) (olMsg *WSOnline, err error) {
	_, message, err := conn.ReadMessage()
	if err != nil {
		return
	}

	if err = proto.Unmarshal(message, x); err != nil {
		return
	}
	if x.Typ != WsMsgType_Online {
		err = fmt.Errorf("invalid online msg type")
		return
	}
	online, ok := x.Payload.(*WsMsg_Online)
	if !ok {
		err = fmt.Errorf("cast to online message failed")
		return
	}
	olMsg = online.Online
	if success := olMsg.Verify(x.Sig); !success {
		err = fmt.Errorf("verfiy signature failed")
		return
	}

	ack := &WSOnlineAck{
		Success: true,
		Seq:     online.Online.UnixTime,
	}
	ackWrap := &WsMsg{
		Typ:     WsMsgType_OnlineACK,
		Payload: &WsMsg_OlAck{ack},
	}
	ackData, err := proto.Marshal(ackWrap)
	if err != nil {
		return
	}
	return olMsg, conn.WriteMessage(websocket.TextMessage, ackData)
}

func (x *WsMsg) Online(conn *websocket.Conn, key *wallet.Key) error {
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
	x.Payload = &WsMsg_Online{online}
	x.Typ = WsMsgType_Online

	xData, err := proto.Marshal(x)
	if err != nil {
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, xData)
	if err != nil {
		return err
	}

	return nil
}

func (x *WSCryptoMsg) MustData() []byte {
	data, _ := proto.Marshal(x)
	return data
}
