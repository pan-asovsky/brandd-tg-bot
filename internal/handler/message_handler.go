package handler

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
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
		return utils.WrapError(err)
	}

	if err := m.bookingSvc.SetPhone(msg.Contact.PhoneNumber, msg.Chat.ID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunction(func() error {
		return m.telegramSvc.ProcessPhone(booking, msg.Chat.ID)
	})
}
