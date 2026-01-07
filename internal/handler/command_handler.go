package handler

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type CommandHandler interface {
	Handle(msg *tgbot.Message) error
}

type commandHandler struct {
	api         *tgbot.BotAPI
	svcProvider *service.Provider
}

func NewCommandHandler(api *tgbot.BotAPI, svcProvider *service.Provider) CommandHandler {
	return &commandHandler{api, svcProvider}
}

func (c *commandHandler) Handle(msg *tgbot.Message) error {
	chatID := msg.Chat.ID
	switch msg.Text {
	case consts.Start:
		return utils.WrapFunctionError(func() error {
			return c.svcProvider.Telegram().SendStartMenu(chatID)
		})
	}

	return nil
}
