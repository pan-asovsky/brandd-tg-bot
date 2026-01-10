package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

type TelegramService interface {
	RequestDate(bookings []entity.AvailableBooking, info *types.UserSessionInfo) error
	RequestZone(zone entity.Zone, info *types.UserSessionInfo) error
	RequestTime(timeslots []entity.Timeslot, info *types.UserSessionInfo) error
	RequestServiceTypes(types []entity.ServiceType, info *types.UserSessionInfo) error
	RequestRimRadius(rims []string, info *types.UserSessionInfo) error
	RequestPreConfirm(booking *entity.Booking, info *types.UserSessionInfo) error
	RequestUserPhone(info *types.UserSessionInfo) error

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
