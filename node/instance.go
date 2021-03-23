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
	MSUserOnline    MessageChannel = "/0.1/Global/user/online"
	MSCryptoPeerMsg MessageChannel = "/0.1/Global/user/crypto_msg"
	MSDebug         MessageChannel = "/0.1/Global/TEST"

	THNOuterMsgReader = "outer message reader thread"
)

var SystemTopics = []MessageChannel{MSUserOnline, MSCryptoPeerMsg, MSDebug}

func userOnlineValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	return pubsub.ValidationAccept
}

func immediateCryptoMsgValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	//service.Inst().InUserTable()//TODO::maybe some easy way to tell the invalid message
	return pubsub.ValidationAccept
}
