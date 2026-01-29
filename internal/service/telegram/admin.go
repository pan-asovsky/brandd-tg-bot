package telegram

import (
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	ikeyboard "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/keyboard"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type adminTelegramService struct {
	tgCommon  itg.TelegramCommonService
	kb        ikeyboard.AdminKeyboardService
	formatter iprovider.MessageFormatterProvider
	dateTime  isvc.DateTimeService
}

func NewTelegramAdminService(
	tgCommon itg.TelegramCommonService,
	kb ikeyboard.AdminKeyboardService,
	msgFmtProvider iprovider.MessageFormatterProvider,
	dateTime isvc.DateTimeService,
) itg.TelegramAdminService {
	return &adminTelegramService{tgCommon: tgCommon, kb: kb, formatter: msgFmtProvider, dateTime: dateTime}
}

func (ats *adminTelegramService) ChoiceMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return ats.tgCommon.SendKeyboardMessage(chatID, admflow.ChoiceContinueFlow, ats.kb.ChoiceFlowKeyboard())
	})
}

func (ats *adminTelegramService) StartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return ats.tgCommon.SendKeyboardMessage(chatID, admflow.AnyMsg, ats.kb.MainMenu())
	})
}

func (ats *adminTelegramService) BookingPreview(chatID int64, booking *entity.Booking) error {
	msg, err := ats.formatter.Booking().BookingPreview(booking)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return ats.tgCommon.SendKeyboardMessageHTMLMode(chatID, msg, ats.kb.BookingInfo(booking.ChatID, booking.ID))
	})
}

func (ats *adminTelegramService) ConfirmAction(chatID int64, info *model.BookingInfo) error {
	return utils.WrapFunctionError(func() error {
		return ats.tgCommon.SendKeyboardMessage(chatID, admflow.ConfirmTerminateBooking, ats.kb.ConfirmationKeyboard(info))
	})
}

func (ats *adminTelegramService) RejectAction(chatID int64, backDirection string) error {
	return utils.WrapFunctionError(func() error {
		return ats.tgCommon.SendKeyboardMessage(chatID, admflow.ActionRejected, ats.kb.BackKeyboard(backDirection))
	})
}

func (ats *adminTelegramService) NoActiveBookings(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		kb := ats.kb.BackKeyboard(admflow.AdminFlowCbk)
		return ats.tgCommon.SendKeyboardMessage(chatID, admflow.NoActiveBookings, kb)
	})
}

func (ats *adminTelegramService) Statistics(chatID int64, stats stat.Stats, p stat.Period) error {
	kb := ats.kb.Statistics(p.Label)
	msg := ats.formatter.Admin().Statistics(stats, p)
	return utils.WrapFunctionError(func() error {
		return ats.tgCommon.SendKeyboardMessageHTMLMode(chatID, msg, kb)
	})
}
