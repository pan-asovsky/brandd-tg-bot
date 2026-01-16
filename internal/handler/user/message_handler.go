package user

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type userMessageHandler struct {
	svcProvider   iprovider.ServiceProvider
	cacheProvider iprovider.CacheProvider
	tgProvider    iprovider.TelegramProvider
}

func NewUserMessageHandler(container provider.Container) ihandler.MessageHandler {
	return &userMessageHandler{svcProvider: container.ServiceProvider, cacheProvider: container.CacheProvider, tgProvider: container.TelegramProvider}
}

func (umh *userMessageHandler) Handle(msg *tgapi.Message) error {
	if msg.Contact != nil {
		return umh.handlePhone(msg.Chat.ID, msg.Contact.PhoneNumber)
	}

	detected, isPhone := umh.svcProvider.Phone().Detect(msg.Text)
	if isPhone {
		return umh.handlePhone(msg.Chat.ID, detected)
	}

	return utils.WrapFunctionError(func() error {
		return umh.tgProvider.Common().SendMessage(msg.Chat.ID, usflow.DontKnowHowToAnswer)
	})
}

func (umh *userMessageHandler) handlePhone(chatID int64, contactPhone string) error {
	if err := umh.tgProvider.Common().RemoveReplyKeyboard(chatID, usflow.UserPhoneSaved); err != nil {
		return utils.WrapError(err)
	}

	phone, err := umh.svcProvider.Phone().Normalize(contactPhone)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = umh.svcProvider.Booking().SetPhone(phone, chatID); err != nil {
		return utils.WrapError(err)
	}

	auto, err := umh.svcProvider.Config().IsAutoConfirm()
	if err != nil {
		return utils.WrapError(err)
	}

	if auto {
		return utils.WrapFunctionError(func() error {
			return umh.handleAutoConfirm(chatID)
		})
	}

	return utils.WrapFunctionError(func() error {
		return umh.handlePendingConfirm(chatID)
	})
}

func (umh *userMessageHandler) handleAutoConfirm(chatID int64) error {
	slot, booking, err := umh.getActiveSlot(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = umh.confirm(chatID, slot); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return umh.notifyAdmins(booking)
	})
}

func (umh *userMessageHandler) getActiveSlot(chatID int64) (*entity.Slot, *entity.Booking, error) {
	var booking *entity.Booking
	var slot *entity.Slot

	booking, err := umh.svcProvider.Booking().FindPending(chatID)
	if err != nil {
		return slot, booking, utils.WrapError(err)
	}

	slot, err = umh.svcProvider.Slot().FindByDateTime(booking.Date, booking.Time)
	if err != nil {
		return slot, booking, utils.WrapError(err)
	}

	return slot, booking, nil
}

func (umh *userMessageHandler) confirm(chatID int64, slot *entity.Slot) error {
	if err := umh.svcProvider.Slot().MarkUnavailable(slot.Date, slot.StartTime); err != nil {
		return utils.WrapError(err)
	}

	if err := umh.svcProvider.Booking().AutoConfirm(chatID); err != nil {
		return utils.WrapError(err)
	}

	if err := umh.svcProvider.Lock().Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	if err := umh.cacheProvider.ServiceType().Clean(chatID); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return umh.tgProvider.User().ProcessConfirm(chatID, slot)
	})
}

func (umh *userMessageHandler) notifyAdmins(booking *entity.Booking) error {
	admins := umh.svcProvider.User().GetActiveAdmins()
	for _, admin := range admins {
		if err := umh.tgProvider.User().NewBookingNotify(admin.ChatID, booking); err != nil {
			return utils.WrapError(err)
		}
	}

	return nil
}

func (umh *userMessageHandler) handlePendingConfirm(chatID int64) error {
	if err := umh.svcProvider.Booking().UpdateStatus(chatID, entity.NotConfirmed); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return umh.tgProvider.User().ProcessPendingConfirm(chatID)
	})
}
