package node

import (
	"context"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-pubsub"
	"sync"
	"time"
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

const (
	P2pChanUserOnOffLine = "/0.1/Global/user/on_offline"
	P2pChanImmediateMsg  = "/0.1/Global/message/immediate"
	P2pChanUnreadMsg     = "/0.1/Global/message/unread"
	P2pChanContactOps    = "/0.1/Global/contact/operation"
	P2pChanContactQuery  = "/0.1/Global/contact/query"
	P2pChanDebug         = "/0.1/Global/TEST"
)

//TODO:: check the peer id's token balance
func userOnlineValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	return pubsub.ValidationAccept
}

//TODO:: check the peer id's token balance
func immediateCryptoMsgValidator(ctx context.Context, peer peer.ID, msg *pubsub.Message) pubsub.ValidationResult {
	//service.Inst().InUserTable()//TODO::maybe some easy way to tell the invalid message
	return pubsub.ValidationAccept
}

//TODO:: to be discussed
func initTopicValidators(ps *pubsub.PubSub) error {

	err := ps.RegisterTopicValidator(P2pChanUserOnOffLine,
		userOnlineValidator,
		pubsub.WithValidatorTimeout(250*time.Millisecond), //TODO::config
		pubsub.WithValidatorConcurrency(_nodeConfig.PsConf.MaxNotifyTopicThread))

	if err != nil {
		return err
	}

	err = ps.RegisterTopicValidator(P2pChanImmediateMsg,
		immediateCryptoMsgValidator,
		pubsub.WithValidatorConcurrency(_nodeConfig.PsConf.MaxNodeTopicThread))
	if err != nil {
		return err
	}

	return nil
}
