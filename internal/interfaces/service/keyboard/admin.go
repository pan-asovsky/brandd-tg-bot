package keyboard

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AdminKeyboardService interface {
	ChoiceFlowKeyboard() tgapi.InlineKeyboardMarkup
	MainMenu() tgapi.InlineKeyboardMarkup
	Bookings() tgapi.InlineKeyboardMarkup
	Statistics() tgapi.InlineKeyboardMarkup
	Settings() tgapi.InlineKeyboardMarkup
}
