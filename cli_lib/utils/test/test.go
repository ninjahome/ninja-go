package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/ninjahome/ninja-go/cli_lib/utils"
)

func main()  {
	ss:=[]string{"aaa","bbb"}


	fmt.Println(utils.StrSlice2String(ss))


	buf:=make([]byte,8)


	var a int64 = 200
	var b int64 = 1000000

	binary.BigEndian.PutUint64(buf,uint64(a))

	fmt.Println(hex.EncodeToString(buf))
	binary.BigEndian.PutUint64(buf,uint64(b))

	fmt.Println(hex.EncodeToString(buf))

}
