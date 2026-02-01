package service

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type bookingService struct {
	repoProvider    iprovider.RepoProvider
	slotService     isvc.SlotService
	priceService    isvc.PriceService
	dateTimeService isvc.DateTimeService
}

func NewBookingService(
	repoProvider iprovider.RepoProvider,
	slotService isvc.SlotService,
	priceService isvc.PriceService,
	dateTimeService isvc.DateTimeService,
) isvc.BookingService {
	return &bookingService{repoProvider: repoProvider, slotService: slotService, priceService: priceService, dateTimeService: dateTimeService}
}

func (bs *bookingService) Create(info *model.UserSessionInfo) (*entity.Booking, error) {
	date, err := time.Parse("2006-01-02", info.Date)
	if err != nil {
		return nil, fmt.Errorf("[create_booking] error parsing date: %v", err)
	}

	booking := &entity.Booking{
		ChatID:     info.ChatID,
		Date:       date,
		Service:    info.Service,
		RimRadius:  info.RimRadius,
		TotalPrice: sql.NullInt64{Int64: info.TotalPrice, Valid: true},
		Status:     entity.Pending,
	}

	start, _ := bs.dateTimeService.ParseToStartEndTime(info.Time)
	booking.Time = start

	return bs.repoProvider.Booking().Save(booking)
}

func (bs *bookingService) Confirm(chatID int64) error {
	return bs.repoProvider.Booking().Confirm(chatID)
}

func (bs *bookingService) AutoConfirm(chatID int64) error {
	return bs.repoProvider.Booking().AutoConfirm(chatID)
}

func (bs *bookingService) SetPhone(phone string, chatID int64) error {
	return bs.repoProvider.Booking().SetPhone(phone, chatID)
}

func (bs *bookingService) FindActiveNotPending(chatID int64) (*entity.Booking, error) {
	return bs.repoProvider.Booking().FindActiveNotPending(chatID)
}

func (bs *bookingService) FindPending(chatID int64) (*entity.Booking, error) {
	return bs.repoProvider.Booking().FindActivePending(chatID)
}

func (bs *bookingService) CancelOldIfExists(chatID int64) error {
	oldBooking, err := bs.FindPending(chatID)
	if err == nil && oldBooking != nil {
		log.Printf("[cancel_old_if_exists] cancelled for chatID: %d", chatID)
		if err = bs.Cancel(chatID); err != nil {
			return utils.WrapError(err)
		}
	}
	return nil
}

func (bs *bookingService) ExistsByChatID(chatID int64) bool {
	return bs.repoProvider.Booking().Exists(chatID)
}

func (bs *bookingService) UpdateStatus(chatID int64, status entity.BookingStatus) error {
	return utils.WrapFunctionError(func() error {
		return bs.repoProvider.Booking().UpdateStatus(chatID, status)
	})
}

func (bs *bookingService) UpdateRimRadius(chatID int64, rimRadius string) error {
	return utils.WrapFunctionError(func() error {
		return bs.repoProvider.Booking().UpdateRimRadius(chatID, rimRadius)
	})
}

func (bs *bookingService) UpdateService(chatID int64, service string) error {
	return utils.WrapFunctionError(func() error {
		return bs.repoProvider.Booking().UpdateService(chatID, service)
	})
}

func (bs *bookingService) RecalculatePrice(chatID int64) error {
	booking, err := bs.FindPending(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	newPrice, err := bs.priceService.Calculate(booking.Service, booking.RimRadius)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return bs.repoProvider.Booking().UpdatePrice(chatID, newPrice)
	})
}

func (bs *bookingService) Cancel(chatID int64) error {
	//booking, err := bs.repoProvider.Booking().FindAnyActive(chatID)
	//if err != nil || booking == nil {
	//	return utils.WrapError(err)
	//}

	// 1.
	// отменяет юзер из меню Мои записи
	// status = confirmed
	// is_active = true
	// confirmed_by != null
	// chatID = booking.ChatID
	// cancelled_by = user(ChatID)

	// 2.
	// отменяют юзер на этапе подтверждения
	// status = pending
	// is_active = true
	// confirmed_by == null
	// chatID = booking.ChatID
	// cancelled_by = user(ChatID)

	//todo: закрывать надо тут в соответствии с логикой, а репой только сохранять результат
	//if entity.Confirmed == booking.Status {
	//	booking.CancelledBy = sql.NullString{
	//		String: fmt.Sprintf("USER(%d)", booking.ChatID),
	//		Valid:  true,
	//	}
	//}
	//booking.Status = entity.Cancelled

	return bs.repoProvider.Booking().Cancel(chatID)
}

func (bs *bookingService) FindAllActive() ([]entity.Booking, error) {
	return bs.repoProvider.Booking().FindAllActive()
}

func (bs *bookingService) FindById(bookingID int64) (*entity.Booking, error) {
	return bs.repoProvider.Booking().Find(bookingID)
}

func (bs *bookingService) Close(info *model.BookingInfo) (*entity.Booking, error) {
	return bs.repoProvider.Booking().Close(info)
}

func (bs *bookingService) parseToStartEndTime(time string) (start, end string) {
	split := strings.Split(time, "-")
	start = split[0]
	end = split[1]
	return start, end
}
