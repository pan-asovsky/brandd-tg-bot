package msg_fmt

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type AdminMessageFormatterService interface {
	NewBookingNotify(booking *entity.Booking) (string, error)
}
