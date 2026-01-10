package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type TelegramService interface {
	RequestDate(bookings []entity.AvailableBooking, info *model.UserSessionInfo) error
	RequestZone(zone entity.Zone, info *model.UserSessionInfo) error
	RequestTime(timeslots []entity.Timeslot, info *model.UserSessionInfo) error
	RequestServiceTypes(types []entity.ServiceType, info *model.UserSessionInfo) error
	RequestRimRadius(rims []string, info *model.UserSessionInfo) error
	RequestPreConfirm(booking *entity.Booking, info *model.UserSessionInfo) error
	RequestUserPhone(info *model.UserSessionInfo) error

	RemoveReplyKeyboard(chatID int64) error

	ProcessConfirm(chatID int64, slot *entity.Slot) error
	ProcessPendingConfirm(chatID int64) error

	SendBookingRestrictionMessage(chatID int64, booking *entity.Booking) error
	SendMyBookingsMessage(chatID int64, fn func() (*entity.Booking, error)) error
	SendStartMenu(chatID int64) error

	SendPreCancelBookingMessage(chatID int64, date, time string) error
	SendCancellationMessage(chatID int64) error
	SendCancelDenyMessage(chatID int64) error

	NewBookingNotify(chatID int64, booking *entity.Booking) error
}
