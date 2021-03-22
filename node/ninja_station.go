package node

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/ninjahome/ninja-go/service"
	"github.com/ninjahome/ninja-go/utils"
)

type NinjaStation struct {
	p2pHost    host.Host
	msgManager *PubSub
	ctxCancel  context.CancelFunc
	ctx        context.Context
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

	ps, err := newPubSub(ctx, h)
	if err != nil {
		panic(err)
	}
	n := &NinjaStation{
		p2pHost:    h,
		msgManager: ps,
		ctx:        ctx,
		ctxCancel:  cancel,
	}
	n.initRpcApis()

	utils.LogInst().Info().Msgf("p2p with id[%s] created addrs:%s", h.ID(), h.Addrs())
	return n
}

func (nt *NinjaStation) Start() error {

	if err := service.Inst().StartService(); err != nil {
		return err
	}

	return nt.msgManager.start()
}

func (nt *NinjaStation) ShutDown() {
}
