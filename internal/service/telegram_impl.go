package service

import (
	"fmt"
	"time"

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
	kb := t.kb.ServiceKeyboard(types, info)
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.ServiceMsg, kb)
	})
}

func (t *telegramService) RequestRimRadius(rims []string, info *types.UserSessionInfo) error {
	kb := t.kb.RimsKeyboard(rims, info)
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, consts.RimMsg, kb)
	})
}

func (t *telegramService) RequestPreConfirm(booking *model.Booking, info *types.UserSessionInfo) error {
	kb := t.kb.ConfirmKeyboard(info)
	msg, err := utils.FmtConfirmMsg(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(info.ChatID, msg, kb)
	})
}

func (t *telegramService) RequestUserPhone(info *types.UserSessionInfo) error {
	kb := t.kb.RequestPhoneKeyboard()
	return utils.WrapFunctionError(func() error {
		return t.sendRequestPhoneMessage(info.ChatID, consts.RequestUserPhone, kb)
	})
}

func (t *telegramService) RemoveReplyKeyboard(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return t.removeReplyKeyboard(chatID)
	})
}

func (t *telegramService) ProcessConfirm(chatID int64, slot *model.Slot) error {
	date, err := utils.FormatDate(slot.Date)
	if err != nil {
		return utils.WrapError(err)
	}

	msg := fmt.Sprintf(consts.ConfirmMsg, date, slot.StartTime)
	return utils.WrapFunctionError(func() error {
		return t.sendMessageHTMLMode(chatID, msg)
	})
}

func (t *telegramService) ProcessPendingConfirm(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return t.sendMessage(chatID, consts.PendingConfirmMsg)
	})
}

func (t *telegramService) SendHelpMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, consts.HelpMessage, t.kb.BackKeyboard())
	})
}

func (t *telegramService) SendBookingRestrictionMessage(chatID int64, booking *model.Booking) error {
	msg, err := utils.FmtBookingRestrictionMsg(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, msg, t.kb.BackKeyboard())
	})
}

func (t *telegramService) SendMyBookingsMessage(chatID int64, fn func() (*model.Booking, error)) error {
	booking, err := fn()
	if err != nil || booking == nil {
		return utils.WrapFunctionError(func() error {
			return t.sendKeyboardMessage(chatID, consts.NoActiveBookings, t.kb.EmptyMyBookingsKeyboard())
		})
	} else {
		return utils.WrapFunctionError(func() error {
			msg, err := utils.FmtMyBookingMsg(booking)
			if err != nil {
				return utils.WrapError(err)
			}
			return t.sendKeyboardMessageHTMLMode(chatID, msg, t.kb.ExistsMyBookingsKeyboard())
		})
	}
}

func (t *telegramService) SendStartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, consts.GreetingMsg, t.kb.GreetingKeyboard())
	})
}

func (t *telegramService) SendCalendar(chatID int64, fn func() (*model.Booking, error)) error {
	booking, err := fn()
	if err != nil || booking == nil {
		return utils.WrapFunctionError(func() error {
			return t.sendKeyboardMessage(chatID, consts.NoBookingsCalendar, t.kb.BackKeyboard())
		})
	}

	startDate, err := utils.ParseDateTimeInMSKZone(booking.Date, booking.Time)
	if err != nil {
		return utils.WrapError(err)
	}
	ics := utils.GenerateICS(startDate, startDate.Add(1*time.Hour))

	return utils.WrapFunctionError(func() error {
		return t.sendCalendarFile(chatID, ics, t.kb.BackKeyboard())
	})
}

func (t *telegramService) SendPreCancelBookingMessage(chatID int64, date, time string) error {
	msg, err := utils.FmtBookingPreCancelMsg(date, time)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, msg, t.kb.BookingCancellationKeyboard())
	})
}

func (t *telegramService) SendCancellationMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, consts.BookingCancelled, t.kb.BackKeyboard())
	})
}

func (t *telegramService) SendCancelDenyMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return t.sendKeyboardMessage(chatID, consts.ThanksForNoLeave, t.kb.BackKeyboard())
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

func (t *telegramService) sendMessageHTMLMode(chatID int64, text string) error {
	msg := api.NewMessage(chatID, text)
	msg.ParseMode = api.ModeHTML
	msg.DisableWebPagePreview = true
	if _, err := t.tgapi.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (t *telegramService) sendKeyboardMessageHTMLMode(chatID int64, text string, kb api.InlineKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ParseMode = api.ModeHTML
	msg.DisableWebPagePreview = true
	msg.ReplyMarkup = kb

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

func (t *telegramService) sendCalendarFile(chatID int64, icsContent string, kb api.InlineKeyboardMarkup) error {
	file := api.FileBytes{
		Name:  "event.ics",
		Bytes: []byte(icsContent),
	}

	doc := api.NewDocument(chatID, file)
	doc.Caption = consts.CalendarReadyMsg
	doc.ParseMode = api.ModeHTML
	doc.ReplyMarkup = kb

	if _, err := t.tgapi.Send(doc); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func printKeyboard(keyboard [][]api.InlineKeyboardButton) {
	for i, row := range keyboard {
		fmt.Printf("Строка %d:\n", i+1)
		for j, btn := range row {
			callbackData := "nil"
			if btn.CallbackData != nil {
				callbackData = *btn.CallbackData
			}
			fmt.Printf("  Кнопка %d: \"%s\" → Callback: \"%s\"\n",
				j+1, btn.Text, callbackData)
		}
	}
}
