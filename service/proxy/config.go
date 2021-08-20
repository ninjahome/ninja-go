package proxy

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ncom "github.com/ninjahome/ninja-go/common"
	"github.com/ninjahome/ninja-go/contract"
	"strconv"
)

const (
	ProxyListenPort = 18088

	infuraUrl   = "https://kovan.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contactAddr = "0x0848abeD6000396fE5852E07ABD468fCafb4f44b"
	tokenAddr   = "0x122938b76c071142ea6b39c34ffc38e5711cada1"
)

var proxyAddr = []string{
	"39.99.198.143:9099",
	"47.113.87.58:9099",
}

type Config struct {
	ListenAddr   string   `json:"listen_addr"`
	ProxyAddr    []string `json:"proxy_addr"`
	EthUrl       string   `json:"eth_url"`
	ContractAddr string   `json:"contract_addr"`
	TokenAddr    string   `json:"token_addr"`
}

var _proxyConfig *Config = nil

func (c *Config) String() string {
	s := fmt.Sprintf("\r\n--------------------Proxy Config--------------------")
	s += fmt.Sprintf("\r\nListen addr: %s", c.ListenAddr)
	s += fmt.Sprintf("\r\nServer Addr: ")
	for i := 0; i < len(c.ProxyAddr); i++ {
		pa := c.ProxyAddr[i]
		s += fmt.Sprintf("\r\n        %s ", pa)
	}

	s += fmt.Sprintf("\r\nEth access node: %s", c.EthUrl)
	s += fmt.Sprintf("\r\ncontract: %s", c.ContractAddr)
	s += fmt.Sprintf("\r\ntoken: %s", c.TokenAddr)

	s += fmt.Sprintf("\r\n-------------------------------------------------------\r\n")

	return s
}

func InitConfig(c *Config) {
	_proxyConfig = c
}

func DefaultConfig() *Config {
	l := ":" + strconv.Itoa(ProxyListenPort)

	var pa []string
	for i := 0; i < len(proxyAddr); i++ {
		pa = append(pa, "http://"+proxyAddr[i])
	}

	return &Config{
		ListenAddr:   l,
		ProxyAddr:    pa,
		EthUrl:       infuraUrl,
		ContractAddr: contactAddr,
		TokenAddr:    tokenAddr,
	}
}

func (c *Config) NewProxyWebServer() *WebProxyServer {
	return NewWebServer(c.ListenAddr, c.ProxyAddr)
}

func GetExpireTimeFromBlockChain(uid string) (int64, error) {
	var (
		c               *ethclient.Client
		err             error
		licenseContract *contract.NinjaChatLicense
		deadline        uint64
		uidaddr         ncom.Address
		userAddr        [32]byte
	)

	if c, err = ethclient.Dial(infuraUrl); err != nil {
		return 0, err
	}

	defer c.Close()

	licenseContract, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr), c)
	if err != nil {
		return 0, err
	}

	uidaddr, err = ncom.HexToAddress(uid)
	if err != nil {
		return 0, err
	}
	userAddr, err = ncom.Naddr2ContractAddr(uidaddr)
	if err != nil {
		return 0, err
	}

	deadline, _, err = licenseContract.GetUserLicense(nil, userAddr)
	if err != nil {
		return 0, err
	}

	return int64(deadline), nil
}
