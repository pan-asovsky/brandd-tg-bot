package admin

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleNoShow(query *tgapi.CallbackQuery) error {
	bookingInfo, err := ach.callbackProvider.AdminCallbackParser().ParseNoShow(query)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		switch bookingInfo.Status {
		case model.PreNoShow:
			return ach.handlePreNoShow(query.Message.Chat.ID, bookingInfo)
		case model.NoShow:
			return ach.handleConfirmClose(query.Message.Chat.ID, bookingInfo)
		default:
			return nil
		}
	})
}

func (ach *adminCallbackHandler) handlePreNoShow(chatID int64, info *model.BookingInfo) error {
	return utils.WrapFunctionError(func() error {
		return ach.tgProvider.Admin().ConfirmAction(chatID, info)
	})
}

func (ach *adminCallbackHandler) handleConfirmClose(chatID int64, info *model.BookingInfo) error {
	booking, err := ach.serviceProvider.Booking().Close(info)
	if err != nil {
		return utils.WrapError(err)
	}

	if err = ach.serviceProvider.Statistics().Add(booking); err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		kb := ach.kbProvider.AdminKeyboard().BackKeyboard(admflow.AdminPrefix + admflow.MenuPrefix + admflow.Bookings)
		return ach.tgProvider.Common().SendKeyboardMessage(chatID, admflow.Closed, kb)
	})
}
