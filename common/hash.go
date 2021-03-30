package common

import (
	"database/sql/driver"
	"encoding/hex"
	"math/big"
	"reflect"
)

var hashT = reflect.TypeOf(Hash{})

type Hash [HashLength]byte

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}
func BigToHash(b *big.Int) Hash { return BytesToHash(b.Bytes()) }

func HexToHash(s string) Hash { panic("implement me") }

func (h Hash) Bytes() []byte { return h[:] }

func (h Hash) Big() *big.Int { return new(big.Int).SetBytes(h[:]) }
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

//func (h *Hash) UnmarshalText(input []byte) error {
//	return hexutil.UnmarshalFixedText("Hash", input, h[:])
//}
//func (h *Hash) UnmarshalJSON(input []byte) error {
//	return hexutil.UnmarshalFixedJSON(hashT, input, h[:])
//}
func (h Hash) Value() (driver.Value, error) {
	return h[:], nil
}

func (h Hash) String() string {
	return h.Hex()
}
func (h Hash) Hex() string { return hex.EncodeToString(h[:]) }
