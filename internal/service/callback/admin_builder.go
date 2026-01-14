package callback

import (
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
)

type adminCallbackBuilderService struct{}

func NewAdminCallbackBuilderService() icallback.AdminCallbackBuilderService {
	return &adminCallbackBuilderService{}
}

func (acbs *adminCallbackBuilderService) StartAdmin() string {
	return "FLOW::ADMIN"
}

func (acbs *adminCallbackBuilderService) StartUser() string {
	return "FLOW::USER"
}
