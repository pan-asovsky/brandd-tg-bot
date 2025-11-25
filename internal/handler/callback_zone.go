package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *callbackHandler) handleZone(cb *api.CallbackQuery, date string) error {
	log.Printf("Handle zone callback: %s", date)
	return nil
}
