package worker

import (
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
)

type StreamGetter func(protocol.ID) (network.Stream, error)

type StreamWorker struct {
	ProtoID protocol.ID
	SGetter StreamGetter
}

func (sw StreamWorker) Stream() (network.Stream, error) {
	return sw.SGetter(sw.ProtoID)
}
