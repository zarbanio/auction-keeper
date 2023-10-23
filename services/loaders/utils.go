package loaders

import (
	"math/big"
	"strings"

	"github.com/zarbanio/auction-keeper/domain"
	"github.com/zarbanio/auction-keeper/domain/math"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

func debtCeiling(line *big.Int) decimal.Decimal {
	return math.Normalize(line, int64(math.RadDecimals))
}

func liquidationRatio(mat *big.Int) decimal.Decimal {
	return math.Normalize(mat, int64(math.RayDecimals))
}

func liquidationPrice(collateralAmount, debtValue, liquidationRatio decimal.Decimal) decimal.Decimal {
	if collateralAmount.Equal(decimal.Zero) {
		return decimal.NewFromInt(1<<63 - 1)
	}
	return debtValue.Mul(liquidationRatio).Div(collateralAmount)
}

func collateralizationRatio(collateralValue, debtValue decimal.Decimal) decimal.Decimal {
	if debtValue.Equal(decimal.Zero) {
		return decimal.NewFromInt(1<<63 - 1)
	}
	return collateralValue.Div(debtValue)
}

func price(par, spot, mat *big.Int) decimal.Decimal {
	p := math.RayMul(math.RayMul(par, spot), mat)
	return math.Normalize(p, int64(math.RayDecimals))
}

func debtValue(art, rate *big.Int) decimal.Decimal {
	mul := math.WadMul(art, rate)
	return math.Normalize(mul, int64(math.RayDecimals))
}

func annualStabilityFee(duty *big.Int) decimal.Decimal {
	return math.Normalize(
		new(big.Int).Sub(math.RayPow(duty, math.SecondsPerYear), math.Ray),
		int64(math.RayDecimals),
	)
}

func IlkToSymbol(name string) domain.Symbol {
	if name == "ETHA" {
		return "ETH"
	}
	if name == "ETHB" {
		return "ETH"
	}
	if name == "DAIA" {
		return "DAI"
	}
	if name == "DAIB" {
		return "DAI"
	}
	s := strings.Split(name, "-")
	return domain.Symbol(s[0])
}

func minSafeCollateralAmount(debtValue, liquidationRatio, price decimal.Decimal) decimal.Decimal {
	if price.Equal(decimal.Zero) {
		return decimal.Zero
	}
	return debtValue.Mul(liquidationRatio).Div(price)
}

func availableToMint(collateralValue, debtValue, liquidationRatio decimal.Decimal) decimal.Decimal {
	maxSafeDebtValue := collateralValue.Div(liquidationRatio)
	if debtValue.LessThan(maxSafeDebtValue) {
		return maxSafeDebtValue.Sub(debtValue)
	}
	return decimal.Zero
}
