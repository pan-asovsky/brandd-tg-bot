package service

import (
	"sort"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	slot "github.com/pan-asovsky/brandd-tg-bot/internal/service/slot"
)

func NewKeyboard() KeyboardService {
	return &service{}
}

type service struct{}

func (s *service) GreetingKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.NewBookingBtn, consts.NewBookingCbk)),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.MyBookingsBtn, consts.MyBookingsCbk)),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.HelpBtn, consts.HelpCbk)),
	)
}

func (s *service) DateKeyboard(bookings []slot.AvailableBooking) tg.InlineKeyboardMarkup {
	row := tg.NewInlineKeyboardRow()
	for _, b := range bookings {
		row = append(row, tg.NewInlineKeyboardButtonData(
			b.Label,
			consts.PrefixDate+b.Date.Format("2006-01-02"),
		))
	}
	return tg.NewInlineKeyboardMarkup(row)
}

func (s *service) ZoneKeyboard(zones model.Zone) tg.InlineKeyboardMarkup {
	keys := make([]string, 0, len(zones))
	for k := range zones {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for i, zoneText := range keys {
		callbackData := consts.PrefixZone + zoneText
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(zoneText, callbackData))

		if i%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	return tg.NewInlineKeyboardMarkup(rows...)
}
