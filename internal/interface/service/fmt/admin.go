package fmt

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type AdminMessageFormatterService interface {
	BookingCreated(booking *entity.Booking) (string, error)
	BookingCancelled(booking *entity.Booking) (string, error)
	BookingCompleted(booking *entity.Booking) (string, error)
}
