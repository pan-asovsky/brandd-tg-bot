package keyboard

import (
	"fmt"
	"sort"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
	keyboard2 "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/keyboard"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"

	"github.com/pan-asovsky/brandd-tg-bot/internal/rules"
)

type userKeyboardService struct {
	cbBuilder icallback.UserCallbackBuilderService
	dateTime  isvc.DateTimeService
}

func NewUserKeyboardService(cbBuilder icallback.UserCallbackBuilderService, dateTime isvc.DateTimeService) keyboard2.UserKeyboardService {
	return &userKeyboardService{cbBuilder: cbBuilder, dateTime: dateTime}
}

func (uks *userKeyboardService) GreetingKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(usflow.NewBookingBtn, uks.cbBuilder.NewBooking())),
		tg.NewInlineKeyboardRow(tg.NewInlineKeyboardButtonData(usflow.MyBookingsBtn, uks.cbBuilder.MyBookings())),
	)
}

func (uks *userKeyboardService) DateKeyboard(bookings []entity.AvailableDate) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton

	row := tg.NewInlineKeyboardRow()
	for _, b := range bookings {
		row = append(row, tg.NewInlineKeyboardButtonData(
			b.Label,
			uks.cbBuilder.Date(b.Date),
		))
	}

	rows = append(rows, row, uks.backKeyboardRow(uks.cbBuilder.Menu()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (uks *userKeyboardService) ZoneKeyboard(zones entity.Zone, date string) tg.InlineKeyboardMarkup {
	keys := make([]string, 0, len(zones))
	for k := range zones {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for x, zone := range keys {
		cb := uks.cbBuilder.Zone(date, zone)
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(zone, cb))

		if x%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, uks.backKeyboardRow(uks.cbBuilder.NewBooking()))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (uks *userKeyboardService) TimeKeyboard(ts []entity.Timeslot, info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	for x, t := range ts {
		time := fmt.Sprintf("%s-%s", t.Start, t.End)
		info.Time = time
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(time, uks.cbBuilder.Time(info)))

		if x%2 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	date, err := uks.dateTime.ParseDate(info.Date, "2006-01-02")
	if err != nil {
		fmt.Printf("[time_keyboard] %v", err)
	}

	rows = append(rows, uks.backKeyboardRow(uks.cbBuilder.Date(date)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (uks *userKeyboardService) ServiceKeyboard(types []entity.ServiceType, info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
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

			cb := uks.cbBuilder.ServiceSelection(t.ServiceCode, info)
			row = append(row, tg.NewInlineKeyboardButtonData(buttonText, cb))
		}

		if x+1 < len(types) {
			t := types[x+1]
			buttonText := t.ServiceName
			if selectedServices[t.ServiceCode] {
				buttonText = "✅ " + buttonText
			}

			cb := uks.cbBuilder.ServiceSelection(t.ServiceCode, info)
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

		controlRow = append(controlRow, tg.NewInlineKeyboardButtonData(usflow.ReadyBtn, uks.cbBuilder.ServiceConfirmation(info)))
		rows = append(rows, controlRow)
	}

	rows = append(rows, uks.backKeyboardRow(uks.cbBuilder.Zone(info.Date, info.Zone)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (uks *userKeyboardService) RimsKeyboard(rims []string, info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
	var rows [][]tg.InlineKeyboardButton
	var currentRow []tg.InlineKeyboardButton

	sort.Strings(rims)
	for x, rim := range rims {
		info.RimRadius = rim
		currentRow = append(currentRow, tg.NewInlineKeyboardButtonData(rim, uks.cbBuilder.Rim(info)))

		if x%3 == 1 {
			rows = append(rows, currentRow)
			currentRow = []tg.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, uks.backKeyboardRow(uks.cbBuilder.Time(info)))
	return tg.NewInlineKeyboardMarkup(rows...)
}

func (uks *userKeyboardService) ConfirmKeyboard(info *model.UserSessionInfo) tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.ConfirmBtn, usflow.ConfirmBookingCbk),
			tg.NewInlineKeyboardButtonData(usflow.RejectBtn, usflow.RejectCbk),
		),
		uks.backKeyboardRow(uks.cbBuilder.ServiceConfirmation(info)),
	)
}

func (uks *userKeyboardService) RequestPhoneKeyboard() tg.ReplyKeyboardMarkup {
	kb := tg.NewReplyKeyboard(
		tg.NewKeyboardButtonRow(
			tg.KeyboardButton{Text: usflow.ShareContactBtn, RequestContact: true},
		),
	)

	kb.ResizeKeyboard = true
	return kb
}

func (uks *userKeyboardService) EmptyMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.BackBtn, uks.cbBuilder.Menu()),
			tg.NewInlineKeyboardButtonData(usflow.NewBookingBtn, uks.cbBuilder.NewBooking()),
		),
	)
}

func (uks *userKeyboardService) ExistsMyBookingsKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.CancelBtn, uks.cbBuilder.PreCancelBooking()),
			tg.NewInlineKeyboardButtonData(usflow.BackBtn, uks.cbBuilder.Menu()),
		),
	)
}

func (uks *userKeyboardService) BookingCancellationKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.NoBtn, uks.cbBuilder.NoCancelBooking()),
			tg.NewInlineKeyboardButtonData(usflow.YesBtn, uks.cbBuilder.CancelBooking()),
		),
	)
}

func (uks *userKeyboardService) BackKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(usflow.BackBtn, uks.cbBuilder.Menu()),
		),
	)
}

func (uks *userKeyboardService) backKeyboardRow(callback string) []tg.InlineKeyboardButton {
	return tg.NewInlineKeyboardRow(
		tg.NewInlineKeyboardButtonData(usflow.BackBtn, callback),
	)
}
