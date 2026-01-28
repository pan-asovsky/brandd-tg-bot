package interfaces

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
)

type BookingRepo interface {
	FindActiveNotPending(chatID int64) (*entity.Booking, error)
	FindActivePending(chatID int64) (*entity.Booking, error)
	FindAnyActive(chatID int64) (*entity.Booking, error)
	Exists(chatID int64) bool
	UpdateRimRadius(chatID int64, rimRadius string) error
	UpdateStatus(chatID int64, status entity.BookingStatus) error
	UpdatePrice(chatID int64, price int64) error
	Save(booking *entity.Booking) (*entity.Booking, error)
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
	AutoConfirm(chatID int64) error
	Cancel(chatID int64) error
	UpdateService(chatID int64, service string) error
	FindAllActive() ([]entity.Booking, error)
	Find(bookingID int64) (*entity.Booking, error)
	Close(info *model.BookingInfo) (*entity.Booking, error)
	ListByPeriod(period stat.Period) ([]entity.Booking, error)
}
