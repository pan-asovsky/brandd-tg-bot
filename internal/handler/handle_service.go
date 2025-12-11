package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleService(q *api.CallbackQuery, cd string) error {
	log.Printf("%s callback: %s", utils.GetCallerName(), cd)

	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return utils.WrapError(err)
	}
	info.ChatID = q.Message.Chat.ID

	rims, err := c.repoProvider.Price().GetAllRimSizes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunction(func() error {
		return c.svcProvider.Telegram().ProcessServiceType(rims, info)
	})
}
