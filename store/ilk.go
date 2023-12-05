package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	"github.com/zarbanio/auction-keeper/domain"
	"github.com/zarbanio/auction-keeper/domain/math"
	"github.com/zarbanio/auction-keeper/x/decimal"
)

type ilkModel struct {
	name                          string
	symbol                        string
	minimumCollateralizationRatio string
	debtCeiling                   string
	debt                          string
	annualStabilityFee            string
	dustLimit                     string
	price                         string
	rate                          string
	join                          string
	median                        string
	gem                           string
	pip                           string
}

func (i *ilkModel) toDomain() *domain.Ilk {
	return &domain.Ilk{
		Name:                          i.name,
		Symbol:                        domain.Symbol(i.symbol),
		MinimumCollateralizationRatio: decimal.MustNewFromString(i.minimumCollateralizationRatio),
		DebtCeiling:                   decimal.MustNewFromString(i.debtCeiling),
		Debt:                          decimal.MustNewFromString(i.debt),
		AnnualStabilityFee:            decimal.MustNewFromString(i.annualStabilityFee),
		DustLimit:                     decimal.MustNewFromString(i.dustLimit),
		Price:                         decimal.MustNewFromString(i.price),
		Rate:                          math.BigIntFromString(i.rate),
		Join:                          common.HexToAddress(i.join),
		Median:                        common.HexToAddress(i.median),
		Gem:                           common.HexToAddress(i.gem),
		Pip:                           common.HexToAddress(i.pip),
	}
}

func (p postgres) CreateOrUpdateIlk(ctx context.Context, ilk domain.Ilk) error {
	query := `
		INSERT INTO ilks (name, symbol, minimum_collateralization_ratio, debt_ceiling, debt, annual_stability_fee, dust_limit, price, rate, "join", median, gem, pip)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		ON CONFLICT (name) DO UPDATE
		SET symbol = $2, 
		minimum_collateralization_ratio = $3, 
		debt_ceiling = $4, 
		debt = $5, 
		annual_stability_fee = $6, 
		dust_limit = $7, 
		price = $8, 
		rate = $9, 
		"join" = $10,
		median = $11,
		gem = $12,
		pip = $13
	`

	_, err := p.conn.Exec(ctx, query,
		ilk.Name,
		ilk.Symbol,
		ilk.MinimumCollateralizationRatio,
		ilk.DebtCeiling,
		ilk.Debt,
		ilk.AnnualStabilityFee,
		ilk.DustLimit,
		ilk.Price,
		ilk.Rate.String(),
		ilk.Join.Hex(),
		ilk.Median.Hex(),
		ilk.Gem.Hex(),
		ilk.Pip.Hex(),
	)
	if err != nil {
		return fmt.Errorf("error creating or updating ilk. %w", err)
	}
	return nil
}

func (p postgres) GetIlkByName(ctx context.Context, name string) (*domain.Ilk, error) {
	query := `SELECT * FROM ilks WHERE name=$1`
	row := p.conn.QueryRow(ctx, query, name)
	var ilk ilkModel
	err := row.Scan(
		&ilk.name,
		&ilk.symbol,
		&ilk.minimumCollateralizationRatio,
		&ilk.debtCeiling,
		&ilk.debt,
		&ilk.annualStabilityFee,
		&ilk.dustLimit,
		&ilk.price,
		&ilk.rate,
		&ilk.join,
		&ilk.median,
		&ilk.gem,
		&ilk.pip,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrIlkNotFound
		}
		return nil, fmt.Errorf("error getting ilk by name. %w", err)
	}
	return ilk.toDomain(), nil
}
