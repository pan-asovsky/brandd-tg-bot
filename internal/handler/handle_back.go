package handler

import (
	"errors"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleBack(query *api.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_back] invalid callback" + query.Data)
	}

	switch payload {
	case consts.Menu:
		return utils.WrapFunctionError(func() error {
			return c.svcProvider.Telegram().SendStartMenu(query.Message.Chat.ID)
		})
	default:
		return nil
	}
}
