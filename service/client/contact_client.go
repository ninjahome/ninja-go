package client

import (
	"bytes"
	"fmt"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/service/contact"
	"github.com/ninjahome/ninja-go/wallet"
	"google.golang.org/protobuf/proto"
	"io"
	"net/http"
)

type ContactCli struct {
	endpoint string
	key      *wallet.Key
}

func NewContactCli(addr string, key *wallet.Key) (*ContactCli, error) {
	if key == nil || !key.IsOpen() {
		return nil, fmt.Errorf("ivnalid key")
	}

	return &ContactCli{
		key:      key,
		endpoint: addr,
	}, nil
}

func (cc *ContactCli) sendRequest(msg *pbs.ContactMsg, path string) (ackMsg *pbs.ContactMsg, err error) {
	reqData, err := proto.Marshal(msg)
	if err != nil {
		return
	}

	r := bytes.NewReader(reqData)
	request, err := http.NewRequest("GET", cc.endpoint+path, r)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	ackMsg = &pbs.ContactMsg{}
	if err = proto.Unmarshal(body, ackMsg); err != nil {
		return
	}
	fmt.Println(ackMsg.String())

	return ackMsg, nil
}

func (cc *ContactCli) makeOpRequest(msg *pbs.ContactMsg) error {

	ackMsg, err := cc.sendRequest(msg, contact.PathOperateContact)
	if err != nil {
		return err
	}
	ack, ok := ackMsg.PayLoad.(*pbs.ContactMsg_OpAck)
	if !ok {
		return fmt.Errorf("invalid operation ack message")
	}

	if !ack.OpAck.Success {
		return fmt.Errorf(ack.OpAck.Msg)
	}

	return nil
}

func (cc *ContactCli) AddContact(cid, nickName, remarks string) error {

	item := &pbs.ContactItem{
		CID:      cid,
		NickName: nickName,
		Remarks:  remarks,
	}

	sig := cc.key.SignData(item.Data())
	request := &pbs.ContactMsg{
		Sig:     sig,
		Typ:     pbs.ContactMsgType_MTAddContact,
		From:    cc.key.Address.String(),
		PayLoad: &pbs.ContactMsg_AddOrUpdate{AddOrUpdate: item},
	}

	return cc.makeOpRequest(request)
}

func (cc *ContactCli) UpdateContact(cid, nickName, remarks string) error {
	return cc.AddContact(cid, nickName, remarks)
}

func (cc *ContactCli) DelContact(cid string) error {
	sig := cc.key.SignData([]byte(cid))
	request := &pbs.ContactMsg{
		Sig:     sig,
		Typ:     pbs.ContactMsgType_MTDeleteContact,
		From:    cc.key.Address.String(),
		PayLoad: &pbs.ContactMsg_DelC{DelC: cid},
	}
	return cc.makeOpRequest(request)
}

func (cc *ContactCli) SyncContact() ([]*pbs.ContactItem, error) {
	//TODO:: only self can query contact
	from := cc.key.Address.String()
	sig := cc.key.SignData([]byte(from))
	request := &pbs.ContactMsg{
		Sig:     sig,
		From:    from,
		PayLoad: &pbs.ContactMsg_Query{Query: from},
	}
	ackMsg, err := cc.sendRequest(request, contact.PathQueryContact)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if ackMsg.Typ == pbs.ContactMsgType_MTAck {
		ack, _ := ackMsg.PayLoad.(*pbs.ContactMsg_OpAck)
		return nil, fmt.Errorf(ack.OpAck.Msg)
	}

	ack, _ := ackMsg.PayLoad.(*pbs.ContactMsg_QueryResult)

	return ack.QueryResult.Contacts, nil
}
