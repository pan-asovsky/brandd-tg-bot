package admin

import (
	"errors"
	"sort"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleMenu(query *tgapi.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_menu]: invalid callback: " + query.Data)
	}

	handlers := map[string]func(q *tgapi.CallbackQuery) error{
		admflow.Bookings: ach.menuBookings,
		admflow.Settings: ach.menuSettings,
	}

	if h, ok := handlers[payload]; ok {
		return h(query)
	}

	if strings.HasPrefix(payload, admflow.PrefixStatistics) {
		query.Data = strings.TrimPrefix(payload, admflow.PrefixStatistics)
		return ach.menuStatistics(query)
	}

	return nil
}

func (ach *adminCallbackHandler) menuBookings(query *tgapi.CallbackQuery) error {
	bookings, err := ach.service.Booking().FindAllActive()
	if err != nil {
		return utils.WrapError(err)
	}

	if bookings == nil || len(bookings) == 0 {
		return utils.WrapFunctionError(func() error {
			return ach.telegram.Admin().NoActiveBookings(query.Message.Chat.ID)
		})
	}

	sort.Slice(bookings, func(i, j int) bool {
		if bookings[i].Date != bookings[j].Date {
			return bookings[i].Date.Before(bookings[j].Date)
		}
		return bookings[i].Time < bookings[j].Time
	})

	return utils.WrapFunctionError(func() error {
		return ach.telegram.Common().SendKeyboardMessageHTMLMode(
			query.Message.Chat.ID,
			admflow.FutureBookings,
			ach.keyboard.AdminKeyboard().Bookings(bookings),
		)
	})
}

func (ach *adminCallbackHandler) menuStatistics(query *tgapi.CallbackQuery) error {
	pf := ach.statistics.PeriodFactory()
	period := pf.FromLabel(stat.Label(query.Data))

	stats, err := ach.statistics.Service().Calculate(period)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return ach.telegram.Admin().Statistics(query.Message.Chat.ID, stats, period)
	})
}

func (ach *adminCallbackHandler) menuSettings(query *tgapi.CallbackQuery) error {
	return utils.WrapFunctionError(func() error {
		return ach.telegram.Common().SendKeyboardMessage(
			query.Message.Chat.ID,
			"Тут настройки",
			ach.keyboard.AdminKeyboard().Settings(),
		)
	})
}
