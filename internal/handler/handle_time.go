package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleTime(query *api.CallbackQuery) error {
	provider := c.svcProvider

	info, err := provider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = c.serviceTypeCache.Clean(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	if err := provider.Lock().Toggle(info); err != nil {
		return utils.WrapError(err)
	}

	types, err := c.pgProvider.Service().GetServiceTypes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return provider.Telegram().RequestServiceTypes(types, info)
	})
}
