package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type BookingService interface {
	Create(data *types.UserSessionInfo) (*model.Booking, error)
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
	AutoConfirm(chatID int64) error
	FindActiveByChatID(chatID int64) (*model.Booking, error)
	UpdateStatus(chatID int64, status model.BookingStatus) error
	ExistsByChatID(chatID int64) bool
	UpdateRimRadius(chatID int64, rimRadius string) error
	UpdateService(chatID int64, service string) error
	RecalculatePrice(chatID int64) error
	Cancel(chatID int64) error
}
