package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleZone(query *api.CallbackQuery) error {
	provider := c.svcProvider

	info, err := provider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = provider.Lock().Clean(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	zones, err := provider.Slot().GetAvailableZones(info.Date)
	if err != nil {
		return utils.WrapError(err)
	}
	//log.Printf("[handle_zone] zone: %v, info.Zone: %s", zones, info.Zone)

	return utils.WrapFunctionError(func() error {
		return provider.Telegram().RequestTime(zones[info.Zone], info)
	})
}
