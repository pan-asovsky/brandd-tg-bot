package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleRim(query *api.CallbackQuery) error {
	provider := c.svcProvider

	info, err := provider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = c.serviceTypeCache.Clean(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	totalPrice, err := provider.Price().Calculate(info.Service, info.RimRadius)
	if err != nil {
		return utils.WrapError(err)
	}
	info.TotalPrice = totalPrice

	exists := provider.Booking().ExistsByChatID(info.ChatID)

	var booking *model.Booking
	if !exists {
		booking, err = provider.Booking().Create(info)
		log.Printf("[handle_rim] (false). created: %v", booking)
		if err != nil {
			return utils.WrapError(err)
		}
	} else {
		if err = provider.Booking().UpdateRimRadius(info.ChatID, info.RimRadius); err != nil {
			return utils.WrapError(err)
		}

		if err = provider.Booking().UpdateService(info.ChatID, info.Service); err != nil {
			return utils.WrapError(err)
		}

		if err = provider.Booking().RecalculatePrice(info.ChatID); err != nil {
			return utils.WrapError(err)
		}

		booking, err = provider.Booking().FindActiveByChatID(info.ChatID)
		log.Printf("[handle_rim] (true) founded: %v", booking)
		if err != nil {
			return utils.WrapError(err)
		}
	}

	return utils.WrapFunctionError(func() error {
		return provider.Telegram().RequestPreConfirm(booking, info)
	})
}
