package postgres

import (
	"database/sql"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type ConfigRepo interface {
	IsAutoConfirm() (bool, error)
}

type configRepo struct {
	db *sql.DB
}

func (c configRepo) IsAutoConfirm() (bool, error) {
	var autoConfirm bool
	if err := c.db.QueryRow(IsAutoConfirm).Scan(&autoConfirm); err != nil {
		return false, utils.WrapError(err)
	}

	return autoConfirm, nil
}
