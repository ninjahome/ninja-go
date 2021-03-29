package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"testing"
	"time"
)

func TestProtoMsgStr(t *testing.T) {
	msg := &pbs.WSCryptoMsg{
		From:     "Alice",
		To:       "Bob",
		PayLoad:  []byte("Hello world"),
		UnixTime: time.Now().Unix(),
	}

	bts, err := proto.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("===>step1:->", string(bts))

	s2 := msg.String()
	fmt.Println("===>step2:->", s2)
}
