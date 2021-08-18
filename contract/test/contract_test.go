package test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	ncom "github.com/ninjahome/ninja-go/common"
	flag "github.com/spf13/pflag"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/ninjahome/ninja-go/contract"
)

const (
	//contactAddr = "0x52996249f64d760ac02c6b82866d92b9e7d02f06"
	contactAddr               = "0x7B133a9BD10F7AE52fa9528b8Bc0f3c34612674c"
	tokenAddr                 = "0x122938b76c071142ea6b39c34ffc38e5711cada1"
	dialerAddr         string = "https://kovan.infura.io/v3/e01a4005bf8b42cca32875c2dc438dba"
	MetaMaskHashPrefix        = "\x19Ethereum Signed Message:\n32"
)

var _cipherTxt = `{
		"cipher": "aes-128-ctr",
		"ciphertext": "3f84fa9dcf9637ee531cd972fa7fcda976e1361f9cee6ee9f5222e2b3d59807d",
		"cipherparams": {
			"iv": "308f21f38eb5f3379168664f1b6a278e"
		},
		"kdf": "scrypt",
		"kdfparams": {
			"dklen": 32,
			"n": 262144,
			"p": 1,
			"r": 8,
			"salt": "e770746f6cf346bd1150b83d5c0a915a9bd6a5e4a6dff4506c6befa97fc5c3d6"
		},
		"mac": "c0881f242b339a17ab71eca7bb4556f1b7e07e8d3fefe16c2af5494841b241ab"
	}`

func GetPrivKey() *ecdsa.PrivateKey {
	j := &keystore.CryptoJSON{}
	json.Unmarshal([]byte(_cipherTxt), j)
	key, err := keystore.DecryptDataV3(*j, "123")
	if err != nil {
		fmt.Println("err is ", err)
		return nil
	}
	priKey := crypto.ToECDSAUnsafe(key)

	return priKey
}

func getRandomId() (rid [32]byte) {
	rand.Read(rid[:])
	return
}

func toPubKeyString(priv *ecdsa.PrivateKey) string {
	pubkey := priv.PublicKey
	return crypto.PubkeyToAddress(pubkey).String()
}

func TestAbi(t *testing.T) {
	parsed, err := abi.JSON(strings.NewReader(contract.NinjaChatLicenseABI))
	if err != nil {
		panic(err)
	}

	for k, _ := range parsed.Methods {
		fmt.Println(k)
	}

	_, exist := parsed.Methods["GenerateLicense"]
	if !exist {
		fmt.Println("method not found")
	} else {
		fmt.Println("method is find")
	}

}

func BalanceEth(balance float64) *big.Int {
	fbalance := new(big.Float)
	fbalance.SetFloat64(balance)
	v := new(big.Float).Mul(fbalance, big.NewFloat(math.Pow10(18)))

	vv := new(big.Int)
	v.Int(vv)

	return vv
}

func TestApprove(t *testing.T) {
	cli, err := ethclient.Dial(dialerAddr)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	var nToken *contract.NinjaToken
	nToken, err = contract.NewNinjaToken(common.HexToAddress(tokenAddr), cli)
	if err != nil {
		panic(err)
	}

	var nid *big.Int
	nid, err = cli.ChainID(context.TODO())
	if err != nil {
		panic(err)
	}

	var transactOpts *bind.TransactOpts

	transactOpts, err = bind.NewKeyedTransactorWithChainID(GetPrivKey(), nid)
	if err != nil {
		panic(err)
	}

	var tx *types.Transaction
	tx, err = nToken.Approve(transactOpts, common.HexToAddress(contactAddr), BalanceEth(100))
	if err != nil {
		panic(err)
	}

	fmt.Println(tx.Hash().String())

}

//run go test -v -run TestGenerateLicense
func TestGenerateLicense(t *testing.T) {

	cli, err := ethclient.Dial(dialerAddr)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	var ncl *contract.NinjaChatLicense
	ncl, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr), cli)
	if err != nil {
		panic(err)
	}

	var nid *big.Int
	nid, err = cli.ChainID(context.TODO())
	if err != nil {
		panic(err)
	}

	var transactOpts *bind.TransactOpts

	transactOpts, err = bind.NewKeyedTransactorWithChainID(GetPrivKey(), nid)
	if err != nil {
		panic(err)
	}

	rid := getRandomId()
	var tx *types.Transaction

	nDays := 30

	tx, err = ncl.GenerateLicense(transactOpts, rid, uint32(nDays))
	if err != nil {
		panic(err)
	}

	addr := toPubKeyString(GetPrivKey())

	fmt.Println("tx:", tx.Hash().String())
	fmt.Println("rid:", hex.EncodeToString(rid[:]))
	fmt.Println("ndays:", nDays)
	fmt.Println("issue:", addr)
	fmt.Println("contract:", contactAddr)

}

var (
	randId, sig, userAddr *string
	nDays                 *int
)

func init() {
	randId = flag.String("randomId", "", "random id")
	nDays = flag.Int("nDays", 0, "license days")
	sig = flag.String("sig", "", "signature")
	userAddr = flag.String("uAddr", "", "user address")

	flag.Parse()

}

//go test -v -run TestCreateLicense -randomId="xx" -nDays=5
func TestCreateLicense(t *testing.T) {

	var (
		abiUint32Type, _ = abi.NewType("uint32", "", nil)
		abiAddrType, _   = abi.NewType("address", "", nil)
		//abiStrType, _    = abi.NewType("string", "", nil)
		abiByte32Type, _ = abi.NewType("bytes32", "", nil)

		abiLicenseDataArgs = abi.Arguments{
			{Type: abiAddrType},
			{Type: abiAddrType},
			{Type: abiByte32Type},
			{Type: abiUint32Type},
		}

		//abiPrefixHashArgs = abi.Arguments{
		//	{Type: abiStrType},
		//	{Type: abiByte32Type},
		//}
	)

	*randId = "200b75c80f13ab1b403efa1bcf06a8543f478369e1c540d0ad844eb5d9cff059"

	if randId == nil {
		fmt.Println("please input random id")
		return
	}

	*nDays = 30

	if nDays == nil {
		fmt.Println("please input time interval")
		return
	}

	ca := common.HexToAddress(contactAddr)
	issue := common.HexToAddress(toPubKeyString(GetPrivKey()))
	ridb, _ := hex.DecodeString(*randId)

	fmt.Println("ridb....", hex.EncodeToString(ridb))

	var rid [32]byte
	copy(rid[:], ridb)

	licenseBytes, err := abiLicenseDataArgs.Pack(
		ca,
		issue,
		rid,
		uint32(*nDays))
	if err != nil {
		panic(err)
	}

	h := crypto.Keccak256Hash(licenseBytes)

	buf := make([]byte, 128)

	n := copy(buf, MetaMaskHashPrefix)

	n += copy(buf[n:], h[:])

	hash4sig := crypto.Keccak256(buf[:n])
	var signature []byte
	signature, err = crypto.Sign(hash4sig, GetPrivKey())
	if err != nil {
		panic(err)
	}

	fmt.Println("contract address:", contactAddr)
	fmt.Println("issue address:", issue.String())
	fmt.Println("random id:", hex.EncodeToString(rid[:]))
	fmt.Println("nDays:", *nDays)
	fmt.Println("sig:", hex.EncodeToString(signature))

	lsig := len(signature)
	if signature[lsig-1] <= 1 {
		signature[lsig-1] = signature[lsig-1] + 27
	}

	buflen := len(issue) + len(rid) + 4 + len(signature)

	buf = make([]byte, buflen)

	buint := make([]byte, 4)
	binary.BigEndian.PutUint32(buint, uint32(*nDays))

	n = copy(buf, issue[:])
	n += copy(buf[n:], rid[:])
	n += copy(buf[n:], buint)
	copy(buf[n:], signature)

	fmt.Println("License for load:", base58.Encode(buf))

}

//go test -v -run TestBindLicense -randomId="xx" -nDays=5 -sig="xxx" -uAddr="xx"
func TestBindLicense(t *testing.T) {

	*randId = "200b75c80f13ab1b403efa1bcf06a8543f478369e1c540d0ad844eb5d9cff059"

	if *randId == "" {
		fmt.Println("please input random id")
		return
	}

	*sig = "2e782782fc49f59d976d70bcb5db962ad81c63fa2554463566c0bcdf1f79259471b39fe1265f55ea07d6321494ac388cd496006ecc7f5e44881cda32f2e3da8001"

	if *sig == "" {
		fmt.Println("please input signature...")
		return
	}

	*userAddr = "93e82eb21e558bd0192c1866071bf0e2aff57e2bd6b1128ffefa52889a25a338c573e6b3b7fcc52b9b98fbf3eee39a34"

	*nDays = 30

	if *nDays == 0 {
		fmt.Println("please input time interval")
		return
	}

	cli, err := ethclient.Dial(dialerAddr)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	var ncl *contract.NinjaChatLicense
	ncl, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr), cli)
	if err != nil {
		panic(err)
	}

	var nid *big.Int
	nid, err = cli.ChainID(context.TODO())
	if err != nil {
		panic(err)
	}

	var transactOpts *bind.TransactOpts
	//transactOpts = bind.NewKeyedTransactor(GetPrivKey())
	transactOpts, err = bind.NewKeyedTransactorWithChainID(GetPrivKey(), nid)
	if err != nil {
		panic(err)
	}

	var (
		issue common.Address
		recv  [32]byte
		rid   [32]byte
		nD    uint32
		s     []byte
	)

	ridb, _ := hex.DecodeString(*randId)
	copy(rid[:], ridb)

	if *userAddr != "" {
		na, _ := ncom.HexToAddress(*userAddr)
		recv, _ = ncom.Naddr2ContractAddr(na)
		fmt.Println("recv:", hex.EncodeToString(recv[:]))
	} else {
		recv = rid
	}

	issue = common.HexToAddress(toPubKeyString(GetPrivKey()))

	nD = uint32(*nDays)

	s, _ = hex.DecodeString(*sig)

	var tx *types.Transaction
	tx, err = ncl.BindLicense(transactOpts, issue, recv, rid, nD, s)
	if err != nil {
		panic(err)
	}

	fmt.Println("tx:", tx.Hash().String())
	fmt.Println("issue:", issue.String())
	fmt.Println("recv:", hex.EncodeToString(recv[:]))
	fmt.Println("random id:", hex.EncodeToString(rid[:]))
	fmt.Println("nDays:", nD)
	fmt.Println("sig:", *sig)

}

func TestGetSetting(t *testing.T) {
	cli, err := ethclient.Dial(dialerAddr)
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	var ncl *contract.NinjaChatLicense
	ncl, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr), cli)
	if err != nil {
		panic(err)
	}

	var taddr, naddr common.Address

	taddr, naddr, err = ncl.GetSettings(nil)
	if err != nil {
		panic(err)
	}

	//fmt.Println(tx.Hash().String())

	fmt.Println("token addr:", taddr.String())
	fmt.Println("ninja addr:", naddr.String())

}

func TestUint32(t *testing.T) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, 30000)

	fmt.Println(hex.EncodeToString(buf))

	l := base58.Decode("2ngSfpc2FbVmi2PtuRt8ZqnQGYNcReyHaWuwXm7cHgyw9KgQHWhiPHAVeAEXEqLzBVvAqxdx18SEtUbXXkoDnqMGPGzKNwmj42WyS3t7Cb2XwX9uud6CvPxiufd6Hhfazsuzf8yS1aEsMK1oNymzXjkiXskVwTxBUcdSki")
	fmt.Println(hex.EncodeToString(l))

}
