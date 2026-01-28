package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type SlotRepo interface {
	IsTodayAvailable() bool
	GetAvailableSlots(date string) ([]entity.Slot, error)
	FindByDateAndTime(date, start string) (*entity.Slot, error)
	MarkUnavailable(date, startTime string) error
	FreeUp(date, startTime string) error
}
