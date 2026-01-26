package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type SlotService interface {
	GetAvailableDates() []entity.AvailableDate
	GetAvailableZones(date string) (entity.Zone, error)
	FindByDateTime(date, start string) (*entity.Slot, error)
	MarkUnavailable(date, startTime string) error
	FreeUp(date, startTime string) error
}
