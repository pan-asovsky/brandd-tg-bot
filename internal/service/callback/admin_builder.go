package callback

import (
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
)

type adminCallbackBuilderService struct{}

func NewAdminCallbackBuilderService() icallback.AdminCallbackBuilderService {
	return &adminCallbackBuilderService{}
}

func (acbs *adminCallbackBuilderService) StartAdmin() string {
	return admflow.AdminFlowCbk
}

func (acbs *adminCallbackBuilderService) StartUser() string {
	return admflow.UserFlowCbk
}

func (acbs *adminCallbackBuilderService) Bookings() string {
	return admflow.BookingsCbk
}
func (acbs *adminCallbackBuilderService) Statistics() string {
	return admflow.StatisticsCbk
}
func (acbs *adminCallbackBuilderService) Settings() string {
	return admflow.SettingsCbk
}

func (acbs *adminCallbackBuilderService) Back(direction string) string {
	return admflow.AdminBackPrefix + direction
}
