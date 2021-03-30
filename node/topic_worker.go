package node

import (
	"context"
	"github.com/libp2p/go-libp2p-core/host"
	discovery "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/ninjahome/ninja-go/utils/thread"
)

type TopicReader func(stop chan struct{}, r *pubsub.Subscription, w *pubsub.Topic)

type TopicWorker struct {
	tid     string
	tWriter *pubsub.Topic
	tReader TopicReader
	sub     *pubsub.Subscription
	thread  *thread.Thread
}

type WorkGroup map[string]*TopicWorker

func (tw *TopicWorker) startWork() error {
	sub, err := tw.tWriter.Subscribe()
	if err != nil {
		return err
	}
	tw.sub = sub
	t := thread.NewThreadWithName(tw.tid, func(stop chan struct{}) {
		tw.tReader(stop, sub, tw.tWriter)
		tw.stopWork()
	})
	tw.thread = t
	t.Run()
	return nil
}

func (tw *TopicWorker) stopWork() {
	tw.thread.Stop()
	tw.tWriter.Close()
	tw.sub.Cancel()
}

func newWorkGroup(ctx context.Context, h host.Host) (WorkGroup, error) {
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

	topics := make(WorkGroup)
	for topID, r := range systemTopics {
		topic, err := ps.Join(topID)
		if err != nil {
			return nil, err
		}

		topics[topID] = &TopicWorker{
			tid:     topID,
			tWriter: topic,
			tReader: r,
		}
	}

	return topics, nil
}
