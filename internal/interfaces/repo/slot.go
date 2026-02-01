package interfaces

import (
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
)

type SlotRepo interface {
	IsTodayAvailable() bool
	GetAvailableSlots(date string) ([]entity.Slot, error)
	FindByDateAndTime(date time.Time, start string) (*entity.Slot, error)
	MarkUnavailable(date time.Time, startTime string) error
	FreeUp(date time.Time, startTime string) error
}
