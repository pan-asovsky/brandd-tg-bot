package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type AdminMessageFormattingService interface {
	NewBookingNotify(booking *entity.Booking) (string, error)
}
