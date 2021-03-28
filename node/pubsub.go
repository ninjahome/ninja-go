package node

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	coreDisc "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-discovery"
	"github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-pubsub"
	pbs "github.com/ninjahome/ninja-go/pbs/service"
	"github.com/ninjahome/ninja-go/utils"
	"sync"
	"time"
)

type PubSub struct {
	ctx    context.Context
	lock   sync.RWMutex
	topics map[MessageChannel]*pubsub.Topic
	dht    *dht.IpfsDHT
	disc   coreDisc.Discovery
	pubSub *pubsub.PubSub
}

func newPubSub(ctx context.Context, h host.Host) (*PubSub, error) {
	dhtOpts, err := _nodeConfig.dhtOpts()

	kademliaDHT, err := dht.New(ctx, h, dhtOpts...)
	if err != nil {
		return nil, err
	}

	disc := discovery.NewRoutingDiscovery(kademliaDHT)

	psOption := _nodeConfig.pubSubOpts(disc)

	ps, err := pubsub.NewGossipSub(ctx, h, psOption...)
	if err != nil {
		return nil, err
	}

	if err := kademliaDHT.Bootstrap(ctx); err != nil {
		return nil, err
	}

	if err := initTopicValidators(ps); err != nil {
		return nil, err
	}

	topics := make(map[MessageChannel]*pubsub.Topic)
	for _, topID := range SystemTopics {
		topic, err := ps.Join(string(topID))
		if err != nil {
			return nil, err
		}
		topics[topID] = topic
	}

	return &PubSub{
		ctx:    ctx,
		dht:    kademliaDHT,
		pubSub: ps,
		disc:   disc,
		topics: topics,
	}, nil
}

func initTopicValidators(ps *pubsub.PubSub) error {

	err := ps.RegisterTopicValidator(P2pChanUserOnline.String(),
		userOnlineValidator,
		pubsub.WithValidatorTimeout(250*time.Millisecond), //TODO::config
		pubsub.WithValidatorConcurrency(_nodeConfig.PsConf.MaxNotifyTopicThread))

	if err != nil {
		return err
	}

	err = ps.RegisterTopicValidator(P2pChanImmediateMsg.String(),
		immediateCryptoMsgValidator,
		pubsub.WithValidatorConcurrency(_nodeConfig.PsConf.MaxNodeTopicThread))
	if err != nil {
		return err
	}

	return nil
}

func (s *PubSub) removeTopic(id MessageChannel) {

	s.lock.Lock()
	defer s.lock.Unlock()

	t, ok := s.topics[id]
	if !ok {
		return
	}
	if err := t.Close(); err != nil {
		utils.LogInst().Warn().Msgf("topic [%s] close failed", id)
	}
	delete(s.topics, id)
	utils.LogInst().Warn().Msgf("remove topic [%s] from system", id)
}

func (s *PubSub) readingMessage(stop chan struct{}, sub *pubsub.Subscription, id MessageChannel, queue chan *pbs.WsMsg) {
	utils.LogInst().Info().Msgf("[pubSub] start reading message for topic[%s]", id)
	defer s.removeTopic(id)

	for {
		msg, err := sub.Next(s.ctx)
		if err != nil {
			utils.LogInst().Warn().Err(err).Send()
			return
		}

		select {
		case <-stop:
			utils.LogInst().Warn().Msg("topic reading thread exit by outer controller")
			return
		default:
			//if msg.ReceivedFrom == cr.self {
			//	continue
			//}
			p2pMsg := &pbs.WsMsg{}
			if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
				utils.LogInst().Warn().Msg("failed parse p2p message")
				continue
			}
			queue <- p2pMsg
		}
	}
}

func (s *PubSub) SendMsg(topic MessageChannel, msgData []byte) error {
	topics := s.topics
	s.lock.RLock()
	defer s.lock.RUnlock()

	t, ok := topics[topic]
	if !ok {
		return fmt.Errorf("no such topic")
	}

	if err := t.Publish(s.ctx, msgData); err != nil {
		return err
	}
	return nil
}

func (s *PubSub) PeersOfTopic(topic string) []peer.ID {
	topics := s.topics
	s.lock.RLock()
	defer s.lock.RUnlock()
	t, ok := topics[MessageChannel(topic)]
	if !ok {
		return nil
	}
	return t.ListPeers()
}
