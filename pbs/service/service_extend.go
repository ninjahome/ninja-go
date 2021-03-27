package service

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/ninjahome/ninja-go/wallet"
	"time"
)

func (x *WSOnline) ReadOnlineFromCli(conn *websocket.Conn) error {
	mt, message, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	if mt != int(SrvMsgType_Online) {
		return fmt.Errorf("first msg must be online noti")
	}
	if err := proto.UnmarshalMerge(message, x); err != nil {
		return err
	}

	//TODO::verify user's balance and signature

	return nil
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
