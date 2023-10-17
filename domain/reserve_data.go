package domain

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/domain/math"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

type RawReserveData struct {
	Id                             string         `json:"id"`
	UnderlyingAsset                common.Address `json:"underlying_asset"`
	Symbol                         string         `json:"symbol"`
	Decimals                       int64          `json:"decimals"`
	IsActive                       bool           `json:"is_active"`
	IsFrozen                       bool           `json:"is_frozen"`
	UsageAsCollateralEnabled       bool           `json:"usage_as_collateral_enabled"`
	ZTokenAddress                  common.Address `json:"ast_token_address"`
	StableDebtTokenAddress         common.Address `json:"stable_debt_token_address"`
	VariableDebtTokenAddress       common.Address `json:"variable_debt_token_address"`
	BorrowingEnabled               bool           `json:"borrowing_enabled"`
	StableBorrowRateEnabled        bool           `json:"stable_borrow_rate_enabled"`
	ReserveFactor                  *big.Int       `json:"reserve_factor"`
	BaseLTVasCollateral            *big.Int       `json:"base_lt_vas_collateral"`
	StableRateSlope1               *big.Int       `json:"stable_rate_slope_1"`
	StableRateSlope2               *big.Int       `json:"stable_rate_slope_2"`
	AverageStableRate              *big.Int       `json:"average_stable_rate"`
	StableDebtLastUpdateTimestamp  int64          `json:"stable_debt_last_update_timestamp"`
	VariableRateSlope1             *big.Int       `json:"variable_rate_slope_1"`
	VariableRateSlope2             *big.Int       `json:"variable_rate_slope_2"`
	LiquidityIndex                 *big.Int       `json:"liquidity_index"`
	ReserveLiquidationThreshold    *big.Int       `json:"reserve_liquidation_threshold"`
	ReserveLiquidationBonus        *big.Int       `json:"reserve_liquidation_bonus"`
	VariableBorrowIndex            *big.Int       `json:"variable_borrow_index"`
	VariableBorrowRate             *big.Int       `json:"variable_borrow_rate"`
	AvailableLiquidity             *big.Int       `json:"available_liquidity"`
	StableBorrowRate               *big.Int       `json:"stable_borrow_rate"`
	LiquidityRate                  *big.Int       `json:"liquidity_rate"`
	TotalPrincipalStableDebt       *big.Int       `json:"total_principal_stable_debt"`
	TotalScaledVariableDebt        *big.Int       `json:"total_scaled_variable_debt"`
	LastUpdateTimestamp            int64          `json:"last_update_timestamp"`
	PriceInMarketReferenceCurrency *big.Int       `json:"price_in_market_reference_currency"`
	InterestRateStrategyAddress    common.Address `json:"interest_rate_strategy_address"`
}

type FormattedReserveData struct {
	Id                          string          `json:"id"`
	UnderlyingAsset             common.Address  `json:"underlying_asset"`
	ZTokenAddress               common.Address  `json:"z_token_address"`
	VariableDebtTokenAddress    common.Address  `json:"variable_debt_token_address"`
	Symbol                      Symbol          `json:"symbol"`
	Decimals                    int64           `json:"decimals"`
	BorrowingEnabled            bool            `json:"borrowing_enabled"`
	IsActive                    bool            `json:"is_active"`
	IsFrozen                    bool            `json:"is_frozen"`
	UsageAsCollateralEnabled    bool            `json:"usage_as_collateral_enabled"`
	ReserveFactor               decimal.Decimal `json:"reserve_factor"`
	BaseLTVasCollateral         decimal.Decimal `json:"base_ltv_as_collateral"`
	ReserveLiquidationThreshold decimal.Decimal `json:"reserve_liquidation_threshold"`
	ReserveLiquidationBonus     decimal.Decimal `json:"reserve_liquidation_bonus"`
	UtilizationRate             decimal.Decimal `json:"utilization_rate"`
	TotalStableDebt             decimal.Decimal `json:"total_stable_debt"`
	TotalVariableDebt           decimal.Decimal `json:"total_variable_debt"`
	TotalDebt                   decimal.Decimal `json:"total_debt"`
	TotalLiquidity              decimal.Decimal `json:"total_liquidity"`
	AvailableLiquidity          decimal.Decimal `json:"available_liquidity"`
	SupplyAPY                   decimal.Decimal `json:"supply_apy"`
	SupplyAPR                   decimal.Decimal `json:"supply_apr"`
	VariableBorrowAPY           decimal.Decimal `json:"variable_borrow_apy"`
	VariableBorrowAPR           decimal.Decimal `json:"variable_borrow_apr"`
	StableBorrowAPY             decimal.Decimal `json:"stable_borrow_apy"`
	StableBorrowAPR             decimal.Decimal `json:"stable_borrow_apr"`
	Price                       decimal.Decimal `json:"price"`
}

func CalculateReserveVariableDebt(r RawReserveData) decimal.Decimal {
	lastUpdate := time.Unix(r.LastUpdateTimestamp, 0)
	mul := math.RayMul(r.TotalScaledVariableDebt, r.VariableBorrowIndex)
	compoundedInterest := math.CalculateCompoundedInterest(r.VariableBorrowRate, time.Since(lastUpdate))
	totalVariableDebt := math.RayMul(mul, compoundedInterest)
	return math.Normalize(totalVariableDebt, r.Decimals)
}

func CalculateReserveStableDebt(r RawReserveData) decimal.Decimal {
	lastUpdate := time.Unix(r.StableDebtLastUpdateTimestamp, 0)
	compoundedInterest := math.CalculateCompoundedInterest(r.StableBorrowRate, time.Since(lastUpdate))
	totalVariableDebt := math.RayMul(r.TotalPrincipalStableDebt, compoundedInterest)
	return math.Normalize(totalVariableDebt, r.Decimals)
}

func CalculateReserveTotalDebt(r RawReserveData) decimal.Decimal {
	variableDebt := CalculateReserveVariableDebt(r)
	stableDebt := CalculateReserveStableDebt(r)
	return variableDebt.Add(stableDebt)
}

func CalculateReserveAvailableLiquidity(r RawReserveData) decimal.Decimal {
	return math.Normalize(r.AvailableLiquidity, r.Decimals)
}

func CalculateReserveTotalLiquidity(r RawReserveData) decimal.Decimal {
	availableLiquidity := math.Normalize(r.AvailableLiquidity, r.Decimals)
	return CalculateReserveTotalDebt(r).Add(availableLiquidity)
}

func CalculateUtilizationRate(r RawReserveData) decimal.Decimal {
	totalLiquidity := CalculateReserveTotalLiquidity(r)

	if totalLiquidity.Equal(decimal.Zero) {
		return decimal.Zero
	}
	return CalculateReserveTotalDebt(r).Div(totalLiquidity)
}

func CalculateSupplyAPY(r RawReserveData) decimal.Decimal {
	div := new(big.Int).Div(r.LiquidityRate, math.SecondsPerYear)
	sum := new(big.Int).Add(div, math.Ray)
	pow := math.RayPow(sum, math.SecondsPerYear)
	apyRay := new(big.Int).Sub(pow, math.Ray)
	return math.Normalize(apyRay, int64(math.RayDecimals))
}

func CalculateSupplyAPR(r RawReserveData) decimal.Decimal {
	return math.Normalize(r.LiquidityRate, int64(math.RayDecimals))
}

func CalculateVariableBorrowAPY(r RawReserveData) decimal.Decimal {
	div := new(big.Int).Div(r.VariableBorrowRate, math.SecondsPerYear)
	sum := new(big.Int).Add(div, math.Ray)
	pow := math.RayPow(sum, math.SecondsPerYear)
	apyRay := new(big.Int).Sub(pow, math.Ray)
	return math.Normalize(apyRay, int64(math.RayDecimals))
}

func CalculateStableBorrowAPY(r RawReserveData) decimal.Decimal {
	div := new(big.Int).Div(r.StableBorrowRate, math.SecondsPerYear)
	sum := new(big.Int).Add(div, math.Ray)
	pow := math.RayPow(sum, math.SecondsPerYear)
	apyRay := new(big.Int).Sub(pow, math.Ray)
	return math.Normalize(apyRay, int64(math.RayDecimals))
}

func CalculateVariableBorrowAPR(r RawReserveData) decimal.Decimal {
	return math.Normalize(r.VariableBorrowRate, int64(math.RayDecimals))
}

func CalculateStableBorrowAPR(r RawReserveData) decimal.Decimal {
	return math.Normalize(r.StableBorrowRate, int64(math.RayDecimals))
}

func CalculatePrice(r RawReserveData, basePrice *big.Int) decimal.Decimal {
	decimals := int64(8)
	priceInMarketReferenceCurrency := new(big.Int).Mul(r.PriceInMarketReferenceCurrency, math.Pow10(decimals+1))
	realPrice := new(big.Int).Div(priceInMarketReferenceCurrency, basePrice)
	roundedUpPrice := new(big.Int).Div(new(big.Int).Add(realPrice, math.Five), math.Ten)
	return math.Normalize(roundedUpPrice, decimals)
}
