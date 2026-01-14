package admin

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
)

type adminMessageHandler struct {
	svcProvider iprovider.ServiceProvider
}

func NewAdminMessageHandler(svcProvider iprovider.ServiceProvider) ihandler.MessageHandler {
	return &adminMessageHandler{svcProvider: svcProvider}
}

func (amh *adminMessageHandler) Handle(msg *tgapi.Message) error {

	return nil
}
