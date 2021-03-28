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

func (x *WsMsg) ReadOnlineFromCli(conn *websocket.Conn) error {
	_, message, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	if err := proto.Unmarshal(message, x); err != nil {
		return err
	}
	if x.Typ != WsMsgType_Online {
		return fmt.Errorf("invalid online msg type")
	}
	online, ok := x.Payload.(*WsMsg_Online)
	if !ok {
		return fmt.Errorf("cast to online message failed")
	}
	if success := online.Online.Verify(x.Sig); !success {
		return fmt.Errorf("verfiy signature failed")
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
		return err
	}

	return conn.WriteMessage(websocket.TextMessage, ackData)
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

func (x *WSCryptoMsg) DBKey() []byte {
	key := fmt.Sprintf(MSGPatternHead, x.To, x.UnixTime)
	return []byte(key)
}
