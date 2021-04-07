package wallet

import (
	"crypto/ed25519"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/ninjahome/bls-wallet/bls"
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
//SignTextWithPassphrase(account Key, passphrase string, hash []byte) ([]byte, error)
//SignTx(key Key, transaction *common.Transaction, chainID *big.Int) (*common.Transaction, error)
//SignTxWithPassphrase(account Key, passphrase string, transaction *common.Transaction, chainID *big.Int) (*common.Transaction, error)

type Key struct {
	ID         uuid.UUID
	Light      bool
	Address    common.Address
	privateKey *bls.SecretKey
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
	sec := GeneratePriKey()
	id := uuid.NewRandom()
	addr, _ := common.PubKeyToAddr(sec.GetPublicKey())
	key := &Key{
		Light:      light,
		ID:         id,
		Address:    addr,
		privateKey: sec,
	}
	return key
}

func (k *Key) Encrypt(auth string) ([]byte, error) {
	if k.Light {
		return EncryptKey(k, auth, LightScryptN, LightScryptP)
	}
	return EncryptKey(k, auth, StandardScryptN, StandardScryptP)
}

func (k *Key) close() {
	k.privateKey = nil
}

func (k *Key) CastEd25519Key() (crypto.PrivKey, error) {
	pri := k.privateKey.Serialize()
	var edPri = ed25519.NewKeyFromSeed(pri)
	return crypto.UnmarshalEd25519PrivateKey(edPri[:])
}

func GeneratePriKey() *bls.SecretKey {
	var sec bls.SecretKey
	sec.SetByCSPRNG()
	return &sec
}

func (k *Key) SignData(msg []byte) []byte {
	sig := k.privateKey.SignByte(msg)
	return sig.Serialize()
}

func VerifyByte(sig *bls.Sign, pub *bls.PublicKey, msg []byte) bool {
	return sig.VerifyByte(pub, msg)
}

func (k *Key) IsOpen() bool {
	return k.privateKey != nil
}

func (k *Key) SharedKey(to string) ([]byte, error) {

	frKey := bls.CastFromSecretKey(k.privateKey)
	pubTo := &bls.PublicKey{}
	if err := pubTo.DeserializeHexStr(to); err != nil {
		return nil, err
	}
	toG1 := bls.CastFromPublicKey(pubTo)
	aesKey := &bls.G1{}
	bls.G1Mul(aesKey, toG1, frKey)
	return aesKey.Serialize()[:32], nil
}

func (k *Key) StoreString(auth string) string {
	bts, _ := k.Encrypt(auth)
	return string(bts)
}

func LoadKeyFromJsonStr(str, auth string) (*Key, error) {
	parsedKey, err := DecryptKey([]byte(str), auth)
	if err != nil {
		return nil, err
	}
	return parsedKey, nil
}
