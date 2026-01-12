package user

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleTime(query *tg.CallbackQuery) error {

	info, err := uch.svcProvider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = uch.cacheProvider.ServiceType().Clean(info.ChatID); err != nil {
		return utils.WrapError(err)
	}

	if err := uch.svcProvider.Lock().Toggle(info); err != nil {
		return utils.WrapError(err)
	}

	types, err := uch.repoProvider.Service().GetServiceTypes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.tgProvider.User().RequestServiceTypes(types, info)
	})
}
