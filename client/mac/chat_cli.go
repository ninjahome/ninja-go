package main

import (
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

type MacChatCli struct {
	receiver   string
	wsCli      *client.WSClient
	contactCli *client.ContactCli
}

func (w MacChatCli) InputMsg(msg *pbs.WSCryptoMsg) error {
	fmt.Println(msg.String())
	return nil
}

func (w MacChatCli) WebSocketClosed() {
	fmt.Println("offline......")
	os.Exit(0)
}

func (w MacChatCli) UnreadMsg(msgs []*pbs.WSCryptoMsg) error {
	for _, msg := range msgs {
		fmt.Println(msg.String())
	}
	return nil
}

func (w MacChatCli) writeFromStdio() {
	term := terminal.NewTerminal(os.Stdin, ">")
	for {
		msg, err := term.ReadLine()
		if err != nil {
			panic(err)
		}
		if err := w.wsCli.Write(w.receiver, []byte(msg)); err != nil {
			panic(err)
		}
	}
}

func (w MacChatCli) contactWindow() {
	term := terminal.NewTerminal(os.Stdin, "*")
	for {
		cmd, err := term.ReadLine()
		if err != nil {
			panic(err)
		}

		switch cmd {
		case "sync":
			contacts := w.contactCli.SyncContact()
			for _, c := range contacts {
				fmt.Println(c.String())
			}
		}
	}
}

func (w MacChatCli) Run() {
	go w.writeFromStdio()
}
