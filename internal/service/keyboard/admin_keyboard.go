package keyboard

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type adminKeyboardService struct {
	callbackBuilding isvc.CallbackBuildingService
	dateTime         isvc.DateTimeService
}

func NewAdminKeyboardService(cbBuilding isvc.CallbackBuildingService, dateTime isvc.DateTimeService) isvc.AdminKeyboardService {
	return &adminKeyboardService{callbackBuilding: cbBuilding, dateTime: dateTime}
}

func (aks *adminKeyboardService) ChoiceFlowKeyboard() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData(admflow.StartUser, aks.callbackBuilding.StartUser()),
			tg.NewInlineKeyboardButtonData(admflow.StartAdmin, aks.callbackBuilding.StartAdmin()),
		),
	)
}
