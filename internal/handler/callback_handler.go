package handler

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CallbackHandler interface {
	Handle(query *tgbot.CallbackQuery) error
}
