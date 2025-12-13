package service

import (
	"database/sql"
	"strings"

	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type bookingService struct {
	pgProvider  *pg.Provider
	slotService SlotService
}

func (b *bookingService) Create(info *types.UserSessionInfo) (*model.Booking, error) {
	booking := &model.Booking{
		ChatID:     info.ChatID,
		Date:       info.Date,
		Service:    info.Service,
		RimRadius:  info.Radius,
		TotalPrice: sql.NullInt64{Int64: info.TotalPrice, Valid: true},
		Status:     model.NotConfirmed,
	}

	start, end := b.parseTime(info.Time)
	if err := b.slotService.MarkUnavailable(info.Date, start, end); err != nil {
		return nil, utils.WrapError(err)
	}

	booking.Time = start

	return b.pgProvider.Booking().Save(booking)
}

func (b *bookingService) Confirm(chatID int64) error {
	return b.pgProvider.Booking().Confirm(chatID)
}

func (b *bookingService) SetPhone(phone string, chatID int64) error {
	return b.pgProvider.Booking().SetPhone(phone, chatID)
}

func (b *bookingService) FindActiveByChatID(chatID int64) (*model.Booking, error) {
	return b.pgProvider.Booking().FindActiveByChatID(chatID)
}

func (b *bookingService) parseTime(time string) (start, end string) {
	split := strings.Split(time, "-")
	start = split[0]
	end = split[1]
	return start, end
}
