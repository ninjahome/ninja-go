package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/ninjahome/bls-wallet/bls"
	ncom "github.com/ninjahome/ninja-go/common"
)

func main()  {
	froms:="0x8cdf71ed43abaf6c969f6c5472b9dffe970330cd680311ad90ebbddec13db449db9108682d9c1d46009472d7d8a13efc"
	tos := "0xb1d6e0c4a0c3a0c74d2c9e13cfa8bc5cfb349c8a5f4746a2c143483a3f352544c72bc87871a7696bd135f42fc076e5aa"
	nDays:=5

	sigs:="8c51a81dc150a9c3279c174da2e229fca0e23f70e2934a8d34f2cf480301f0d551b91226e83fb943bc6e7d104f3c170d09f88deefce2a4cbdbd0e6ac6ee514cc083ffb346b9836bab7d16440054a6eb5c795a9d7bc1cec36e815b5b3504e6c51"


	from,_:=ncom.HexToAddress(froms)
	to,_:=ncom.HexToAddress(tos)

	sigb,_:=hex.DecodeString(sigs)

	buf:=make([]byte,1024)
	n:=copy(buf,from[:])
	n+=copy(buf[n:],to[:])
	binary.BigEndian.PutUint32(buf[n:],uint32(nDays))

	n+=4

	sig := &bls.Sign{}

	if err := sig.Deserialize(sigb); err != nil {
		fmt.Println(err)
		return
	}
	p := &bls.PublicKey{}
	if err := p.Deserialize(from[:]); err != nil {
		fmt.Println(err)
		return
	}

	b:=sig.VerifyByte(p, buf[:n])

	fmt.Println(b)
}
