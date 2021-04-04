package main

import (
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"strconv"
)

type MacChatCli struct {
	receiver   string
	wsCli      *client.WSClient
	contactCli *client.ContactCli
}

func (w MacChatCli) ImmediateMessage(msg *pbs.WSCryptoMsg) error {
	fmt.Printf("\r\n%s\r\n>", string(msg.PayLoad))
	fmt.Printf("\r\n%s\r\n>", msg.From)
	return nil
}

func (w MacChatCli) WebSocketClosed() {
	fmt.Println("offline......")
	os.Exit(0)
}

func (w MacChatCli) UnreadMsg(msgs *pbs.WSUnreadAck) error {

	fmt.Printf("\r\n%s\r\n>", msgs.NodeID)
	fmt.Printf("\r\n%s\r\n>", msgs.Receiver)

	for _, msg := range msgs.Payload {
		fmt.Printf("\r\n%s\r\n>", msg.From)
		fmt.Printf("\r\n%s\r\n>", string(msg.PayLoad))
	}
	fmt.Print("\r\n>")
	return nil
}

func (w MacChatCli) writeFromStdio(term *terminal.Terminal, state *terminal.State) {
	defer terminal.Restore(0, state)

	if err := w.wsCli.Online(); err != nil {
		panic(err)
	}
	if err := w.wsCli.PullMsg(0); err != nil {
		panic(err)
	}

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
func (w MacChatCli) contactWindow(term *terminal.Terminal, state *terminal.State) {
	defer terminal.Restore(0, state)

	for {
		cmd, err := term.ReadLine()
		if err != nil {
			panic(err)
		}
		no, err := strconv.Atoi(cmd)
		if err != nil {
			continue
		}
		switch no {
		case 1:
			contacts, err := w.contactCli.SyncContact()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println("contact query result:", len(contacts))
			for _, c := range contacts {
				fmt.Println(c.String())
			}

		case 2:
			cid, err := term.ReadLine()
			if err != nil {
				panic(err)
			}
			nickName, err := term.ReadLine()
			if err != nil {
				panic(err)
			}

			if err := w.contactCli.AddContact(cid, nickName, ""); err != nil {
				panic(err)
			}
		case 3:
			cid, err := term.ReadLine()
			if err != nil {
				panic(err)
			}
			nickName, err := term.ReadLine()
			if err != nil {
				panic(err)
			}

			if err := w.contactCli.UpdateContact(cid, nickName, ""); err != nil {
				panic(err)
			}

		case 4:
			cid, err := term.ReadLine()
			if err != nil {
				panic(err)
			}
			if err := w.contactCli.DelContact(cid); err != nil {
				panic(err)
			}
		default:
			fmt.Println("unknown command")
		}
	}
}

func (w MacChatCli) Run() error {

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return err
	}
	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	term := terminal.NewTerminal(screen, ">")

	//go w.writeFromStdio(term, oldState)
	go w.contactWindow(term, oldState)
	return nil
}
