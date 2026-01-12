package user

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type userMessageHandler struct {
	tgapi         *tg.BotAPI
	svcProvider   p.ServiceProvider
	cacheProvider p.CacheProvider
	tgProvider    p.TelegramProvider
}

func NewUserMessageHandler(tgapi *tg.BotAPI, svcProvider p.ServiceProvider, cacheProvider p.CacheProvider, tgProvider p.TelegramProvider) i.MessageHandler {
	return &userMessageHandler{tgapi: tgapi, svcProvider: svcProvider, cacheProvider: cacheProvider, tgProvider: tgProvider}
}

func (m *userMessageHandler) Handle(msg *tg.Message) error {
	if msg.Contact != nil {
		return m.handlePhone(msg.Chat.ID, msg.Contact.PhoneNumber)
	}

	detected, isPhone := m.svcProvider.Phone().Detect(msg.Text)
	if isPhone {
		return m.handlePhone(msg.Chat.ID, detected)
	}

	message := tg.NewMessage(msg.Chat.ID, usflow.DontKnowHowToAnswer)
	if _, err := m.tgapi.Send(message); err != nil {
		return err
	}
	return nil
}

func (m *userMessageHandler) handlePhone(chatID int64, contactPhone string) error {
	if err := m.tgProvider.Common().RemoveReplyKeyboard(chatID, usflow.UserPhoneSaved); err != nil {
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

func (m *userMessageHandler) handleAutoConfirm(chatID int64) error {
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

func (m *userMessageHandler) getActiveSlot(chatID int64) (*entity.Slot, *entity.Booking, error) {
	var booking *entity.Booking
	var slot *entity.Slot

	booking, err := m.svcProvider.Booking().FindPending(chatID)
	if err != nil {
		return slot, booking, utils.WrapError(err)
	}

	slot, err = m.svcProvider.Slot().FindByDateAndTime(booking.Date, booking.Time)
	if err != nil {
		return slot, booking, utils.WrapError(err)
	}

	return slot, booking, nil
}

func (m *userMessageHandler) confirm(chatID int64, slot *entity.Slot) error {
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
		return m.tgProvider.User().ProcessConfirm(chatID, slot)
	})
}

func (m *userMessageHandler) notifyAdmins(booking *entity.Booking) error {
	admins := m.svcProvider.User().GetActiveAdmins()
	for _, admin := range admins {
		if err := m.tgProvider.User().NewBookingNotify(admin.ChatID, booking); err != nil {
			return utils.WrapError(err)
		}
	}

	return nil
}

func (m *userMessageHandler) handlePendingConfirm(chatID int64) error {
	if err := m.svcProvider.Booking().UpdateStatus(chatID, entity.NotConfirmed); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return m.tgProvider.User().ProcessPendingConfirm(chatID)
	})
}
