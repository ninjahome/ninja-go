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

var keyStr = `
{
        "address": "a653234466b9cc858f238b49771298a853be717a112ee280f8b6d4fe6bf01c71cfe5334435be4a8f43cc436002763de5",
        "crypto": {
                "cipher": "aes-128-ctr",
                "ciphertext": "f3ab5c896960e87b550f2a9daf10e31d827a35202303abbd1e43a98dde033c66",
                "cipherParams": {
                        "iv": "84170a60bd5bd82a6380d74dd2ba2908"
                },
                "kdf": "scrypt",
                "kdfParams": {
                        "dklen": 32,
                        "n": 262144,
                        "p": 1,
                        "r": 8,
                        "salt": "90b189ff698a2a3f753c31cc19b8d29ee48dcab1ffac74d0be06a263ec4bfb9e"
                },
                "mac": "0b945ee59581c17ec227a8b064994330e0882c6f418abbd18156d5f4372cef82"
        },
        "id": "b57ec371-2e95-4038-b3e0-45b104519556",
        "version": 1
}
`

func main() {

	ss := &MacChatCli{}
	key, err := wallet.LoadKeyFromJsonStr(keyStr, "123")
	if err != nil {
		panic(err)
	}
	fmt.Printf("key is loaded :=>%t\n", key.IsOpen())

	ws, err := client.NewWSClient("167.179.78.33:6666", key, ss) //202.182.101.145//167.179.78.33//127.0.0.1
	if err != nil {
		panic(err)
	}
	ss.wsCli = ws

	c, err := client.NewContactCli("167.179.78.33:6667", key) //202.182.101.145
	if err != nil {
		panic(err)
	}

	ss.contactCli = c

	if err := ss.Run(); err != nil {
		panic(err)
	}

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
