package common

import (
	"encoding/hex"
	"github.com/herumi/bls-eth-go-binary/bls"
)

const (
	AddressLength = 48
	HashLength    = 32
)

type Address [AddressLength]byte

var (
	InvalidAddr Address
)

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
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
	addr.SetBytes(bts)
	return
}

func (a *Address) UnmarshalJSON(data []byte) error {
	bts, err := hex.DecodeString(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	a.SetBytes(bts)
	return nil
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

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func PubKeyToAddr(p *bls.PublicKey) Address {
	pubBytes := p.Serialize()
	return BytesToAddress(pubBytes)
}

func AddrToPub(a *Address) (*bls.PublicKey, error) {
	pub := &bls.PublicKey{}
	if err := pub.Deserialize(a[:]); err != nil {
		return nil, err
	}

	return pub, nil
}
