package tg_svc

import (
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	tg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type telegramUserService struct {
	kb             i.KeyboardService
	dateTime       i.DateTimeService
	msgFmtProvider p.MessageFormattingProvider
	tgCommon       tg.TelegramCommonService
}

func (tcs *telegramUserService) RequestDate(bookings []entity.AvailableBooking, info *model.UserSessionInfo) error {
	kb := tcs.kb.DateKeyboard(bookings)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.DateMsg, kb)
	})
}

func (tcs *telegramUserService) RequestZone(zone entity.Zone, info *model.UserSessionInfo) error {
	kb := tcs.kb.ZoneKeyboard(zone, info.Date)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.ZoneMsg, kb)
	})
}

func (tcs *telegramUserService) RequestTime(timeslots []entity.Timeslot, info *model.UserSessionInfo) error {
	kb := tcs.kb.TimeKeyboard(timeslots, info)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.TimeMsg, kb)
	})
}

func (tcs *telegramUserService) RequestServiceTypes(types []entity.ServiceType, info *model.UserSessionInfo) error {
	kb := tcs.kb.ServiceKeyboard(types, info)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.ServiceMsg, kb)
	})
}

func (tcs *telegramUserService) RequestRimRadius(rims []string, info *model.UserSessionInfo) error {
	kb := tcs.kb.RimsKeyboard(rims, info)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.RimMsg, kb)
	})
}

func (tcs *telegramUserService) RequestPreConfirm(booking *entity.Booking, info *model.UserSessionInfo) error {
	kb := tcs.kb.ConfirmKeyboard(info)
	msg, err := tcs.msgFmtProvider.Booking().PreConfirm(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, msg, kb)
	})
}

func (tcs *telegramUserService) RequestUserPhone(info *model.UserSessionInfo) error {
	kb := tcs.kb.RequestPhoneKeyboard()
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendRequestPhoneMessage(info.ChatID, usflow.RequestUserPhone, kb)
	})
}

func (tcs *telegramUserService) ProcessConfirm(chatID int64, slot *entity.Slot) error {
	date, err := tcs.dateTime.FormatDate(slot.Date, "2006-01-02", "02.01.2006")
	if err != nil {
		return utils.WrapError(err)
	}

	msg := tcs.msgFmtProvider.Booking().Confirm(date, slot.StartTime)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendMessageHTMLMode(chatID, msg)
	})
}

func (tcs *telegramUserService) ProcessPendingConfirm(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendMessage(chatID, usflow.PendingConfirmMsg)
	})
}

func (tcs *telegramUserService) SendBookingRestrictionMessage(chatID int64, booking *entity.Booking) error {
	msg, err := tcs.msgFmtProvider.Booking().Restriction(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, msg, tcs.kb.BackKeyboard())
	})
}

func (tcs *telegramUserService) SendMyBookingsMessage(chatID int64, fn func() (*entity.Booking, error)) error {
	booking, err := fn()
	if err != nil || booking == nil {
		return utils.WrapFunctionError(func() error {
			return tcs.tgCommon.SendKeyboardMessage(chatID, usflow.NoActiveBookings, tcs.kb.EmptyMyBookingsKeyboard())
		})
	} else {
		return utils.WrapFunctionError(func() error {
			msg, err := tcs.msgFmtProvider.Booking().My(booking)
			if err != nil {
				return utils.WrapError(err)
			}
			return tcs.tgCommon.SendKeyboardMessageHTMLMode(chatID, msg, tcs.kb.ExistsMyBookingsKeyboard())
		})
	}
}

func (tcs *telegramUserService) StartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, usflow.GreetingMsg, tcs.kb.GreetingKeyboard())
	})
}

func (tcs *telegramUserService) SendPreCancelBookingMessage(chatID int64, date, time string) error {
	msg, err := tcs.msgFmtProvider.Booking().PreCancel(date, time)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, msg, tcs.kb.BookingCancellationKeyboard())
	})
}

func (tcs *telegramUserService) SendCancellationMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, usflow.BookingCancelled, tcs.kb.BackKeyboard())
	})
}

func (tcs *telegramUserService) SendCancelDenyMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, usflow.ThanksForNoLeave, tcs.kb.BackKeyboard())
	})
}

func (tcs *telegramUserService) NewBookingNotify(chatID int64, booking *entity.Booking) error {
	msg, err := tcs.msgFmtProvider.Admin().NewBookingNotify(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendMessageHTMLMode(chatID, msg)
	})
}
