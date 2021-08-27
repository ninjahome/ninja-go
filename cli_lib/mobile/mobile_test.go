package chatLib

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/ninjahome/bls-wallet/bls"
	ncom "github.com/ninjahome/ninja-go/common"
	"github.com/ninjahome/ninja-go/contract"
	"github.com/ninjahome/ninja-go/wallet"
	"testing"
	"time"
)

func TestDecodeBase58(t *testing.T) {
	s := DecodeLicense("22FVcbpbffH7R7xzJgZBL1imo4WXphCudeCSFM48voSgwCD9pXWu4qQcSkxeoTH5JGegTQUED4bT7nmwGRHuwyckGZKtbVECtp3sCkC3ECCuwGKy1DsLJvHfLdVwsmS7okLmbCE3aoThbGzPE4BW2dkiorb42rbv3Nh8ca")

	fmt.Println(s)
}

var __cipherTxt = `
{
	"address": "8cdf71ed43abaf6c969f6c5472b9dffe970330cd680311ad90ebbddec13db449db9108682d9c1d46009472d7d8a13efc",
	"crypto": {
		"cipher": "aes-128-ctr",
		"ciphertext": "bae096af6b412b43f62ad7613f05ca87918a94772f57d99bbe9f50558e3f8165",
		"cipherParams": {
			"iv": "a55997b58a88a2e94c98509b17e07a45"
		},
		"kdf": "scrypt",
		"kdfParams": {
			"dklen": 32,
			"n": 262144,
			"p": 1,
			"r": 8,
			"salt": "d39af5f3a26d2b5bbecfc5b340bb2bae9b4cff950886f792d4f05e811d009db7"
		},
		"mac": "89bb0af8d339277a59252324246f348347308dec9e85a7decccea9c49a27f32d"
	},
	"id": "04cabf04-dacd-4631-9866-f569044cf299",
	"version": 1
}
`

func TestImportLicense(t *testing.T) {
	key, err := wallet.LoadKeyFromJsonStr(__cipherTxt, "123")
	if err != nil {
		panic(err)
	}
	_inst.key = key
	s := ImportLicense("2ngSfpc2FbVmi2PtuRt8ZqnQGYNcQMoF8X65diPmpZfRzLKmtAjPAyN6VW3tEi7wWjD1QSHXERKxh2k7wjzQLGFxWetZBdhLTrMwJTiudyXYgA5UF3YJb2DDUuz7EELeKnogJhCxsJSm1mvLpTuR7Pw6jK6nt3o5kjWqL7")
	fmt.Println(s)
}

func TestNewWallet(t *testing.T) {
	s := NewWallet("123")
	fmt.Println(s)
}

func TestRandomSrvList(t *testing.T) {
	srvs := RandomSrvList()

	for i := 0; i < len(srvs); i++ {
		fmt.Println(srvs[i])
	}
}

func TestNinjaAddr2LicenseAddr(t *testing.T) {
	ninstraddr := "b1d6e0c4a0c3a0c74d2c9e13cfa8bc5cfb349c8a5f4746a2c143483a3f352544c72bc87871a7696bd135f42fc076e5aa"
	//ninstraddr:="93e82eb21e558bd0192c1866071bf0e2aff57e2bd6b1128ffefa52889a25a338c573e6b3b7fcc52b9b98fbf3eee39a34"

	fmt.Println(NinjaAddr2LicenseAddr(ninstraddr))

}

func TestGetExpireTime(t *testing.T) {
	key, err := wallet.LoadKeyFromJsonStr(__cipherTxt, "123")
	if err != nil {
		panic(err)
	}
	_inst.key = key

	addr, _ := ncom.Naddr2ContractAddr(key.Address)

	fmt.Println(hex.EncodeToString(addr[:]))

	expireTime := GetExpireTime()

	fmt.Println(expireTime)

	fmt.Println(time.Unix(expireTime, 0).String())

}

func TestIsValidLicense(t *testing.T) {
	l := "2ngSfpc2FbVmi2PtuRt8ZqnQGYNc6Y9ys6EwA6xYLzWZMPQeyt6dSBJbUdnbFg5ETMUr8gS9Dc3o4noPojSrz3VsJQRTGTzdt9JsNzzEmgJSk4Km93B9paNjzcjpnkB3f9XTAZS4T7aMh8xoCxejzdcxp9C5ybFr6RXwNs"
	s := IsValidLicense(l)
	if s == ValidTrue {
		fmt.Println("license have been used")
	} else if s == ValidFalse {
		fmt.Println("license is valid")
	} else {
		fmt.Println(s)
	}

}
func TestTransferLicense(t *testing.T) {
	key, err := wallet.LoadKeyFromJsonStr(__cipherTxt, "123")
	if err != nil {
		panic(err)
	}
	_inst.key = key
	ret := TransferLicense("0xb1d6e0c4a0c3a0c74d2c9e13cfa8bc5cfb349c8a5f4746a2c143483a3f352544c72bc87871a7696bd135f42fc076e5aa", 5)
	fmt.Println(ret)
}

func TestSign(t *testing.T) {
	froms := "0x8cdf71ed43abaf6c969f6c5472b9dffe970330cd680311ad90ebbddec13db449db9108682d9c1d46009472d7d8a13efc"
	tos := "0xb1d6e0c4a0c3a0c74d2c9e13cfa8bc5cfb349c8a5f4746a2c143483a3f352544c72bc87871a7696bd135f42fc076e5aa"
	nDays := 5

	from, _ := ncom.HexToAddress(froms)
	to, _ := ncom.HexToAddress(tos)

	buf := make([]byte, 1024)
	n := copy(buf, from[:])
	n += copy(buf[n:], to[:])
	binary.BigEndian.PutUint32(buf[n:], uint32(nDays))

	n += 4

	key, err := wallet.LoadKeyFromJsonStr(__cipherTxt, "123")
	if err != nil {
		panic(err)
	}
	_inst.key = key

	fmt.Println("address:", _inst.key.Address.String())

	fmt.Println("msg:", hex.EncodeToString(buf[:n]))

	sig := _inst.key.SignData(buf[:n])

	fmt.Println("sig:", hex.EncodeToString(sig))
}

func TestVerifySign(t *testing.T) {
	froms := "0x8cdf71ed43abaf6c969f6c5472b9dffe970330cd680311ad90ebbddec13db449db9108682d9c1d46009472d7d8a13efc"
	tos := "0xb1d6e0c4a0c3a0c74d2c9e13cfa8bc5cfb349c8a5f4746a2c143483a3f352544c72bc87871a7696bd135f42fc076e5aa"
	nDays := 5

	sigs := "8c51a81dc150a9c3279c174da2e229fca0e23f70e2934a8d34f2cf480301f0d551b91226e83fb943bc6e7d104f3c170d09f88deefce2a4cbdbd0e6ac6ee514cc083ffb346b9836bab7d16440054a6eb5c795a9d7bc1cec36e815b5b3504e6c51"

	from, _ := ncom.HexToAddress(froms)
	to, _ := ncom.HexToAddress(tos)

	sigb, _ := hex.DecodeString(sigs)

	buf := make([]byte, 1024)
	n := copy(buf, from[:])
	n += copy(buf[n:], to[:])
	binary.BigEndian.PutUint32(buf[n:], uint32(nDays))

	n += 4

	sig := &bls.Sign{}

	if err := sig.Deserialize(sigb); err != nil {
		fmt.Println(err)
		return
	}
	p := &bls.PublicKey{}
	if err := p.Deserialize(from[:]); err != nil {
		fmt.Println(err)
		return
	}

	b := sig.VerifyByte(p, buf[:n])

	fmt.Println(b)

}

func TestGetBootsTrapList(t *testing.T)  {
	lst,err:=contract.GetBootsTrapList()
	if err!=nil{
		panic(err)
	}

	for i:=0;i<len(lst);i++{
		fmt.Println(lst[i].WSHostString())
	}
}

func TestGetEthConfig(t *testing.T)  {
	ta,c,u,err:=contract.GetEthConfig()
	if err!=nil{
		panic(err)
	}

	fmt.Println(ta.String())
	fmt.Println(c.String())
	fmt.Println(string(u))

}