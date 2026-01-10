package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type commandHandler struct {
	tgapi       *tg.BotAPI
	svcProvider *service.Provider
}

func NewCommandHandler(tgapi *tg.BotAPI, svcProvider *service.Provider) i.MessageHandler {
	return &commandHandler{tgapi, svcProvider}
}

func (c *commandHandler) Handle(msg *tg.Message) error {
	chatID := msg.Chat.ID
	switch msg.Text {
	case consts.Start:
		return utils.WrapFunctionError(func() error {
			return c.svcProvider.Telegram().SendStartMenu(chatID)
		})
	}

	return nil
}
