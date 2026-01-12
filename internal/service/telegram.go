package service

//import (
//	"fmt"
//
//	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
//	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
//	service3 "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
//	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
//
//	msgfmt "github.com/pan-asovsky/brandd-tg-bot/internal/service/msg_fmt"
//	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
//)
//
//type telegramService struct {
//	kb             service3.KeyboardService
//	dateTime       service3.DateTimeService
//	msgFmtProvider msgfmt.messageFormattingProviderService
//	tgapi          *api.BotAPI
//}
//
//func (t *telegramService) RequestDate(bookings []entity.AvailableBooking, info *model.UserSessionInfo) error {
//	kb := t.kb.DateKeyboard(bookings)
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(info.ChatID, usflow.DateMsg, kb)
//	})
//}
//
//func (t *telegramService) RequestZone(zone entity.Zone, info *model.UserSessionInfo) error {
//	kb := t.kb.ZoneKeyboard(zone, info.Date)
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(info.ChatID, usflow.ZoneMsg, kb)
//	})
//}
//
//func (t *telegramService) RequestTime(timeslots []entity.Timeslot, info *model.UserSessionInfo) error {
//	kb := t.kb.TimeKeyboard(timeslots, info)
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(info.ChatID, usflow.TimeMsg, kb)
//	})
//}
//
//func (t *telegramService) RequestServiceTypes(types []entity.ServiceType, info *model.UserSessionInfo) error {
//	kb := t.kb.ServiceKeyboard(types, info)
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(info.ChatID, usflow.ServiceMsg, kb)
//	})
//}
//
//func (t *telegramService) RequestRimRadius(rims []string, info *model.UserSessionInfo) error {
//	kb := t.kb.RimsKeyboard(rims, info)
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(info.ChatID, usflow.RimMsg, kb)
//	})
//}
//
//func (t *telegramService) RequestPreConfirm(booking *entity.Booking, info *model.UserSessionInfo) error {
//	kb := t.kb.ConfirmKeyboard(info)
//	msg, err := t.msgFmtProvider.Booking().PreConfirm(booking)
//	if err != nil {
//		return utils.WrapError(err)
//	}
//
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(info.ChatID, msg, kb)
//	})
//}
//
//func (t *telegramService) RequestUserPhone(info *model.UserSessionInfo) error {
//	kb := t.kb.RequestPhoneKeyboard()
//	return utils.WrapFunctionError(func() error {
//		return t.sendRequestPhoneMessage(info.ChatID, usflow.RequestUserPhone, kb)
//	})
//}
//
//func (t *telegramService) RemoveReplyKeyboard(chatID int64) error {
//	return utils.WrapFunctionError(func() error {
//		return t.removeReplyKeyboard(chatID)
//	})
//}
//
//func (t *telegramService) ProcessConfirm(chatID int64, slot *entity.Slot) error {
//	date, err := t.dateTime.FormatDate(slot.Date, "2006-01-02", "02.01.2006")
//	if err != nil {
//		return utils.WrapError(err)
//	}
//
//	msg := t.msgFmtProvider.Booking().Confirm(date, slot.StartTime)
//	return utils.WrapFunctionError(func() error {
//		return t.sendMessageHTMLMode(chatID, msg)
//	})
//}
//
//func (t *telegramService) ProcessPendingConfirm(chatID int64) error {
//	return utils.WrapFunctionError(func() error {
//		return t.sendMessage(chatID, usflow.PendingConfirmMsg)
//	})
//}
//
//func (t *telegramService) SendBookingRestrictionMessage(chatID int64, booking *entity.Booking) error {
//	msg, err := t.msgFmtProvider.Booking().Restriction(booking)
//	if err != nil {
//		return utils.WrapError(err)
//	}
//
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(chatID, msg, t.kb.BackKeyboard())
//	})
//}
//
//func (t *telegramService) SendMyBookingsMessage(chatID int64, fn func() (*entity.Booking, error)) error {
//	booking, err := fn()
//	if err != nil || booking == nil {
//		return utils.WrapFunctionError(func() error {
//			return t.sendKeyboardMessage(chatID, usflow.NoActiveBookings, t.kb.EmptyMyBookingsKeyboard())
//		})
//	} else {
//		return utils.WrapFunctionError(func() error {
//			msg, err := t.msgFmtProvider.Booking().My(booking)
//			if err != nil {
//				return utils.WrapError(err)
//			}
//			return t.sendKeyboardMessageHTMLMode(chatID, msg, t.kb.ExistsMyBookingsKeyboard())
//		})
//	}
//}
//
//func (t *telegramService) StartMenu(chatID int64) error {
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(chatID, usflow.GreetingMsg, t.kb.GreetingKeyboard())
//	})
//}
//
//func (t *telegramService) SendPreCancelBookingMessage(chatID int64, date, time string) error {
//	msg, err := t.msgFmtProvider.Booking().PreCancel(date, time)
//	if err != nil {
//		return utils.WrapError(err)
//	}
//
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(chatID, msg, t.kb.BookingCancellationKeyboard())
//	})
//}
//
//func (t *telegramService) SendCancellationMessage(chatID int64) error {
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(chatID, usflow.BookingCancelled, t.kb.BackKeyboard())
//	})
//}
//
//func (t *telegramService) SendCancelDenyMessage(chatID int64) error {
//	return utils.WrapFunctionError(func() error {
//		return t.sendKeyboardMessage(chatID, usflow.ThanksForNoLeave, t.kb.BackKeyboard())
//	})
//}
//
//func (t *telegramService) NewBookingNotify(chatID int64, booking *entity.Booking) error {
//	msg, err := t.msgFmtProvider.Admin().NewBookingNotify(booking)
//	if err != nil {
//		return utils.WrapError(err)
//	}
//
//	return utils.WrapFunctionError(func() error {
//		return t.sendMessageHTMLMode(chatID, msg)
//	})
//}
//
//func (t *telegramService) sendKeyboardMessage(chatID int64, text string, kb api.InlineKeyboardMarkup) error {
//	msg := api.NewMessage(chatID, text)
//	msg.ReplyMarkup = kb
//
//	if _, err := t.tgapi.Send(msg); err != nil {
//		return utils.WrapError(err)
//	}
//
//	return nil
//}
//
//func (t *telegramService) sendEditedKeyboard(chatID int64, messageID int, kb api.InlineKeyboardMarkup) error {
//	msg := api.NewEditMessageReplyMarkup(chatID, messageID, kb)
//	if _, err := t.tgapi.Send(msg); err != nil {
//		return utils.WrapError(err)
//	}
//
//	return nil
//}
//
//func (t *telegramService) sendMessage(chatID int64, text string) error {
//	msg := api.NewMessage(chatID, text)
//	if _, err := t.tgapi.Send(msg); err != nil {
//		return utils.WrapError(err)
//	}
//
//	return nil
//}
//
//func (t *telegramService) sendMessageHTMLMode(chatID int64, text string) error {
//	msg := api.NewMessage(chatID, text)
//	msg.ParseMode = api.ModeHTML
//	msg.DisableWebPagePreview = true
//	if _, err := t.tgapi.Send(msg); err != nil {
//		return utils.WrapError(err)
//	}
//
//	return nil
//}
//
//func (t *telegramService) sendKeyboardMessageHTMLMode(chatID int64, text string, kb api.InlineKeyboardMarkup) error {
//	msg := api.NewMessage(chatID, text)
//	msg.ParseMode = api.ModeHTML
//	msg.DisableWebPagePreview = true
//	msg.ReplyMarkup = kb
//
//	if _, err := t.tgapi.Send(msg); err != nil {
//		return utils.WrapError(err)
//	}
//
//	return nil
//}
//
//func (t *telegramService) sendRequestPhoneMessage(chatID int64, text string, kb api.ReplyKeyboardMarkup) error {
//	msg := api.NewMessage(chatID, text)
//	msg.ReplyMarkup = kb
//
//	if _, err := t.tgapi.Send(msg); err != nil {
//		return utils.WrapError(err)
//	}
//
//	return nil
//}
//
//func (t *telegramService) removeReplyKeyboard(chatID int64) error {
//	msg := api.NewMessage(chatID, usflow.UserPhoneSaved)
//	msg.ReplyMarkup = api.ReplyKeyboardRemove{RemoveKeyboard: true}
//	if _, err := t.tgapi.Send(msg); err != nil {
//		return utils.WrapError(err)
//	}
//
//	return nil
//}
//
//func printKeyboard(keyboard [][]api.InlineKeyboardButton) {
//	for x, row := range keyboard {
//		fmt.Printf("Строка %d:\n", x+1)
//		for j, btn := range row {
//			callbackData := "nil"
//			if btn.CallbackData != nil {
//				callbackData = *btn.CallbackData
//			}
//			fmt.Printf("  Кнопка %d: \"%s\" → Callback: \"%s\"\n",
//				j+1, btn.Text, callbackData)
//		}
//	}
//}
