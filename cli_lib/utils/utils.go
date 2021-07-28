package utils

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
)

func buf2int32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func ID2IconIdx(id string, mod int) uint32 {
	s := sha256.New()
	s.Write([]byte(id))

	h := s.Sum(nil)

	sum := buf2int32(h)

	for i := 1; i < len(h)/4; i++ {
		sum = sum ^ buf2int32(h[i*4:])
	}

	return sum % uint32(mod)
}




func StrSlice2String(ss []string) string {
	j,_:=json.Marshal(ss)

	return string(j)
}