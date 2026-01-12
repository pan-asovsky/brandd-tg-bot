package admin

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
)

type adminCallbackHandler struct{}

func NewAdminCallbackHandler() i.CallbackHandler {
	return &adminCallbackHandler{}
}

func (a *adminCallbackHandler) Handle(callback *tg.CallbackQuery) error {

	return nil
}
