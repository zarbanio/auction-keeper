package math

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestCalculateCompoundedInterest(t *testing.T) {
	actual := CalculateCompoundedInterest(new(big.Int).Mul(big.NewInt(23), Pow10(16)), time.Hour*24*7)
	assert.Equal(t, "1000000000004410958903948800", actual.String())
}

func TestGetCompoundedBalance(t *testing.T) {
	principalBalance := new(big.Int).Mul(big.NewInt(78), Pow10(18))
	reserveIndex := new(big.Int).Mul(big.NewInt(8000), Pow10(18))
	reserveRate := new(big.Int).Mul(big.NewInt(19), Pow10(16))
	duration := time.Hour * 24 * 7
	actual := GetCompoundedBalance(principalBalance, reserveIndex, reserveRate, duration)
	assert.Equal(t, "624000000002274", actual.String())
}

func TestCalculateLinearInterest(t *testing.T) {
	rate := new(big.Int).Mul(big.NewInt(4), Pow10(16))
	duration := time.Hour * 24 * 7
	actual := CalculateLinearInterest(rate, duration)
	assert.Equal(t, "1000000000000767123287671233", actual.String())
}

func TestGetReserveNormalizedIncome(t *testing.T) {
	rate := new(big.Int).Mul(big.NewInt(4), Pow10(16))
	reserveIndex := new(big.Int).Mul(big.NewInt(8000), Pow10(18))
	duration := time.Hour * 24 * 365 * 10
	actual := GetReserveNormalizedIncome(rate, reserveIndex, duration)
	expected, ok := new(big.Int).SetString("7999999996800000000000", 10)
	if !ok {
		t.Fatal()
	}
	diff := new(big.Int).Sub(actual, expected)
	absDiff := new(big.Int).Abs(diff)
	assert.Equal(t, Zero, new(big.Int).Div(absDiff, Pow10(14)))
}

func TestGetLinearBalance(t *testing.T) {
	rate := new(big.Int).Mul(big.NewInt(4), Pow10(16))
	reserveIndex := new(big.Int).Mul(big.NewInt(8000), Pow10(18))
	balance := new(big.Int).Mul(big.NewInt(54), Pow10(18))
	duration := time.Hour * 24 * 365
	actual := GetLinearBalance(balance, reserveIndex, rate, duration)
	expected, ok := new(big.Int).SetString("431999999982720", 10)
	if !ok {
		t.Fatal()
	}
	diff := new(big.Int).Sub(actual, expected)
	absDiff := new(big.Int).Abs(diff)
	assert.Equal(t, Zero, new(big.Int).Div(absDiff, Pow10(12)))
}

func TestGetCompoundedStableBalance(t *testing.T) {
	rate := new(big.Int).Mul(big.NewInt(6), Pow10(16))
	balance := new(big.Int).Mul(big.NewInt(7), Pow10(16))
	duration := time.Hour * 24 * 299
	actual := GetCompoundedStableBalance(balance, rate, duration)
	expected, ok := new(big.Int).SetString("69999999996559452", 10)
	if !ok {
		t.Fatal()
	}
	diff := new(big.Int).Sub(actual, expected)
	absDiff := new(big.Int).Abs(diff)
	assert.Equal(t, Zero, new(big.Int).Div(absDiff, Pow10(12)))
}

func TestCalculateHealthFactor(t *testing.T) {
	collateralBalanceETH := new(big.Int).Mul(big.NewInt(7), Pow10(16))
	borrowBalanceETH := new(big.Int).Mul(big.NewInt(3), Pow10(16))
	currentLiquidationThreshold := new(big.Int).Mul(big.NewInt(75), Pow10(4))
	actual := CalculateHealthFactor(collateralBalanceETH, borrowBalanceETH, currentLiquidationThreshold)
	assert.Equal(t, "175", actual.String())
}

func TestCalculateAvailableBorrowsETH(t *testing.T) {
	collateralBalanceETH := new(big.Int).Mul(big.NewInt(7), Pow10(16))
	borrowBalanceETH := new(big.Int).Mul(big.NewInt(3), Pow10(16))
	currentLtv := new(big.Int).Mul(big.NewInt(60), Pow10(4))
	actual := CalculateAvailableBorrowsETH(collateralBalanceETH, borrowBalanceETH, currentLtv)
	assert.Equal(t, "4170000000000000000", actual.String())
}

func TestCalculateAverageRate(t *testing.T) {
	index0 := new(big.Int).Mul(big.NewInt(8000), Pow10(18))
	index1 := new(big.Int).Mul(big.NewInt(9000), Pow10(18))
	duration := time.Hour * 24 * 10
	actual := CalculateAverageRate(index0, index1, duration)
	expected, ok := new(big.Float).SetString("4.56250000000012848")
	if !ok {
		t.Fatal()
	}
	diff := new(big.Float).Sub(actual, expected)
	absDiff := new(big.Float).Abs(diff)
	powMinus12 := new(big.Float).Quo(OneFloat, new(big.Float).SetInt(Pow10(12)))
	assert.Negative(t, absDiff.Cmp(powMinus12))
}
