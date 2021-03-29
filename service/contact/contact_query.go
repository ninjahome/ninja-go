package contact

import (
	"github.com/libp2p/go-libp2p-pubsub"
	"github.com/ninjahome/ninja-go/utils"
	"net/http"
)

func (s *Service) queryContact(w http.ResponseWriter, r *http.Request) {

}

func (s *Service) ContactQueryFromP2pNetwork(stop chan struct{}, r *pubsub.Subscription, w *pubsub.Topic) {
	s.contactQuery = w

	for true {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("contact query channel exit by outer controller")
			return
		default:
			msg, err := r.Next(s.ctx)
			if err != nil {
				utils.LogInst().Warn().Err(err).Send()
				return
			}

			if msg.ReceivedFrom.String() == s.id {
				continue
			}
		}
	}
}
