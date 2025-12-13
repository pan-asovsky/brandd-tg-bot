package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleServiceSelect(q *api.CallbackQuery, cd string) error {
	log.Printf("[service_select] callback: %s", cd)
	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return utils.WrapError(err)
	}
	info.ChatID = q.Message.Chat.ID

	selected, err := c.sessionRepo.Toggle(info.ChatID, info.Service)
	if err != nil {
		return utils.WrapError(err)
	}
	info.SelectedServices = selected
	log.Println("[service_select] services to callback:", info.SelectedServices)

	types, err := c.pgProvider.Service().GetServiceTypes()
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		//return c.svcProvider.Telegram().ProcessServiceSelect(types, info)
		return c.svcProvider.Telegram().RequestServiceTypes(types, info)
	})
}

func (c *callbackHandler) handleServiceConfirm(q *api.CallbackQuery, cd string) error {
	log.Printf("[service_confirm] callback: %s", cd)
	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return utils.WrapError(err)
	}
	info.ChatID = q.Message.Chat.ID

	rims, err := c.pgProvider.Price().GetAllRimSizes()
	if err != nil {
		return utils.WrapError(err)
	}
	log.Println("[service_confirm] session info:", info)

	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().RequestRimRadius(rims, info)
	})
}
