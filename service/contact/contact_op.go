package contact

import (
	"encoding/json"
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/syndtr/goleveldb/leveldb"
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
	if err != nil && err != leveldb.ErrNotFound {
		utils.LogInst().Warn().Str("Load book db", err.Error()).Send()
		return err
	}

	book := make(Book)
	if len(data) > 0 {
		if err := json.Unmarshal(data, &book); err != nil {
			utils.LogInst().Warn().Str("unmarshal book err", err.Error()).Send()
			return err
		}
	}

	if err := op(&book); err != nil {
		utils.LogInst().Warn().Str("operate book err", err.Error()).Send()
		return err
	}

	newData, err := json.Marshal(book)
	if err != nil {
		utils.LogInst().Warn().Str("Marshal book err", err.Error()).Send()
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
	utils.LogInst().Debug().
		Str("From", msg.From).
		Str("add contact cid", item.String()).
		Send()

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
	utils.LogInst().Debug().Str("From", msg.From).Str("remove contact cid", cid).Send()
	return s.operation(msg.From, func(book *Book) error {
		delete(*book, cid)
		return nil
	})
}
