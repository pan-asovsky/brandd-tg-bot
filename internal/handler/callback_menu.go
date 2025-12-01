package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleMenu(q *api.CallbackQuery, cd string) error {
	//log.Printf("[handle_menu] callback: %s", cd)

	bookings := c.slot.GetAvailableBookings()
	kb := c.kb.DateKeyboard(bookings)
	utils.SendKeyboardMessage(q.Message.Chat.ID, consts.DateMsg, kb, c.api)

	return nil
}
