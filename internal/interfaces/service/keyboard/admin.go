package keyboard

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type AdminKeyboardService interface {
	ChoiceFlowKeyboard() tgapi.InlineKeyboardMarkup
	MainMenu() tgapi.InlineKeyboardMarkup
	Bookings(bookings []entity.Booking) tgapi.InlineKeyboardMarkup
	Statistics() tgapi.InlineKeyboardMarkup
	Settings() tgapi.InlineKeyboardMarkup
	BookingInfo(userChatID int64, bookingID int64) tgapi.InlineKeyboardMarkup
	ConfirmationKeyboard(info *model.BookingInfo) tgapi.InlineKeyboardMarkup
}
