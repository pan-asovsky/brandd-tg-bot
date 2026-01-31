package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type priceRepo struct {
	pool *pgxpool.Pool
}

func NewPriceRepo(p *pgxpool.Pool) irepo.PriceRepo {
	return &priceRepo{pool: p}
}

func (pr *priceRepo) GetAllRimSizes() ([]string, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	rows, err := pr.pool.Query(ctx, GetAllRimSizes)
	if err != nil {
		return nil, fmt.Errorf("[get_all_rim_sizes] query error: %w", err)
	}
	defer rows.Close()

	var rimSizes []string
	for rows.Next() {
		var rimSize string
		if err = rows.Scan(&rimSize); err != nil {
			return nil, fmt.Errorf("[get_all_rim_sizes] scan error: %w", err)
		}
		rimSizes = append(rimSizes, rimSize)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[get_all_rim_sizes] rows error: %w", err)
	}

	return rimSizes, nil
}

func (pr *priceRepo) GetSetPrice(svc string, radius string) (int64, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	var price entity.Price
	if err := pr.pool.QueryRow(ctx, GetPricePerSet, svc, radius).Scan(&price.PricePerSet); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("[get_set_price] not founded for %s %s %w", svc, radius, err)
		}
		return 0, utils.WrapError(err)
	}

	return price.PricePerSet, nil
}
