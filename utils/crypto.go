package utils

import (
	"crypto/aes"
	"crypto/rand"
	"golang.org/x/crypto/sha3"
	"hash"
	"io"
)

type Salt [aes.BlockSize]byte

type KeccakState interface {
	hash.Hash
	Read([]byte) (int, error)
}

// Keccak256 calculates and returns the Keccak256 hash of the input data.
func Keccak256(data ...[]byte) []byte {
	b := make([]byte, 32)
	d := sha3.NewLegacyKeccak256().(KeccakState)
	for _, b := range data {
		_, _ = d.Write(b)
	}
	_, _ = d.Read(b)
	return b
}

func NewSalt() (s Salt, err error) {
	iv := make([]byte, aes.BlockSize) //[:]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}
	copy(s[:], iv)
	return
}
