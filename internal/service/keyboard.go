package service

import (
	"fmt"
	"sort"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/rules"
)

type keyboardService struct {
	callbackBuilding i.CallbackBuildingService
	dateTime         i.DateTimeService
}

func (s *keyboardService) GreetingKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.NewBookingBtn, s.callbackBuilding.NewBooking())),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(consts.MyBookingsBtn, s.callbackBuilding.MyBookings())),
	)
}

func (s *keyboardService) DateKeyboard(bookings []entity.AvailableBooking) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton

	row := tg.NewInlineKeyboardRow()
	for _, b := range bookings {
		row = append(row, tg.NewInlineKeyboardButtonData(
			b.Label,
			s.callbackBuilding.Date(b.Date),
		))
	}

	rows = append(rows, row, s.backKeyboardRow(s.callbackBuilding.Menu()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) ZoneKeyboard(zones entity.Zone, date string) tg.InlineKeyboardMarkup {
	keys := make([]string, 0, len(zones))
	for k := range zones {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for x, zone := range keys {
		cb := s.callbackBuilding.Zone(date, zone)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(zone, cb))

		if x%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackBuilding.NewBooking()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) TimeKeyboard(ts []entity.Timeslot, info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for x, t := range ts {
		time := fmt.Sprintf("%s-%s", t.Start, t.End)
		info.Time = time
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(time, s.callbackBuilding.Time(info)))

		if x%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	date, err := s.dateTime.ParseDate(info.Date, "2006-01-02")
	if err != nil {
		fmt.Printf("[time_keyboard] %v", err)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackBuilding.Date(date)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) ServiceKeyboard(types []entity.ServiceType, info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	selectedServices := info.SelectedServices
	var rows [][]tg.InlineKeyboardButton

	for x := 0; x < len(types); x += 2 {
		var row []tg.InlineKeyboardButton

		if x < len(types) {
			t := types[x]
			buttonText := t.ServiceName
			if selectedServices[t.ServiceCode] {
				buttonText = "✅ " + buttonText
			}

			cb := s.callbackBuilding.ServiceSelection(t.ServiceCode, info)
			row = append(row, tg.NewInlineKeyboardButtonData(buttonText, cb))
		}

		if x+1 < len(types) {
			t := types[x+1]
			buttonText := t.ServiceName
			if selectedServices[t.ServiceCode] {
				buttonText = "✅ " + buttonText
			}

			cb := s.callbackBuilding.ServiceSelection(t.ServiceCode, info)
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

		controlRow = append(controlRow, tg.NewInlineKeyboardButtonData(consts.ReadyBtn, s.callbackBuilding.ServiceConfirmation(info)))
		rows = append(rows, controlRow)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackBuilding.Zone(info.Date, info.Zone)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) RimsKeyboard(rims []string, info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	sort.Strings(rims)
	for x, rim := range rims {
		info.RimRadius = rim
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(rim, s.callbackBuilding.Rim(info)))

		if x%3 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, s.backKeyboardRow(s.callbackBuilding.Time(info)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (s *keyboardService) ConfirmKeyboard(info *types.UserSessionInfo) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.ConfirmBtn, consts.ConfirmBookingCbk),
			tg.NewInlineKeyboardButtonData(consts.RejectBtn, consts.RejectCbk),
		),
		s.backKeyboardRow(s.callbackBuilding.ServiceConfirmation(info)),
	)
}

func (s *keyboardService) RequestPhoneKeyboard() tg.ReplyKeyboardMarkup {
	kb := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.KeyboardButton{Text: consts.ShareContactBtn, RequestContact: true},
		),
	)

	kb.ResizeKeyboard = true
	return kb
}

func (s *keyboardService) EmptyMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.BackBtn, s.callbackBuilding.Menu()),
			tg.NewInlineKeyboardButtonData(consts.NewBookingBtn, s.callbackBuilding.NewBooking()),
		),
	)
}

func (s *keyboardService) ExistsMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.CancelBtn, s.callbackBuilding.PreCancelBooking()),
			tg.NewInlineKeyboardButtonData(consts.BackBtn, s.callbackBuilding.Menu()),
		),
	)
}

func (s *keyboardService) BookingCancellationKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.NoBtn, s.callbackBuilding.NoCancelBooking()),
			tg.NewInlineKeyboardButtonData(consts.YesBtn, s.callbackBuilding.CancelBooking()),
		),
	)
}

func (s *keyboardService) BackKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(consts.BackBtn, s.callbackBuilding.Menu()),
		),
	)
}

func (s *keyboardService) backKeyboardRow(callback string) []tg.InlineKeyboardButton {
	return tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(consts.BackBtn, callback),
	)
}
