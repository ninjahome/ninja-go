package chatLib

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/forgoer/openssl"
	"github.com/ninjahome/bls-wallet/bls"
	ncom "github.com/ninjahome/ninja-go/common"
	"github.com/ninjahome/ninja-go/contract"
	"github.com/ninjahome/ninja-go/service/client"
	"github.com/ninjahome/ninja-go/wallet"
	"testing"
	"time"
)

func TestDecodeBase58(t *testing.T) {
	s := DecodeLicense("2ds8ttQErMP2CaiakrW27hpJhaco6cnvAHrrcWL8Zx9k64KHntagAqH1avqGdDRNCvS2Hdr7fS5e3MG8kTF6GYKmbYtCLtujp6T4pq85JoBMKULkmN4dqKh6NV1dhVF3g4y3GpmqRYxYCFJdwNDhGhzBxppUzd1kY7BV1U")

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
	client.InitDefaultBootsNode()
	key, err := wallet.LoadKeyFromJsonStr(__cipherTxt, "123")
	if err != nil {
		panic(err)
	}
	_inst.key = key
	s := ImportLicense("2ngSfpc2FbVmi2PtuRt8ZqnQGYNcTAqF4pbgiKePX8vPJQWJQFM75CFkZHXN8euyLsBSk5ozeMg5mTtYZpETo5DDRqPjBRaUt6VfE5qRBnEayPLD4sZWDbYc3g4HspDVW5tfJH4S1MJGdZhb83pATvFF4SMjYihmtnycbC")
	fmt.Println(s)
}

func TestNewWallet(t *testing.T) {
	s := NewWallet("123")
	fmt.Println(s)
}

func TestRandomSrvList(t *testing.T) {

	client.InitDefaultBootsNode()

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

	InitEth()

	expireTime := GetExpireTime()

	fmt.Println(expireTime)

	fmt.Println(time.Unix(expireTime, 0).String())

}

func TestIsValidLicense(t *testing.T) {
	InitEth()
	l := "2ds8ttQErMP2CaiakrW27hpJhaco6cnvAHrrcWL8Zx9k64KHntagAqH1avqGdDRNCvS2Hdr7fS5e3MG8kTF6GYKmbYtCLtujp6T4pq85JoBMKULkmN4dqKh6NV1dhVF3g4y3GpmqRYxYCFJdwNDhGhzBxppUzd1kY7BV1U"
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
	client.InitDefaultBootsNode()
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

func TestGetBootsTrapList(t *testing.T) {
	lst, err := contract.GetBootsTrapList()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(lst); i++ {
		fmt.Println(lst[i].WSHostString())
	}
}

func TestGetEthConfig(t *testing.T) {
	ta, c, u, err := contract.GetEthConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(ta.String())
	fmt.Println(c.String())
	fmt.Println(string(u))

}

var _groupInfo = `
{"group_name":"嘿嘿",
"owner_id":"a656a1a460f39b9c147ac5ac92c4829182c142e5f59b12b9f918feb30e5d07129a908ed2c1f22ab6b06649d49b740560",
"group_id":"HMn9j\/KRjWZvzTzF7ISGhek+p1sZf1e8Ews\/bIDlzok=",
"ban_talking":false,"nick_name":["🎈 8","xxf","豆豆","晓芙"],"member_id":["93e82eb21e558bd0192c1866071bf0e2aff57e2bd6b1128ffefa52889a25a338c573e6b3b7fcc52b9b98fbf3eee39a34","8c8aae805033609b279ab903e6f2dcaddb153ca64fb2c455166f01d08f65c7556a0b6e68f55c778023d81f8e79ff8b0d","a3eb5e8f0fba490cb65296e7a863fd85cdf6ab846aca283cb6e120ab4ca4ba09db1a424e3b5c1ea11d69f75bdc895067","a656a1a460f39b9c147ac5ac92c4829182c142e5f59b12b9f918feb30e5d07129a908ed2c1f22ab6b06649d49b740560"]}
`

func TestDecodeGroupInfo(t *testing.T) {
	gi := &GroupInfo{}
	if err := json.Unmarshal([]byte(_groupInfo), gi); err != nil {
		panic(err)
	}
	if j, err := json.MarshalIndent(*gi, "", "\t"); err != nil {
		panic(err)
	} else {
		fmt.Println(string(j))
	}

}

func TestEncrypt(t *testing.T) {
	src := "hello,world"

	key := []byte("12345123451234512345123451234512")

	dst, err := openssl.AesECBEncrypt([]byte(src), key, openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}

	src1 := dst

	fmt.Println(base64.StdEncoding.EncodeToString(dst))

	dst, err = openssl.AesECBDecrypt(src1, key, openssl.PKCS7_PADDING)

	fmt.Println(string(dst))

	fmt.Println(base64.StdEncoding.EncodeToString(src1))

}
