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
	str := "fed1w02d3aae09wxk66w68ecqzf7esnzzay8eu9aqc"
	addr, err := HexToAddress(str)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("success=>", addr[:])
}
