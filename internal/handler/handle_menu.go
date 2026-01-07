package handler

import (
	"errors"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleMenu(query *api.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_menu]: invalid callback: " + query.Data)
	}

	switch payload {
	//case consts.New:
	//	return c.handleNew(query)
	//case consts.My:
	//	return c.handleMy(query)
	case consts.Help:
		return c.handleHelp(query)
	case consts.Calendar:
		return c.HandleCalendar(query)
	}
	return nil
}

//func (c *callbackHandler) handleNew(q *api.CallbackQuery) error {
//	info := &types.UserSessionInfo{ChatID: q.Message.Chat.ID}
//
//	booking, err := c.svcProvider.Booking().FindActiveByChatID(info.ChatID)
//	if booking != nil && err == nil {
//		return c.svcProvider.Telegram().SendBookingRestrictionMessage(info.ChatID, booking)
//	}
//
//	bookings := c.svcProvider.Slot().GetAvailableBookings()
//	return utils.WrapFunctionError(func() error {
//		return c.svcProvider.Telegram().RequestDate(bookings, info)
//	})
//}
//
//func (c *callbackHandler) handleMy(q *api.CallbackQuery) error {
//	chatID := q.Message.Chat.ID
//	return utils.WrapFunctionError(func() error {
//		return c.svcProvider.Telegram().SendMyBookingsMessage(chatID, func() (*model.Booking, error) {
//			return c.svcProvider.Booking().FindActiveByChatID(chatID)
//		})
//	})
//}

func (c *callbackHandler) handleHelp(q *api.CallbackQuery) error {
	return utils.WrapFunctionError(func() error { return c.svcProvider.Telegram().SendHelpMessage(q.Message.Chat.ID) })
}

func (c *callbackHandler) HandleCalendar(q *api.CallbackQuery) error {
	chatID := q.Message.Chat.ID
	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().SendCalendar(chatID, func() (*model.Booking, error) {
			return c.svcProvider.Booking().FindActiveByChatID(chatID)
		})
	})
}
