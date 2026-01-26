package admin

import (
	"errors"
	"log"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
)

func (ach *adminCallbackHandler) handleBack(query *tgapi.CallbackQuery) error {
	log.Printf("[handle_back] callback: %s", query.Data)
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_back] invalid callback: " + query.Data)
	}

	switch payload {
	case admflow.Flow:
		return ach.telegram.Admin().ChoiceMenu(query.Message.Chat.ID)
	case admflow.Menu:
		return ach.telegram.Admin().StartMenu(query.Message.Chat.ID)
	default:
		return errors.New("[handle_back] invalid callback: " + query.Data)
	}
}
