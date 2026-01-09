package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type SlotService interface {
	GetAvailableBookings() []model.AvailableBooking
	GetAvailableZones(date string) (model.Zone, error)
	FindByDateAndTime(date, start string) (*model.Slot, error)
	MarkUnavailable(date, startTime string) error
}
