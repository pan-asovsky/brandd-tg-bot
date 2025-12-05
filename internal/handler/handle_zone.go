package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleZone(q *api.CallbackQuery, cd string) {
	info := utils.ParseZoneCallback(cd)
	info.ChatID = q.Message.Chat.ID

	zones := c.slot.GetAvailableZones(info.Date)
	c.telegram.ProcessZone(zones[info.Zone], info)
}
