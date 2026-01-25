package admin

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleComplete(query *tgapi.CallbackQuery) error {
	log.Printf("[handle_complete] callback: %s", query.Data)

	bookingInfo, err := ach.callbackProvider.AdminCallbackParser().ParseComplete(query)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		switch bookingInfo.Status {
		case model.PreCompleted:
			return ach.handlePreComplete(query.Message.Chat.ID, bookingInfo)
		case model.Completed:
			return ach.handleConfirmClose(query.Message.Chat.ID, bookingInfo)
		default:
			return nil
		}
	})
}

func (ach *adminCallbackHandler) handlePreComplete(chatID int64, info *model.BookingInfo) error {
	log.Printf("[handle_pre_complete] info: %v", info)
	return utils.WrapFunctionError(func() error {
		return ach.tgProvider.Admin().ConfirmAction(chatID, info)
	})
}
