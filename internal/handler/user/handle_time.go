package user

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleTime(query *tg.CallbackQuery) error {

	info, err := uch.callback.UserCallbackParser().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = uch.cache.ServiceType().Clean(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	if err := uch.service.Lock().Toggle(info); err != nil {
		return utils.WrapError(err)
	}

	types, err := uch.repo.Service().GetServiceTypes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.telegram.User().RequestServiceTypes(types, info)
	})
}
