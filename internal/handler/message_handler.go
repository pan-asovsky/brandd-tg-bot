package handler

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler interface {
	Handle(msg *tgbot.Message) error
}
