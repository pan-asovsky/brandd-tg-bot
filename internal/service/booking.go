package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type BookingService interface {
	Create(data *types.UserSessionInfo) (*model.Booking, error)
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
	FindActiveByChatID(chatID int64) (*model.Booking, error)
}
