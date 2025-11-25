package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	rd "github.com/pan-asovsky/brandd-tg-bot/internal/repository/redis"
)

func (c *callbackHandler) handleDate(cb *api.CallbackQuery, date string) error {
	log.Printf("Handle date callback: %s", date)

	zones, err := c.slot.GetAvailableZones(date)
	if err != nil {
		log.Printf("Error getting zones: %s", err)
	}

	if err = c.cache.SetZones(rd.KeyForDate(date), zones); err != nil {
		log.Printf("Error caching zones: %s", err)
	}

	kb := c.kb.ZoneKeyboard(zones)
	msg := api.NewMessage(cb.Message.Chat.ID, consts.ZoneChoosingMsg)
	msg.ReplyMarkup = kb

	if _, err := c.api.Send(msg); err != nil {
		log.Printf("Error sending message to chat: %s", err)
	}

	return nil
}
