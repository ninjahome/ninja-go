package utils

import (
	"fmt"
	"testing"
)

func TestSizeUint(t *testing.T) {
	var x1 ByteSize = 1<<10 + 1<<8
	fmt.Println(x1.String())
	var x2 ByteSize = 1<<20 + 1<<18 + 1<<19
	fmt.Println(x2.String())
	var x3 ByteSize = 1<<30 + 1<<28
	fmt.Println(x3.String())
	var x4 ByteSize = 1<<40 + 1<<38
	fmt.Println(x4.String())
}
