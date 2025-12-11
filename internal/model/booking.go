package model

import (
	"database/sql"
	"time"
)

type BookingStatus string

const (
	Pending      BookingStatus = "PENDING"
	NotConfirmed BookingStatus = "NOT_CONFIRMED"
	Confirmed    BookingStatus = "CONFIRMED"
	Completed    BookingStatus = "COMPLETED"
	Cancelled    BookingStatus = "CANCELLED"
	NoShow       BookingStatus = "NO_SHOW"
)

type Booking struct {
	ID         int64          `db:"id"`
	ChatID     int64          `db:"chat_id"`
	UserPhone  sql.NullString `db:"user_phone"`
	Date       string         `db:"slot_id"`
	Time       string         `db:"time"`
	Service    string         `db:"service"`
	RimRadius  string         `db:"rim_radius"`
	TotalPrice sql.NullInt64  `db:"total_price"`
	Status     BookingStatus  `db:"status"`
	IsActive   bool           `db:"is_active"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
}
