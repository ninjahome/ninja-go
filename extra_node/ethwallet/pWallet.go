package ethwallet

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pborman/uuid"

	"io/ioutil"
)

func (pw *PWallet) SignKey() *ecdsa.PrivateKey {
	return pw.key.MainPriKey
}

func (pw *PWallet) MainAddress() common.Address {
	return pw.MainAddr
}

func (pw *PWallet) SignJson(v interface{}) ([]byte, error) {
	rawBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	hash := crypto.Keccak256(rawBytes)
	return crypto.Sign(hash, pw.key.MainPriKey)
}

func (pw *PWallet) Sign(v []byte) ([]byte, error) {
	return crypto.Sign(v, pw.key.MainPriKey)
}

func (pw *PWallet) VerifySig(message, signature []byte) bool {
	hash := crypto.Keccak256Hash(message)
	pk := crypto.FromECDSAPub(&pw.key.MainPriKey.PublicKey)
	signature = signature[:len(signature)-1]
	return crypto.VerifySignature(pk, hash[:], signature)
}

func (pw *PWallet) IsOpen() bool {
	return pw.key != nil
}

func (pw *PWallet) SaveToPath(wPath string) error {
	bytes, err := json.MarshalIndent(pw, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(wPath, bytes, 0644)
}

func (pw *PWallet) Open(auth string) error {

	keyBytes, err := keystore.DecryptDataV3(pw.Crypto, auth)
	if err != nil {
		return err
	}

	key := &WalletKey{
		MainPriKey: crypto.ToECDSAUnsafe(keyBytes),
	}
	pw.key = key
	return nil
}

func (pw *PWallet) Close() {
	pw.key = nil
}

func (pw *PWallet) String() string {
	b, e := json.Marshal(pw)
	if e != nil {
		return ""
	}
	return string(b)
}

func (pw *PWallet) ExportEth(auth, eAuth, path string) error {

	keyBytes, err := keystore.DecryptDataV3(pw.Crypto, auth)
	if err != nil {
		return err
	}
	key := crypto.ToECDSAUnsafe(keyBytes)

	ethKey := &keystore.Key{
		Address:    crypto.PubkeyToAddress(key.PublicKey),
		PrivateKey: key,
	}

	var id uuid.UUID
	id = uuid.NewRandom()

	ethKey.Id = id

	newJson, err := keystore.EncryptKey(ethKey, eAuth, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return fmt.Errorf("error encrypting with new password: %v", err)
	}
	if err := ioutil.WriteFile(path, newJson, 0644); err != nil {
		return fmt.Errorf("error writing new keyfile to disk: %v", err)
	}

	return nil
}
