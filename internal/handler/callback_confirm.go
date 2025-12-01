package handler

import (
	"fmt"
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleConfirm(q *api.CallbackQuery, cd string) error {
	log.Printf("[handle_confirm] callback: %s", cd)

	auto, err := c.cfgRepo.IsAutoConfirm()
	if err != nil {
		return err
	}

	if auto {
		msg := fmt.Sprintf(consts.ConfirmMsg, "30.02.2026", "00:00") //todo: сюда надо тянуть данные callback
		utils.SendMessage(q.Message.Chat.ID, msg, c.api)
	} else {
		utils.SendMessage(q.Message.Chat.ID, consts.PendingConfirmMsg, c.api)
	}

	return nil
}
