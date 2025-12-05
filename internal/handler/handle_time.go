package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleTime(q *api.CallbackQuery, cd string) {
	//log.Printf("[handle_time] callback: %s", cd)
	t, _, d := utils.ParseTimeCallback(cd)

	chatID := q.Message.Chat.ID
	if err := c.lockSvc.Toggle(chatID, d, t); err != nil {
		log.Fatalf("[handle_time] error: %v", err)
	}

	types, err := c.svcRepo.GetServiceTypes()
	if err != nil {
		log.Fatalf("[handle_time] error: %v", err)
	}

	kb := c.kb.ServiceKeyboard(types, t, d)
	utils.SendKeyboardMsg(chatID, consts.ServiceMsg, kb, c.api)
}
