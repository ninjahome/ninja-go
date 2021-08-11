package chatLib

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	ncom "github.com/ninjahome/ninja-go/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/ninja-go/contract"
)

const(
	infuraUrl= "https://ropsten.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contactAddr = "0xf08192dcfc78f5e5ea74c83147378e1f4b8fe560"
	tokenAddr = "0x122938b76c071142ea6b39c34ffc38e5711cada1"
)

func GetExpireTime(addr string) int64  {

	var (
		c *ethclient.Client
		err error
		licenseContact *contract.NinjaChatLicense
		deadline uint64
		userAddr [32]byte
		naddr ncom.Address
	)

	naddr,err=ncom.HexToAddress(addr)
	if err!=nil{
		fmt.Println(err)
		return 0
	}

	h:=sha256.New()
	_,err = h.Write(naddr[:])
	if err!=nil{
		fmt.Println(err)
		return 0
	}

	h32:=h.Sum(nil)
	copy(userAddr[:],h32)

	if c,err=ethclient.Dial(infuraUrl);err!=nil{
		fmt.Println(err)
		return 0
	}

	defer c.Close()

	licenseContact, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr),c)
	if err!=nil{
		fmt.Println(err)
		return 0
	}

	deadline,_,err=licenseContact.GetUserLicense(nil,userAddr)
	if err!=nil{
		fmt.Println(err)
		return 0
	}

	return int64(deadline)
}

type ChatLicenseContent struct {
	IssueAddr string	`json:"issue_addr"`
	RandomId  string    `json:"random_id"`
	NDays     int       `json:"n_days"`
}

type ChatLicense struct {
	Content *ChatLicenseContent `json:"content"`
	Signature string 			`json:"signature"`
}


func ImportLicense(licenseB58 string, selfAddr string) error  {
	license:=base58.Decode(licenseB58)

	cl:=&ChatLicense{}
	if err:=json.Unmarshal(license,cl);err!=nil{
		return err
	}

	if b:=IsValidNinjaAddr(selfAddr);!b{
		return errors.New("address not correct")
	}

	//todo...

	return nil
}
