package cmd

import (
	"context"
	"fmt"
	"github.com/ninjahome/ninja-go/node"
	pbs "github.com/ninjahome/ninja-go/pbs/cmd"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type cmdService struct{}

func (c *cmdService) WebSocketInfo(ctx context.Context, req *pbs.WSInfoReq) (*pbs.CommonResponse, error) {
	result := websocket.Inst().DebugInfo(req.Online, req.Local, req.UserAddr)
	return &pbs.CommonResponse{
		Msg: result,
	}, nil
}

func (c *cmdService) ShowAllThreads(ctx context.Context, req *pbs.ThreadGroup) (*pbs.CommonResponse, error) {
	var result = ""
	if req.List {
		result = thread.Inst().AllThread()
	}

	if req.ThreadName != "" {
		result += thread.Inst().ThreadInfo(req.ThreadName)
	}

	return &pbs.CommonResponse{
		Msg: result,
	}, nil
}

const (
	DefaultCmdPort = 8848
	ThreadName     = "Internal Rpc Cmd Thread"
)

var (
	_instance = &cmdService{}
)

func (c *cmdService) P2PShowPeers(_ context.Context, peer *pbs.ShowPeer) (*pbs.CommonResponse, error) {
	result := node.Inst().DebugTopicPeers(peer.Topic)
	return &pbs.CommonResponse{
		Msg: result,
	}, nil
}

func (c *cmdService) P2PSendTopicMsg(_ context.Context, msg *pbs.TopicMsg) (*pbs.CommonResponse, error) {
	result := node.Inst().DebugTopicMsg(msg.Topic, msg.Msg)
	return &pbs.CommonResponse{
		Msg: result,
	}, nil
}

func StartCmdRpc(_ chan struct{}) {
	var address = fmt.Sprintf("127.0.0.1:%d", DefaultCmdPort)
	l, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	cmdServer := grpc.NewServer()

	pbs.RegisterCmdServiceServer(cmdServer, _instance)

	reflection.Register(cmdServer)
	utils.LogInst().Info().Msg("command rpc thread start success......")
	if err := cmdServer.Serve(l); err != nil {
		panic(err)
	}
}

func DialToCmdService() pbs.CmdServiceClient {

	var address = fmt.Sprintf("127.0.0.1:%d", DefaultCmdPort)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pbs.NewCmdServiceClient(conn)

	return client
}
