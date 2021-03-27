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

var rpcCmd = &cobra.Command{
	Use:   "service",
	Short: "service",
	Long:  `TODO::.`,
	Run:   rpcUsage,
}

var (
	topic   string
	msgBody string
)

func init() {
	pushCmd.Flags().StringVarP(&topic, "topic", "t", string(node.P2pChanDebug),
		"ninja debug push -t [TOPIC]")
	pushCmd.Flags().StringVarP(&msgBody, "message", "m", "",
		"ninja debug push -t [TOPIC] -m \"[MESSAGE]\"")
	DebugCmd.AddCommand(pushCmd)

	showPeerCmd.Flags().StringVarP(&topic, "topic", "t", string(node.P2pChanDebug),
		"ninja debug peers -t [TOPIC]")
	DebugCmd.AddCommand(showPeerCmd)

	DebugCmd.AddCommand(rpcCmd)
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

func rpcUsage(c *cobra.Command, _ []string) {
	_ = c.Usage()
}
