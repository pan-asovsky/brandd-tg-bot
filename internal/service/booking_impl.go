package service

import (
	"log"
	"strings"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
)

type bookingService struct {
	repoProvider *pg.Provider
}

func (b *bookingService) Create(info *types.UserSessionInfo) error {
	booking := &model.Booking{
		ChatID:    info.ChatID,
		RimRadius: info.Radius,
		Status:    model.NotConfirmed,
		CreatedAt: time.Now(),
	}

	booking.SlotID = b.getSlotID(info)
	booking.ServiceTypeID = b.getServiceTypeID(info)

	if err := b.repoProvider.Booking().Save(booking); err != nil {
		return err
	}

	//log.Printf("[booking]: %+v", booking)
	return nil
}

func (b *bookingService) Confirm(chatID int64) error {
	return nil
}

func (b *bookingService) SetPhone(phone string, chatID int64) error {
	booking, err := b.repoProvider.Booking().FindActiveByChatID(chatID)
	if err != nil {
		return err
	}
	booking.UserPhone = phone
	return nil
}

func (b *bookingService) getSlotID(info *types.UserSessionInfo) int64 {
	start, end := b.parseTime(info.Time)
	slot, err := b.repoProvider.Slot().FindByDateAndTime(info.Date, start, end)
	if err != nil {
		log.Fatal(err)
	}
	return slot.ID
}

func (b *bookingService) getServiceTypeID(info *types.UserSessionInfo) int64 {
	svcType, err := b.repoProvider.Service().FindByCode(info.Service)
	if err != nil {
		log.Fatal(err)
	}
	return svcType.ID
}

func (b *bookingService) parseTime(time string) (start, end string) {
	split := strings.Split(time, "-")
	start = split[0]
	end = split[1]
	//log.Printf("[parse_time]: time: %s, start: %s, end: %s", time, start, end)
	return start, end
}
