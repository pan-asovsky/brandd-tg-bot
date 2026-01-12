package admin

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
)

type adminMessageHandler struct {
	svcProvider p.ServiceProvider
}

func NewAdminMessageHandler(svcProvider p.ServiceProvider) i.MessageHandler {
	return &adminMessageHandler{svcProvider: svcProvider}
}

func (amh *adminMessageHandler) Handle(msg *tg.Message) error {

	return nil
}
