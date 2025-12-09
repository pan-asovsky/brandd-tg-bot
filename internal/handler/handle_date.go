package handler

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleDate(q *api.CallbackQuery, cd string) error {
	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return fmt.Errorf("[hanle_date] %w", err)
	}
	info.ChatID = q.Message.Chat.ID

	zones, err := c.svcProvider.Slot().GetAvailableZones(cd)
	if err != nil {
		return fmt.Errorf("[hanle_date] %w", err)
	}

	if err := c.svcProvider.Telegram().ProcessDate(zones, info); err != nil {
		return fmt.Errorf("[hanle_date] %w", err)
	}

	return nil
}
