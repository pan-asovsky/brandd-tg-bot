package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type TelegramService interface {
	ProcessMenu(bookings []AvailableBooking, info *types.UserSessionInfo) error
	ProcessDate(zone model.Zone, info *types.UserSessionInfo) error
	ProcessZone(timeslots []model.Timeslot, info *types.UserSessionInfo) error
	ProcessTime(types []model.ServiceType, info *types.UserSessionInfo) error
	ProcessServiceType(rims []string, info *types.UserSessionInfo) error
	ProcessRimRadius(info *types.UserSessionInfo) error
	ProcessPhone(booking *model.Booking, chatID int64) error
	ProcessConfirm(chatID int64, slot *model.Slot) error
	ProcessPendingConfirm(chatID int64) error
}
