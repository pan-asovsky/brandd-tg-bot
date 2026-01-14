package user

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleServiceSelect(query *tg.CallbackQuery) error {
	info, err := uch.callbackProvider.UserCallbackParser().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if len(info.Service) > 0 {
		selected, err := uch.cacheProvider.ServiceType().Toggle(info.ChatID, info.Service)
		if err != nil {
			return utils.WrapError(err)
		}
		info.SelectedServices = selected
	}

	types, err := uch.repoProvider.Service().GetServiceTypes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().RequestServiceTypes(types, info)
	})
}

func (uch *userCallbackHandler) handleServiceConfirm(query *tg.CallbackQuery) error {
	info, err := uch.callbackProvider.UserCallbackParser().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	rims, err := uch.repoProvider.Price().GetAllRimSizes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegramProvider.User().RequestRimRadius(rims, info)
	})
}
