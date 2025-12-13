package handler

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type MessageHandler interface {
	Handle(msg *tgbot.Message) error
}

type messageHandler struct {
	api         *tgbot.BotAPI
	svcProvider *service.Provider
}

func NewMessageHandler(api *tgbot.BotAPI, svcProvider *service.Provider) MessageHandler {
	return &messageHandler{api: api, svcProvider: svcProvider}
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
	if err := m.svcProvider.Booking().SetPhone(msg.Contact.PhoneNumber, msg.Chat.ID); err != nil {
		return utils.WrapError(err)
	}

	chatID := msg.Chat.ID
	auto, err := m.svcProvider.Config().IsAutoConfirm()
	if err != nil {
		return utils.WrapError(err)
	}

	if auto {
		return utils.WrapFunctionError(func() error {
			return m.handleAutoConfirm(chatID)
		})
	}

	return utils.WrapFunctionError(func() error {
		return m.handlePendingConfirm(chatID)
	})
}

func (m *messageHandler) handleAutoConfirm(chatID int64) error {
	slot, err := m.getActiveSlot(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return m.confirmAndNotify(chatID, slot)
	})
}

func (m *messageHandler) getActiveSlot(chatID int64) (*model.Slot, error) {
	booking, err := m.svcProvider.Booking().FindActiveByChatID(chatID)
	if err != nil {
		return nil, utils.WrapError(err)
	}
	return utils.WrapFunction(func() (*model.Slot, error) {
		return m.svcProvider.Slot().FindByDateAndTime(booking.Date, booking.Time)
	})
}

func (m *messageHandler) confirmAndNotify(chatID int64, slot *model.Slot) error {
	if err := m.svcProvider.Booking().Confirm(chatID); err != nil {
		return utils.WrapError(err)
	}
	return utils.WrapFunctionError(func() error {
		return m.svcProvider.Telegram().ProcessConfirm(chatID, slot)
	})
}

func (m *messageHandler) handlePendingConfirm(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return m.svcProvider.Telegram().ProcessPendingConfirm(chatID)
	})
}
