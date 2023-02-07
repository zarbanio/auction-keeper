package math

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/shopspring/decimal"
	"math/big"
	"sync"
)

var cache map[int64]*big.Int
var once sync.Once

func Pow10(n int64) *big.Int {
	once.Do(func() {
		cache = make(map[int64]*big.Int)
	})
	if v, ok := cache[n]; ok {
		return v
	}
	cache[n] = math.BigPow(10, n)
	return cache[n]
}

func Normalize(n *big.Int, decimals int64) decimal.Decimal {
	return decimal.NewFromBigInt(n, 0).Div(decimal.NewFromBigInt(Pow10(decimals), 0)).Truncate(8)
}

func BigIntFromString(str string) *big.Int {
	ret, _ := new(big.Int).SetString(str, 10)
	return ret
}

func DecimalFromString(str string) decimal.Decimal {
	ret, _ := decimal.NewFromString(str)
	return ret
}
