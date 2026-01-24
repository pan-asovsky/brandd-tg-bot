package admin

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (ach *adminCallbackHandler) handleReject(query *tgapi.CallbackQuery) error {
	log.Printf("[handle_reject] callback: %s", query.Data)

	return nil
}
