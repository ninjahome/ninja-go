package wallet

import (
	"math"
	"math/big"
)

func BalanceHuman(balance *big.Int) float64 {
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	v := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	vv, _ := v.Float64()
	return vv
}

func BalanceEth(balance float64) *big.Int {
	fbalance := new(big.Float)
	fbalance.SetFloat64(balance)
	v := new(big.Float).Mul(fbalance, big.NewFloat(math.Pow10(18)))

	vv := new(big.Int)
	v.Int(vv)

	return vv
}
