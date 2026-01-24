package admin

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleNoShow(query *tgapi.CallbackQuery) error {
	log.Printf("[handle_no_show] callback: %s", query.Data)

	bookingInfo, err := ach.callbackProvider.AdminCallbackParser().ParseNoShow(query)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		switch bookingInfo.Status {
		case model.PreNoShow:
			return ach.handlePreNoShow(query.Message.Chat.ID, bookingInfo)
		case model.NoShow:
			return ach.handleConfirmNoShow(query.Message.Chat.ID, bookingInfo)
		default:
			return nil
		}
	})
}

func (ach *adminCallbackHandler) handlePreNoShow(chatID int64, info *model.BookingInfo) error {
	log.Printf("[handle_pre_no_show] info: %v", info)
	return utils.WrapFunctionError(func() error {
		return ach.tgProvider.Admin().ConfirmNoShow(chatID, info)
	})
}

func (ach *adminCallbackHandler) handleConfirmNoShow(chatID int64, info *model.BookingInfo) error {
	log.Printf("[handle_confirm_no_show] info: %v", info)
	return nil
}
