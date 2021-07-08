package node

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/ninjahome/ninja-go/node/worker"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"sync"
	"time"
)

var _instance *NinjaNode
var once sync.Once

func Inst() *NinjaNode {
	once.Do(func() {
		_instance = newNode()
	})
	return _instance
}

type NinjaNode struct {
	nodeID    string
	p2pHost   host.Host
	tWorkers  worker.WorkGroup
	pubSubs   *pubsub.PubSub
	ctxCancel context.CancelFunc
	ctx       context.Context
}

var systemTopics map[string]worker.TopicReader

func newNode() *NinjaNode {
	if _nodeConfig == nil {
		panic("Please init p2p _nodeConfig first")
	}

	opts := _nodeConfig.initOptions()
	ctx, cancel := context.WithCancel(context.Background())
	h, err := libp2p.New(ctx, opts...)
	if err != nil {
		panic(err)
	}
	_nodeConfig.initStreamWorker(h)
	ps, err := newPubSub(ctx, h)
	if err != nil {
		panic(err)
	}
	n := &NinjaNode{
		nodeID:    h.ID().String(),
		p2pHost:   h,
		pubSubs:   ps,
		ctx:       ctx,
		ctxCancel: cancel,
	}

	systemTopics[P2pChanDebug] = n.DebugPeerMsg

	contact.FGetSrvPort = GetSrvPost

	utils.LogInst().Info().Str("NodeID", h.ID().String()).
		Msgf("address:%s", h.Addrs())
	return n
}

func (nt *NinjaNode) RandomPeer(protocID protocol.ID) (network.Stream, error) {
	peers := nt.p2pHost.Network().Peers()
	for _, pid := range peers {
		stream, err := nt.p2pHost.NewStream(nt.ctx, pid, protocID)
		if err == nil {
			utils.LogInst().Info().Str("Selected Peer", pid.String()).Send()
			return stream, nil
		}
	}
	return nil, fmt.Errorf("no valid peer id")
}

func (nt *NinjaNode) Start() error {

	workers := make(worker.WorkGroup)
	if err := workers.StartUp(nt.ctx, nt.pubSubs, systemTopics, _nodeConfig.WorkerStartTimeOut); err != nil {
		return err
	}
	nt.tWorkers = workers

	stream, err := nt.RandomPeer(StreamSyncOnline)
	if err != nil {
		utils.LogInst().Warn().Msg("got random stream failed may be i'm genesis......")
	} else {
		if err := websocket.Inst().SyncOnlineSetFromPeerNodes(stream); err != nil {
			return err
		}
	}

	stream, err = nt.RandomPeer(StreamSyncDevTokens)
	if err != nil {
		utils.LogInst().Warn().Msg("got devtokens random stream failed may be i'm genesis......")
	} else {
		if err := websocket.Inst().SyncDevInfoFromPeerNodes(stream); err != nil {
			return err
		}
	}

	websocket.Inst().StartService(nt.nodeID)

	contactSyncWorker := &worker.StreamWorker{
		ProtoID: StreamContactQuery,
		SGetter: nt.RandomPeer,
	}
	contact.Inst().StartService(nt.nodeID, contactSyncWorker)

	utils.LogInst().Info().Msg(">>>>>>>>>>>>>>>>>>>>>>>>>>Node start success......")

	return nil
}

func (nt *NinjaNode) ShutDown() {
	websocket.Inst().ShutDown()
	contact.Inst().ShutDown()

	if nt.tWorkers == nil {
		return
	}
	nt.tWorkers.StopWork()
	nt.tWorkers = nil

	nt.ctxCancel()
	_ = nt.p2pHost.Close()
	time.Sleep(100 * time.Millisecond)
}

func (nt *NinjaNode) PeersOfTopic(topic string) []peer.ID {
	w, ok := nt.tWorkers[(topic)]
	if !ok {
		return nil
	}
	//tw, ok := w.(*worker.TopicWorker)
	//if !ok{
	//	return nil
	//}
	return w.PeersOfTopic()
}
