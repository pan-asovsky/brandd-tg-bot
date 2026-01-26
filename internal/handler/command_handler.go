package handler

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constant"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interface/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type commandHandler struct {
	tgProvider  iprovider.TelegramProvider
	svcProvider iprovider.ServiceProvider
}

func NewCommandHandler(tgProvider iprovider.TelegramProvider, svcProvider iprovider.ServiceProvider) ihandler.MessageHandler {
	return &commandHandler{tgProvider, svcProvider}
}

func (ch *commandHandler) Handle(msg *tgapi.Message) error {
	chatID := msg.Chat.ID
	exists, role := ch.svcProvider.User().GetRole(chatID)

	switch msg.Text {
	case consts.Start:
		if !exists {
			return utils.WrapFunctionError(func() error {
				return ch.tgProvider.User().StartMenu(chatID)
			})
		}

		switch role {
		case "admin":
			return utils.WrapFunctionError(func() error {
				return ch.tgProvider.Admin().ChoiceMenu(chatID)
			})
		}
	}

	return nil
}
