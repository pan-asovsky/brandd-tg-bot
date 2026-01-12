package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type commandHandler struct {
	tgProvider  p.TelegramProvider
	svcProvider p.ServiceProvider
}

func NewCommandHandler(tgProvider p.TelegramProvider, svcProvider p.ServiceProvider) i.MessageHandler {
	return &commandHandler{tgProvider, svcProvider}
}

func (ch *commandHandler) Handle(msg *tg.Message) error {
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
				return ch.tgProvider.Admin().StartMenu(chatID)
			})
		}
	}

	return nil
}
