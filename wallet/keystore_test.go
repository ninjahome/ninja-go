package wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/ninjahome/ninja-go/common"
	"testing"
)

func TestKeyPath(t *testing.T) {
	var auth = "123"
	key := NewKey()
	fmt.Println(key.Address)
	fmt.Println(hex.EncodeToString(key.Address[:]))
	ks := NewLightKeyStore("key_dir", key.Light)
	if err := ks.StoreKey(key, auth); err != nil {
		t.Fatal(err)
	}
	fmt.Println("case 3 success=>")
}

func TestLoadKey(t *testing.T) {
	var auth = "123"
	ks := NewLightKeyStore("key_dir", false)
	path := "UTC--2021-03-03T03-29-47.812245000Z--fed1rrnt6q0rq8esdq07ca9z7a4egd80ls6ckvz50w"
	addr, err := common.HexToAddress("fed1rrnt6q0rq8esdq07ca9z7a4egd80ls6ckvz50w")
	if err != nil {
		t.Fatal(err)
	}
	key, err := ks.GetKey(addr, path, auth)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(key)
	if key.Address != addr {
		t.Fatal("load key failed")
	}
}
