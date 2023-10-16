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
