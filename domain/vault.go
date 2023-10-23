package domain

import (
	"encoding/json"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

type Vault struct {
	Id                     int64           `json:"id"`
	Owner                  common.Address  `json:"owner"`
	Urn                    common.Address  `json:"urn"`
	LiquidationPrice       decimal.Decimal `json:"liquidation_price"`
	CollateralLocked       decimal.Decimal `json:"collateral_locked"`
	CollateralizationRatio decimal.Decimal `json:"collateralization_ratio"`
	Debt                   decimal.Decimal `json:"debt"`
	AvailableToWithdraw    decimal.Decimal `json:"available_to_withdraw"`
	AvailableToMint        decimal.Decimal `json:"available_to_mint"`
	IlkName                string          `json:"-"`
}

type Ilk struct {
	Name                          string          `json:"name"`
	Symbol                        Symbol          `json:"symbol"`
	MinimumCollateralizationRatio decimal.Decimal `json:"minimum_collateralization_ratio"`
	DebtCeiling                   decimal.Decimal `json:"debt_ceiling"`
	Debt                          decimal.Decimal `json:"debt"`
	AnnualStabilityFee            decimal.Decimal `json:"stability_fee"`
	DustLimit                     decimal.Decimal `json:"dust_limit"`
	Price                         decimal.Decimal `json:"price"`
	LiquidationPenalty            decimal.Decimal `json:"liquidation_penalty"`
	Hole                          decimal.Decimal `json:"hole"`
	Dirt                          decimal.Decimal `json:"dirt"`
	Rate                          *big.Int        `json:"rate"`
	Join                          common.Address  `json:"join"`
	Median                        common.Address  `json:"median"`
	Gem                           common.Address  `json:"gem"`
	Clipper                       common.Address  `json:"clipper"`
	Pip                           common.Address  `json:"pip"`
}

type VaultEventType string

var (
	VaultRepayEvent       VaultEventType = "Vault_Repay"
	VaultDepositEvent     VaultEventType = "Vault_Deposit"
	VaultWithdrawEvent    VaultEventType = "Vault_Withdraw"
	VaultMintEvent        VaultEventType = "Vault_Mint"
	VaultLiquidationEvent VaultEventType = "Vault_Liquidation"
)

type VaultEvent struct {
	Type            VaultEventType   `json:"type"`
	TxHash          common.Hash      `json:"tx_hash"`
	Timestamp       time.Time        `json:"timestamp"`
	DeltaCollateral *decimal.Decimal `json:"delta_collateral,omitempty"`
	DeltaDebt       *decimal.Decimal `json:"delta_debt,omitempty"`
	IlkName         string           `json:"ilkName"`
}

type Currency struct {
	symbol Symbol
	values map[Symbol]decimal.Decimal
}

func NewCurrency(symbol Symbol, value decimal.Decimal) *Currency {
	return &Currency{symbol: symbol, values: map[Symbol]decimal.Decimal{symbol: value}}
}

func NewCurrencyFromJson(data []byte) *Currency {
	var v *Currency
	_ = json.Unmarshal(data, &v)
	return v
}

func (c *Currency) SetValueInSymbol(symbol Symbol, value decimal.Decimal) *Currency {
	c.values[symbol] = value
	return c
}

func (c *Currency) ChangeDefaultSymbol(symbol Symbol) {
	c.symbol = symbol
}

func (c *Currency) Value() decimal.Decimal {
	return c.values[c.symbol]
}

func (c *Currency) Symbol() Symbol {
	return c.symbol
}

func (c *Currency) ValueInSymbol(symbol Symbol) decimal.Decimal {
	return c.values[symbol]
}

func (c *Currency) ChangeValue(f func(current decimal.Decimal) decimal.Decimal) {
	c.values[c.Symbol()] = f(c.Value())
}

func (c *Currency) String() string {
	data, _ := c.MarshalJSON()
	return string(data)
}

func (c *Currency) MarshalJSON() ([]byte, error) {
	jsn := map[string]interface{}{}
	for k, v := range c.values {
		jsn[string(k)] = v
	}
	return json.Marshal(&jsn)
}

func (c *Currency) UnmarshalJSON(body []byte) (err error) {
	jsn := map[string]interface{}{}
	err = json.Unmarshal(body, &jsn)
	if err != nil {
		return err
	}
	c.values = map[Symbol]decimal.Decimal{}
	for k, v := range jsn {
		std, err := decimal.NewFromString(v.(string))
		if err != nil {
			return err
		}
		c.values[Symbol(k)] = std
		c.ChangeDefaultSymbol(Symbol(k))
	}
	return nil
}

func Bytes32ToString(id [32]byte) string {
	return string(removeZeroBytes(id[:]))
}

func removeZeroBytes(arr []byte) []byte {
	var res []byte
	for _, elem := range arr {
		if elem != 0 {
			res = append(res, elem)
		}
	}
	return res
}
