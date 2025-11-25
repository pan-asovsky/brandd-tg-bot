package model

import "time"

type Price struct {
	ID            int64     `db:"id"`
	RimSize       int       `db:"rim_size"`
	ServiceTypeID int64     `db:"service_type_id"`
	PricePerWheel int64     `db:"price_per_wheel"`
	PricePerSet   int64     `db:"price_per_set"`
	IsActive      bool      `db:"is_active"`
	CreatedAt     time.Time `db:"created_at"`
}
