package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleService(q *api.CallbackQuery, cd string) error {
	//log.Printf("[handle_service] callback: %s", cd)

	//todo:
	svc, t, d := utils.ParseServiceCallback(cd)

	rims := c.priceRepo.GetAllRimSizes()
	kb := c.kb.RimsKeyboard(rims, svc, t, d)
	utils.SendKeyboardMsg(q.Message.Chat.ID, consts.RimMsg, kb, c.api)

	return nil
}
