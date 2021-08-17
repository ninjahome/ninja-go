package chatLib

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ncom "github.com/ninjahome/ninja-go/common"
	"github.com/ninjahome/ninja-go/contract"
	"github.com/ninjahome/ninja-go/extra_node/webmsg"
	"github.com/ninjahome/ninja-go/extra_node/webserver"
	"github.com/ninjahome/ninja-go/service/client"
	"github.com/ninjahome/ninja-go/service/proxy"
	"github.com/ninjahome/ninja-go/service/proxy/httputil"
	"strconv"
	"strings"
)

const (
	infuraUrl   = "https://kovan.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contactAddr = "0x7B133a9BD10F7AE52fa9528b8Bc0f3c34612674c"
	tokenAddr   = "0x122938b76c071142ea6b39c34ffc38e5711cada1"
)

func GetExpireTime() int64 {

	var (
		c              *ethclient.Client
		err            error
		licenseContact *contract.NinjaChatLicense
		deadline       uint64
		userAddr       [32]byte
	)

	if !(_inst.key != nil && _inst.key.IsOpen()) {
		fmt.Println(errors.New("wallet not opened"))
		return 0
	}

	if userAddr, err = ncom.Naddr2ContractAddr(_inst.key.Address); err != nil {
		fmt.Println(0)
		return 0
	}

	if c, err = ethclient.Dial(infuraUrl); err != nil {
		fmt.Println(err)
		return 0
	}

	defer c.Close()

	licenseContact, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr), c)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	deadline, _, err = licenseContact.GetUserLicense(nil, userAddr)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int64(deadline)
}

type ChatLicenseContent struct {
	IssueAddr string `json:"issue_addr"`
	RandomId  string `json:"random_id"`
	NDays     int    `json:"n_days"`
}

type ChatLicense struct {
	Content   *ChatLicenseContent `json:"content"`
	Signature string              `json:"signature"`
}

func ImportLicense(licenseB58 string) string {
	if !(_inst.key != nil && _inst.key.IsOpen()) {
		fmt.Println(errors.New("wallet not opened"))
		return ""
	}
	var (
		userAddr  [32]byte
		issueAddr common.Address
		randId    [32]byte
		nDays     uint32
		j, sig    []byte
		ret       string
		code      int
		err       error
	)

	license := base58.Decode(licenseB58)

	if len(license) < 20+32+4+65 {
		return ""
	}

	n := copy(issueAddr[:], license)
	n += copy(randId[:], license[n:])
	nDays = binary.BigEndian.Uint32(license[n:])
	sig = license[n+4:]

	userAddr, err = ncom.Naddr2ContractAddr(_inst.key.Address)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	msg := &webmsg.LicenseBind{
		IssueAddr: issueAddr[:],
		UserAddr:  userAddr[:],
		NDays:     int32(nDays),
		RandomId:  randId[:],
		Signature: sig,
	}

	j, err = json.Marshal(*msg)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(j))

	srvs := client.DefaultBootWsService
	for i := 0; i < len(srvs); i++ {
		url := bootNode2HttpAddr(srvs[i])
		ret, code, err = httputil.NewHttpPost(nil, false, 2, 20).
			ProtectPost(url, string(j))
		if err != nil {
			fmt.Println(url, err)
			continue
		}

		if code != 200 {
			fmt.Println(url, "post failed",ret)
			continue
		}

		fmt.Println(url, "post success")

		return ret
	}

	return ""
}

func DecodeLicense(licenseB58 string) string {
	var (
		issueAddr common.Address
		randId    [32]byte
		nDays     uint32
		j, sig    []byte
	)
	license := base58.Decode(licenseB58)
	if len(license) < 20+32+4+65 {
		return ""
	}

	n := copy(issueAddr[:], license)
	n += copy(randId[:], license[n:])
	nDays = binary.BigEndian.Uint32(license[n:])
	sig = license[n+4:]

	clc := &ChatLicenseContent{
		IssueAddr: hex.EncodeToString(issueAddr[:]),
		RandomId:  hex.EncodeToString(randId[:]),
		NDays:     int(nDays),
	}

	cl := &ChatLicense{
		Content:   clc,
		Signature: hex.EncodeToString(sig),
	}

	j, _ = json.Marshal(*cl)

	return string(j)

}

func bootNode2HttpAddr(addr string) string {
	arr := strings.Split(addr, ":")

	return "http://" + arr[0] + ":" + strconv.Itoa(proxy.ProxyListenPort) + webserver.LicenseAddPath
}
