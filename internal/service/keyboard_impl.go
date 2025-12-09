package service

import (
	"fmt"
	"sort"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type keyboardService struct{}

func (s *keyboardService) GreetingKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.NewBookingBtn, consts.NewBookingCbk)),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.MyBookingsBtn, consts.MyBookingsCbk)),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.HelpBtn, consts.HelpCbk)),
	)
}

func (s *keyboardService) DateKeyboard(bookings []AvailableBooking) tg.InlineKeyboardMarkup {
	row := tg.NewInlineKeyboardRow()
	for _, b := range bookings {
		row = append(row, tg.NewInlineKeyboardButtonData(
			b.Label,
			consts.PrefixDate+b.Date.Format("2006-01-02"),
		))
	}
	return tg.NewInlineKeyboardMarkup(row)
}

func (s *keyboardService) ZoneKeyboard(zones model.Zone, date string) tg.InlineKeyboardMarkup {
	keys := make([]string, 0, len(zones))
	for k := range zones {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for i, zoneText := range keys {
		cb := fmt.Sprintf("%s%s/%s", consts.PrefixZone, zoneText, date)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(zoneText, cb))

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

func (s *keyboardService) TimeKeyboard(ts []model.Timeslot, info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for i, t := range ts {
		txt := fmt.Sprintf("%s-%s", t.Start, t.End)
		cb := fmt.Sprintf("%s%s/%s/%s", consts.PrefixTime, txt, info.Zone, info.Date)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(txt, cb))

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

func (s *keyboardService) ServiceKeyboard(types []model.ServiceType, time, date string) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for i, t := range types {
		cb := fmt.Sprintf("%s%s/%s/%s", consts.PrefixService, t.ServiceCode, time, date)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(t.ServiceName, cb))

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

func (s *keyboardService) RimsKeyboard(rims []string, svc, time, date string) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	sort.Strings(rims)
	for i, rim := range rims {
		cb := fmt.Sprintf("%s%s/%s/%s/%s", consts.PrefixRim, rim, svc, time, date)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(rim, cb))

		if i%3 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) ConfirmKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.ConfirmBtn, consts.ConfirmCbk),
			tg.NewInlineKeyboardButtonData(consts.RejectBtn, consts.RejectCbk),
		),
	)
}

func (s *keyboardService) RequestPhoneKeyboard() tg.ReplyKeyboardMarkup {
	kb := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.KeyboardButton{Text: consts.ShareContactBtn, RequestContact: true},
		),
	)
	kb.ResizeKeyboard = true
	kb.OneTimeKeyboard = true
	return kb
}
