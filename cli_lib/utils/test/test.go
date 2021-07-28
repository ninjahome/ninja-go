package main

import (
	"fmt"
	"github.com/ninjahome/ninja-go/cli_lib/utils"
)

func main()  {
	ss:=[]string{"aaa","bbb"}


	fmt.Println(utils.StrSlice2String(ss))
}
