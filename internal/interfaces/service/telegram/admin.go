package tg

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type TelegramAdminService interface {
	ChoiceMenu(chatID int64) error
	StartMenu(chatID int64) error
	BookingPreview(chatID int64, booking *entity.Booking) error
	ConfirmNoShow(chatID int64, bookingInfo *model.BookingInfo) error
}
