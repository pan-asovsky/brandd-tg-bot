package handler

import (
	"fmt"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleConfirm(query *api.CallbackQuery) error {
	userChoice, ok := strings.CutPrefix(query.Data, "CONFIRM::")
	if !ok {
		return fmt.Errorf("[handle_confirm] invalid callback: %s", query.Data)
	}
	switch userChoice {
	case consts.Yes:
		return c.handleYes(query)
	case consts.No:
		return c.handleNo(query)
	}

	return nil
}

func (c *callbackHandler) handleYes(query *api.CallbackQuery) error {
	info := &types.UserSessionInfo{ChatID: query.Message.Chat.ID}

	if err := c.cleanSession(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().RequestUserPhone(info)
	})
}

func (c *callbackHandler) handleNo(query *api.CallbackQuery) error {
	info := &types.UserSessionInfo{ChatID: query.Message.Chat.ID}

	if err := c.cleanSession(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	if err := c.svcProvider.Booking().Cancel(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().SendCancellationMessage(info.ChatID)
	})
}

func (c *callbackHandler) cleanSession(chatID int64) error {
	if err := c.serviceTypeCache.Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Lock().Clean(chatID)
	})
}
