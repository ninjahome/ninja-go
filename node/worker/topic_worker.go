package worker

import (
	"context"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/ninjahome/ninja-go/utils/thread"
)

type TopicReader func(*TopicWorker)

type TopicWorker struct {
	ctx     context.Context
	tid     string
	Pub     *pubsub.Topic
	tReader TopicReader
	Sub     *pubsub.Subscription
	thread  *thread.Thread
}

type WorkGroup map[string]*TopicWorker

func (tw *TopicWorker) StartWork() error {
	sub, err := tw.Pub.Subscribe()
	if err != nil {
		return err
	}
	tw.Sub = sub

	t := thread.NewThreadWithName(tw.tid, func(_ chan struct{}) {
		utils.LogInst().Info().Msgf("......subscribe topic[%s] thread success!", tw.tid)
		tw.tReader(tw)
		tw.StopWork()
	})
	tw.thread = t
	t.Run()
	return nil
}

func (tw *TopicWorker) WriteData(data []byte) error {
	return tw.Pub.Publish(tw.ctx, data)
}

func (tw *TopicWorker) StopWork() {
	tw.Pub.Close()
	tw.Sub.Cancel()
}

func (tw *TopicWorker) PeersOfTopic() []peer.ID {
	return tw.Pub.ListPeers()
}

func NewWorker(ctx context.Context, topID string, topic *pubsub.Topic, r TopicReader) *TopicWorker {
	return &TopicWorker{
		ctx:     ctx,
		tid:     topID,
		Pub:     topic,
		tReader: r,
	}
}
