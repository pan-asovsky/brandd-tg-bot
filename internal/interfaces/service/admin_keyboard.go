package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AdminKeyboardService interface {
	AdminGreetingKeyboard() tg.InlineKeyboardMarkup
}
