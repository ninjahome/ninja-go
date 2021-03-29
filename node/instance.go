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
	P2pChanUserOnOffLine MessageChannel = "/0.1/Global/user/on_offline"
	P2pChanImmediateMsg  MessageChannel = "/0.1/Global/message/immediate"
	P2pChanUnreadMsg     MessageChannel = "/0.1/Global/message/unread"
	P2pChanContactMsg    MessageChannel = "/0.1/Global/message/unread"
	P2pChanDebug         MessageChannel = "/0.1/Global/TEST"

	THNOuterMsgReader = "outer message reader thread"
)

var SystemTopics = []MessageChannel{
	P2pChanUserOnOffLine,
	P2pChanImmediateMsg,
	P2pChanUnreadMsg,
	P2pChanContactMsg,
	P2pChanDebug,
}

//TODO:: check the peer id's token balance
func userOnlineValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	return pubsub.ValidationAccept
}

//TODO:: check the peer id's token balance
func immediateCryptoMsgValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	//service.Inst().InUserTable()//TODO::maybe some easy way to tell the invalid message
	return pubsub.ValidationAccept
}
