package handler

import (
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	kb "github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type CommandHandler interface {
	Handle(msg *tgbot.Message) error
}

type commandHandler struct {
	api *tgbot.BotAPI
	kb  kb.KeyboardService
}

func NewCommandHandler(api *tgbot.BotAPI, kb kb.KeyboardService) CommandHandler {
	return &commandHandler{api, kb}
}

func (c *commandHandler) Handle(msg *tgbot.Message) error {
	switch msg.Text {
	case consts.Start:
		c.handleStart(msg.Chat.ID)
	}

	return nil
}

func (c *commandHandler) handleStart(chatID int64) {
	message := tgbot.NewMessage(chatID, consts.GreetingMsg)
	message.ReplyMarkup = c.kb.GreetingKeyboard()

	if _, err := c.api.Send(message); err != nil {
		log.Printf("Error handling start command for chat %d: %v", chatID, err)
	}
}
