package handler

import (
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type messageHandler struct {
	api           *tgbot.BotAPI
	svcProvider   *service.Provider
	cacheProvider *cache.Provider
}

func NewMessageHandler(api *tgbot.BotAPI, svcProvider *service.Provider, cacheProvider *cache.Provider) i.MessageHandler {
	return &messageHandler{api: api, svcProvider: svcProvider, cacheProvider: cacheProvider}
}

func (m *messageHandler) Handle(msg *tgbot.Message) error {
	if msg.Contact != nil {
		return m.handlePhone(msg.Chat.ID, msg.Contact.PhoneNumber)
	}

	detected, isPhone := m.svcProvider.Phone().Detect(msg.Text)
	if isPhone {
		return m.handlePhone(msg.Chat.ID, detected)
	}

	message := tgbot.NewMessage(msg.Chat.ID, consts.DontKnowHowToAnswer)
	if _, err := m.api.Send(message); err != nil {
		return err
	}
	return nil
}

func (m *messageHandler) handlePhone(chatID int64, contactPhone string) error {
	if err := m.svcProvider.Telegram().RemoveReplyKeyboard(chatID); err != nil {
		return utils.WrapError(err)
	}

	phone, err := m.svcProvider.Phone().Normalize(contactPhone)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = m.svcProvider.Booking().SetPhone(phone, chatID); err != nil {
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
	slot, booking, err := m.getActiveSlot(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = m.confirm(chatID, slot); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return m.notifyAdmins(booking)
	})
}

func (m *messageHandler) getActiveSlot(chatID int64) (*entity.Slot, *entity.Booking, error) {
	var booking *entity.Booking
	var slot *entity.Slot

	booking, err := m.svcProvider.Booking().FindPending(chatID)
	if err != nil {
		return slot, booking, utils.WrapError(err)
	}

	log.Printf("[get_active_slot] booking: %v", booking)

	slot, err = m.svcProvider.Slot().FindByDateAndTime(booking.Date, booking.Time)
	if err != nil {
		return slot, booking, utils.WrapError(err)
	}

	log.Printf("[get_active_slot] slot: %v", slot)

	return slot, booking, nil
}

func (m *messageHandler) confirm(chatID int64, slot *entity.Slot) error {
	if err := m.svcProvider.Slot().MarkUnavailable(slot.Date, slot.StartTime); err != nil {
		return utils.WrapError(err)
	}

	if err := m.svcProvider.Booking().AutoConfirm(chatID); err != nil {
		return utils.WrapError(err)
	}

	if err := m.svcProvider.Lock().Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	if err := m.cacheProvider.ServiceType().Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return m.svcProvider.Telegram().ProcessConfirm(chatID, slot)
	})
}

func (m *messageHandler) notifyAdmins(booking *entity.Booking) error {
	admins := m.svcProvider.User().GetActiveAdmins()
	for _, admin := range admins {
		if err := m.svcProvider.Telegram().NewBookingNotify(admin.ChatID, booking); err != nil {
			return utils.WrapError(err)
		}
	}

	return nil
}

func (m *messageHandler) handlePendingConfirm(chatID int64) error {
	if err := m.svcProvider.Booking().UpdateStatus(chatID, entity.NotConfirmed); err != nil {
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
