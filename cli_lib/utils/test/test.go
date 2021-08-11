package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	chatLib "github.com/ninjahome/ninja-go/cli_lib/mobile"

)

type EthAccount struct {
	PrivKey *ecdsa.PrivateKey `json:"-"`
	PubKey  *ecdsa.PublicKey  `json:"-"`
	EAddr   common.Address    `json:"-"`
	SAddr   string            `json:"s_addr"`
}

func NewEthAccount() (acct *EthAccount, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	acct = &EthAccount{}
	acct.PrivKey = key
	acct.PubKey = (key.Public()).(*ecdsa.PublicKey)
	acct.EAddr = crypto.PubkeyToAddress(*acct.PubKey)
	acct.SAddr = acct.EAddr.String()

	return acct, nil
}
func main() {
	//ss := []string{"aaa", "bbb"}
	//
	//fmt.Println(utils.StrSlice2String(ss))
	//
	//buf := make([]byte, 8)
	//
	//var a int64 = 200
	//var b int64 = 1000000
	//
	//binary.BigEndian.PutUint64(buf, uint64(a))
	//
	//fmt.Println(hex.EncodeToString(buf))
	//binary.BigEndian.PutUint64(buf, uint64(b))
	//
	//fmt.Println(hex.EncodeToString(buf))


	key,_:=NewEthAccount()

	var rid [32]byte

	rand.Read(rid[:])

	l:=&chatLib.ChatLicense{}

	c:=&chatLib.ChatLicenseContent{}



	c.IssueAddr = hex.EncodeToString(key.EAddr[:])
	c.RandomId = hex.EncodeToString(rid[:])
	c.NDays = 120

	l.Content = c

	j,_:=json.Marshal(*c)

	fmt.Println("for signature data, not include []:")
	s:=fmt.Sprintf("[%s]",string(j))
	fmt.Println(s)

	hash:=crypto.Keccak256Hash(j)

	sig,_:=crypto.Sign(hash[:],key.PrivKey)

	l.Signature = hex.EncodeToString(sig)

	fmt.Println("signature:")

	j,_ = json.Marshal(*l)

	fmt.Println(string(j))

	fmt.Println("encode to base58:")
	fmt.Println(base58.Encode(j))

}
