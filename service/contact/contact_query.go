package contact

import (
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p-pubsub"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/utils"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

func (s *Service) loadFromDb(from string) (Book, error) {
	s.contactLock.Lock()
	defer s.contactLock.Unlock()

	key := []byte(fmt.Sprintf(DBPatternHead, from))
	data, err := s.dataBase.Get(key, nil)
	if err != nil {
		return nil, err
	}

	book := make(Book)
	if err := json.Unmarshal(data, &book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) queryContact(w http.ResponseWriter, r *http.Request) {

	msg := &pbs.ContactMsg{}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write(pbs.ErrAck(err.Error()))
		return
	}

	if err = proto.Unmarshal(data, msg); err != nil {
		w.Write(pbs.ErrAck(err.Error()))
		return
	}

	query, ok := msg.PayLoad.(*pbs.ContactMsg_Query)
	if !ok {
		w.Write(pbs.ErrAck("invalid query message payload"))
		return
	}

	book, err := s.loadFromDb(query.Query)
	if err != nil {
		w.Write(pbs.ErrAck(err.Error()))
		return
	}

	contactArr := make([]*pbs.ContactItem, 0)
	for _, c := range book {
		contactArr = append(contactArr, &pbs.ContactItem{
			CID:      c.CID,
			NickName: c.NickName,
			Remarks:  c.Remarks,
		})
	}
	result := &pbs.ContactMsg{
		Typ:     pbs.ContactMsgType_MTContactList,
		PayLoad: &pbs.ContactMsg_QueryResult{QueryResult: &pbs.ContactList{Contacts: contactArr}},
	}

	w.Write(result.Data())
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
