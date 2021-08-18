package chatLib

import (
	"fmt"
	"github.com/ninjahome/ninja-go/wallet"
	"testing"
	"time"
)

func TestDecodeBas58(t *testing.T) {
	s := DecodeLicense("2ngSfpc2FbVmi2PtuRt8ZqnQGYNcReyHaWuwXm7cHgyw9KgQHWhiPHAVeAEXEqLzBVvAqxdx18SEtUbXXkoDnqMGPGzKNwmj42WyS3t7Cb2XwX9uud6CvPxiufd6Hhfazsuzf8yS1aEsMK1oNymzXjkiXskVwTxBUcdSki")

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
	s := ImportLicense("2ngSfpc2FbVmi2PtuRt8ZqnQGYNcN1e1bYUCLsMTY9o4Do15yHEhZTY3hJqUxfKXepXoVHY4LWFwcL9Fe9JypVbZECycYg44VnozdHZxM8cR55BU5NV9TFHUW1uwrcy4dx88TJ4LtBRfivn3ocvM4XbwM8d9o8pu8T2hpn")
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
	ninstraddr := "a656a1a460f39b9c147ac5ac92c4829182c142e5f59b12b9f918feb30e5d07129a908ed2c1f22ab6b06649d49b740560"
	//ninstraddr:="93e82eb21e558bd0192c1866071bf0e2aff57e2bd6b1128ffefa52889a25a338c573e6b3b7fcc52b9b98fbf3eee39a34"

	fmt.Println(NinjaAddr2LicenseAddr(ninstraddr))

}

func TestGetExpireTime(t *testing.T) {
	key, err := wallet.LoadKeyFromJsonStr(__cipherTxt, "123")
	if err != nil {
		panic(err)
	}
	_inst.key = key

	expireTime := GetExpireTime()

	fmt.Println(time.Unix(expireTime, 0).String())

}
