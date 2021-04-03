package worker

import (
	"context"
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/ninjahome/ninja-go/utils"
	"sync"
	"time"
)

const (
	TopicPeerNoCheckPeriods = 500 * time.Millisecond
)

type Worker interface {
	Stop()
}
type WorkGroup map[string]*TopicWorker

func (wg *WorkGroup) StartUp(ctx context.Context, ps *pubsub.PubSub, topics map[string]TopicReader, timeOut time.Duration) error {

	var grp sync.WaitGroup

	for topID, r := range topics {
		grp.Add(1)
		topic, err := ps.Join(topID)
		if err != nil {
			return err
		}

		w := newTopicWorker(ctx, topID, topic, r)

		if err := w.startWork(func() {
			delete(*wg, w.tid)
		}); err != nil {
			return err
		}

		go wg.checkPeerNo(&grp, w, timeOut)
		(*wg)[w.tid] = w
	}

	grp.Wait()
	return nil
}

func (wg *WorkGroup) checkPeerNo(grp *sync.WaitGroup, tw *TopicWorker, timeOut time.Duration) {
	defer grp.Done()

	checker := time.NewTicker(TopicPeerNoCheckPeriods)
	var tryTimes = 0
	tryTimeOut := time.NewTimer(timeOut)
	for {
		select {
		case <-checker.C:
			tryTimes++
			if len(tw.Pub.ListPeers()) > 0 {
				utils.LogInst().Info().Str("topic", tw.tid).
					Int("found success peers:", len(tw.Pub.ListPeers())).
					Send()
				return
			}
			utils.LogInst().Info().Str("topic", tw.tid).
				Int("tryTimes:", tryTimes).
				Send()
		case <-tryTimeOut.C:
			utils.LogInst().Error().Str("Topic Sync Result", "timeout")
			return
		}
	}
}

func (wg *WorkGroup) StopWork() {
	for _, t := range *wg {
		t.Stop()
	}
}
