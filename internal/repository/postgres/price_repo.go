package postgres

import (
	"database/sql"
	"fmt"
)

type PriceRepo interface {
	GetAllRimSizes() ([]string, error)
}

type priceRepo struct {
	db *sql.DB
}

func (pr priceRepo) GetAllRimSizes() ([]string, error) {
	rows, err := pr.db.Query(GetAllRimSizes)
	if err != nil {
		return nil, fmt.Errorf("[get_all_rim_sizes] query error: %w", err)
	}
	defer rows.Close()

	var rimSizes []string
	for rows.Next() {
		var rimSize string
		if err := rows.Scan(&rimSize); err != nil {
			return nil, fmt.Errorf("[get_all_rim_sizes] scan error: %w", err)
		}
		rimSizes = append(rimSizes, rimSize)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[get_all_rim_sizes] rows error: %w", err)
	}

	return rimSizes, nil
}
