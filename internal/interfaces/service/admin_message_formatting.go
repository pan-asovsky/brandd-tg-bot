package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type AdminMessageFormattingService interface {
	NewBookingNotify(booking *model.Booking) (string, error)
}
