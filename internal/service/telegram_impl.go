package service

import (
	"fmt"

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
	if err := t.sendKeyboardMessage(info.ChatID, consts.DateMsg, kb); err != nil {
		return fmt.Errorf("[process_menu] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessDate(zone model.Zone, info *types.UserSessionInfo) error {
	kb := t.kb.ZoneKeyboard(zone, info.Date)
	if err := t.sendKeyboardMessage(info.ChatID, consts.ZoneMsg, kb); err != nil {
		return fmt.Errorf("[process_date] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessZone(timeslots []model.Timeslot, info *types.UserSessionInfo) error {
	kb := t.kb.TimeKeyboard(timeslots, info)
	if err := t.sendKeyboardMessage(info.ChatID, consts.TimeMsg, kb); err != nil {
		return fmt.Errorf("[process_zone] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessTime(types []model.ServiceType, info *types.UserSessionInfo) error {
	kb := t.kb.ServiceKeyboard(types, info.Time, info.Date)
	if err := t.sendKeyboardMessage(info.ChatID, consts.ServiceMsg, kb); err != nil {
		return fmt.Errorf("[process_time] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessServiceType(rims []string, info *types.UserSessionInfo) error {
	kb := t.kb.RimsKeyboard(rims, info.Service, info.Time, info.Date)
	if err := t.sendKeyboardMessage(info.ChatID, consts.RimMsg, kb); err != nil {
		return fmt.Errorf("[process_] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessRimRadius(info *types.UserSessionInfo) error {
	kb := t.kb.RequestPhoneKeyboard()
	if err := t.sendRequestPhoneMessage(info.ChatID, consts.RequestUserPhone, kb); err != nil {
		return fmt.Errorf("[process_rim_radius] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessPhone(booking *model.Booking, chatID int64) error {
	if err := t.removeReplyKeyboard(chatID); err != nil {
		return fmt.Errorf("[process_phone] %w", err)
	}

	kb := t.kb.ConfirmKeyboard()
	if err := t.sendKeyboardMessage(chatID, utils.FmtRimMsg(booking), kb); err != nil {
		return fmt.Errorf("[process_phone] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessConfirm(chatID int64, slot *model.Slot) error {
	msg := fmt.Sprintf(consts.ConfirmMsg, slot.Date, fmt.Sprintf("%s-%s", slot.StartTime, slot.EndTime))
	if err := t.sendMessage(chatID, msg); err != nil {
		return fmt.Errorf("[process_confirm] %w", err)
	}
	return nil
}

func (t *telegramService) ProcessPendingConfirm(chatID int64) error {
	if err := t.sendMessage(chatID, consts.PendingConfirmMsg); err != nil {
		return fmt.Errorf("[process_pending_confirm] %w", err)
	}
	return nil
}

func (t *telegramService) sendKeyboardMessage(chatID int64, text string, kb api.InlineKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := t.tgapi.Send(msg); err != nil {
		return fmt.Errorf("[send_keyboard_message] %w", err)
	}
	return nil
}

func (t *telegramService) sendMessage(chatID int64, text string) error {
	msg := api.NewMessage(chatID, text)
	if _, err := t.tgapi.Send(msg); err != nil {
		return fmt.Errorf("[send_message] %w", err)
	}
	return nil
}

func (t *telegramService) sendRequestPhoneMessage(chatID int64, text string, kb api.ReplyKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := t.tgapi.Send(msg); err != nil {
		return fmt.Errorf("[send_request_phone_message] %w", err)
	}
	return nil
}

func (t *telegramService) removeReplyKeyboard(chatID int64) error {
	msg := api.NewMessage(chatID, "\n")
	msg.ReplyMarkup = api.ReplyKeyboardRemove{RemoveKeyboard: true}
	if _, err := t.tgapi.Send(msg); err != nil {
		return fmt.Errorf("[remove_reply_keyboard] %w", err)
	}

	return nil
}
