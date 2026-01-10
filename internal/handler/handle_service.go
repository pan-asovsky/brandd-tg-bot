package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleServiceSelect(query *api.CallbackQuery) error {
	svcProvider := c.svcProvider

	info, err := svcProvider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if len(info.Service) > 0 {
		selected, err := c.cacheProvider.ServiceType().Toggle(info.ChatID, info.Service)
		if err != nil {
			return utils.WrapError(err)
		}
		info.SelectedServices = selected
	}

	types, err := c.pgProvider.Service().GetServiceTypes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return svcProvider.Telegram().RequestServiceTypes(types, info)
	})
}

func (c *callbackHandler) handleServiceConfirm(query *api.CallbackQuery) error {
	svcProvider := c.svcProvider

	info, err := svcProvider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	rims, err := c.pgProvider.Price().GetAllRimSizes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return svcProvider.Telegram().RequestRimRadius(rims, info)
	})
}
