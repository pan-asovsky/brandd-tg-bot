package telegram

import (
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type userTelegramService struct {
	kb             isvc.UserKeyboardService
	dateTime       isvc.DateTimeService
	msgFmtProvider iprovider.MessageFormatterProvider
	tgCommon       itg.TelegramCommonService
}

func NewTelegramUserService(
	kb isvc.UserKeyboardService,
	dateTime isvc.DateTimeService,
	msgFmtProvider iprovider.MessageFormatterProvider,
	tgCommon itg.TelegramCommonService,
) itg.TelegramUserService {
	return &userTelegramService{kb: kb, dateTime: dateTime, msgFmtProvider: msgFmtProvider, tgCommon: tgCommon}
}

func (tcs *userTelegramService) RequestDate(bookings []entity.AvailableDate, info *model.UserSessionInfo) error {
	kb := tcs.kb.DateKeyboard(bookings)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.DateMsg, kb)
	})
}

func (tcs *userTelegramService) RequestZone(zone entity.Zone, info *model.UserSessionInfo) error {
	kb := tcs.kb.ZoneKeyboard(zone, info.Date)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.ZoneMsg, kb)
	})
}

func (tcs *userTelegramService) RequestTime(timeslots []entity.Timeslot, info *model.UserSessionInfo) error {
	kb := tcs.kb.TimeKeyboard(timeslots, info)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.TimeMsg, kb)
	})
}

func (tcs *userTelegramService) RequestServiceTypes(types []entity.ServiceType, info *model.UserSessionInfo) error {
	kb := tcs.kb.ServiceKeyboard(types, info)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.ServiceMsg, kb)
	})
}

func (tcs *userTelegramService) RequestRimRadius(rims []string, info *model.UserSessionInfo) error {
	kb := tcs.kb.RimsKeyboard(rims, info)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, usflow.RimMsg, kb)
	})
}

func (tcs *userTelegramService) RequestPreConfirm(booking *entity.Booking, info *model.UserSessionInfo) error {
	kb := tcs.kb.ConfirmKeyboard(info)
	msg, err := tcs.msgFmtProvider.Booking().PreConfirm(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(info.ChatID, msg, kb)
	})
}

func (tcs *userTelegramService) RequestUserPhone(info *model.UserSessionInfo) error {
	kb := tcs.kb.RequestPhoneKeyboard()
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendRequestPhoneMessage(info.ChatID, usflow.RequestUserPhone, kb)
	})
}

func (tcs *userTelegramService) ProcessConfirm(chatID int64, slot *entity.Slot) error {
	date, err := tcs.dateTime.FormatDate(slot.Date, "2006-01-02", "02.01.2006")
	if err != nil {
		return utils.WrapError(err)
	}

	msg := tcs.msgFmtProvider.Booking().Confirm(date, slot.StartTime)
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendMessageHTMLMode(chatID, msg)
	})
}

func (tcs *userTelegramService) ProcessPendingConfirm(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendMessage(chatID, usflow.PendingConfirmMsg)
	})
}

func (tcs *userTelegramService) SendBookingRestrictionMessage(chatID int64, booking *entity.Booking) error {
	msg, err := tcs.msgFmtProvider.Booking().Restriction(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, msg, tcs.kb.BackKeyboard())
	})
}

func (tcs *userTelegramService) SendMyBookingsMessage(chatID int64, fn func() (*entity.Booking, error)) error {
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

func (tcs *userTelegramService) StartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, usflow.GreetingMsg, tcs.kb.GreetingKeyboard())
	})
}

func (tcs *userTelegramService) SendPreCancelBookingMessage(chatID int64, date, time string) error {
	msg, err := tcs.msgFmtProvider.Booking().PreCancel(date, time)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, msg, tcs.kb.BookingCancellationKeyboard())
	})
}

func (tcs *userTelegramService) SendCancellationMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, usflow.BookingCancelled, tcs.kb.BackKeyboard())
	})
}

func (tcs *userTelegramService) SendCancelDenyMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendKeyboardMessage(chatID, usflow.ThanksForNoLeave, tcs.kb.BackKeyboard())
	})
}

func (tcs *userTelegramService) NewBookingNotify(chatID int64, booking *entity.Booking) error {
	msg, err := tcs.msgFmtProvider.Admin().NewBookingNotify(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return tcs.tgCommon.SendMessageHTMLMode(chatID, msg)
	})
}
