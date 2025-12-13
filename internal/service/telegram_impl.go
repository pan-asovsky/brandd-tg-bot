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

func (t *telegramService) RequestDate(bookings []AvailableBooking, info *types.UserSessionInfo) error {
	kb := t.kb.DateKeyboard(bookings)
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.DateMsg, kb)
	})
}

func (t *telegramService) RequestZone(zone model.Zone, info *types.UserSessionInfo) error {
	kb := t.kb.ZoneKeyboard(zone, info.Date)
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.ZoneMsg, kb)
	})
}

func (t *telegramService) RequestTime(timeslots []model.Timeslot, info *types.UserSessionInfo) error {
	kb := t.kb.TimeKeyboard(timeslots, info)
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.TimeMsg, kb)
	})
}

func (t *telegramService) RequestServiceTypes(types []model.ServiceType, info *types.UserSessionInfo) error {
	kb := t.kb.ServiceKeyboardV2(types, info)
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.ServiceMsg, kb)
	})
}

func (t *telegramService) RequestRimRadius(rims []string, info *types.UserSessionInfo) error {
	kb := t.kb.RimsKeyboard(rims, info.Service, info.Time, info.Date)
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.RimMsg, kb)
	})
}

func (t *telegramService) RequestUserPhone(info *types.UserSessionInfo) error {
	kb := t.kb.RequestPhoneKeyboard()
	return utils.WrapFunctionError(func() error {
		return t.sendRequestPhoneMessage(info.ChatID, consts.RequestUserPhone, kb)
	})
}

func (t *telegramService) RequestPreConfirm(booking *model.Booking, chatID int64) error {
	kb := t.kb.ConfirmKeyboard()
	msg, err := utils.FmtConfirmMsg(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, msg, kb)
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

	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, msg, kb)
	})
}

func (t *telegramService) ProcessConfirm(chatID int64, slot *model.Slot) error {
	date := strings.ReplaceAll(slot.Date, "-", ".")
	msg := fmt.Sprintf(consts.ConfirmMsg, date, slot.StartTime)
	return utils.WrapFunctionError(func() error {
		return t.sendMessage(chatID, msg)
	})
}

func (t *telegramService) ProcessPendingConfirm(chatID int64) error {
	return utils.WrapFunctionError(func() error {
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

func (t *telegramService) sendEditedKeyboard(chatID int64, messageID int, kb api.InlineKeyboardMarkup) error {
	msg := api.NewEditMessageReplyMarkup(chatID, messageID, kb)
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
