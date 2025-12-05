package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleDate(q *api.CallbackQuery, cd string) {
	info := utils.ParseDateCallback(cd)
	info.ChatID = q.Message.Chat.ID

	zones := c.slot.GetAvailableZones(cd)
	c.telegram.ProcessDate(zones, info)
}
