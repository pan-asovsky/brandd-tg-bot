package callback

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type AdminCallbackBuilderService interface {
	StartUser() string
	StartAdmin() string
	BookingsMenu() string
	Booking(bookingID int64) string
	Statistics() string
	Settings() string
	Back(direction string) string
	Chat(userChatID int64) string
	PreComplete(userChatID int64, bookingID int64) string
	PreNoShow(userChatID int64, bookingID int64) string
	Confirm(info *model.BookingInfo) string
	Reject(info *model.BookingInfo) string
}
