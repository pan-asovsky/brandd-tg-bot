package admin

import (
	"errors"
	"sort"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleMenu(query *tgapi.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_menu]: invalid callback: " + query.Data)
	}

	switch payload {
	case admflow.Bookings:
		return ach.menuBookings(query)
	case admflow.Statistics:
		return ach.menuStatistics(query)
	case admflow.Settings:
		return ach.menuSettings(query)
	}

	return nil
}

func (ach *adminCallbackHandler) menuBookings(query *tgapi.CallbackQuery) error {
	bookings, err := ach.serviceProvider.Booking().FindAllActive()
	if err != nil {
		return utils.WrapError(err)
	}

	sort.Slice(bookings, func(i, j int) bool {
		if bookings[i].Date != bookings[j].Date {
			return bookings[i].Date < bookings[j].Date
		}
		return bookings[i].Time < bookings[j].Time
	})

	return utils.WrapFunctionError(func() error {
		kb := ach.keyboardProvider.AdminKeyboard().Bookings(bookings)
		return ach.tgProvider.Common().SendKeyboardMessage(query.Message.Chat.ID, admflow.FutureBookings, kb)
	})
}

func (ach *adminCallbackHandler) menuStatistics(query *tgapi.CallbackQuery) error {
	return utils.WrapFunctionError(func() error {
		return ach.tgProvider.Common().SendKeyboardMessage(query.Message.Chat.ID, "Тут какая-нибудь статистика", ach.keyboardProvider.AdminKeyboard().Statistics())
	})
}

func (ach *adminCallbackHandler) menuSettings(query *tgapi.CallbackQuery) error {
	return utils.WrapFunctionError(func() error {
		return ach.tgProvider.Common().SendKeyboardMessage(query.Message.Chat.ID, "Тут настройки", ach.keyboardProvider.AdminKeyboard().Settings())
	})
}
