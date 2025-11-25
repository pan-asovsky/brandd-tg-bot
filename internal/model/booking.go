package model

import "time"

type BookingStatus string

const (
	Pending   BookingStatus = "PENDING"
	Confirmed BookingStatus = "CONFIRMED"
	Completed BookingStatus = "COMPLETED"
	Cancelled BookingStatus = "CANCELLED"
	NoShow    BookingStatus = "NO_SHOW"
)

type Booking struct {
	ID            int64         `db:"id"`
	TelegramID    int64         `db:"telegram_id"`
	UserName      string        `db:"user_name"`
	UserPhone     string        `db:"user_phone"`
	SlotID        int64         `db:"slot_id"`
	ServiceTypeID int64         `db:"service_type_id"`
	RimSize       int           `db:"rim_size"`
	WheelCount    int64         `db:"wheel_count"`
	TotalPrice    int64         `db:"total_price"`
	Status        BookingStatus `db:"status"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`

	Slot        *Slot        `db:"-"`
	ServiceType *ServiceType `db:"-"`
}
