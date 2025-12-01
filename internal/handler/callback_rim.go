package handler

import (
	"fmt"
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleRim(q *api.CallbackQuery, cd string) error {
	log.Printf("[handle_rim] callback: %s", cd)
	r, svc, time, date := utils.ParseRimCallback(cd)

	msg := fmt.Sprintf(consts.PreConfirmMsg, date, consts.Time[time], time, consts.ServiceNames[svc], r)
	kb := c.kb.ConfirmKeyboard()
	utils.SendKeyboardMessage(q.Message.Chat.ID, msg, kb, c.api)

	return nil
}
