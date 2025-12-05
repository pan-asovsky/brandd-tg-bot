package handler

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleRim(q *api.CallbackQuery, cd string) error {
	//log.Printf("[handle_rim] callback: %s", cd)
	info, err := utils.ParseRimCallback(cd)
	if err != nil {
		return err
	}
	info.ChatID = q.Message.Chat.ID

	if err = c.bookingSvc.Create(info); err != nil {
		return fmt.Errorf("[handle_rim] booking create error: %w", err)
	}

	kb := c.kb.RequestPhoneKeyboard()
	utils.SendRequestPhoneMsg(q.Message.Chat.ID, consts.RequestUserPhone, kb, c.api)

	return nil
}
