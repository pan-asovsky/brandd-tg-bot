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
			tg.NewInlineKeyboardButtonData(admflow.StartUser, aks.cbBuilder.StartUser()),
			tg.NewInlineKeyboardButtonData(admflow.StartAdmin, aks.cbBuilder.StartAdmin()),
		),
	)
}
