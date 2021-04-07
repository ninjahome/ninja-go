package common

import (
	"fmt"
	"testing"
)

func TestValidAddress(t *testing.T) {
	var addr Address
	if addr != InvalidAddr {
		t.Fatal("invalid address")
	}
	fmt.Println(addr[:])
	fmt.Println(InvalidAddr[:])
}

func TestDecodeAddress(t *testing.T) {
	str := "958019a0b9fe3c8df83e6fbbfdc85a834628194a58116e68efd91640eb9b2beddac3f3c03a041ea1015d9be8b961df2a"
	addr, err := HexToAddress(str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("success=>", addr[:])
}
