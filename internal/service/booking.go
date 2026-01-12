package service

import (
	"database/sql"
	"strings"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type bookingService struct {
	repoProvider p.RepoProvider
	slotService  i.SlotService
	priceService i.PriceService
}

func (b *bookingService) Create(info *model.UserSessionInfo) (*entity.Booking, error) {
	booking := &entity.Booking{
		ChatID:     info.ChatID,
		Date:       info.Date,
		Service:    info.Service,
		RimRadius:  info.RimRadius,
		TotalPrice: sql.NullInt64{Int64: info.TotalPrice, Valid: true},
		Status:     entity.Pending,
	}

	//todo: parse func!!!
	start, _ := b.parseTime(info.Time)
	booking.Time = start

	return b.repoProvider.Booking().Save(booking)
}

func (b *bookingService) Confirm(chatID int64) error {
	return b.repoProvider.Booking().Confirm(chatID)
}

func (b *bookingService) AutoConfirm(chatID int64) error {
	return b.repoProvider.Booking().AutoConfirm(chatID)
}

func (b *bookingService) SetPhone(phone string, chatID int64) error {
	return b.repoProvider.Booking().SetPhone(phone, chatID)
}

func (b *bookingService) FindActiveNotPending(chatID int64) (*entity.Booking, error) {
	return b.repoProvider.Booking().FindActiveNotPending(chatID)
}

func (b *bookingService) FindPending(chatID int64) (*entity.Booking, error) {
	return b.repoProvider.Booking().FindPending(chatID)
}

func (b *bookingService) ExistsByChatID(chatID int64) bool {
	return b.repoProvider.Booking().Exists(chatID)
}

func (b *bookingService) UpdateStatus(chatID int64, status entity.BookingStatus) error {
	return utils.WrapFunctionError(func() error {
		return b.repoProvider.Booking().UpdateStatus(chatID, status)
	})
}

func (b *bookingService) UpdateRimRadius(chatID int64, rimRadius string) error {
	return utils.WrapFunctionError(func() error {
		return b.repoProvider.Booking().UpdateRimRadius(chatID, rimRadius)
	})
}

func (b *bookingService) UpdateService(chatID int64, service string) error {
	return utils.WrapFunctionError(func() error {
		return b.repoProvider.Booking().UpdateService(chatID, service)
	})
}

func (b *bookingService) RecalculatePrice(chatID int64) error {
	booking, err := b.FindPending(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	newPrice, err := b.priceService.Calculate(booking.Service, booking.RimRadius)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return b.repoProvider.Booking().UpdatePrice(chatID, newPrice)
	})
}

func (b *bookingService) Cancel(chatID int64) error {
	return b.repoProvider.Booking().Cancel(chatID)
}

func (b *bookingService) parseTime(time string) (start, end string) {
	split := strings.Split(time, "-")
	start = split[0]
	end = split[1]
	return start, end
}
