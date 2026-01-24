package fmt

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type BookingMessageFormatterService interface {
	Confirm(date, startTime string) string
	PreConfirm(booking *entity.Booking) (string, error)
	My(booking *entity.Booking) (string, error)
	Restriction(booking *entity.Booking) (string, error)
	PreCancel(date, time string) (string, error)
	BookingPreview(booking *entity.Booking) (string, error)
}
