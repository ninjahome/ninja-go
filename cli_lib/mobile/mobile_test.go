package chatLib

import (
	"fmt"
	"github.com/ninjahome/ninja-go/wallet"
	"testing"
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


func TestImportLicense(t *testing.T)  {
	key, err := wallet.LoadKeyFromJsonStr(__cipherTxt, "123")
	if err != nil {
		panic(err)
	}
	_inst.key = key
	s:=ImportLicense("2ngSfpc2FbVmi2PtuRt8ZqnQGYNcdTaFymncnLSSo95J9zSy5TX1TqTnK91WfkFAGfUSs6PUraxe6mCmNnQypmnb6YMCBnf1ynf9ngfQouAhxrinMoexKAXp2S3QeqyLXKmSmmcVqrSeZYXTAX3GgQek23RbcKZzKagsnz")
	fmt.Println(s)
}

func TestNewWallet(t *testing.T)  {
	s:=NewWallet("123")
	fmt.Println(s)
}