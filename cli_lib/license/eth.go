package licenseLib

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/ninja-go/contract"
	"github.com/ninjahome/ninja-go/extra_node/ethwallet"
	"math/big"
)

const (
	infuraUrl   = "https://kovan.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	contactAddr = "0x0848abeD6000396fE5852E07ABD468fCafb4f44b"
	tokenAddr   = "0x122938b76c071142ea6b39c34ffc38e5711cada1"
)

var _ethWallet ethwallet.Wallet

func NewWallet(auth string) string {
	if wallet, err := ethwallet.NewWallet(auth); err != nil {
		return ""
	} else {
		return wallet.String()
	}
}

func LoadWallet(walletStr, auth string) bool {
	if w, err := ethwallet.LoadWalletByData(walletStr); err != nil {
		return false
	} else {
		if err = w.Open(auth); err != nil {
			return false
		}
		_ethWallet = w
	}

	return true
}

type EthBalance struct {
	Approve string `json:"approve"`
	Eth     string `json:"eth"`
	Ncc     string `json:"ncc"`
}

func bigInt2str(b *big.Int) string {
	if b == nil {
		return ""
	}

	return b.String()
}

func Balance() string {
	appr, eth, ncc, err := _balance()
	if err != nil {
		return ""
	}

	eb := &EthBalance{
		Approve: bigInt2str(appr),
		Eth:     bigInt2str(eth),
		Ncc:     bigInt2str(ncc),
	}

	j, _ := json.Marshal(*eb)

	return string(j)
}

func _balance() (approve, eth, ncc *big.Int, err error) {

	if _ethWallet == nil {
		err = errors.New("wallet have not been opened")
		return
	}

	var cli *ethclient.Client
	cli, err = ethclient.Dial(infuraUrl)
	if err != nil {
		return
	}
	defer cli.Close()

	eth, err = cli.BalanceAt(context.TODO(), _ethWallet.MainAddress(), nil)
	if err != nil {
		fmt.Println("eth balance error")
	}
	var token *contract.NinjaToken
	token, err = contract.NewNinjaToken(common.HexToAddress(tokenAddr), cli)
	if err != nil {
		fmt.Println("token client error")
	} else {
		approve, err = token.Allowance(nil, _ethWallet.MainAddress(), common.HexToAddress(contactAddr))
		if err != nil {
			fmt.Println("token allowance error")
		}

		ncc, err = token.BalanceOf(nil, _ethWallet.MainAddress())
		if err != nil {
			fmt.Println("token ncc error")
		}
	}

	return
}

func getRandomId() (rid [32]byte) {
	rand.Read(rid[:])
	return
}

func toPubKeyString(priv *ecdsa.PrivateKey) string {
	pubkey := priv.PublicKey
	return crypto.PubkeyToAddress(pubkey).String()
}

func generateLicense(nDays int) (rid [32]byte, txh common.Hash, err error) {
	if _ethWallet == nil {
		err = errors.New("wallet have not been opened")
		return
	}
	var (
		cli          *ethclient.Client
		ncl          *contract.NinjaChatLicense
		nid          *big.Int
		transactOpts *bind.TransactOpts
		tx           *types.Transaction
	)
	cli, err = ethclient.Dial(infuraUrl)
	if err != nil {
		return
	}
	defer cli.Close()

	ncl, err = contract.NewNinjaChatLicense(common.HexToAddress(contactAddr), cli)
	if err != nil {
		return
	}

	nid, err = cli.ChainID(context.TODO())
	if err != nil {
		return
	}

	transactOpts, err = bind.NewKeyedTransactorWithChainID(_ethWallet.SignKey(), nid)
	if err != nil {
		return
	}

	rid = getRandomId()

	tx, err = ncl.GenerateLicense(transactOpts, rid, uint32(nDays))
	if err != nil {
		return
	}

	txh = tx.Hash()

	return
}
