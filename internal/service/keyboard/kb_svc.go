package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	slot "github.com/pan-asovsky/brandd-tg-bot/internal/service/slot"
)

type KeyboardService interface {
	GreetingKeyboard() tg.InlineKeyboardMarkup
	DateKeyboard([]slot.AvailableBooking) tg.InlineKeyboardMarkup
	ZoneKeyboard(zone model.Zone, date string) tg.InlineKeyboardMarkup
	TimeKeyboard([]model.Timeslot) tg.InlineKeyboardMarkup
}
