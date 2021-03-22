package wallet

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/ninjahome/ninja-go/common"
	"github.com/ninjahome/ninja-go/utils"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	keyHeaderKDF = "scrypt"

	// StandardScryptN is the N parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptN = 1 << 18

	// StandardScryptP is the P parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptP = 1

	// LightScryptN is the N parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptN = 1 << 12

	// LightScryptP is the P parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptP = 6

	scryptR     = 8
	scryptDKLen = 32
)

var (
	ErrDecrypt = fmt.Errorf("could not decrypt key with given password")
)

type KeyStore struct {
	keysDirPath string
	scryptN     int
	scryptP     int
}

func NewKeyStore(dir string) *KeyStore {
	return NewLightKeyStore(dir, false)
}

func NewLightKeyStore(dir string, isLight bool) *KeyStore {
	if isLight {
		return &KeyStore{
			keysDirPath: dir,
			scryptN:     LightScryptN,
			scryptP:     LightScryptP,
		}
	} else {
		return &KeyStore{
			keysDirPath: dir,
			scryptN:     StandardScryptN,
			scryptP:     StandardScryptP,
		}
	}
}

func keyFileFormat(keyAddr common.Address) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", utils.ToISO8601(ts), keyAddr.String())
}

func (ks KeyStore) ValidKeyFiles() (map[string]string, error) {
	ret := make(map[string]string)
	entries, err := os.ReadDir(ks.keysDirPath)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		addr, ok := ks.IsKeyFile(e.Name())
		if !ok {
			continue
		}
		ret[addr] = e.Name()
	}
	if len(ret) == 0 {
		return nil, fmt.Errorf("no valid key files under [%s]", ks.keysDirPath)
	}
	return ret, nil
}

func (ks KeyStore) IsKeyFile(filename string) (string, bool) {
	filePath := ks.joinPath(filename)
	keyJson, err := os.ReadFile(filePath)
	if err != nil {
		return "", false
	}

	k := new(encryptedKeyJSON)
	if err := json.Unmarshal(keyJson, k); err != nil {
		return "", false
	}
	if keyID := uuid.Parse(k.ID); keyID == nil {
		return "", false
	}
	if _, err := common.HexToAddress(k.Address); err != nil {
		return "", false
	}
	return k.Address, true
}

func (ks KeyStore) GetRawKey(filename, auth string) (*Key, error) {
	filePath := ks.joinPath(filename)
	keyJson, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	key, err := DecryptKey(keyJson, auth)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func (ks KeyStore) GetKey(addr common.Address, filename, auth string) (*Key, error) {
	key, err := ks.GetRawKey(filename, auth)
	if err != nil {
		return nil, err
	}
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have wallet %x, want %x", key.Address, addr)
	}
	return key, nil
}

func (ks KeyStore) StoreKey(key *Key, auth string) error {
	keyJson, err := EncryptKey(key, auth, ks.scryptN, ks.scryptP)
	if err != nil {
		return err
	}
	path := ks.joinPath(keyFileFormat(key.Address))
	fmt.Println(path)
	return utils.WriteToFile(path, keyJson)
}

func (ks KeyStore) joinPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(ks.keysDirPath, filename)
}

func EncryptKey(key *Key, auth string, scryptN, scryptP int) ([]byte, error) {
	keyBytes := key.PrivateKey.Serialize()
	if len(keyBytes) != PriKeyLen {
		return nil, fmt.Errorf("invalid private key in Key structure")
	}
	cryptoStruct, err := EncryptData(keyBytes, []byte(auth), scryptN, scryptP)
	if err != nil {
		return nil, err
	}
	encryptedKeyJSON := encryptedKeyJSON{
		key.Address.String(),
		cryptoStruct,
		key.ID.String(),
		version,
	}
	return json.MarshalIndent(encryptedKeyJSON, "", "\t")
}

func EncryptData(data, auth []byte, scryptN, scryptP int) (CryptoJSON, error) {
	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	derivedKey, err := scrypt.Key(auth, salt, scryptN, scryptR, scryptP, scryptDKLen)
	if err != nil {
		return CryptoJSON{}, err
	}
	encryptKey := derivedKey[:16]

	iv := make([]byte, aes.BlockSize) // 16
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	cipherText, err := aesCTRXOR(encryptKey, data, iv)
	if err != nil {
		return CryptoJSON{}, err
	}
	mac := utils.Keccak256(derivedKey[16:32], cipherText)

	scryptParamsJSON := make(map[string]interface{}, 5)
	scryptParamsJSON["n"] = scryptN
	scryptParamsJSON["r"] = scryptR
	scryptParamsJSON["p"] = scryptP
	scryptParamsJSON["dklen"] = scryptDKLen
	scryptParamsJSON["salt"] = hex.EncodeToString(salt)
	cipherParamsJSON := cipherParamsJSON{
		IV: hex.EncodeToString(iv),
	}

	cryptoStruct := CryptoJSON{
		Cipher:       "aes-128-ctr",
		CipherText:   hex.EncodeToString(cipherText),
		CipherParams: cipherParamsJSON,
		KDF:          keyHeaderKDF,
		KDFParams:    scryptParamsJSON,
		MAC:          hex.EncodeToString(mac),
	}
	return cryptoStruct, nil
}

func DecryptKey(keyJson []byte, auth string) (*Key, error) {
	k := new(encryptedKeyJSON)
	if err := json.Unmarshal(keyJson, k); err != nil {
		return nil, err
	}

	keyID := uuid.Parse(k.ID)
	if keyID == nil {
		return nil, fmt.Errorf("invalid key uuid")
	}

	keyBytes, err := DecryptData(k.Crypto, auth)
	if err != nil {
		return nil, err
	}
	var sec bls.SecretKey
	if err = sec.Deserialize(keyBytes); err != nil {
		return nil, err
	}
	pub := sec.GetPublicKey()
	return &Key{
		ID:         keyID,
		Address:    common.PubKeyToAddr(pub),
		PrivateKey: &sec,
	}, nil
}
func DecryptData(cj CryptoJSON, auth string) ([]byte, error) {
	if cj.Cipher != "aes-128-ctr" {
		return nil, fmt.Errorf("cipher not supported: %v", cj.Cipher)
	}
	mac, err := hex.DecodeString(cj.MAC)
	if err != nil {
		return nil, err
	}

	iv, err := hex.DecodeString(cj.CipherParams.IV)
	if err != nil {
		return nil, err
	}

	cipherText, err := hex.DecodeString(cj.CipherText)
	if err != nil {
		return nil, err
	}

	derivedKey, err := getKDFKey(cj, auth)
	if err != nil {
		return nil, err
	}

	calculatedMAC := utils.Keccak256(derivedKey[16:32], cipherText)
	if !bytes.Equal(calculatedMAC, mac) {
		return nil, ErrDecrypt
	}

	plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)
	if err != nil {
		return nil, err
	}
	return plainText, err
}

func getKDFKey(cryptoJSON CryptoJSON, auth string) ([]byte, error) {
	authArray := []byte(auth)
	salt, err := hex.DecodeString(cryptoJSON.KDFParams["salt"].(string))
	if err != nil {
		return nil, err
	}
	dkLen := ensureInt(cryptoJSON.KDFParams["dklen"])

	if cryptoJSON.KDF == keyHeaderKDF {
		n := ensureInt(cryptoJSON.KDFParams["n"])
		r := ensureInt(cryptoJSON.KDFParams["r"])
		p := ensureInt(cryptoJSON.KDFParams["p"])
		return scrypt.Key(authArray, salt, n, r, p, dkLen)

	} else if cryptoJSON.KDF == "pbkdf2" {
		c := ensureInt(cryptoJSON.KDFParams["c"])
		prf := cryptoJSON.KDFParams["prf"].(string)
		if prf != "hmac-sha256" {
			return nil, fmt.Errorf("unsupported PBKDF2 PRF: %s", prf)
		}
		key := pbkdf2.Key(authArray, salt, c, dkLen, sha256.New)
		return key, nil
	}

	return nil, fmt.Errorf("unsupported KDF: %s", cryptoJSON.KDF)
}
func ensureInt(x interface{}) int {
	res, ok := x.(int)
	if !ok {
		res = int(x.(float64))
	}
	return res
}

func aesCTRXOR(key, inText, iv []byte) ([]byte, error) {
	// AES-128 is selected due to size of encryptKey.
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesBlock, iv)
	outText := make([]byte, len(inText))
	stream.XORKeyStream(outText, inText)
	return outText, err
}
