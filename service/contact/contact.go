package contact

import (
	"encoding/json"
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"net/http"
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

func (s *Service) save(from string, item *pbs.ContactItem) error {

	return s.operation(from, func(book *Book) error {
		(*book)[item.CID] = &BookItem{
			CID:      item.CID,
			NickName: item.NickName,
			Remarks:  item.Remarks,
			UpTime:   time.Now(),
		}
		return nil
	})
}

func (s *Service) del(from, cid string) error {

	return s.operation(from, func(book *Book) error {
		delete(*book, cid)
		return nil
	})
}

func (s *Service) queryContact(w http.ResponseWriter, r *http.Request) {
}
