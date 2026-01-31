package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type configRepo struct {
	pool *pgxpool.Pool
}

func NewConfigRepo(p *pgxpool.Pool) irepo.ConfigRepo {
	return &configRepo{pool: p}
}

func (c configRepo) IsAutoConfirm() (bool, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	var autoConfirm bool
	if err := c.pool.QueryRow(ctx, IsAutoConfirm).Scan(&autoConfirm); err != nil {
		return false, utils.WrapError(err)
	}

	return autoConfirm, nil
}
