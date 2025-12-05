package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *callbackHandler) handleMenu(q *api.CallbackQuery, cd string) error {
	//log.Printf("[handle_menu] callback: %s", cd)

	bookings := c.slot.GetAvailableBookings()
	c.telegramSvc.ProcessMenu(bookings, q.Message.Chat.ID)

	//kb := c.kb.DateKeyboard(bookings)
	//utils.SendKeyboardMsg(q.Message.Chat.ID, consts.DateMsg, kb, c.api)

	return nil
}
