package handler

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *callbackHandler) handleService(q *api.CallbackQuery, cd string) error {
	log.Printf("[handle_service] callback: %s", cd)
	return nil
}
