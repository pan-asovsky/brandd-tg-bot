package keyboard

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
)

type adminKeyboardService struct {
	cbBuilder icallback.AdminCallbackBuilderService
	dateTime  isvc.DateTimeService
}

func NewAdminKeyboardService(cbBuilder icallback.AdminCallbackBuilderService, dateTime isvc.DateTimeService) isvc.AdminKeyboardService {
	return &adminKeyboardService{cbBuilder: cbBuilder, dateTime: dateTime}
}

func (aks *adminKeyboardService) ChoiceFlowKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.StartUserBtn, aks.cbBuilder.StartUser()),
			tg.NewInlineKeyboardButtonData(admflow.StartAdminBtn, aks.cbBuilder.StartAdmin()),
		),
	)
}

func (aks *adminKeyboardService) MainMenu() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.BookingsBtn, aks.cbBuilder.Bookings()),
			tg.NewInlineKeyboardButtonData(admflow.StatisticsBtn, aks.cbBuilder.Statistics()),
			tg.NewInlineKeyboardButtonData(admflow.SettingsBtn, aks.cbBuilder.Settings()),
		),
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.BackBtn, aks.cbBuilder.Back(admflow.Flow)),
		),
	)
}

func (aks *adminKeyboardService) Bookings() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.BackBtn, aks.cbBuilder.Back(admflow.Menu)),
		),
	)
}

func (aks *adminKeyboardService) Statistics() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.BackBtn, aks.cbBuilder.Back(admflow.Menu)),
		),
	)
}

func (aks *adminKeyboardService) Settings() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.BackBtn, aks.cbBuilder.Back(admflow.Menu)),
		),
	)
}
