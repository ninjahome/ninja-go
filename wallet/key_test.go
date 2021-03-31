package wallet

import (
	"crypto/ed25519"
	"fmt"
	"github.com/herumi/bls-eth-go-binary/bls"
	"github.com/ninjahome/ninja-go/common"
	"testing"
)

func init() {
	if err := bls.Init(bls.BLS12_381); err != nil {
		panic(err)
	}
	if err := bls.SetETHmode(bls.EthModeDraft07); err != nil {
		panic(err)
	}
}
func TestPublicKey(t *testing.T) {

	key := NewLightKey(false)
	fmt.Println(key.Address)
	strAddr := key.Address.String()
	fmt.Println(strAddr)

	pub2, err := common.HexToAddress(strAddr)
	if err != nil {
		t.Fatal(err)
	}
	if pub2 != key.Address {
		t.Fatal("failed to convert address")
	}
}
func TestG1Pub(t *testing.T) {
	pri_a := GeneratePriKey()
	pub_a := pri_a.GetPublicKey()

	pri_b := GeneratePriKey()
	pub_b := pri_b.GetPublicKey()

	G1_a := bls.CastFromPublicKey(pub_a)
	Fr_a := bls.CastFromSecretKey(pri_a)

	G1_b := bls.CastFromPublicKey(pub_b)
	Fr_b := bls.CastFromSecretKey(pri_b)
	key_ab := &bls.G1{}
	bls.G1Mul(key_ab, G1_b, Fr_a)

	key_ba := &bls.G1{}
	bls.G1Mul(key_ba, G1_a, Fr_b)

	fmt.Printf("key a->b:=>[%x]\n", key_ab.Serialize())
	fmt.Printf("key b->a:=>[%x]\n", key_ba.Serialize())

}

func TestKeyNew(t *testing.T) {
	k := NewKey()
	pub := k.privateKey.GetPublicKey()
	fmt.Printf("case 1 success=> pub:%x\n", pub.Serialize())

}

func TestKeyAuth(t *testing.T) {
	var auth = "123"
	key := NewKey()

	cipherTxt := key.StoreString(auth)
	fmt.Println(cipherTxt)
	parsedKey, err := LoadKeyFromJsonStr(cipherTxt, auth)
	if err != nil {
		t.Fatal(err)
	}
	if key.Address != parsedKey.Address {
		t.Fatal("address is not same")
	}

	if !key.privateKey.IsEqual(parsedKey.privateKey) {
		t.Fatal("private key is not same")
	}
	fmt.Println("case 2 success=>")
}

func TestCastEdKey(t *testing.T) {
	k := NewKey()
	pri := k.privateKey.Serialize()
	var edPri = ed25519.NewKeyFromSeed(pri)
	t.Logf("edkey:%x\n", edPri)

	var edPriS = ed25519.NewKeyFromSeed(pri)
	t.Logf("edkey:%x\n", edPriS)

	msg := "hello world"
	sig := ed25519.Sign(edPri, []byte(msg))
	t.Logf("sig 1:%x\n", sig)
	edPub := edPri.Public().(ed25519.PublicKey)
	if !ed25519.Verify(edPub, []byte(msg), sig) {
		t.Fatal("convert 1 failed")
	}

	sig2 := ed25519.Sign(edPriS, []byte(msg))
	t.Logf("sig 2:%x\n", sig2)
	edPub2 := edPriS.Public().(ed25519.PublicKey)
	if !ed25519.Verify(edPub2, []byte(msg), sig2) {
		t.Fatal("convert 2 failed")
	}
	t.Logf("TestCastEdKey success=>:")
}

func TestCastP2PKey(t *testing.T) {
	k := NewKey()
	p2pKey, err := k.CastEd25519Key()
	if err != nil {
		t.Fatal(err)
	}
	bts, err := p2pKey.Bytes()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("fisrt time get:[%x]\n", bts)
	for i := 0; i < 20; i++ {
		_, err = k.CastEd25519Key()
		if err != nil {
			t.Fatal(err)
		}
	}
	p2pKey, err = k.CastEd25519Key()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("22th get:[%x]\n", bts)
}

func TestPriKeyToHexStr(t *testing.T) {
	k := NewKey()
	t.Log(k.privateKey.SerializeToHexStr())
	t.Log(k.privateKey.GetHexString())
}

func TestHexStrToPriKey(t *testing.T) {

	keyStrs := []string{
		"066c6b1a28955a9089670d1e1386484f7370ef7b4f725876e72d82438de06c9e",
		"66c6b1a28955a9089670d1e1386484f7370ef7b4f725876e72d82438de06c9e",
	}

	for _, keyStr := range keyStrs {
		var key bls.SecretKey
		if err := key.SetHexString(keyStr); err != nil {
			t.Fatal(err)
		}
		t.Log(key.SerializeToHexStr())
		t.Log(key.GetHexString())
		if err := key.DeserializeHexStr(keyStr); err != nil {
			t.Fatal(err)
		}
		t.Log(key.SerializeToHexStr())
		t.Log(key.GetHexString())
	}
}
