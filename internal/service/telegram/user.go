package tg_svc

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type telegramUserService struct {
	kb             i.KeyboardService
	dateTime       i.DateTimeService
	msgFmtProvider p.MessageFormattingProvider
	botAPI         *api.BotAPI
}

func (us *telegramUserService) RequestDate(bookings []entity.AvailableBooking, info *model.UserSessionInfo) error {
	kb := us.kb.DateKeyboard(bookings)
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(info.ChatID, usflow.DateMsg, kb)
	})
}

func (us *telegramUserService) RequestZone(zone entity.Zone, info *model.UserSessionInfo) error {
	kb := us.kb.ZoneKeyboard(zone, info.Date)
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(info.ChatID, usflow.ZoneMsg, kb)
	})
}

func (us *telegramUserService) RequestTime(timeslots []entity.Timeslot, info *model.UserSessionInfo) error {
	kb := us.kb.TimeKeyboard(timeslots, info)
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(info.ChatID, usflow.TimeMsg, kb)
	})
}

func (us *telegramUserService) RequestServiceTypes(types []entity.ServiceType, info *model.UserSessionInfo) error {
	kb := us.kb.ServiceKeyboard(types, info)
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(info.ChatID, usflow.ServiceMsg, kb)
	})
}

func (us *telegramUserService) RequestRimRadius(rims []string, info *model.UserSessionInfo) error {
	kb := us.kb.RimsKeyboard(rims, info)
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(info.ChatID, usflow.RimMsg, kb)
	})
}

func (us *telegramUserService) RequestPreConfirm(booking *entity.Booking, info *model.UserSessionInfo) error {
	kb := us.kb.ConfirmKeyboard(info)
	msg, err := us.msgFmtProvider.Booking().PreConfirm(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(info.ChatID, msg, kb)
	})
}

func (us *telegramUserService) RequestUserPhone(info *model.UserSessionInfo) error {
	kb := us.kb.RequestPhoneKeyboard()
	return utils.WrapFunctionError(func() error {
		return us.sendRequestPhoneMessage(info.ChatID, usflow.RequestUserPhone, kb)
	})
}

func (us *telegramUserService) ProcessConfirm(chatID int64, slot *entity.Slot) error {
	date, err := us.dateTime.FormatDate(slot.Date, "2006-01-02", "02.01.2006")
	if err != nil {
		return utils.WrapError(err)
	}

	msg := us.msgFmtProvider.Booking().Confirm(date, slot.StartTime)
	return utils.WrapFunctionError(func() error {
		return us.sendMessageHTMLMode(chatID, msg)
	})
}

func (us *telegramUserService) ProcessPendingConfirm(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return us.sendMessage(chatID, usflow.PendingConfirmMsg)
	})
}

func (us *telegramUserService) SendBookingRestrictionMessage(chatID int64, booking *entity.Booking) error {
	msg, err := us.msgFmtProvider.Booking().Restriction(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(chatID, msg, us.kb.BackKeyboard())
	})
}

func (us *telegramUserService) SendMyBookingsMessage(chatID int64, fn func() (*entity.Booking, error)) error {
	booking, err := fn()
	if err != nil || booking == nil {
		return utils.WrapFunctionError(func() error {
			return us.sendKeyboardMessage(chatID, usflow.NoActiveBookings, us.kb.EmptyMyBookingsKeyboard())
		})
	} else {
		return utils.WrapFunctionError(func() error {
			msg, err := us.msgFmtProvider.Booking().My(booking)
			if err != nil {
				return utils.WrapError(err)
			}
			return us.sendKeyboardMessageHTMLMode(chatID, msg, us.kb.ExistsMyBookingsKeyboard())
		})
	}
}

func (us *telegramUserService) StartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(chatID, usflow.GreetingMsg, us.kb.GreetingKeyboard())
	})
}

func (us *telegramUserService) SendPreCancelBookingMessage(chatID int64, date, time string) error {
	msg, err := us.msgFmtProvider.Booking().PreCancel(date, time)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(chatID, msg, us.kb.BookingCancellationKeyboard())
	})
}

func (us *telegramUserService) SendCancellationMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(chatID, usflow.BookingCancelled, us.kb.BackKeyboard())
	})
}

func (us *telegramUserService) SendCancelDenyMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return us.sendKeyboardMessage(chatID, usflow.ThanksForNoLeave, us.kb.BackKeyboard())
	})
}

func (us *telegramUserService) NewBookingNotify(chatID int64, booking *entity.Booking) error {
	msg, err := us.msgFmtProvider.Admin().NewBookingNotify(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return us.sendMessageHTMLMode(chatID, msg)
	})
}

func (us *telegramUserService) sendKeyboardMessage(chatID int64, text string, kb api.InlineKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := us.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (us *telegramUserService) sendEditedKeyboard(chatID int64, messageID int, kb api.InlineKeyboardMarkup) error {
	msg := api.NewEditMessageReplyMarkup(chatID, messageID, kb)
	if _, err := us.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (us *telegramUserService) sendMessage(chatID int64, text string) error {
	msg := api.NewMessage(chatID, text)
	if _, err := us.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (us *telegramUserService) sendMessageHTMLMode(chatID int64, text string) error {
	msg := api.NewMessage(chatID, text)
	msg.ParseMode = api.ModeHTML
	msg.DisableWebPagePreview = true
	if _, err := us.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (us *telegramUserService) sendKeyboardMessageHTMLMode(chatID int64, text string, kb api.InlineKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ParseMode = api.ModeHTML
	msg.DisableWebPagePreview = true
	msg.ReplyMarkup = kb

	if _, err := us.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (us *telegramUserService) sendRequestPhoneMessage(chatID int64, text string, kb api.ReplyKeyboardMarkup) error {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := us.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (us *telegramUserService) removeReplyKeyboard(chatID int64) error {
	msg := api.NewMessage(chatID, usflow.UserPhoneSaved)
	msg.ReplyMarkup = api.ReplyKeyboardRemove{RemoveKeyboard: true}
	if _, err := us.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func printKeyboard(keyboard [][]api.InlineKeyboardButton) {
	for x, row := range keyboard {
		fmt.Printf("Строка %d:\n", x+1)
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
