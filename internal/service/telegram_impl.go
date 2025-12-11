package service

import (
	"fmt"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type telegramService struct {
	kb    KeyboardService
	tgapi *api.BotAPI
}

func (t *telegramService) ProcessMenu(bookings []AvailableBooking, info *types.UserSessionInfo) error {
	kb := t.kb.DateKeyboard(bookings)
	return utils.WrapFunction(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.DateMsg, kb)
	})
}

func (t *telegramService) ProcessDate(zone model.Zone, info *types.UserSessionInfo) error {
	kb := t.kb.ZoneKeyboard(zone, info.Date)
	return utils.WrapFunction(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.ZoneMsg, kb)
	})
}

func (t *telegramService) ProcessZone(timeslots []model.Timeslot, info *types.UserSessionInfo) error {
	kb := t.kb.TimeKeyboard(timeslots, info)
	return utils.WrapFunction(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.TimeMsg, kb)
	})
}

func (t *telegramService) ProcessTime(types []model.ServiceType, info *types.UserSessionInfo) error {
	kb := t.kb.ServiceKeyboard(types, info.Time, info.Date)
	return utils.WrapFunction(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.ServiceMsg, kb)
	})
}

func (t *telegramService) ProcessServiceType(rims []string, info *types.UserSessionInfo) error {
	kb := t.kb.RimsKeyboard(rims, info.Service, info.Time, info.Date)
	return utils.WrapFunction(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.RimMsg, kb)
	})
}

func (t *telegramService) ProcessRimRadius(info *types.UserSessionInfo) error {
	kb := t.kb.RequestPhoneKeyboard()
	return utils.WrapFunction(func() error {
		return t.sendRequestPhoneMessage(info.ChatID, consts.RequestUserPhone, kb)
	})
}

func (t *telegramService) ProcessPhone(booking *model.Booking, chatID int64) error {
	if err := t.removeReplyKeyboard(chatID); err != nil {
		return utils.WrapError(err)
	}

	kb := t.kb.ConfirmKeyboard()
	msg, err := utils.FmtConfirmMsg(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunction(func() error {
		return t.sendKeyboardMessage(chatID, msg, kb)
	})
}

func (t *telegramService) ProcessConfirm(chatID int64, slot *model.Slot) error {
	date := strings.ReplaceAll(slot.Date, "-", ".")
	msg := fmt.Sprintf(consts.ConfirmMsg, date, slot.StartTime)
	return utils.WrapFunction(func() error {
		return t.sendMessage(chatID, msg)
	})
}

func (t *telegramService) ProcessPendingConfirm(chatID int64) error {
	return utils.WrapFunction(func() error {
		return t.sendMessage(chatID, consts.PendingConfirmMsg)
	})
}

func (t *telegramService) sendKeyboardMessage(chatID int64, text string, kb api.InlineKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := t.tgapi.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (t *telegramService) sendMessage(chatID int64, text string) error {
	msg := api.NewMessage(chatID, text)
	if _, err := t.tgapi.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (t *telegramService) sendRequestPhoneMessage(chatID int64, text string, kb api.ReplyKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := t.tgapi.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (t *telegramService) removeReplyKeyboard(chatID int64) error {
	msg := api.NewMessage(chatID, consts.UserPhoneSaved)
	msg.ReplyMarkup = api.ReplyKeyboardRemove{RemoveKeyboard: true}
	if _, err := t.tgapi.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}
