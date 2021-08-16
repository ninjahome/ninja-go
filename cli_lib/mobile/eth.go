package chatLib

import (
	"encoding/base64"
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
	contactAddr = "0x52996249f64d760ac02c6b82866d92b9e7d02f06"
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

	license := base58.Decode(licenseB58)

	cl := &ChatLicense{}
	if err := json.Unmarshal(license, cl); err != nil {
		fmt.Println(err)
		return ""
	}

	var (
		userAddr                  [32]byte
		issueAddr, randId, sig, j []byte
		ret                       string
		code                      int
		err                       error
	)

	//issueAddr:=common.HexToAddress(cl.Content.IssueAddr)
	//issueAddr, err = hex.DecodeString(cl.Content.IssueAddr)
	issueAddr,err = base64.StdEncoding.DecodeString(cl.Content.IssueAddr)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	//randId, err = hex.DecodeString(cl.Content.RandomId)
	randId, err = base64.StdEncoding.DecodeString(cl.Content.RandomId)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	//sig, err = hex.DecodeString(cl.Signature)
	sig, err = base64.StdEncoding.DecodeString(cl.Signature)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	userAddr, err = ncom.Naddr2ContractAddr(_inst.key.Address)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	msg := &webmsg.LicenseBind{
		IssueAddr: issueAddr,
		UserAddr:  userAddr[:],
		NDays:     int32(cl.Content.NDays),
		RandomId:  randId,
		Signature: sig,
	}

	j, err = json.Marshal(*msg)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	srvs := client.DefaultBootWsService
	for i := 0; i < len(srvs); i++ {
		url := bootNode2HttpAddr(srvs[i])
		ret, code, err = httputil.NewHttpPost(nil, false, 2, 2).
			ProtectPost(url, string(j))
		if err != nil {
			fmt.Println(url, err)
			continue
		}

		if code != 200 {
			fmt.Println(url, "post failed")
			continue
		}

		fmt.Println(url, "post success")

		return ret
	}

	return ""
}

func bootNode2HttpAddr(addr string) string {
	arr := strings.Split(addr, ":")

	return "http://" + arr[0] + ":" + strconv.Itoa(proxy.ProxyListenPort) + webserver.LicenseAddPath
}
