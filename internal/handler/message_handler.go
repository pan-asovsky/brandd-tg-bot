package handler

import (
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type messageHandler struct {
	api              *tgbot.BotAPI
	svcProvider      *service.Provider
	serviceTypeCache cache.ServiceTypeCache
}

func NewMessageHandler(api *tgbot.BotAPI, svcProvider *service.Provider, serviceTypeCache cache.ServiceTypeCache) i.MessageHandler {
	return &messageHandler{api: api, svcProvider: svcProvider, serviceTypeCache: serviceTypeCache}
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
	chatID := msg.Chat.ID

	if err := m.svcProvider.Telegram().RemoveReplyKeyboard(chatID); err != nil {
		return utils.WrapError(err)
	}

	if err := m.svcProvider.Booking().SetPhone(msg.Contact.PhoneNumber, chatID); err != nil {
		return utils.WrapError(err)
	}

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

	if err = m.confirm(chatID, slot); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return m.notifyAdmins(chatID)
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

func (m *messageHandler) confirm(chatID int64, slot *model.Slot) error {
	if err := m.svcProvider.Slot().MarkUnavailable(slot.Date, slot.StartTime); err != nil {
		return utils.WrapError(err)
	}

	if err := m.svcProvider.Booking().AutoConfirm(chatID); err != nil {
		return utils.WrapError(err)
	}

	if err := m.svcProvider.Lock().Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	if err := m.serviceTypeCache.Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return m.svcProvider.Telegram().ProcessConfirm(chatID, slot)
	})
}

func (m *messageHandler) notifyAdmins(chatID int64) error {
	booking, err := m.svcProvider.Booking().FindActiveByChatID(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	admins := m.svcProvider.User().GetActiveAdmins()
	for _, admin := range admins {
		if err = m.svcProvider.Telegram().NewBookingNotify(admin.ChatID, booking); err != nil {
			return utils.WrapError(err)
		}
	}

	return nil
}

func (m *messageHandler) handlePendingConfirm(chatID int64) error {
	if err := m.svcProvider.Booking().UpdateStatus(chatID, model.NotConfirmed); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return m.svcProvider.Telegram().ProcessPendingConfirm(chatID)
	})
}

func (m *messageHandler) cleanup(chatID int64, messageID int) {
	if _, err := m.api.Request(tgbot.NewDeleteMessage(chatID, messageID)); err != nil {
		log.Printf("error delete previous message %d: %v", messageID, err)
	}
}
