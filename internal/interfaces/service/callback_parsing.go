package service

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type CallbackParsingService interface {
	Parse(query *api.CallbackQuery) (*model.UserSessionInfo, error)
}
