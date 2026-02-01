package fmt

import (
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
)

type UserMessageFormatterService interface {
	Confirm(date, startTime string) string
	PreConfirm(booking *entity.Booking) (string, error)
	My(booking *entity.Booking) (string, error)
	Restriction(booking *entity.Booking) (string, error)
	PreCancel(date time.Time, time string) (string, error)
	BookingPreview(booking *entity.Booking) (string, error)
}
