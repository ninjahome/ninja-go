package test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"net"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/ninja-go/contract"
	"testing"
)

const(
	accessEthNodeUrl = "https://kovan.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contractAddr = "0x6D048EA9ca9876b0F3775657b3B4eBee61DfCb54"
)


//
//var _cipherTxt = `{
//		"cipher":"aes-128-ctr",
//		"ciphertext":"c79da2f1df233563b0c1c993a2a8209a4834d8f3452cadab0cbb398cde30973d",
//		"cipherparams":{
//			"iv":"8fe749bfec3d41d05b980b01255bf87c"
//		},
//		"kdf":"scrypt",
//		"kdfparams":{
//			"dklen":32,
//			"n":262144,
//			"p":1,
//			"r":8,
//			"salt":"9e2d547fd6eafe616d522c7b18960795847ba96c5d0ba3c2f1c4c093cb181fd3"
//		},
//		"mac":"65db7be298d916ca9ef27d9f5377b708059f76366c0a78311074b1187cc870a4"
//	}
//`
//
//func GetPrivKey() *ecdsa.PrivateKey {
//	j := &keystore.CryptoJSON{}
//	json.Unmarshal([]byte(_cipherTxt), j)
//	key, err := keystore.DecryptDataV3(*j, "1qaz2wsx")
//	if err != nil {
//		fmt.Println("err is ", err)
//		return nil
//	}
//	priKey := crypto.ToECDSAUnsafe(key)
//
//	return priKey
//}
//
//func getRandomId() (rid [32]byte) {
//	rand.Read(rid[:])
//	return
//}
//
//func toPubKeyString(priv *ecdsa.PrivateKey) string {
//	pubkey := priv.PublicKey
//	return crypto.PubkeyToAddress(pubkey).String()
//}

func TestSetEthConfig(t *testing.T)  {

	infuraUrl   := "https://kovan.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	cAddr := "0x291F8A7353E460416095602e7BEc53a12cb5F0cC"
	tAddr   := "0x122938b76c071142ea6b39c34ffc38e5711cada1"

	cli,err:=ethclient.Dial(accessEthNodeUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	elc,err:=contract.NewNinjaConfig(common.HexToAddress(contractAddr),cli)
	if err!=nil{
		panic(err)
	}

	chainId,err:=cli.ChainID(context.TODO())
	if err!= nil{
		panic(err)
	}

	transopt,err:=bind.NewKeyedTransactorWithChainID(GetPrivKey(),chainId)
	if err!= nil{
		panic(err)
	}

	tx,err:=elc.LicenseConfigSet(transopt,common.HexToAddress(tAddr),common.HexToAddress(cAddr),[]byte(infuraUrl))
	if err!=nil{
		panic(err)
	}

	fmt.Println(tx.Hash().String())
}

func TestPrintEthConfig(t *testing.T)  {
	cli,err:=ethclient.Dial(accessEthNodeUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	elc,err:=contract.NewNinjaConfig(common.HexToAddress(contractAddr),cli)
	if err!=nil{
		panic(err)
	}

	tk,c,u,err:=elc.GetLicenseConfig(nil)
	if err!=nil{
		panic(err)
	}

	fmt.Println("token:",tk.String())
	fmt.Println("contract:",c.String())
	fmt.Println("url:",string(u))

}

func TestAddBootstrapNode(t *testing.T)  {
	cli,err:=ethclient.Dial(accessEthNodeUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	elc,err:=contract.NewNinjaConfig(common.HexToAddress(contractAddr),cli)
	if err!=nil{
		panic(err)
	}

	chainId,err:=cli.ChainID(context.TODO())
	if err!= nil{
		panic(err)
	}

	transopt,err:=bind.NewKeyedTransactorWithChainID(GetPrivKey(),chainId)
	if err!= nil{
		panic(err)
	}

	ip:="18.183.238.197"
	var port1 uint16 = 16666
	var port2 uint16 = 19999
	var port3 uint16 = 18088

	fmt.Println(ip)

	ipb:=net.ParseIP(ip).To4()

	var ipa [4]byte
	copy(ipa[:],ipb)

	//fmt.Println(ipa[0],ipa[1],ipa[2],ipa[3])

	tx,err:=elc.AddBootsTrap(transopt,ipa,port1,port2,port3,0,0,0)
	if err!=nil{
		panic(err)
	}
	fmt.Println(tx.Hash().String())
}

func TestListBootsTrapNode(t *testing.T)  {
	cli,err:=ethclient.Dial(accessEthNodeUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	elc,err:=contract.NewNinjaConfig(common.HexToAddress(contractAddr),cli)
	if err!=nil{
		panic(err)
	}

	l,err:=elc.GetIpAddrList(nil)
	if err!=nil{
		panic(err)
	}

	for i:=0;i<len(l);i++{
		if l[i][0] == 0{
			break
		}
		ip:=net.IPv4(l[i][0],l[i][1],l[i][2],l[i][3])

		port1,port2,port3,_,_,_,err:=elc.GetIPPort(nil,l[i])
		if err!=nil{
			fmt.Println(err)
		}

		fmt.Println(ip.String(),port1,port2,port3)
	}

}

func TestDelBootsTrapNode(t *testing.T)  {
	cli,err:=ethclient.Dial(accessEthNodeUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	elc,err:=contract.NewNinjaConfig(common.HexToAddress(contractAddr),cli)
	if err!=nil{
		panic(err)
	}

	chainId,err:=cli.ChainID(context.TODO())
	if err!= nil{
		panic(err)
	}

	transopt,err:=bind.NewKeyedTransactorWithChainID(GetPrivKey(),chainId)
	if err!= nil{
		panic(err)
	}

	ipaddr:="47.113.87.58"
	ipb:=net.ParseIP(ipaddr).To4()

	var ip [4]byte

	copy(ip[:],ipb)

	tx,err:=elc.DelBootsTrap(transopt,ip)
	if err!=nil{
		panic(err)
	}

	fmt.Println(tx.Hash().String())
}

