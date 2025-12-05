package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type TelegramService interface {
	ProcessMenu(bookings []AvailableBooking, chatID int64)
	ProcessDate(zone model.Zone, info *types.UserSessionInfo)
	ProcessZone(timeslots []model.Timeslot, info *types.UserSessionInfo)
	ProcessTime()
	ProcessServiceType()
	ProcessRimRadius()
	ProcessPhone()
}
