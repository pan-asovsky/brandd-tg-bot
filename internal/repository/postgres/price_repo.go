package postgres

import (
	"database/sql"
	"log"
)

type PriceRepo interface {
	GetAllRimSizes() []string
}

type priceRepo struct {
	db *sql.DB
}

func NewPriceRepo(db *sql.DB) PriceRepo {
	return &priceRepo{db: db}
}

func (pr priceRepo) GetAllRimSizes() []string {
	rows, err := pr.db.Query(GetAllRimSizes)
	if err != nil {
		log.Printf("[get_all_rim_sizes] error: %v", err)
		return nil
	}
	defer rows.Close()

	var rimSizes []string
	for rows.Next() {
		var rimSize string
		if err := rows.Scan(&rimSize); err != nil {
			log.Printf("[get_all_rim_sizes] scan error: %v", err)
			return nil
		}
		rimSizes = append(rimSizes, rimSize)
	}

	if err := rows.Err(); err != nil {
		log.Printf("[get_all_rim_sizes] rows error: %v", err)
		return nil
	}

	return rimSizes
}
