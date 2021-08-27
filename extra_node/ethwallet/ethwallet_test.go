package ethwallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestNewWalletFromPrivateBytes(t *testing.T) {
	priv:="a23da23da23da23d37fa472ad5e9769515ea8d80ad136a2a725e5a3db717130f99536394a69c5301"
	auth:="1qaz2wsx"

	w,err:=NewWalletFromPrivateBytes(auth,priv[16:])
	if err!=nil{
		panic(err)
	}

	fmt.Println(w.String())
}

func TestPrintEthAddr(t *testing.T)  {
	fmt.Println(toPubKeyString(GetPrivKey()))
}

var _cipherTxt = `{
		"cipher":"aes-128-ctr",
		"ciphertext":"c79da2f1df233563b0c1c993a2a8209a4834d8f3452cadab0cbb398cde30973d",
		"cipherparams":{
			"iv":"8fe749bfec3d41d05b980b01255bf87c"
		},
		"kdf":"scrypt",
		"kdfparams":{
			"dklen":32,
			"n":262144,
			"p":1,
			"r":8,
			"salt":"9e2d547fd6eafe616d522c7b18960795847ba96c5d0ba3c2f1c4c093cb181fd3"
		},
		"mac":"65db7be298d916ca9ef27d9f5377b708059f76366c0a78311074b1187cc870a4"
	}
`

func GetPrivKey() *ecdsa.PrivateKey {
	j := &keystore.CryptoJSON{}
	json.Unmarshal([]byte(_cipherTxt), j)
	key, err := keystore.DecryptDataV3(*j, "1qaz2wsx")
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
