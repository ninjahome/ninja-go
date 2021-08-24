package licenseLib

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ncom "github.com/ninjahome/ninja-go/common"
	"github.com/ninjahome/ninja-go/contract"
	"math/big"
)

type GenLicenseResult struct {
	NDays    int    `json:"n_days"`
	RandomId string `json:"random_id"`
	Tx       string `json:"tx"`
}

func GenLicense(nDays int) string {
	rid, txh, err := generateLicense(nDays)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	glr := &GenLicenseResult{
		NDays:    nDays,
		RandomId: "0x" + hex.EncodeToString(rid[:]),
		Tx:       txh.String(),
	}

	j, _ := json.Marshal(*glr)

	return string(j)
}

func decodeHex(hexstr string) ([]byte, error) {
	if hexstr == "" {
		return nil, errors.New("string is empty")
	}

	if len(hexstr) >= 2 {
		if hexstr[:2] == "0x" {
			hexstr = hexstr[2:]
		}
	}

	return hex.DecodeString(hexstr)
}

func CreateLicense(randomId string, nDays int) string {
	var (
		abiUint32Type, _ = abi.NewType("uint32", "", nil)
		abiAddrType, _   = abi.NewType("address", "", nil)
		abiByte32Type, _ = abi.NewType("bytes32", "", nil)

		abiLicenseDataArgs = abi.Arguments{
			{Type: abiAddrType},
			{Type: abiAddrType},
			{Type: abiByte32Type},
			{Type: abiUint32Type},
		}
	)

	if randomId == "" || nDays == 0 {
		fmt.Println("parameter error")
		return ""
	}

	var (
		ridb []byte
		err  error
	)

	if ridb, err = decodeHex(randomId); err != nil {
		fmt.Println(err)
		return ""
	}
	var rid [32]byte
	copy(rid[:], ridb)

	ca := common.HexToAddress(contactAddr)
	issue := _ethWallet.MainAddress()

	licenseBytes, err := abiLicenseDataArgs.Pack(
		ca,
		issue,
		rid,
		uint32(nDays))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	h := crypto.Keccak256Hash(licenseBytes)

	buf := make([]byte, 128)

	n := copy(buf, []byte("\x19Ethereum Signed Message:\n32"))

	n += copy(buf[n:], h[:])

	hash4sig := crypto.Keccak256(buf[:n])
	var signature []byte
	signature, err = crypto.Sign(hash4sig, _ethWallet.SignKey())
	if err != nil {
		panic(err)
	}

	fmt.Println("contract address:", contactAddr)
	fmt.Println("issue address:", issue.String())
	fmt.Println("random id:", hex.EncodeToString(rid[:]))
	fmt.Println("nDays:", nDays)
	fmt.Println("sig:", hex.EncodeToString(signature))

	lsig := len(signature)
	if signature[lsig-1] <= 1 {
		signature[lsig-1] = signature[lsig-1] + 27
	}

	buflen := len(issue) + len(rid) + 4 + len(signature)

	buf = make([]byte, buflen)

	buint := make([]byte, 4)
	binary.BigEndian.PutUint32(buint, uint32(nDays))

	n = copy(buf, issue[:])
	n += copy(buf[n:], rid[:])
	n += copy(buf[n:], buint)
	copy(buf[n:], signature)

	b58 := base58.Encode(buf)

	fmt.Println("License for load:", b58)

	return b58
}

type ChargeUserResult struct {
	Result bool   `json:"result"`
	Tx     string `json:"tx"`
}

func ChargeUserLicense(nDays int, userAddr string) string {
	cli, err := ethclient.Dial(infuraUrl)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer cli.Close()

	var ncl *contract.NinjaChatLicense
	ncl, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr), cli)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var nid *big.Int
	nid, err = cli.ChainID(context.TODO())
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var transactOpts *bind.TransactOpts
	transactOpts, err = bind.NewKeyedTransactorWithChainID(_ethWallet.SignKey(), nid)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	var ua ncom.Address
	if ua, err = ncom.HexToAddress(userAddr); err != nil {
		fmt.Println(err)
		return ""
	}
	var cua [32]byte
	if cua, err = ncom.Naddr2ContractAddr(ua); err != nil {
		fmt.Println(err)
		return ""
	}

	var tx *types.Transaction
	tx, err = ncl.ChargeUser(transactOpts, cua, uint32(nDays))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	cur := &ChargeUserResult{
		Result: true,
		Tx:     tx.Hash().String(),
	}

	fmt.Println(tx.Hash().String())

	j, _ := json.Marshal(*cur)

	return string(j)
}
