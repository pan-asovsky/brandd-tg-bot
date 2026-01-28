package tg

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
)

type TelegramAdminService interface {
	ChoiceMenu(chatID int64) error
	StartMenu(chatID int64) error
	BookingPreview(chatID int64, booking *entity.Booking) error
	ConfirmAction(chatID int64, bookingInfo *model.BookingInfo) error
	RejectAction(chatID int64, backDirection string) error
	NoActiveBookings(chatID int64) error
	Statistics(chatID int64, stats stat.Stats, p stat.Period) error
}
