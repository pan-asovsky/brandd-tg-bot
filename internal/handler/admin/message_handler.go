package admin

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type adminMessageHandler struct {
	tgProvider  iprovider.TelegramProvider
	svcProvider iprovider.ServiceProvider
}

func NewAdminMessageHandler(tgProvider iprovider.TelegramProvider, svcProvider iprovider.ServiceProvider) ihandler.MessageHandler {
	return &adminMessageHandler{tgProvider: tgProvider, svcProvider: svcProvider}
}

func (amh *adminMessageHandler) Handle(msg *tgapi.Message) error {
	return utils.WrapFunctionError(func() error {
		return amh.tgProvider.Common().SendMessage(msg.Chat.ID, usflow.DontKnowHowToAnswer)
	})
}
