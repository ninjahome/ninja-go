package client

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	pbs "github.com/ninjahome/ninja-go/pbs/contact"
	"github.com/ninjahome/ninja-go/wallet"
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

func (cc *ContactCli) makeOpRequest(msg *pbs.ContactMsg) error {

	reqData, err := proto.Marshal(msg)
	if err != nil {
		return err
	}

	r := bytes.NewReader(reqData)
	request, err := http.NewRequest("GET", cc.endpoint, r)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	ackMsg := &pbs.ContactMsg{}

	if err := proto.Unmarshal(body, ackMsg); err != nil {
		return err
	}
	ack, ok := ackMsg.PayLoad.(*pbs.ContactMsg_OpAck)
	if !ok {
		return fmt.Errorf("invalid operation ack message")
	}

	if ack.OpAck.Success {
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

	itemData, err := proto.Marshal(item)
	if err != nil {
		return err
	}
	sig := cc.key.SignData(itemData)

	request := &pbs.ContactMsg{
		Sig:     sig,
		PayLoad: &pbs.ContactMsg_AddC{AddC: item},
	}

	return cc.makeOpRequest(request)
}

func (cc *ContactCli) UpdateContact(cid, nickName, remarks string) error {
	return cc.AddContact(cid, nickName, remarks)
}

func (cc *ContactCli) DelContact(cid, nickName, remarks string) error {
	queryID := cc.key.Address
	sig := cc.key.SignData(queryID[:])
	request := &pbs.ContactMsg{
		Sig:     sig,
		PayLoad: &pbs.ContactMsg_DelC{DelC: queryID.String()},
	}
	return cc.makeOpRequest(request)
}
