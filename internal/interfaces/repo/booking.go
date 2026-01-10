package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type BookingRepo interface {
	FindActiveNotPending(chatID int64) (*entity.Booking, error)
	FindPending(chatID int64) (*entity.Booking, error)
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
}
