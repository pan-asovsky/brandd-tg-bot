package handler

import (
	"fmt"
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
)

func (c *callbackHandler) handleMenu(cb *api.CallbackQuery, data string) error {
	log.Printf("Handle menu callback data: %s", data)

	bookings := c.slot.GetAvailableBookings()
	kb := c.kb.DateKeyboard(bookings)
	msg := api.NewMessage(cb.Message.Chat.ID, consts.DateChoosingMsg)
	msg.ReplyMarkup = kb

	if _, err := c.api.Send(msg); err != nil {
		return fmt.Errorf("error sending message to chat %d: %v", cb.Message.Chat.ID, err)
	}

	return nil
}
