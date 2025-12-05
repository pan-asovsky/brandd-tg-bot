package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type KeyboardService interface {
	GreetingKeyboard() tg.InlineKeyboardMarkup
	DateKeyboard([]AvailableBooking) tg.InlineKeyboardMarkup
	ZoneKeyboard(zone model.Zone, date string) tg.InlineKeyboardMarkup
	TimeKeyboard(ts []model.Timeslot, zone, date string) tg.InlineKeyboardMarkup
	ServiceKeyboard(types []model.ServiceType, time, date string) tg.InlineKeyboardMarkup
	RimsKeyboard(rims []string, svc, t, d string) tg.InlineKeyboardMarkup
	ConfirmKeyboard() tg.InlineKeyboardMarkup
	RequestPhoneKeyboard() tg.ReplyKeyboardMarkup
}
