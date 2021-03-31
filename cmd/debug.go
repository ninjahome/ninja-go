package cmd

import (
	"context"
	"fmt"
	"github.com/ninjahome/ninja-go/node"
	pbs "github.com/ninjahome/ninja-go/pbs/cmd"
	"github.com/spf13/cobra"
)

var DebugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug ",
	Long:  `TODO::.`,
	Run:   debug,
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "ninja debug push -t [TOPIC] -m [MESSAGE]",
	Long:  `TODO::.`,
	Run:   p2pAction,
}

var showPeerCmd = &cobra.Command{
	Use:   "peers",
	Short: "ninja debug showPeer -t [TOPIC]",
	Long:  `TODO::.`,
	Run:   showPeerAction,
}

var websocketCmd = &cobra.Command{
	Use:   "ws",
	Short: "ninja debug ws -u [User Address] -l -o",
	Long:  `TODO::.`,
	Run:   webSocketAction,
}

var threadCmd = &cobra.Command{
	Use:   "thread",
	Short: "ninja debug thread",
	Long:  `TODO::.`,
	Run:   threadAction,
}

var (
	topic   string
	msgBody string
	user    string
	local   bool
	online  bool
	thName  string
)

func init() {
	pushCmd.Flags().StringVarP(&topic, "topic", "t", node.P2pChanDebug,
		"ninja debug push -t [TOPIC]")
	pushCmd.Flags().StringVarP(&msgBody, "message", "m", "",
		"ninja debug push -t [TOPIC] -m \"[MESSAGE]\"")
	DebugCmd.AddCommand(pushCmd)

	showPeerCmd.Flags().StringVarP(&topic, "topic", "t", node.P2pChanDebug,
		"ninja debug peers -t [TOPIC]")
	DebugCmd.AddCommand(showPeerCmd)

	websocketCmd.Flags().StringVarP(&user, "user", "u", "",
		"ninja debug ws -u [User Address]")

	websocketCmd.Flags().BoolVarP(&local, "local", "l", false,
		"ninja debug ws -l")

	websocketCmd.Flags().BoolVarP(&online, "online", "o", false,
		"ninja debug ws -o")

	DebugCmd.AddCommand(websocketCmd)

	threadCmd.Flags().BoolVarP(&local, "list", "l", false,
		"ninja debug thread -l")

	threadCmd.Flags().StringVarP(&thName, "tname", "n", "",
		"ninja debug thread -n [Thread Name]")

	DebugCmd.AddCommand(threadCmd)
}

func debug(c *cobra.Command, _ []string) {
	_ = c.Usage()
}

func p2pAction(c *cobra.Command, _ []string) {
	if topic == "" || msgBody == "" {
		_ = c.Usage()
		return
	}

	cli := DialToCmdService()
	rsp, err := cli.P2PSendTopicMsg(context.Background(), &pbs.TopicMsg{
		Topic: topic,
		Msg:   msgBody,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Msg)
}

func showPeerAction(c *cobra.Command, _ []string) {
	if topic == "" {
		_ = c.Usage()
		return
	}
	cli := DialToCmdService()
	rsp, err := cli.P2PShowPeers(context.Background(), &pbs.ShowPeer{
		Topic: topic,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("peers result:=>", rsp.Msg)
}

func webSocketAction(c *cobra.Command, _ []string) {
	cli := DialToCmdService()
	rsp, err := cli.WebSocketInfo(context.Background(), &pbs.WSInfoReq{
		Online:   online,
		Local:    local,
		UserAddr: user,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Msg)
}

func threadAction(c *cobra.Command, _ []string) {

	cli := DialToCmdService()
	rsp, err := cli.ShowAllThreads(context.Background(), &pbs.ThreadGroup{
		List:       local,
		ThreadName: thName,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(rsp.Msg)
}
