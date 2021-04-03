package main

import (
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
)

type MacChatCli struct {
	receiver   string
	wsCli      *client.WSClient
	contactCli *client.ContactCli
}

func (w MacChatCli) InputMsg(msg *pbs.WSCryptoMsg) error {
	fmt.Println(string(msg.PayLoad))
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
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return
	}
	defer terminal.Restore(0, oldState)
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	term := terminal.NewTerminal(screen, ">")
	for {
		msg, err := term.ReadLine()
		if err != nil {
			panic(err)
		}
		if msg == "" {
			continue
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

func (w MacChatCli) Run() error {
	if err := w.wsCli.Online(); err != nil {
		return err
	}
	go w.writeFromStdio()
	//go w.contactWindow()
	return nil
}
