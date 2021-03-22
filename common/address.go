package common

import (
	"fmt"
	"github.com/btcsuite/btcutil/bech32"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/ninjahome/ninja-go/utils"
)

const (
	NinjaAddressHRP = "nin"
	AddressLength   = 20
	HashLength      = 32
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

func convertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", fmt.Errorf("convertBits failed:%s", err)
	}
	return bech32.Encode(hrp, converted)

}
func decodeAndConvert(addrStr string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(addrStr)
	if err != nil {
		return "", nil, fmt.Errorf("decode bech32 failed:%s", err)
	}
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, fmt.Errorf("convert bits failed:%s", err)
	}
	return hrp, converted, nil
}

func IsFedAddress(s string) bool {
	hrp, bytes, err := decodeAndConvert(s)
	if err != nil || (hrp != NinjaAddressHRP) || len(bytes) != AddressLength {
		return false
	}
	return true
}
func HexToAddress(s string) (addr Address, err error) {
	hrp, bytes, err := decodeAndConvert(s)
	if err != nil || (hrp != NinjaAddressHRP) || len(bytes) != AddressLength {
		return addr, err
	}
	return BytesToAddress(bytes), nil
}

func (a Address) Hex() string {
	str, err := convertAndEncode(NinjaAddressHRP, a[:])
	if err != nil {
		return ""
	}
	return str
}

func (a Address) String() string {
	return a.Hex()
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func PubKeyToAddr(p *bls.PublicKey) Address {
	pubBytes := p.Serialize()
	return BytesToAddress(utils.Keccak256(pubBytes[1:])[12:])
}
