package user

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleRim(query *tg.CallbackQuery) error {
	provider := uch.serviceProvider

	info, err := uch.callbackProvider.UserCallbackParser().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = uch.cacheProvider.ServiceType().Clean(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	totalPrice, err := provider.Price().Calculate(info.Service, info.RimRadius)
	if err != nil {
		return utils.WrapError(err)
	}
	info.TotalPrice = totalPrice

	exists := provider.Booking().ExistsByChatID(info.ChatID)

	var booking *entity.Booking
	if !exists {
		booking, err = provider.Booking().Create(info)
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

		booking, err = provider.Booking().FindPending(info.ChatID)
		if err != nil {
			return utils.WrapError(err)
		}
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().RequestPreConfirm(booking, info)
	})
}
