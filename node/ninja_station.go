package node

import (
	"context"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	pbs2 "github.com/ninjahome/ninja-go/pbs/contact"
	pbs "github.com/ninjahome/ninja-go/pbs/websocket"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/service/websocket"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
	"google.golang.org/protobuf/proto"
)

type NinjaStation struct {
	nodeID                 string
	p2pHost                host.Host
	pubSub                 *PubSub
	ctxCancel              context.CancelFunc
	ctx                    context.Context
	threads                map[string]*thread.Thread
	readInFromPeerMsgQueue chan *pbs.WsMsg
	outToPeerMsgQueue      chan *pbs.WsMsg
	contactMsgQueue        chan *pbs2.ContactMsg
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
		nodeID:                 h.ID().String(),
		p2pHost:                h,
		pubSub:                 ps,
		ctx:                    ctx,
		ctxCancel:              cancel,
		readInFromPeerMsgQueue: make(chan *pbs.WsMsg, _nodeConfig.MaxMsgQueueSize),
		outToPeerMsgQueue:      make(chan *pbs.WsMsg, _nodeConfig.MaxMsgQueueSize),
		contactMsgQueue:        make(chan *pbs2.ContactMsg, _nodeConfig.MaxMsgQueueSize),
	}
	utils.LogInst().Info().Msgf("p2p with id[%s] created addrs:%s", h.ID(), h.Addrs())
	return n
}

func (nt *NinjaStation) Start() error {
	websocket.Inst().StartService(nt.nodeID, nt.outToPeerMsgQueue)
	contact.Inst().StartService(nt.nodeID, nt.contactMsgQueue)

	t := thread.NewThreadWithName(THNOuterMsgReader, nt.waitMsgWork)
	nt.threads[THNOuterMsgReader] = t
	t.Run()

	for id, topic := range nt.pubSub.topics {
		sub, err := topic.Subscribe()
		if err != nil {
			return err
		}
		t := thread.NewThreadWithName(id.String(), func(stop chan struct{}) {
			nt.pubSub.readingMessage(stop, sub, id, nt.readInFromPeerMsgQueue)
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

		case msg := <-nt.contactMsgQueue:
			if err := nt.procContactMsg(msg); err != nil {
				utils.LogInst().Warn().Err(err).Send()
			}

		case <-stop:
			utils.LogInst().Warn().Msg("node outer message thread exit")
			return
		}
	}
}

func (nt *NinjaStation) procOuterChMsg(msg *pbs.WsMsg) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	switch msg.Typ {
	case pbs.WsMsgType_ImmediateMsg:
		return nt.pubSub.SendMsg(P2pChanImmediateMsg, data)

	case pbs.WsMsgType_Online, pbs.WsMsgType_Offline:
		return nt.pubSub.SendMsg(P2pChanUserOnOffLine, data)

	case pbs.WsMsgType_PullUnread, pbs.WsMsgType_UnreadAck:
		return nt.pubSub.SendMsg(P2pChanUnreadMsg, data)

	default:
		utils.LogInst().Warn().Msgf("unknown to output peer to peer msg type:[%d]", msg.Typ)
	}
	return nil
}

func (nt *NinjaStation) procInputChMsg(msg *pbs.WsMsg) error {

	switch msg.Typ {

	case pbs.WsMsgType_Online:
		return websocket.Inst().OnlineFromOtherPeer(msg)
	case pbs.WsMsgType_Offline:
		return websocket.Inst().OfflineFromOtherPeer(msg)
	case pbs.WsMsgType_ImmediateMsg:
		return websocket.Inst().PeerImmediateCryptoMsg(msg)
	case pbs.WsMsgType_PullUnread:
		return websocket.Inst().PeerUnreadMsg(msg)
	case pbs.WsMsgType_UnreadAck:
		return websocket.Inst().PeerUnreadAckMsg(msg)
	default:
		utils.LogInst().Warn().Msgf("unknown read in peer to peer msg type:[%d]", msg.Typ)
	}

	return nil
}
func (nt *NinjaStation) procContactMsg(msg *pbs2.ContactMsg) error {
	return nil
}
