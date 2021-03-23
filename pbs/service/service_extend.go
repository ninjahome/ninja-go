package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

func (x *WSOnline) FullFill(conn *websocket.Conn) error {
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
