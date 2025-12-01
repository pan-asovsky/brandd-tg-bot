package handler

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type messageHandler struct {
	api *tgbotapi.BotAPI
}

func NewMessageHandler(api *tgbotapi.BotAPI) MessageHandler {
	return &messageHandler{api: api}
}

func (m *messageHandler) Handle(msg *tgbotapi.Message) error {
	message := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	if _, err := m.api.Send(message); err != nil {
		return err
	}
	return nil
}
