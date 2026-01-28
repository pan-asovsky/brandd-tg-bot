package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/callback"
)

type callbackProvider struct{}

func NewCallbackProvider() iprovider.CallbackProvider {
	return &callbackProvider{}
}

func (cp *callbackProvider) UserCallbackParser() icallback.UserCallbackParserService {
	return callback.NewUserCallbackParserService()
}

func (cp *callbackProvider) AdminCallbackParser() icallback.AdminCallbackParserService {
	return callback.NewAdminCallbackParserService()
}

func (cp *callbackProvider) UserCallbackBuilder() icallback.UserCallbackBuilderService {
	return callback.NewUserCallbackBuilderService()
}

func (cp *callbackProvider) AdminCallbackBuilder() icallback.AdminCallbackBuilderService {
	return callback.NewAdminCallbackBuilderService()
}
