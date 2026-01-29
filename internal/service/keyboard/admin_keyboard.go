package keyboard

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
	ikeyboard "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/keyboard"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
)

type adminKeyboardService struct {
	cbBuilder icallback.AdminCallbackBuilderService
	dateTime  isvc.DateTimeService
}

func NewAdminKeyboardService(cbBuilder icallback.AdminCallbackBuilderService, dateTime isvc.DateTimeService) ikeyboard.AdminKeyboardService {
	return &adminKeyboardService{cbBuilder: cbBuilder, dateTime: dateTime}
}

func (aks *adminKeyboardService) ChoiceFlowKeyboard() tgapi.InlineKeyboardMarkup {
	return tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.StartUserBtn, aks.cbBuilder.StartUser()),
			tgapi.NewInlineKeyboardButtonData(admflow.StartAdminBtn, aks.cbBuilder.StartAdmin()),
		),
	)
}

func (aks *adminKeyboardService) MainMenu() tgapi.InlineKeyboardMarkup {
	return tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.BookingsBtn, aks.cbBuilder.BookingsMenu()),
			tgapi.NewInlineKeyboardButtonData(admflow.StatisticsBtn, aks.cbBuilder.Statistics(stat.Today)),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.SettingsBtn, aks.cbBuilder.Settings()),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.BackBtn, aks.cbBuilder.Back(admflow.Flow)),
		),
	)
}

func (aks *adminKeyboardService) Bookings(bookings []entity.Booking) tgapi.InlineKeyboardMarkup {
	var rows [][]tgapi.InlineKeyboardButton
	var currentRow []tgapi.InlineKeyboardButton

	for x, booking := range bookings {
		view, err := aks.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
		if err != nil {
			continue
		}
		currentRow = append(currentRow, tgapi.NewInlineKeyboardButtonData(view, aks.cbBuilder.Booking(booking.ID)))

		if len(currentRow) == 2 || x == len(bookings)-1 {
			rows = append(rows, currentRow)
			currentRow = []tgapi.InlineKeyboardButton{}
		}
	}

	if len(currentRow) > 0 {
		rows = append(rows, currentRow)
	}

	rows = append(rows, aks.backKeyboardRow(aks.cbBuilder.Back(admflow.Menu)))
	return tgapi.NewInlineKeyboardMarkup(rows...)
}

func (aks *adminKeyboardService) Statistics(label stat.Label) tgapi.InlineKeyboardMarkup {
	//todo: keyboard [Сегодня], [Вчера], [Неделя], [Месяц]
	// по умолчанию выбрано [Сегодня], эта кнопка не видна
	// далее по callback кнопки меняются, исключая текущую

	todayBtn := tgapi.NewInlineKeyboardButtonData(admflow.TodayBtn, aks.cbBuilder.Statistics(stat.Today))
	yesterdayBtn := tgapi.NewInlineKeyboardButtonData(admflow.YesterdayBtn, aks.cbBuilder.Statistics(stat.Yesterday))
	weekBtn := tgapi.NewInlineKeyboardButtonData(admflow.WeekBtn, aks.cbBuilder.Statistics(stat.Week))
	monthBtn := tgapi.NewInlineKeyboardButtonData(admflow.MonthBtn, aks.cbBuilder.Statistics(stat.Month))

	var buttons = map[stat.Label]tgapi.InlineKeyboardButton{
		stat.Today:     todayBtn,
		stat.Yesterday: yesterdayBtn,
		stat.Week:      weekBtn,
		stat.Month:     monthBtn,
	}

	var row []tgapi.InlineKeyboardButton
	for l := range buttons {
		if l != label {
			row = append(row, buttons[l])
		}
	}

	var rows [][]tgapi.InlineKeyboardButton
	rows = append(rows, row)
	rows = append(rows, aks.backKeyboardRow(aks.cbBuilder.Back(admflow.Menu)))

	return tgapi.NewInlineKeyboardMarkup(rows...)
}

func (aks *adminKeyboardService) Settings() tgapi.InlineKeyboardMarkup {
	return tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.BackBtn, aks.cbBuilder.Back(admflow.Menu)),
		),
	)
}

func (aks *adminKeyboardService) BookingInfo(userChatID int64, bookingID int64) tgapi.InlineKeyboardMarkup {
	return tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonURL(admflow.StartChatBtn, aks.cbBuilder.Chat(userChatID)),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.CompleteBookingBtn, aks.cbBuilder.PreComplete(userChatID, bookingID)),
			tgapi.NewInlineKeyboardButtonData(admflow.NoShowBtn, aks.cbBuilder.PreNoShow(userChatID, bookingID)),
		),
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.BackBtn, admflow.BookingsCbk),
		),
	)
}

func (aks *adminKeyboardService) ConfirmationKeyboard(info *model.BookingInfo) tgapi.InlineKeyboardMarkup {
	return tgapi.NewInlineKeyboardMarkup(
		tgapi.NewInlineKeyboardRow(
			tgapi.NewInlineKeyboardButtonData(admflow.ConfirmBtn, aks.cbBuilder.Confirm(info)),
			tgapi.NewInlineKeyboardButtonData(admflow.RejectBtn, aks.cbBuilder.Reject(info)),
		),
	)
}

func (aks *adminKeyboardService) BackKeyboard(backDirection string) tgapi.InlineKeyboardMarkup {
	return tgapi.NewInlineKeyboardMarkup(
		aks.backKeyboardRow(backDirection),
	)
}

func (aks *adminKeyboardService) backKeyboardRow(callback string) []tgapi.InlineKeyboardButton {
	return tgapi.NewInlineKeyboardRow(
		tgapi.NewInlineKeyboardButtonData(admflow.BackBtn, callback),
	)
}
