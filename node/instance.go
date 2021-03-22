package node

import (
	"context"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-pubsub"
	"sync"
)

type NinjaNetwork interface {
	Start() error
	ShutDown()
}

var _instance NinjaNetwork
var once sync.Once

func Inst() NinjaNetwork {
	once.Do(func() {
		_instance = newStation()
	})
	return _instance
}

type MessageChannel string

func (mc MessageChannel) String() string {
	return string(mc)
}

const (
	MSNotify  MessageChannel = "/0.1/Global/notify"
	MSNodeMsg MessageChannel = "/0.1/Global/NODE"
	MSDebug   MessageChannel = "/0.1/Global/TEST"
)

var SystemTopics = []MessageChannel{MSNotify, MSNodeMsg, MSDebug}

func notifyMsgValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	return pubsub.ValidationAccept
}

func nodeMsgValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	return pubsub.ValidationAccept
}
