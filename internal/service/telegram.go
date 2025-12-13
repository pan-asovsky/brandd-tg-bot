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
	RequestPreConfirm(booking *model.Booking, chatID int64) error
	RequestUserPhone(info *types.UserSessionInfo) error

	ProcessPhone(booking *model.Booking, chatID int64) error
	ProcessConfirm(chatID int64, slot *model.Slot) error
	ProcessPendingConfirm(chatID int64) error
}
