package admin

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
)

type adminCallbackHandler struct{}

func NewAdminCallbackHandler() ihandler.CallbackHandler {
	return &adminCallbackHandler{}
}

func (a *adminCallbackHandler) Handle(callback *tgapi.CallbackQuery) error {

	return nil
}
