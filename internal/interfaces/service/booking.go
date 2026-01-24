package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type BookingService interface {
	Create(data *model.UserSessionInfo) (*entity.Booking, error)
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
	AutoConfirm(chatID int64) error
	FindActiveNotPending(chatID int64) (*entity.Booking, error)
	FindPending(chatID int64) (*entity.Booking, error)
	CancelOldIfExists(chatID int64) error
	UpdateStatus(chatID int64, status entity.BookingStatus) error
	ExistsByChatID(chatID int64) bool
	UpdateRimRadius(chatID int64, rimRadius string) error
	UpdateService(chatID int64, service string) error
	RecalculatePrice(chatID int64) error
	Cancel(chatID int64) error
	FindAllActive() ([]entity.Booking, error)
	Find(bookingID int64) (*entity.Booking, error)
}
