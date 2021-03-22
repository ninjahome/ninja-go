package wallet

import (
	"crypto/ed25519"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/ninjahome/ninja-go/common"
	"github.com/pborman/uuid"
)

const (
	version   = 1
	PriKeyLen = 32
)

type FedKey bls.SecretKey

//SignData(account Key, mimeType string, data []byte) ([]byte, error)
//SignDataWithPassphrase(key Key, passphrase, mimeType string, data []byte) ([]byte, error)
//SignText(key Key, text []byte) ([]byte, error)
//SignTextWithPassphrase(account Key, passphrase string, hash []byte) ([]byte, error)
//SignTx(key Key, transaction *common.Transaction, chainID *big.Int) (*common.Transaction, error)
//SignTxWithPassphrase(account Key, passphrase string, transaction *common.Transaction, chainID *big.Int) (*common.Transaction, error)

type Key struct {
	ID         uuid.UUID
	Light      bool
	Address    common.Address
	PrivateKey *bls.SecretKey
}

type encryptedKeyJSON struct {
	Address string     `json:"address"`
	Crypto  CryptoJSON `json:"crypto"`
	ID      string     `json:"id"`
	Version int        `json:"version"`
}

type CryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams cipherParamsJSON       `json:"cipherParams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfParams"`
	MAC          string                 `json:"mac"`
}
type cipherParamsJSON struct {
	IV string `json:"iv"`
}

func NewKey() *Key {
	return NewLightKey(false)
}

func NewLightKey(light bool) *Key {
	sec := GenerateKey()
	id := uuid.NewRandom()
	key := &Key{
		Light:      light,
		ID:         id,
		Address:    common.PubKeyToAddr(sec.GetPublicKey()),
		PrivateKey: sec,
	}
	return key
}

func (k *Key) Encrypt(auth string) ([]byte, error) {
	if k.Light {
		return EncryptKey(k, auth, LightScryptN, LightScryptP)
	} else {
		return EncryptKey(k, auth, StandardScryptN, StandardScryptP)
	}
}

func (k *Key) isOpen() bool {
	return k.PrivateKey == nil
}

func (k *Key) close() {
	k.PrivateKey = nil
}

func (k *Key) CastP2pKey() (crypto.PrivKey, error) {
	pri := k.PrivateKey.Serialize()
	var edPri = ed25519.NewKeyFromSeed(pri)
	return crypto.UnmarshalEd25519PrivateKey(edPri[:])
}

func GenerateKey() *bls.SecretKey {
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	return &sec
}
