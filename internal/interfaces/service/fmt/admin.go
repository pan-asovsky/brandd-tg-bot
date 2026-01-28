package fmt

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
)

type AdminMessageFormatterService interface {
	BookingCreated(booking *entity.Booking) (string, error)
	BookingCancelled(booking *entity.Booking) (string, error)
	BookingCompleted(booking *entity.Booking) (string, error)
	Statistics(stats stat.Stats, p stat.Period) string
}
