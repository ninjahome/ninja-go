package contract

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"net"
)

const (
	accessEthNodeUrl = "https://kovan.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contractAddr     = "0x6D048EA9ca9876b0F3775657b3B4eBee61DfCb54"
)

func GetEthConfig() (tAddr, cAddr common.Address, accessUrl []byte, err error) {
	cli, err := ethclient.Dial(accessEthNodeUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer cli.Close()

	var elc *NinjaConfig
	elc, err = NewNinjaConfig(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return
	}

	return elc.GetLicenseConfig(nil)
}

type BootsTrapNode struct {
	NetIp [4]byte
	Port  [6]uint16
}

const (
	wsPort int = iota
	p2pPort
	httpPort
)

func (btn *BootsTrapNode) WSHostString() string {
	return fmt.Sprintf("%s:%d", net.IPv4(btn.NetIp[0], btn.NetIp[1], btn.NetIp[2], btn.NetIp[3]).String(), btn.Port[wsPort])
}

func (btn *BootsTrapNode) HttpHostString() string {
	return fmt.Sprintf("%s:%d", net.IPv4(btn.NetIp[0], btn.NetIp[1], btn.NetIp[2], btn.NetIp[3]).String(), btn.Port[httpPort])
}

func (btn *BootsTrapNode) P2pHostString() string {
	return fmt.Sprintf("%s:%d", net.IPv4(btn.NetIp[0], btn.NetIp[1], btn.NetIp[2], btn.NetIp[3]).String(), btn.Port[p2pPort])
}

func GetBootsTrapList() ([]*BootsTrapNode, error) {

	var boots []*BootsTrapNode

	cli, err := ethclient.Dial(accessEthNodeUrl)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	var elc *NinjaConfig

	elc, err = NewNinjaConfig(common.HexToAddress(contractAddr), cli)
	if err != nil {
		return nil, err
	}

	var l [][4]byte
	l, err = elc.GetIpAddrList(nil)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(l); i++ {
		if l[i][0] == 0 {
			break
		}
		bt := &BootsTrapNode{}
		bt.NetIp = l[i]

		bt.Port[0], bt.Port[1], bt.Port[2], bt.Port[3], bt.Port[4], bt.Port[5], err = elc.GetIPPort(nil, l[i])
		if err != nil {
			continue
		}

		boots = append(boots, bt)

	}

	return boots, nil
}
