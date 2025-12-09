package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleMenu(q *api.CallbackQuery, cd string) error {
	info := &types.UserSessionInfo{ChatID: q.Message.Chat.ID}

	bookings := c.svcProvider.Slot().GetAvailableBookings()
	if err := c.svcProvider.Telegram().ProcessMenu(bookings, info); err != nil {
		return utils.Error(err)
	}
	return nil
}
