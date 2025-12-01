package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackHandler interface {
	Handle(query *tgbotapi.CallbackQuery) error
}
