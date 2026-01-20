package telegram

import (
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/keyboard"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type userTelegramService struct {
	kb             keyboard.UserKeyboardService
	dateTime       isvc.DateTimeService
	msgFmtProvider iprovider.MessageFormatterProvider
	tgCommon       itg.TelegramCommonService
}

func NewTelegramUserService(
	kb keyboard.UserKeyboardService,
	dateTime isvc.DateTimeService,
	msgFmtProvider iprovider.MessageFormatterProvider,
	tgCommon itg.TelegramCommonService,
) itg.TelegramUserService {
	return &userTelegramService{kb: kb, dateTime: dateTime, msgFmtProvider: msgFmtProvider, tgCommon: tgCommon}
}

func (uts *userTelegramService) RequestDate(bookings []entity.AvailableDate, info *model.UserSessionInfo) error {
	kb := uts.kb.DateKeyboard(bookings)
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(info.ChatID, usflow.DateMsg, kb)
	})
}

func (uts *userTelegramService) RequestZone(zone entity.Zone, info *model.UserSessionInfo) error {
	kb := uts.kb.ZoneKeyboard(zone, info.Date)
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(info.ChatID, usflow.ZoneMsg, kb)
	})
}

func (uts *userTelegramService) RequestTime(timeslots []entity.Timeslot, info *model.UserSessionInfo) error {
	kb := uts.kb.TimeKeyboard(timeslots, info)
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(info.ChatID, usflow.TimeMsg, kb)
	})
}

func (uts *userTelegramService) RequestServiceTypes(types []entity.ServiceType, info *model.UserSessionInfo) error {
	kb := uts.kb.ServiceKeyboard(types, info)
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(info.ChatID, usflow.ServiceMsg, kb)
	})
}

func (uts *userTelegramService) RequestRimRadius(rims []string, info *model.UserSessionInfo) error {
	kb := uts.kb.RimsKeyboard(rims, info)
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(info.ChatID, usflow.RimMsg, kb)
	})
}

func (uts *userTelegramService) RequestPreConfirm(booking *entity.Booking, info *model.UserSessionInfo) error {
	kb := uts.kb.ConfirmKeyboard(info)
	msg, err := uts.msgFmtProvider.Booking().PreConfirm(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(info.ChatID, msg, kb)
	})
}

func (uts *userTelegramService) RequestUserPhone(info *model.UserSessionInfo) error {
	kb := uts.kb.RequestPhoneKeyboard()
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendRequestPhoneMessage(info.ChatID, usflow.RequestUserPhone, kb)
	})
}

func (uts *userTelegramService) ProcessConfirm(chatID int64, slot *entity.Slot) error {
	date, err := uts.dateTime.FormatDate(slot.Date, "2006-01-02", "02.01.2006")
	if err != nil {
		return utils.WrapError(err)
	}

	msg := uts.msgFmtProvider.Booking().Confirm(date, slot.StartTime)
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendMessageHTMLMode(chatID, msg)
	})
}

func (uts *userTelegramService) ProcessPendingConfirm(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendMessage(chatID, usflow.PendingConfirmMsg)
	})
}

func (uts *userTelegramService) SendBookingRestrictionMessage(chatID int64, booking *entity.Booking) error {
	msg, err := uts.msgFmtProvider.Booking().Restriction(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(chatID, msg, uts.kb.BackKeyboard())
	})
}

func (uts *userTelegramService) SendMyBookingsMessage(chatID int64, fn func() (*entity.Booking, error)) error {
	booking, err := fn()
	if err != nil || booking == nil {
		return utils.WrapFunctionError(func() error {
			return uts.tgCommon.SendKeyboardMessage(chatID, usflow.NoActiveBookings, uts.kb.EmptyMyBookingsKeyboard())
		})
	} else {
		return utils.WrapFunctionError(func() error {
			msg, err := uts.msgFmtProvider.Booking().My(booking)
			if err != nil {
				return utils.WrapError(err)
			}
			return uts.tgCommon.SendKeyboardMessageHTMLMode(chatID, msg, uts.kb.ExistsMyBookingsKeyboard())
		})
	}
}

func (uts *userTelegramService) StartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(chatID, usflow.GreetingMsg, uts.kb.GreetingKeyboard())
	})
}

func (uts *userTelegramService) SendPreCancelBookingMessage(chatID int64, date, time string) error {
	msg, err := uts.msgFmtProvider.Booking().PreCancel(date, time)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(chatID, msg, uts.kb.BookingCancellationKeyboard())
	})
}

func (uts *userTelegramService) SendCancellationMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(chatID, usflow.BookingCancelled, uts.kb.BackKeyboard())
	})
}

func (uts *userTelegramService) SendCancelDenyMessage(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendKeyboardMessage(chatID, usflow.ThanksForNoLeave, uts.kb.BackKeyboard())
	})
}

func (uts *userTelegramService) NewBookingNotify(chatID int64, booking *entity.Booking) error {
	msg, err := uts.msgFmtProvider.Admin().NewBookingNotify(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return uts.tgCommon.SendMessageHTMLMode(chatID, msg)
	})
}
