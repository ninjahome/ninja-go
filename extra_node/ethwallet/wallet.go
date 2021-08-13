package ethwallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
)

const (
	WalletVersion = 1
)

type Wallet interface {
	SignKey() *ecdsa.PrivateKey
	MainAddress() common.Address
	SignJson(v interface{}) ([]byte, error)
	Sign(v []byte) ([]byte, error)
	VerifySig(message, signature []byte) bool
	Open(auth string) error
	IsOpen() bool
	SaveToPath(wPath string) error
	String() string
	Close()
	ExportEth(auth, eAuth, path string) error
}

type WalletKey struct {
	MainPriKey *ecdsa.PrivateKey
}

type PWallet struct {
	Version  int                 `json:"version"`
	MainAddr common.Address      `json:"mainAddress"`
	Crypto   keystore.CryptoJSON `json:"crypto"`
	key      *WalletKey          `json:"-"`
}

func NewWallet(auth string) (Wallet, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	keyBytes := math.PaddedBigBytes(privateKeyECDSA.D, 32)
	cryptoStruct, err := keystore.EncryptDataV3(keyBytes, []byte(auth), keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return nil, err
	}

	obj := &PWallet{
		Version:  WalletVersion,
		MainAddr: crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		Crypto:   cryptoStruct,
		key: &WalletKey{
			MainPriKey: privateKeyECDSA,
		},
	}
	return obj, nil
}

func NewWalletFromPrivateBytes(auth string, hexpriv string) (Wallet, error) {
	var privateKeyECDSA *ecdsa.PrivateKey
	var err error
	if privateKeyECDSA, err = crypto.HexToECDSA(hexpriv); err != nil {
		return nil, err
	}

	keyBytes := math.PaddedBigBytes(privateKeyECDSA.D, 32)
	cryptoStruct, err := keystore.EncryptDataV3(keyBytes, []byte(auth), keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return nil, err
	}

	obj := &PWallet{
		Version:  WalletVersion,
		MainAddr: crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		Crypto:   cryptoStruct,
		key: &WalletKey{
			MainPriKey: privateKeyECDSA,
		},
	}
	return obj, nil

}

func LoadWallet(wPath string) (Wallet, error) {
	jsonStr, err := ioutil.ReadFile(wPath)
	if err != nil {
		return nil, err
	}

	w := new(PWallet)
	if err := json.Unmarshal(jsonStr, w); err != nil {
		return nil, err
	}
	return w, nil
}

func LoadWalletByData(jsonStr string) (Wallet, error) {
	w := new(PWallet)
	if err := json.Unmarshal([]byte(jsonStr), w); err != nil {
		return nil, err
	}
	return w, nil
}

func VerifyJsonSig(mainAddr common.Address, sig []byte, v interface{}) bool {
	return mainAddr == RecoverJson(sig, v)
}

func VerifyAbiSig(mainAddr common.Address, sig []byte, msg []byte) bool {
	signer, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	return mainAddr == crypto.PubkeyToAddress(*signer)
}
func RecoverJson(sig []byte, v interface{}) common.Address {
	data, err := json.Marshal(v)
	if err != nil {
		return common.Address{}
	}
	hash := crypto.Keccak256(data)
	signer, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return common.Address{}
	}
	address := crypto.PubkeyToAddress(*signer)
	return address
}
