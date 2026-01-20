package keyboard

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type UserKeyboardService interface {
	GreetingKeyboard() tgapi.InlineKeyboardMarkup
	DateKeyboard([]entity.AvailableDate) tgapi.InlineKeyboardMarkup
	ZoneKeyboard(zone entity.Zone, date string) tgapi.InlineKeyboardMarkup
	TimeKeyboard(ts []entity.Timeslot, info *model.UserSessionInfo) tgapi.InlineKeyboardMarkup
	ServiceKeyboard(types []entity.ServiceType, info *model.UserSessionInfo) tgapi.InlineKeyboardMarkup
	RimsKeyboard(rims []string, info *model.UserSessionInfo) tgapi.InlineKeyboardMarkup
	ConfirmKeyboard(info *model.UserSessionInfo) tgapi.InlineKeyboardMarkup
	RequestPhoneKeyboard() tgapi.ReplyKeyboardMarkup
	EmptyMyBookingsKeyboard() tgapi.InlineKeyboardMarkup
	ExistsMyBookingsKeyboard() tgapi.InlineKeyboardMarkup
	BackKeyboard() tgapi.InlineKeyboardMarkup
	BookingCancellationKeyboard() tgapi.InlineKeyboardMarkup
}
