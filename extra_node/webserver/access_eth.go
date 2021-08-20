package webserver

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"

	"github.com/ninjahome/ninja-go/contract"
	"github.com/ninjahome/ninja-go/extra_node/config"
)

func Bind(issueAddr common.Address, userAddr [32]byte, randomId [32]byte, nDays int32, signature []byte, privKey *ecdsa.PrivateKey) (tx []byte, err error) {

	c := config.GetExtraConfig()
	var cli *ethclient.Client
	cli, err = ethclient.Dial(c.EthUrl)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	var ncl *contract.NinjaChatLicense
	ncl, err = contract.NewNinjaChatLicense(common.HexToAddress(c.LicenseContract), cli)
	if err != nil {
		return nil, err
	}

	var nid *big.Int
	nid, err = cli.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}

	var transactOpts *bind.TransactOpts
	transactOpts, err = bind.NewKeyedTransactorWithChainID(privKey, nid)
	if err != nil {
		return nil, err
	}
	var txs *types.Transaction
	txs, err = ncl.BindLicense(transactOpts, issueAddr, userAddr, randomId, uint32(nDays), signature)
	if err != nil {
		return nil, err
	}

	hash := txs.Hash()

	return hash[:], nil
}

func TransferLicense(fromAddr, toAddr [32]byte, nDays int, privKey *ecdsa.PrivateKey) (tx []byte, err error) {
	c := config.GetExtraConfig()
	var cli *ethclient.Client
	cli, err = ethclient.Dial(c.EthUrl)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	var ncl *contract.NinjaChatLicense
	ncl, err = contract.NewNinjaChatLicense(common.HexToAddress(c.LicenseContract), cli)
	if err != nil {
		return nil, err
	}

	var nid *big.Int
	nid, err = cli.ChainID(context.TODO())
	if err != nil {
		return nil, err
	}

	var transactOpts *bind.TransactOpts
	transactOpts, err = bind.NewKeyedTransactorWithChainID(privKey, nid)
	if err != nil {
		return nil, err
	}
	var txs *types.Transaction
	txs, err = ncl.TransferLicense(transactOpts, fromAddr, toAddr, uint32(nDays))
	if err != nil {
		return nil, err
	}

	hash := txs.Hash()

	return hash[:], nil
}
