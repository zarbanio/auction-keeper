package math

import (
	"math/big"
	"time"
)

func CalculateAssetAPY(rate *big.Int) *big.Int {
	ratePerSecond := RayDiv(rate, new(big.Int).Mul(Ray, SecondsPerYear)) // rate per second in Ray unit
	assetAPYInRay := new(big.Int).Sub(BinomialApproximatedRayPow(ratePerSecond, SecondsPerYear), Ray)

	return assetAPYInRay
}

func CalculateCompoundedInterest(rate *big.Int, duration time.Duration) *big.Int {
	ratePerSecond := new(big.Int).Div(rate, SecondsPerYear)
	return BinomialApproximatedRayPow(ratePerSecond, big.NewInt(int64(duration.Seconds())))
}

func GetCompoundedBalance(principalBalance, reserveIndex, reserveRate *big.Int, duration time.Duration) *big.Int {
	if principalBalance.Cmp(Zero) == 0 {
		return principalBalance
	}

	compoundedInterest := CalculateCompoundedInterest(reserveRate, duration)
	cumulatedInterest := RayMul(compoundedInterest, reserveIndex)
	principalBalanceRay := WadToRay(principalBalance)

	return RayToWad(RayMul(principalBalanceRay, cumulatedInterest))
}

func CalculateLinearInterest(rate *big.Int, duration time.Duration) *big.Int {
	timeDelta := WadToRay(big.NewInt(int64(duration.Seconds())))
	timeDeltaInSeconds := RayDiv(timeDelta, WadToRay(SecondsPerYear))
	mul := RayMul(rate, timeDeltaInSeconds)
	return new(big.Int).Add(mul, Ray)
}

func GetReserveNormalizedIncome(rate, index *big.Int, duration time.Duration) *big.Int {
	if rate.Cmp(Zero) == 0 {
		return index
	}

	cumulatedInterest := CalculateLinearInterest(rate, duration)
	return RayMul(cumulatedInterest, index)
}

func GetLinearBalance(balance, index, rate *big.Int, duration time.Duration) *big.Int {
	mul := RayMul(WadToRay(balance), GetReserveNormalizedIncome(rate, index, duration))
	return RayToWad(mul)
}

func GetCompoundedStableBalance(principalBalance, userStableRate *big.Int, duration time.Duration) *big.Int {
	if principalBalance.Cmp(Zero) == 0 {
		return principalBalance
	}

	cumulatedInterest := CalculateCompoundedInterest(userStableRate, duration)
	principalBalanceRay := WadToRay(principalBalance)
	return RayToWad(RayMul(principalBalanceRay, cumulatedInterest))
}

func CalculateHealthFactor(collateralBalanceETH, borrowBalanceETH, currentLiquidationThreshold *big.Int) *big.Int {
	if borrowBalanceETH.Cmp(Zero) == 0 {
		return MinusOne
	}
	mul := new(big.Int).Mul(collateralBalanceETH, currentLiquidationThreshold)
	div := new(big.Int).Div(mul, Pow10(int64(LoanToValuePrecision)))
	return new(big.Int).Div(div, borrowBalanceETH)
}

func CalculateAvailableBorrowsETH(collateralBalanceETH, borrowBalanceETH, currentLtv *big.Int) *big.Int {
	if currentLtv.Cmp(Zero) == 0 {
		return Zero
	}
	mul := new(big.Int).Mul(collateralBalanceETH, currentLtv)
	div := new(big.Int).Div(mul, Pow10(int64(LoanToValuePrecision)))
	availableBorrowsETH := new(big.Int).Sub(div, borrowBalanceETH)

	if availableBorrowsETH.Cmp(Zero) == 1 {
		return availableBorrowsETH
	}
	return Zero
}

func CalculateAverageRate(index0, index1 *big.Int, duration time.Duration) *big.Float {
	div := new(big.Float).Quo(new(big.Float).SetInt(index1), new(big.Float).SetInt(index0))
	sub := new(big.Float).Sub(div, OneFloat)
	timeDelta := big.NewFloat(duration.Seconds())
	div = new(big.Float).Quo(sub, timeDelta)
	return new(big.Float).Mul(div, SecondsPerYearFloat)
}
