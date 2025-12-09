package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
)

type bookingService struct {
	repoProvider *pg.Provider
	slotService  SlotService
}

func (b *bookingService) Create(info *types.UserSessionInfo) error {
	booking := &model.Booking{
		ChatID:    info.ChatID,
		Date:      info.Date,
		RimRadius: info.Radius,
		Status:    model.NotConfirmed,
		CreatedAt: time.Now(),
	}

	start, end := b.parseTime(info.Time)
	if err := b.slotService.MarkUnavailable(info.Date, start, end); err != nil {
		return fmt.Errorf("[create_booking] %w", err)
	}

	booking.Time = start

	return b.repoProvider.Booking().Save(booking)
}

func (b *bookingService) Confirm(chatID int64) error {
	return b.repoProvider.Booking().Confirm(chatID)
}

func (b *bookingService) SetPhone(phone string, chatID int64) error {
	return b.repoProvider.Booking().SetPhone(phone, chatID)
}

func (b *bookingService) FindActiveByChatID(chatID int64) (*model.Booking, error) {
	return b.repoProvider.Booking().FindActiveByChatID(chatID)
}

func (b *bookingService) parseTime(time string) (start, end string) {
	split := strings.Split(time, "-")
	start = split[0]
	end = split[1]
	return start, end
}
