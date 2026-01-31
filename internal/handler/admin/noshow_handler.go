package admin

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	notif "github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleNoShow(query *tgapi.CallbackQuery) error {
	bookingInfo, err := ach.callback.AdminCallbackParser().ParseNoShow(query)
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
		return ach.telegram.Admin().ConfirmAction(chatID, info)
	})
}

// handleConfirmClose todo: что-то надо подумать, чтобы корректно отправлять уведомления
func (ach *adminCallbackHandler) handleConfirmClose(chatID int64, info *model.BookingInfo) error {
	booking, err := ach.service.Booking().Close(info)
	if err != nil {
		return utils.WrapError(err)
	}

	//todo: что это делает в no_show?
	switch info.Status {
	case model.Completed:
		if err = ach.notification.Service().Notify(notif.Event{
			Type: notif.BookingCompleted,
			Data: booking,
		}); err != nil {
			return utils.WrapError(err)
		}
	}

	return utils.WrapFunctionError(func() error {
		kb := ach.keyboard.AdminKeyboard().BackKeyboard(admflow.AdminPrefix + admflow.MenuPrefix + admflow.Bookings)
		return ach.telegram.Common().SendKeyboardMessage(chatID, admflow.Closed, kb)
	})
}
