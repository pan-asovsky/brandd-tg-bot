package admin

import (
	"errors"
	"log"
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

	//todo: how to improve?
	if admflow.Bookings == payload {
		return ach.menuBookings(query)
	} else if strings.Contains(payload, admflow.PrefixStatistics) {
		parts := strings.Split(payload, "::")
		query.Data = parts[len(parts)-1]
		return ach.menuStatistics(query)
	} else if admflow.Settings == payload {
		return ach.menuSettings(query)
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
			return bookings[i].Date < bookings[j].Date
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
	log.Printf("[handle_statistics] callback: %s", query.Data)
	pf := ach.statistics.PeriodFactory()
	period := pf.FromLabel(stat.Label(query.Data))
	log.Printf("[handle_statistics] query.Data: %s, period: %s", query.Data, period.Label)

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
