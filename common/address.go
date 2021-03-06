package common

import (
	"encoding/hex"
	"fmt"
	"github.com/ninjahome/bls-wallet/bls"
)

const (
	AddressLength = 48
	HashLength    = 32
)

type Address [AddressLength]byte

var (
	InvalidAddr Address
)

func (a *Address) SetBytes(b []byte) error {
	if len(b) != len(a) {
		return fmt.Errorf("invalid byte for address")
	}
	copy(a[:], b)
	return nil
}

//
//func convertAndEncode(hrp string, data []byte) (string, error) {
//	converted, err := bech32.ConvertBits(data, 8, 5, true)
//	if err != nil {
//		return "", fmt.Errorf("convertBits failed:%s", err)
//	}
//	return bech32.Encode(hrp, converted)
//
//}
//func decodeAndConvert(addrStr string) (string, []byte, error) {
//	hrp, data, err := bech32.Decode(addrStr)
//	if err != nil {
//		return "", nil, fmt.Errorf("decode bech32 failed:%s", err)
//	}
//	converted, err := bech32.ConvertBits(data, 5, 8, false)
//	if err != nil {
//		return "", nil, fmt.Errorf("convert bits failed:%s", err)
//	}
//	return hrp, converted, nil
//}
//
//func IsNinJaAddress(s string) bool {
//	hrp, bytes, err := decodeAndConvert(s)
//	if err != nil || (hrp != NinjaAddressHRP) || len(bytes) != AddressLength {
//		return false
//	}
//	return true
//}

func HexToAddress(s string) (addr Address, err error) {
	bts, err := hex.DecodeString(s)
	if err != nil {
		return
	}
	err = addr.SetBytes(bts)
	return
}

func (a *Address) UnmarshalJSON(data []byte) error {
	bts, err := hex.DecodeString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	return a.SetBytes(bts)
}

func (a *Address) MarshalText() ([]byte, error) {
	return []byte(a.Hex()), nil
}

func (a *Address) Hex() string {
	return hex.EncodeToString(a[:])
}

func (a *Address) String() string {
	return a.Hex()
}

func PubKeyToAddr(p *bls.PublicKey) (addr Address, err error) {
	pubBytes := p.Serialize()
	err = addr.SetBytes(pubBytes)
	return
}

func AddrToPub(a *Address) (*bls.PublicKey, error) {
	pub := &bls.PublicKey{}
	b := make([]byte, len(a))
	copy(b, a[:])
	if err := pub.Deserialize(b); err != nil {
		return nil, err
	}
	return pub, nil
}
