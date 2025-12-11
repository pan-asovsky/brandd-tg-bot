package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleTime(q *api.CallbackQuery, cd string) error {
	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return utils.WrapError(err)
	}
	info.ChatID = q.Message.Chat.ID

	if err := c.svcProvider.Lock().Toggle(info); err != nil {
		return utils.WrapError(err)
	}

	types, err := c.repoProvider.Service().GetServiceTypes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunction(func() error {
		return c.svcProvider.Telegram().ProcessTime(types, info)
	})
}
