package repository

import (
	"database/sql"

	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interface/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type configRepo struct {
	db *sql.DB
}

func NewConfigRepo(db *sql.DB) irepo.ConfigRepo {
	return &configRepo{db: db}
}

func (c configRepo) IsAutoConfirm() (bool, error) {
	var autoConfirm bool
	if err := c.db.QueryRow(IsAutoConfirm).Scan(&autoConfirm); err != nil {
		return false, utils.WrapError(err)
	}

	return autoConfirm, nil
}
