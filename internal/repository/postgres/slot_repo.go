package postgres

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type SlotRepository interface {
	IsTodayAvailable() bool
	GetAvailableSlots(date string) ([]model.Slot, error)
	FindByZone(zone model.Zone) []model.Slot
	FindByTimeslot(timeslot model.Timeslot) []model.Slot
}
