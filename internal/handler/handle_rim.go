package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleRim(q *api.CallbackQuery, cd string) error {
	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return utils.WrapError(err)
	}
	info.ChatID = q.Message.Chat.ID

	totalPrice, err := c.svcProvider.Price().Calculate(info.Service, info.Radius)
	if err != nil {
		return utils.WrapError(err)
	}
	info.TotalPrice = totalPrice

	booking, err := c.svcProvider.Booking().Create(info)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().RequestPreConfirm(booking, info.ChatID)
	})
}
