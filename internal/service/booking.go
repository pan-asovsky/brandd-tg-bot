package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

type BookingService interface {
	Create(data *types.UserSessionInfo) error
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
}
