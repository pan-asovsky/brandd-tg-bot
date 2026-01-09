package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type BookingRepo interface {
	FindActiveByChatID(telegramID int64) (*model.Booking, error)
	ExistsByChatID(id int64) bool
	UpdateRimRadius(chatID int64, rimRadius string) error
	UpdateStatus(chatID int64, status model.BookingStatus) error
	UpdatePrice(chatID int64, price int64) error
	Save(booking *model.Booking) (*model.Booking, error)
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
	AutoConfirm(chatID int64) error
	Cancel(chatID int64) error
	UpdateService(chatID int64, service string) error
}
