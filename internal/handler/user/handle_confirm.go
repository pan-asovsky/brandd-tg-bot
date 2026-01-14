package user

import (
	"fmt"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleConfirm(query *tgapi.CallbackQuery) error {
	userChoice, ok := strings.CutPrefix(query.Data, "CONFIRM::")
	if !ok {
		return fmt.Errorf("[handle_confirm] invalid callback: %s", query.Data)
	}
	switch userChoice {
	case consts.Yes:
		return uch.handleYes(query)
	case consts.No:
		return uch.handleNo(query)
	}

	return nil
}

func (uch *userCallbackHandler) handleYes(query *tgapi.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: query.Message.Chat.ID}

	if err := uch.cleanSession(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().RequestUserPhone(info)
	})
}

func (uch *userCallbackHandler) handleNo(query *tgapi.CallbackQuery) error {
	info := &model.UserSessionInfo{ChatID: query.Message.Chat.ID}

	if err := uch.cleanSession(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	if err := uch.serviceProvider.Booking().Cancel(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().SendCancellationMessage(info.ChatID)
	})
}

func (uch *userCallbackHandler) cleanSession(chatID int64) error {
	if err := uch.cacheProvider.ServiceType().Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.serviceProvider.Lock().Clean(chatID)
	})
}
