package handler

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleZone(q *api.CallbackQuery, cd string) error {
	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return utils.Error(err)
	}
	info.ChatID = q.Message.Chat.ID

	zones, err := c.svcProvider.Slot().GetAvailableZones(info.Date)
	if err != nil {
		return fmt.Errorf("[handle_zone] %w", err)
	}

	if err := c.svcProvider.Telegram().ProcessZone(zones[info.Zone], info); err != nil {
		return fmt.Errorf("[handle_zone] %w", err)
	}

	return nil
}
