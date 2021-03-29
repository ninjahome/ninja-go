package contact

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/libp2p/go-libp2p-pubsub"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/utils"
	"time"
)

type BookItem struct {
	CID      string
	NickName string
	Remarks  string
	UpTime   time.Time
}
type Book map[string]*BookItem
type BookOpFunc func(book *Book) error

func (s *Service) operation(from string, op BookOpFunc) error {
	s.contactLock.Lock()
	defer s.contactLock.Unlock()

	key := []byte(fmt.Sprintf(DBPatternHead, from))
	data, err := s.dataBase.Get(key, nil)
	if err != nil {
		return err
	}

	book := make(Book)
	if err := json.Unmarshal(data, &book); err != nil {
		return err
	}
	if err := op(&book); err != nil {
		return err
	}

	newData, err := json.Marshal(book)
	if err != nil {
		return err
	}
	return s.dataBase.Put(key, newData, nil)
}

func (s *Service) saveContact(msg *pbs.ContactMsg) error {
	body, ok := msg.PayLoad.(*pbs.ContactMsg_AddOrUpdate)
	if !ok {
		return ErBodyCast
	}

	item := body.AddOrUpdate
	if !msg.Verify(item.Data()) {
		return ErVerifyFailed
	}

	return s.operation(msg.From, func(book *Book) error {
		(*book)[item.CID] = &BookItem{
			CID:      item.CID,
			NickName: item.NickName,
			Remarks:  item.Remarks,
			UpTime:   time.Now(),
		}
		return nil
	})
}

func (s *Service) delContact(msg *pbs.ContactMsg) error {
	body, ok := msg.PayLoad.(*pbs.ContactMsg_DelC)
	if !ok {
		return ErBodyCast
	}

	cid := body.DelC
	if !msg.Verify([]byte(cid)) {
		return ErVerifyFailed
	}

	return s.operation(msg.From, func(book *Book) error {
		delete(*book, cid)
		return nil
	})
}

func (s *Service) ContactOperationToP2pNetwork(stop chan struct{}, r *pubsub.Subscription, w *pubsub.Topic) {
	s.contactOpWriter = w

	for true {
		select {
		case <-stop:
			utils.LogInst().Warn().Msg("contact operation channel exit by outer controller")
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
			p2pMsg := &pbs.ContactMsg{}
			if err := proto.UnmarshalMerge(msg.Data, p2pMsg); err != nil {
				utils.LogInst().Warn().Err(err).Send()
				continue
			}
			if err := s.procContactOperation(p2pMsg); err != nil {
				utils.LogInst().Warn().Err(err).Send()
			}
		}
	}
}
