package node

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/ninjahome/ninja-go/node/worker"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"sync"
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
	workers   worker.WorkGroup
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

	ps, err := newWorkGroup(ctx, h)
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
	utils.LogInst().Info().Msgf("p2p with id[%s] created addrs:%s", h.ID(), h.Addrs())
	return n
}

func (nt *NinjaNode) Start() error {
	websocket.Inst().StartService(nt.nodeID, nt.ctx)
	contact.Inst().StartService(nt.nodeID, nt.ctx)

	workers := make(worker.WorkGroup)
	for topID, r := range systemTopics {
		topic, err := nt.pubSubs.Join(topID)
		if err != nil {
			return err
		}
		w := worker.NewWorker(nt.ctx, topID, topic, r)
		workers[topID] = w
		if err := w.StartWork(); err != nil {
			return err
		}
	}
	nt.workers = workers

	return nil
}

func (nt *NinjaNode) ShutDown() {
	if nt.workers == nil {
		return
	}
	for _, t := range nt.workers {
		t.StopWork()
	}
	nt.workers = nil
	nt.ctxCancel()
	_ = nt.p2pHost.Close()
}

func (nt *NinjaNode) PeersOfTopic(topic string) []peer.ID {
	w, ok := nt.workers[(topic)]
	if !ok {
		return nil
	}
	return w.PeersOfTopic()
}
