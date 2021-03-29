package node

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
)

type NinjaStation struct {
	nodeID    string
	p2pHost   host.Host
	workers   WorkGroup
	ctxCancel context.CancelFunc
	ctx       context.Context
}

func newStation() *NinjaStation {
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
	n := &NinjaStation{
		nodeID:    h.ID().String(),
		p2pHost:   h,
		workers:   ps,
		ctx:       ctx,
		ctxCancel: cancel,
	}
	utils.LogInst().Info().Msgf("p2p with id[%s] created addrs:%s", h.ID(), h.Addrs())
	return n
}

func (nt *NinjaStation) Start() error {
	websocket.Inst().StartService(nt.nodeID, nt.ctx)
	contact.Inst().StartService(nt.nodeID, nt.ctx)

	for _, worker := range nt.workers {
		if err := worker.startWork(); err != nil {
			return err
		}
	}

	return nil
}

func (nt *NinjaStation) ShutDown() {
	if nt.workers == nil {
		return
	}
	for _, t := range nt.workers {
		t.stopWork()
	}
	nt.workers = nil
	nt.ctxCancel()
	_ = nt.p2pHost.Close()
}

func (nt *NinjaStation) PeersOfTopic(topic string) []peer.ID {
	worker, ok := nt.workers[(topic)]
	if !ok {
		return nil
	}
	return worker.tWriter.ListPeers()
}
