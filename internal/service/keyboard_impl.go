package service

import (
	"fmt"
	"sort"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/rules"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type keyboardService struct {
	callbackService BuildCallbackService
}

func (s *keyboardService) GreetingKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		//tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.NewBookingBtn, consts.NewBookingCbk)),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.NewBookingBtn, s.callbackService.NewBooking())),
		//tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.MyBookingsBtn, consts.MyBookingsCbk)),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.MyBookingsBtn, s.callbackService.MyBookings())),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.CalendarBtn, consts.CalendarCbk)),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.HelpBtn, consts.HelpCbk)),
	)
}

func (s *keyboardService) DateKeyboard(bookings []AvailableBooking) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton

	row := tg.NewInlineKeyboardRow()
	for _, b := range bookings {
		row = append(row, tg.NewInlineKeyboardButtonData(
			b.Label,
			s.callbackService.Date(b.Date),
		))
	}

	rows = append(rows, row, s.backKeyboardRow(s.callbackService.Menu()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) ZoneKeyboard(zones model.Zone, date string) tg.InlineKeyboardMarkup {
	keys := make([]string, 0, len(zones))
	for k := range zones {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for i, zone := range keys {
		cb := s.callbackService.Zone(date, zone)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(zone, cb))

		if i%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackService.NewBooking()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) TimeKeyboard(ts []model.Timeslot, info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for i, t := range ts {
		time := fmt.Sprintf("%s-%s", t.Start, t.End)
		info.Time = time
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(time, s.callbackService.Time(info)))

		if i%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	date, err := utils.ParseDate(info.Date)
	if err != nil {
		fmt.Printf("[time_keyboard] %v", err)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackService.Date(date)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) ServiceKeyboard(types []model.ServiceType, info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	selectedServices := info.SelectedServices
	var rows [][]tg.InlineKeyboardButton

	for i := 0; i < len(types); i += 2 {
		var row []tg.InlineKeyboardButton

		if i < len(types) {
			t := types[i]
			buttonText := t.ServiceName
			if selectedServices[t.ServiceCode] {
				buttonText = "✅ " + buttonText
			}

			cb := s.callbackService.ServiceSelection(t.ServiceCode, info)
			row = append(row, tg.NewInlineKeyboardButtonData(buttonText, cb))
		}

		if i+1 < len(types) {
			t := types[i+1]
			buttonText := t.ServiceName
			if selectedServices[t.ServiceCode] {
				buttonText = "✅ " + buttonText
			}

			cb := s.callbackService.ServiceSelection(t.ServiceCode, info)
			row = append(row, tg.NewInlineKeyboardButtonData(buttonText, cb))
		}

		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	var controlRow []tg.InlineKeyboardButton
	if len(selectedServices) > 0 {
		var selectedTrue []string
		for svcCode, selected := range selectedServices {
			if selected {
				selectedTrue = append(selectedTrue, svcCode)
			}
		}

		sort.Strings(selectedTrue)

		serviceRules := rules.ServiceRules{}
		service := serviceRules.MapServices(selectedTrue)
		info.Service = service

		controlRow = append(controlRow, tg.NewInlineKeyboardButtonData(consts.ReadyBtn, s.callbackService.ServiceConfirmation(info)))
		rows = append(rows, controlRow)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackService.Zone(info.Date, info.Zone)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) RimsKeyboard(rims []string, info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	sort.Strings(rims)
	for i, rim := range rims {
		info.RimRadius = rim
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(rim, s.callbackService.Rim(info)))

		if i%3 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackService.Time(info)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) ConfirmKeyboard(info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.ConfirmBtn, consts.ConfirmBookingCbk),
			tg.NewInlineKeyboardButtonData(consts.RejectBtn, consts.RejectCbk),
		),
		s.backKeyboardRow(s.callbackService.ServiceConfirmation(info)),
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

func (s *keyboardService) EmptyMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.NewBookingBtn, s.callbackService.NewBooking()),
			tg.NewInlineKeyboardButtonData(consts.BackBtn, s.callbackService.Menu()),
		),
	)
}

func (s *keyboardService) ExistsMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.CancelBtn, s.callbackService.PreCancelBooking()),
			tg.NewInlineKeyboardButtonData(consts.BackBtn, s.callbackService.Menu()),
		),
	)
}

func (s *keyboardService) BookingCancellationKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.NoBtn, s.callbackService.NoCancelBooking()),
			tg.NewInlineKeyboardButtonData(consts.YesBtn, s.callbackService.CancelBooking()),
		),
	)
}

func (s *keyboardService) BackKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.BackBtn, s.callbackService.Menu()),
		),
	)
}

func (s *keyboardService) backKeyboardRow(callback string) []tg.InlineKeyboardButton {
	return tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(consts.BackBtn, fmt.Sprintf(callback)),
	)
}
