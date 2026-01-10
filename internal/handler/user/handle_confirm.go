package user

import (
	"fmt"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *userCallbackHandler) handleConfirm(query *tg.CallbackQuery) error {
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

func (c *userCallbackHandler) handleYes(query *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: query.Message.Chat.ID}

	if err := c.cleanSession(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().RequestUserPhone(info)
	})
}

func (c *userCallbackHandler) handleNo(query *tg.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: query.Message.Chat.ID}

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

func (c *userCallbackHandler) cleanSession(chatID int64) error {
	if err := c.cacheProvider.ServiceType().Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Lock().Clean(chatID)
	})
}
