package math

import (
	"math/big"
	"time"
)

func CalculateCompoundedInterest(rate *big.Int, duration time.Duration) *big.Int {
	ratePerSecond := new(big.Int).Div(rate, SecondsPerYear)
	return BinomialApproximatedRayPow(ratePerSecond, big.NewInt(int64(duration.Seconds())))
}
