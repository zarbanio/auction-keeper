package math

import (
	"math/big"

	"github.com/zarbanio/auction-keeper/x/decimal"
)

func BigIntFromString(str string) *big.Int {
	ret, _ := new(big.Int).SetString(str, 10)
	return ret
}

func Normalize(n *big.Int, decimals int64) decimal.Decimal {
	a := decimal.NewFromBigInt(n)
	b := decimal.NewFromBigInt(Pow10(decimals))
	return a.Div(b)
}

// BigMin returns the smaller of x or y.
func BigMin(x, y *big.Int) *big.Int {
	if x.Cmp(y) > 0 {
		return y
	}
	return x
}
