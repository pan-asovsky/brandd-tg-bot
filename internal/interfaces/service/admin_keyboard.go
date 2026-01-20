package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AdminKeyboardService interface {
	ChoiceFlowKeyboard() tg.InlineKeyboardMarkup
	MainMenu() tg.InlineKeyboardMarkup
	Bookings() tg.InlineKeyboardMarkup
	Statistics() tg.InlineKeyboardMarkup
	Settings() tg.InlineKeyboardMarkup
}
