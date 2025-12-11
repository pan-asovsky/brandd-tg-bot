package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleConfirm(q *api.CallbackQuery, _ string) error {
	chatID := q.Message.Chat.ID
	auto, err := c.repoProvider.Config().IsAutoConfirm()
	if err != nil {
		return utils.WrapError(err)
	}

	if auto {
		return utils.WrapFunction(func() error {
			return c.handleAutoConfirm(chatID)
		})
	}
	return utils.WrapFunction(func() error {
		return c.handlePendingConfirm(chatID)
	})
}

func (c *callbackHandler) handleAutoConfirm(chatID int64) error {
	slot, err := c.getSlot(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunction(func() error {
		return c.confirmAndNotify(chatID, slot)
	})
}

func (c *callbackHandler) getSlot(chatID int64) (*model.Slot, error) {
	booking, err := c.svcProvider.Booking().FindActiveByChatID(chatID)
	if err != nil {
		return nil, utils.WrapError(err)
	}
	return utils.WrapErrorWithValue(c.svcProvider.Slot().FindByDateAndTime(booking.Date, booking.Time))
}

func (c *callbackHandler) confirmAndNotify(chatID int64, slot *model.Slot) error {
	if err := c.svcProvider.Booking().Confirm(chatID); err != nil {
		return utils.WrapError(err)
	}
	return utils.WrapFunction(func() error {
		return c.svcProvider.Telegram().ProcessConfirm(chatID, slot)
	})
}

func (c *callbackHandler) handlePendingConfirm(chatID int64) error {
	return utils.WrapFunction(func() error {
		return c.svcProvider.Telegram().ProcessPendingConfirm(chatID)
	})
}
