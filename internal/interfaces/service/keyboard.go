package service

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type KeyboardService interface {
	GreetingKeyboard() tg.InlineKeyboardMarkup
	DateKeyboard([]entity.AvailableBooking) tg.InlineKeyboardMarkup
	ZoneKeyboard(zone entity.Zone, date string) tg.InlineKeyboardMarkup
	TimeKeyboard(ts []entity.Timeslot, info *model.UserSessionInfo) tg.InlineKeyboardMarkup
	ServiceKeyboard(types []entity.ServiceType, info *model.UserSessionInfo) tg.InlineKeyboardMarkup
	RimsKeyboard(rims []string, info *model.UserSessionInfo) tg.InlineKeyboardMarkup
	ConfirmKeyboard(info *model.UserSessionInfo) tg.InlineKeyboardMarkup
	RequestPhoneKeyboard() tg.ReplyKeyboardMarkup
	EmptyMyBookingsKeyboard() tg.InlineKeyboardMarkup
	ExistsMyBookingsKeyboard() tg.InlineKeyboardMarkup
	BackKeyboard() tg.InlineKeyboardMarkup
	BookingCancellationKeyboard() tg.InlineKeyboardMarkup
}
