package service

import (
	"fmt"
	"sort"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"

	"github.com/pan-asovsky/brandd-tg-bot/internal/rules"
)

type keyboardService struct {
	callbackBuilding isvc.CallbackBuildingService
	dateTime         isvc.DateTimeService
}

func (ks *keyboardService) GreetingKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(usflow.NewBookingBtn, ks.callbackBuilding.NewBooking())),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(usflow.MyBookingsBtn, ks.callbackBuilding.MyBookings())),
	)
}

func (ks *keyboardService) DateKeyboard(bookings []entity.AvailableBooking) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton

	row := tg.NewInlineKeyboardRow()
	for _, b := range bookings {
		row = append(row, tg.NewInlineKeyboardButtonData(
			b.Label,
			ks.callbackBuilding.Date(b.Date),
		))
	}

	rows = append(rows, row, ks.backKeyboardRow(ks.callbackBuilding.Menu()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (ks *keyboardService) ZoneKeyboard(zones entity.Zone, date string) tg.InlineKeyboardMarkup {
	keys := make([]string, 0, len(zones))
	for k := range zones {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for x, zone := range keys {
		cb := ks.callbackBuilding.Zone(date, zone)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(zone, cb))

		if x%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, ks.backKeyboardRow(ks.callbackBuilding.NewBooking()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (ks *keyboardService) TimeKeyboard(ts []entity.Timeslot, info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for x, t := range ts {
		time := fmt.Sprintf("%ks-%ks", t.Start, t.End)
		info.Time = time
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(time, ks.callbackBuilding.Time(info)))

		if x%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	date, err := ks.dateTime.ParseDate(info.Date, "2006-01-02")
	if err != nil {
		fmt.Printf("[time_keyboard] %v", err)
	}

	rows = append(rows, ks.backKeyboardRow(ks.callbackBuilding.Date(date)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (ks *keyboardService) ServiceKeyboard(types []entity.ServiceType, info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
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

			cb := ks.callbackBuilding.ServiceSelection(t.ServiceCode, info)
			row = append(row, tg.NewInlineKeyboardButtonData(buttonText, cb))
		}

		if x+1 < len(types) {
			t := types[x+1]
			buttonText := t.ServiceName
			if selectedServices[t.ServiceCode] {
				buttonText = "✅ " + buttonText
			}

			cb := ks.callbackBuilding.ServiceSelection(t.ServiceCode, info)
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

		controlRow = append(controlRow, tg.NewInlineKeyboardButtonData(usflow.ReadyBtn, ks.callbackBuilding.ServiceConfirmation(info)))
		rows = append(rows, controlRow)
	}

	rows = append(rows, ks.backKeyboardRow(ks.callbackBuilding.Zone(info.Date, info.Zone)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (ks *keyboardService) RimsKeyboard(rims []string, info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	sort.Strings(rims)
	for x, rim := range rims {
		info.RimRadius = rim
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(rim, ks.callbackBuilding.Rim(info)))

		if x%3 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, ks.backKeyboardRow(ks.callbackBuilding.Time(info)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (ks *keyboardService) ConfirmKeyboard(info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.ConfirmBtn, usflow.ConfirmBookingCbk),
			tg.NewInlineKeyboardButtonData(usflow.RejectBtn, usflow.RejectCbk),
		),
		ks.backKeyboardRow(ks.callbackBuilding.ServiceConfirmation(info)),
	)
}

func (ks *keyboardService) RequestPhoneKeyboard() tg.ReplyKeyboardMarkup {
	kb := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.KeyboardButton{Text: usflow.ShareContactBtn, RequestContact: true},
		),
	)

	kb.ResizeKeyboard = true
	return kb
}

func (ks *keyboardService) EmptyMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.BackBtn, ks.callbackBuilding.Menu()),
			tg.NewInlineKeyboardButtonData(usflow.NewBookingBtn, ks.callbackBuilding.NewBooking()),
		),
	)
}

func (ks *keyboardService) ExistsMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.CancelBtn, ks.callbackBuilding.PreCancelBooking()),
			tg.NewInlineKeyboardButtonData(usflow.BackBtn, ks.callbackBuilding.Menu()),
		),
	)
}

func (ks *keyboardService) BookingCancellationKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.NoBtn, ks.callbackBuilding.NoCancelBooking()),
			tg.NewInlineKeyboardButtonData(usflow.YesBtn, ks.callbackBuilding.CancelBooking()),
		),
	)
}

func (ks *keyboardService) BackKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.BackBtn, ks.callbackBuilding.Menu()),
		),
	)
}

func (ks *keyboardService) AdminGreetingKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.StartUser, ks.callbackBuilding.StartUser()),
			tg.NewInlineKeyboardButtonData(admflow.StartAdmin, ks.callbackBuilding.StartAdmin()),
		),
	)
}

func (ks *keyboardService) backKeyboardRow(callback string) []tg.InlineKeyboardButton {
	return tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(usflow.BackBtn, callback),
	)
}
