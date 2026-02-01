package service

import (
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
)

type SlotService interface {
	GetAvailableDates() []entity.AvailableDate
	GetAvailableZones(date string) (entity.Zone, error)
	FindByDateTime(date time.Time, start string) (*entity.Slot, error)
	MarkUnavailable(date time.Time, startTime string) error
	FreeUp(date time.Time, startTime string) error
}
