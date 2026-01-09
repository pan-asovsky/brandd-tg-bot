package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type BookingMessageFormattingService interface {
	Confirm(date, startTime string) string
	PreConfirm(booking *model.Booking) (string, error)
	My(booking *model.Booking) (string, error)
	Restriction(booking *model.Booking) (string, error)
	PreCancel(date, time string) (string, error)
}
