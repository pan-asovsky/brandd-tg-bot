package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type TelegramService interface {
	RequestDate(bookings []AvailableBooking, info *types.UserSessionInfo) error
	RequestZone(zone model.Zone, info *types.UserSessionInfo) error
	RequestTime(timeslots []model.Timeslot, info *types.UserSessionInfo) error
	RequestServiceTypes(types []model.ServiceType, info *types.UserSessionInfo) error
	RequestRimRadius(rims []string, info *types.UserSessionInfo) error
	RequestPreConfirm(booking *model.Booking, info *types.UserSessionInfo) error
	RequestUserPhone(info *types.UserSessionInfo) error

	RemoveReplyKeyboard(chatID int64) error

	ProcessConfirm(chatID int64, slot *model.Slot) error
	ProcessPendingConfirm(chatID int64) error

	SendHelpMessage(chatID int64) error
	SendBookingRestrictionMessage(chatID int64, booking *model.Booking) error
	SendMyBookingsMessage(chatID int64, fn func() (*model.Booking, error)) error
	SendStartMenu(chatID int64) error

	SendCalendar(chatID int64, fn func() (*model.Booking, error)) error

	SendPreCancelBookingMessage(chatID int64, date, time string) error
	SendCancellationMessage(chatID int64) error
	SendCancelDenyMessage(chatID int64) error
}
