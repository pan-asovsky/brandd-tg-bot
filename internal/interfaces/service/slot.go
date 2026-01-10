package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type SlotService interface {
	GetAvailableBookings() []entity.AvailableBooking
	GetAvailableZones(date string) (entity.Zone, error)
	FindByDateAndTime(date, start string) (*entity.Slot, error)
	MarkUnavailable(date, startTime string) error
	FreeUp(date, startTime string) error
}
