package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleRim(q *api.CallbackQuery, cd string) {
	//log.Printf("[handle_rim] callback: %s", cd)
	info := utils.ParseRimCallback(cd)
	info.ChatID = q.Message.Chat.ID

	if err := c.bookingSvc.Create(info); err != nil {
		log.Fatalf("[handle_rim] booking create error: %v", err)
	}

	kb := c.kb.RequestPhoneKeyboard()
	utils.SendRequestPhoneMsg(q.Message.Chat.ID, consts.RequestUserPhone, kb, c.api)
}
