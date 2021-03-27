package node

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/service"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"google.golang.org/protobuf/proto"
)

type NinjaStation struct {
	p2pHost                host.Host
	pubSub                 *PubSub
	ctxCancel              context.CancelFunc
	ctx                    context.Context
	threads                map[string]*thread.Thread
	readInFromPeerMsgQueue chan *pbs.P2PMsg
	outToPeerMsgQueue      chan *pbs.P2PMsg
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
		p2pHost:                h,
		pubSub:                 ps,
		ctx:                    ctx,
		ctxCancel:              cancel,
		readInFromPeerMsgQueue: make(chan *pbs.P2PMsg, 1024), //TODO::config
		outToPeerMsgQueue:      make(chan *pbs.P2PMsg, 1024),
	}
	utils.LogInst().Info().Msgf("p2p with id[%s] created addrs:%s", h.ID(), h.Addrs())
	return n
}

func (nt *NinjaStation) Start() error {
	service.Inst().StartService(nt.outToPeerMsgQueue)

	t := thread.NewThreadWithName(THNOuterMsgReader, func(stop chan struct{}) {
		nt.waitMsgWork(stop)
	})
	nt.threads[THNOuterMsgReader] = t
	t.Run()

	for id, topic := range nt.pubSub.topics {
		sub, err := topic.Subscribe()
		if err != nil {
			return err
		}
		t := thread.NewThreadWithName(id.String(), func(stop chan struct{}) {
			utils.LogInst().Info().Msgf("[pubSub] start reading message for topic[%s]", id)

			nt.pubSub.readingMessage(stop, sub, nt.readInFromPeerMsgQueue)

			defer nt.pubSub.removeTopic(id)

		})
		nt.threads[id.String()] = t
		t.Run()
	}
	return nil
}

func (nt *NinjaStation) ShutDown() {
	for _, t := range nt.threads {
		t.Stop()
	}
	nt.threads = nil
	//TODO pubSub destroy
}

func (nt *NinjaStation) waitMsgWork(stop chan struct{}) {

	for {
		select {
		case msg := <-nt.readInFromPeerMsgQueue:
			if err := nt.procInputChMsg(msg); err != nil {
				utils.LogInst().Warn().Err(err).Send()
			}

		case msg := <-nt.outToPeerMsgQueue:
			if err := nt.procOuterChMsg(msg); err != nil {
				utils.LogInst().Warn().Err(err).Send()
			}

		case <-stop:
			utils.LogInst().Warn().Msg("node outer message thread exit")
			return
		}
	}
}

func (nt *NinjaStation) procOuterChMsg(msg *pbs.P2PMsg) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	switch msg.MsgTyp {
	case pbs.P2PMsgType_P2pCryptoMsg:
		return nt.pubSub.SendMsg(P2pChanCryptoMsg, data)

	case pbs.P2PMsgType_P2pOnline:
		return nt.pubSub.SendMsg(P2pChanUserOnline, data)
	default:
		utils.LogInst().Warn().Msgf("unknown to output peer to peer msg type:[%d]", msg.MsgTyp)
	}
	return nil
}

func (nt *NinjaStation) procInputChMsg(msg *pbs.P2PMsg) error {

	switch msg.MsgTyp {

	case pbs.P2PMsgType_P2pOnline:
		body, ok := msg.Payload.(*pbs.P2PMsg_Online)
		if !ok {
			return fmt.Errorf("this is not a valid online p2p message")
		}
		return service.Inst().OnlineFromOtherPeer(body.Online)

	case pbs.P2PMsgType_P2pCryptoMsg:
		body, ok := msg.Payload.(*pbs.P2PMsg_Msg)
		if !ok {
			return fmt.Errorf("this is not a valid p2p crypto message")
		}
		return service.Inst().PeerImmediateCryptoMsg(body.Msg)

	default:
		utils.LogInst().Warn().Msgf("unknown read in peer to peer msg type:[%d]", msg.MsgTyp)
	}

	return nil
}
