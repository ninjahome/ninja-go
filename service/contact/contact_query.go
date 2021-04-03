package contact

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/ninjahome/ninja-go/node/worker"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	pbsS "github.com/ninjahome/ninja-go/pbs/stream"
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

func (s *Service) ContactQueryFromP2pNetwork(stream network.Stream) {
	defer stream.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	bts, err := rw.ReadBytes(SyncStreamDelim)
	if err != nil {
		return
	}
	streamMsg := &pbsS.StreamMsg{}
	if err := proto.Unmarshal(bts, streamMsg); err != nil {
		utils.LogInst().Warn().Msg("failed parse p2p message")
		return
	}
	body, ok := streamMsg.Payload.(*pbsS.StreamMsg_ContactSync)
	if !ok {
		return
	}

	//if body.ContactSync.SeqVer

	book, err := s.loadFromDb(body.ContactSync.UID)
	if err != nil {
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

	resp := &pbsS.StreamMsg{
		MTyp:    pbsS.StreamMType_MTContactAck,
		Payload: &pbsS.StreamMsg_ContactAck{ContactAck: &pbsS.ContactAck{Contacts: contactArr}},
	}

	data := resp.Data()
	data = append(data, SyncStreamDelim)
	if _, err := rw.Write(data); err != nil {
		return
	}
}

func (s *Service) ContactOperationFromP2pNetwork(w *worker.TopicWorker) {
	s.contactOperateWorker = w
	for {
		msg, err := w.ReadMsg()
		if err != nil {
			utils.LogInst().Warn().Msgf("contact operation thread exit:=>%s", err)
			return
		}

		if msg.ReceivedFrom.String() == s.id {
			continue
		}

		p2pMsg := &pbs.ContactMsg{}
		if err := proto.Unmarshal(msg.Data, p2pMsg); err != nil {
			utils.LogInst().Warn().Err(err).Send()
			continue
		}
		if err := s.procContactOperation(p2pMsg); err != nil {
			utils.LogInst().Warn().Err(err).Send()
		}
	}
}
