package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

//func (c *callbackHandler) handleConfirm(q *api.CallbackQuery, cd string) error {
//	chatID := q.Message.Chat.ID
//	if c.repoProvider.Config().IsAutoConfirm() {
//		booking, err := c.svcProvider.Booking().FindActiveByChatID(chatID)
//		if err != nil {
//			return fmt.Errorf("[handle_confirm] %w", err)
//		}
//
//		slot, err := c.svcProvider.Slot().FindByDateAndTime(booking.Date, booking.Time)
//		if err != nil {
//			return fmt.Errorf("[handle_confirm] %w", err)
//		}
//
//		if err := c.svcProvider.Booking().Confirm(chatID); err != nil {
//			return fmt.Errorf("[handle_confirm] %w", err)
//		}
//
//		c.svcProvider.Telegram().ProcessConfirm(chatID, slot)
//	} else {
//		c.svcProvider.Telegram().ProcessPendingConfirm(chatID)
//	}
//
//	return nil
//}

func (c *callbackHandler) handleConfirm(q *api.CallbackQuery, cd string) error {
	chatID := q.Message.Chat.ID

	if c.repoProvider.Config().IsAutoConfirm() {
		return c.handleAutoConfirm(chatID)
	}

	return c.handlePendingConfirm(chatID)
}

func (c *callbackHandler) handleAutoConfirm(chatID int64) error {
	slot, err := c.getSlot(chatID)
	if err != nil {
		return utils.Error(err)
	}

	if err := c.confirmAndNotify(chatID, slot); err != nil {
		return utils.Error(err)
	}

	return nil
}

func (c *callbackHandler) getSlot(chatID int64) (*model.Slot, error) {
	booking, err := c.svcProvider.Booking().FindActiveByChatID(chatID)
	if err != nil {
		return nil, utils.Error(err)
	}

	slot, err := c.svcProvider.Slot().FindByDateAndTime(booking.Date, booking.Time)
	if err != nil {
		return nil, utils.Error(err)
	}

	return slot, nil
}

func (c *callbackHandler) confirmAndNotify(chatID int64, slot *model.Slot) error {
	if err := c.svcProvider.Booking().Confirm(chatID); err != nil {
		return utils.Error(err)
	}

	if err := c.svcProvider.Telegram().ProcessConfirm(chatID, slot); err != nil {
		return utils.Error(err)
	}

	return nil
}

func (c *callbackHandler) handlePendingConfirm(chatID int64) error {
	if err := c.svcProvider.Telegram().ProcessPendingConfirm(chatID); err != nil {
		return utils.Error(err)
	}
	return nil
}
