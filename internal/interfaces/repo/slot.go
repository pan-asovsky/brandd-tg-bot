package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type SlotRepo interface {
	IsTodayAvailable() bool
	GetAvailableSlots(date string) ([]model.Slot, error)
	FindByDateAndTime(date, start string) (*model.Slot, error)
	MarkUnavailable(date, startTime string) error
}
