package main

import (
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
)

type MacChatCli struct {
	wsCli      *client.WSClient
	contactCli *client.ContactCli
}

func (w MacChatCli) InputMsg(msg *pbs.WSCryptoMsg) error {
	panic("implement me")
}

func (w MacChatCli) DidClosed() {
	panic("implement me")
}

func (w MacChatCli) UnreadMsg(msgs []*pbs.WSCryptoMsg) error {
	panic("implement me")
}
