package handler

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleRim(q *api.CallbackQuery, cd string) error {
	info, err := utils.GetSessionInfo(cd)
	if err != nil {
		return fmt.Errorf("[handle_rim] %w", err)
	}
	info.ChatID = q.Message.Chat.ID

	if err := c.svcProvider.Booking().Create(info); err != nil {
		return fmt.Errorf("[handle_rim] %w", err)
	}

	if err := c.svcProvider.Telegram().ProcessRimRadius(info); err != nil {
		return fmt.Errorf("[handle_rim] %w", err)
	}

	return nil
}
