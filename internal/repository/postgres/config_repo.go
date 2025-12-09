package postgres

import (
	"database/sql"
	"log"
)

type ConfigRepo interface {
	IsAutoConfirm() bool
}

type configRepo struct {
	db *sql.DB
}

func (c configRepo) IsAutoConfirm() bool {
	var autoConfirm bool
	if err := c.db.QueryRow(IsAutoConfirm).Scan(&autoConfirm); err != nil {
		log.Fatalf("[is_auto_confirm] error: %v", err)
	}
	return autoConfirm
}
