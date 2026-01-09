package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleDate(query *api.CallbackQuery) error {
	provider := c.svcProvider

	info, err := provider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	zones, err := provider.Slot().GetAvailableZones(info.Date)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return provider.Telegram().RequestZone(zones, info)
	})
}
