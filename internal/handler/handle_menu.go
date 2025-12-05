package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *callbackHandler) handleMenu(q *api.CallbackQuery, cd string) {
	bookings := c.slot.GetAvailableBookings()
	c.telegram.ProcessMenu(bookings, q.Message.Chat.ID)
}
