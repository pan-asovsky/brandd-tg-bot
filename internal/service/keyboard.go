package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type KeyboardService interface {
	GreetingKeyboard() tg.InlineKeyboardMarkup
	DateKeyboard([]AvailableBooking) tg.InlineKeyboardMarkup
	ZoneKeyboard(zone model.Zone, date string) tg.InlineKeyboardMarkup
	TimeKeyboard(ts []model.Timeslot, info *types.UserSessionInfo) tg.InlineKeyboardMarkup
	ServiceKeyboard(types []model.ServiceType, time, date string) tg.InlineKeyboardMarkup
	ServiceKeyboardV2(types []model.ServiceType, info *types.UserSessionInfo) tg.InlineKeyboardMarkup
	RimsKeyboard(rims []string, svc, time, date string) tg.InlineKeyboardMarkup
	ConfirmKeyboard() tg.InlineKeyboardMarkup
	RequestPhoneKeyboard() tg.ReplyKeyboardMarkup
}
