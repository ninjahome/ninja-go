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

func (tw *TopicWorker) ID() string {
	return tw.tid
}
func (tw *TopicWorker) startWork() error {
	sub, err := tw.Pub.Subscribe()
	if err != nil {
		return err
	}
	tw.Sub = sub

	t := thread.NewThreadWithName(tw.tid, func(_ chan struct{}) {
		utils.LogInst().Info().Msgf("......topic[%s] thread start success!", tw.tid)
		tw.tReader(tw)
		tw.Stop()
		utils.LogInst().Warn().Msgf("......topic[%s] thread exit!", tw.tid)
	})
	tw.thread = t
	t.Run()
	return nil
}

func (tw *TopicWorker) WriteData(data []byte) error {
	return tw.Pub.Publish(tw.ctx, data)
}

func (tw *TopicWorker) Stop() {
	tw.Pub.Close()
	tw.Sub.Cancel()
}

func (tw *TopicWorker) PeersOfTopic() []peer.ID {
	return tw.Pub.ListPeers()
}

func (tw *TopicWorker) BroadCast(data []byte) error {
	return tw.Pub.Publish(tw.ctx, data)
}

func (tw *TopicWorker) ReadMsg() (*pubsub.Message, error) {
	return tw.Sub.Next(tw.ctx)
}

func newTopicWorker(ctx context.Context, topID string, topic *pubsub.Topic, r TopicReader) *TopicWorker {
	w := &TopicWorker{
		ctx:     ctx,
		tid:     topID,
		Pub:     topic,
		tReader: r,
	}
	return w
}
