package user

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleDate(query *tg.CallbackQuery) error {
	info, err := uch.svcProvider.CallbackParsing().Parse(query)
	if err != nil {
		return utils.WrapError(err)
	}

	zones, err := uch.svcProvider.Slot().GetAvailableZones(info.Date)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uch.tgProvider.User().RequestZone(zones, info)
	})
}
