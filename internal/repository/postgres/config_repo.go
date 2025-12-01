package postgres

import (
	"database/sql"
	"fmt"
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
		return false, fmt.Errorf("[is_auto_confirm] error: %+v", err)
	}
	return autoConfirm, nil
}
