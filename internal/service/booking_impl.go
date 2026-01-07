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
	pgProvider   *pg.Provider
	slotService  SlotService
	priceService PriceService
}

func (b *bookingService) Create(info *types.UserSessionInfo) (*model.Booking, error) {
	booking := &model.Booking{
		ChatID:     info.ChatID,
		Date:       info.Date,
		Service:    info.Service,
		RimRadius:  info.RimRadius,
		TotalPrice: sql.NullInt64{Int64: info.TotalPrice, Valid: true},
		Status:     model.Pending,
	}

	start, _ := b.parseTime(info.Time)
	booking.Time = start

	return b.pgProvider.Booking().Save(booking)
}

func (b *bookingService) Confirm(chatID int64) error {
	return b.pgProvider.Booking().Confirm(chatID)
}

func (b *bookingService) AutoConfirm(chatID int64) error {
	return b.pgProvider.Booking().AutoConfirm(chatID)
}

func (b *bookingService) SetPhone(phone string, chatID int64) error {
	return b.pgProvider.Booking().SetPhone(phone, chatID)
}

func (b *bookingService) FindActiveByChatID(chatID int64) (*model.Booking, error) {
	return b.pgProvider.Booking().FindActiveByChatID(chatID)
}

func (b *bookingService) ExistsByChatID(chatID int64) bool {
	return b.pgProvider.Booking().ExistsByChatID(chatID)
}

func (b *bookingService) UpdateStatus(chatID int64, status model.BookingStatus) error {
	return utils.WrapFunctionError(func() error {
		return b.pgProvider.Booking().UpdateStatus(chatID, status)
	})
}

func (b *bookingService) UpdateRimRadius(chatID int64, rimRadius string) error {
	return b.pgProvider.Booking().UpdateRimRadius(chatID, rimRadius)
}

func (b *bookingService) RecalculatePrice(chatID int64) error {
	booking, err := b.pgProvider.Booking().FindActiveByChatID(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	newPrice, err := b.priceService.Calculate(booking.Service, booking.RimRadius)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return b.pgProvider.Booking().UpdatePrice(chatID, newPrice)
	})
}

func (b *bookingService) Cancel(chatID int64) error {
	return b.pgProvider.Booking().Cancel(chatID)
}

func (b *bookingService) parseTime(time string) (start, end string) {
	split := strings.Split(time, "-")
	start = split[0]
	end = split[1]
	return start, end
}
