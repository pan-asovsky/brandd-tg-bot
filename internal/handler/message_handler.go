package handler

import (
	"fmt"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type MessageHandler interface {
	Handle(msg *tgbot.Message) error
}

type messageHandler struct {
	api         *tgbot.BotAPI
	bookingSvc  service.BookingService
	telegramSvc service.TelegramService
}

func NewMessageHandler(api *tgbot.BotAPI, bookingSvc service.BookingService, telegramSvc service.TelegramService) MessageHandler {
	return &messageHandler{api: api, bookingSvc: bookingSvc, telegramSvc: telegramSvc}
}

func (m *messageHandler) Handle(msg *tgbot.Message) error {
	if msg.Contact != nil {
		return m.handlePhone(msg)
	}

	message := tgbot.NewMessage(msg.Chat.ID, consts.DontKnowHowToAnswer)
	if _, err := m.api.Send(message); err != nil {
		return err
	}
	return nil
}

func (m *messageHandler) handlePhone(msg *tgbot.Message) error {
	booking, err := m.bookingSvc.FindActiveByChatID(msg.Chat.ID)
	if err != nil {
		return fmt.Errorf("[handle_phone] %w", err)
	}
	if err := m.bookingSvc.SetPhone(msg.Contact.PhoneNumber, msg.Chat.ID); err != nil {
		return fmt.Errorf("[handle_phone] %w", err)
	}
	if err := m.telegramSvc.ProcessPhone(booking, msg.Chat.ID); err != nil {
		return fmt.Errorf("[handle_phone] %w", err)
	}
	return nil
}
