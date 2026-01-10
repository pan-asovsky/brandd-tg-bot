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

func (c *userCallbackHandler) handleBooking(query *tg.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_booking]: invalid callback: " + query.Data)
	}

	switch payload {
	case consts.New:
		return c.handleNew(query)
	case consts.My:
		return c.handleMy(query)
	case consts.PreCancel:
		return c.handlePreCancel(query)
	case consts.Cancel:
		return c.handleCancel(query)
	case consts.NoCancel:
		return c.handleNoCancel(query)
	default:
		return errors.New("[handle_booking]: invalid callback: " + query.Data)
	}
}

func (c *userCallbackHandler) handleNew(q *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: q.Message.Chat.ID}

	booking, err := c.svcProvider.Booking().FindActiveNotPending(info.ChatID)
	if booking != nil && err == nil {
		return c.svcProvider.Telegram().SendBookingRestrictionMessage(info.ChatID, booking)
	}

	bookings := c.svcProvider.Slot().GetAvailableBookings()
	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().RequestDate(bookings, info)
	})
}

func (c *userCallbackHandler) handleMy(q *tg.CallbackQuery) error {
	chatID := q.Message.Chat.ID
	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().SendMyBookingsMessage(chatID, func() (*entity.Booking, error) {
			return c.svcProvider.Booking().FindActiveNotPending(chatID)
		})
	})
}

func (c *userCallbackHandler) handlePreCancel(q *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: q.Message.Chat.ID}
	booking, err := c.svcProvider.Booking().FindActiveNotPending(q.Message.Chat.ID)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().SendPreCancelBookingMessage(info.ChatID, booking.Date, booking.Time)
	})
}

func (c *userCallbackHandler) handleCancel(q *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: q.Message.Chat.ID}
	if err := c.svcProvider.Booking().Cancel(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().SendCancellationMessage(info.ChatID)
	})
}

func (c *userCallbackHandler) handleNoCancel(query *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: query.Message.Chat.ID}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().SendCancelDenyMessage(info.ChatID)
	})
}
