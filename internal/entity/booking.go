package entity

import (
	"database/sql"
	"time"
)

type AvailableDate struct {
	Date  time.Time
	Label string
}
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
	ID          int64          `db:"id"`
	ChatID      int64          `db:"chat_id"`
	UserPhone   sql.NullString `db:"user_phone"`
	Date        time.Time      `db:"date"`
	Time        string         `db:"time"`
	Service     string         `db:"service"`
	RimRadius   string         `db:"rim_radius"`
	TotalPrice  sql.NullInt64  `db:"total_price"`
	Status      BookingStatus  `db:"status"`
	ConfirmedBy sql.NullString `db:"confirmed_by"`
	CancelledBy sql.NullString `db:"cancelled_by"`
	Notes       sql.NullString `db:"notes"`
	Active      bool           `db:"is_active"`
	CreatedAt   time.Time      `db:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at"`
}

func (b *Booking) IsActive() bool {
	return b.Status == Confirmed
}

func (b *Booking) IsCancelled() bool {
	return b.Status == Cancelled
}

func (b *Booking) IsCompleted() bool {
	return b.Status == Completed
}

func (b *Booking) IsNoShow() bool {
	return b.Status == NoShow
}

func (b *Booking) IsPending() bool {
	return b.Status == Pending || b.Status == NotConfirmed
}
