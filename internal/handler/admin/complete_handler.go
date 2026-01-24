package admin

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (ach *adminCallbackHandler) handleComplete(query *tgapi.CallbackQuery) error {
	log.Printf("[handle_complete] callback: %s", query.Data)

	return nil
}
