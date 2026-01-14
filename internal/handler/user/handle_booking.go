package user

import (
	"errors"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleBooking(query *tg.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_booking]: invalid callback: " + query.Data)
	}

	switch payload {
	case consts.New:
		return uch.handleNew(query)
	case consts.My:
		return uch.handleMy(query)
	case consts.PreCancel:
		return uch.handlePreCancel(query)
	case consts.Cancel:
		return uch.handleCancel(query)
	case consts.NoCancel:
		return uch.handleNoCancel(query)
	default:
		return errors.New("[handle_booking]: invalid callback: " + query.Data)
	}
}

func (uch *userCallbackHandler) handleNew(q *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: q.Message.Chat.ID}

	if err := uch.serviceProvider.Booking().CancelOldIfExists(q.Message.Chat.ID); err != nil {
		return utils.WrapError(err)
	}

	booking, err := uch.serviceProvider.Booking().FindActiveNotPending(info.ChatID)
	if booking != nil && err == nil {
		return uch.telegramProvider.User().SendBookingRestrictionMessage(info.ChatID, booking)
	}

	bookings := uch.serviceProvider.Slot().GetAvailableDates()
	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().RequestDate(bookings, info)
	})
}

func (uch *userCallbackHandler) handleMy(q *tg.CallbackQuery) error {
	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().SendMyBookingsMessage(q.Message.Chat.ID, func() (*entity.Booking, error) {
			return uch.serviceProvider.Booking().FindActiveNotPending(q.Message.Chat.ID)
		})
	})
}

func (uch *userCallbackHandler) handlePreCancel(q *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: q.Message.Chat.ID}
	booking, err := uch.serviceProvider.Booking().FindActiveNotPending(q.Message.Chat.ID)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().SendPreCancelBookingMessage(info.ChatID, booking.Date, booking.Time)
	})
}

func (uch *userCallbackHandler) handleCancel(q *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: q.Message.Chat.ID}

	booking, err := uch.serviceProvider.Booking().FindActiveNotPending(info.ChatID)
	if err != nil {
		return utils.WrapError(err)
	}

	err = uch.serviceProvider.Slot().FreeUp(booking.Date, booking.Time)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = uch.serviceProvider.Booking().Cancel(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().SendCancellationMessage(info.ChatID)
	})
}

func (uch *userCallbackHandler) handleNoCancel(query *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: query.Message.Chat.ID}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().SendCancelDenyMessage(info.ChatID)
	})
}
