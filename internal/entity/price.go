package entity

import "time"

type Price struct {
	ID              int64     `db:"id"`
	RimSize         int       `db:"rim_size"`
	ServiceTypeCode string    `db:"service_type_id"`
	PricePerWheel   int64     `db:"price_per_wheel"`
	PricePerSet     int64     `db:"price_per_set"`
	IsActive        bool      `db:"is_active"`
	CreatedAt       time.Time `db:"created_at"`
}
