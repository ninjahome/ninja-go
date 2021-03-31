package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/client"
	websocket2 "github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/wallet"
	"google.golang.org/protobuf/proto"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func main() {
	ss := &MacChatCli{}
	keyStr := ""
	key, err := wallet.DecryptKey([]byte(keyStr), "123")
	if err != nil {
		panic(err)
	}
	ws, err := client.NewWSClient("202.182.101.145:6666", key, ss)
	if err != nil {
		panic(err)
	}
	ss.wsCli = ws

	c, err := client.NewContactCli("202.182.101.145:6666", key)
	if err != nil {
		panic(err)
	}

	ss.contactCli = c

	ss.Run()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt
}

func test1() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("localhost:%d", websocket2.DefaultWsPort), Path: websocket2.CPUserOnline}
	fmt.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	online := &pbs.WSOnline{
		UID:      "111111",
		UnixTime: time.Now().Unix(),
	}
	data, err := proto.Marshal(online)
	if err != nil {
		panic(err)
	}

	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		panic(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				return
			}
			fmt.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				fmt.Println("write:", err)
				return
			}
		case <-interrupt:
			fmt.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				fmt.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
