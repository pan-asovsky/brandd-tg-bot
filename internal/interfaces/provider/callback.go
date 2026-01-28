package provider

import icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"

type CallbackProvider interface {
	UserCallbackParser() icallback.UserCallbackParserService
	AdminCallbackParser() icallback.AdminCallbackParserService
	UserCallbackBuilder() icallback.UserCallbackBuilderService
	AdminCallbackBuilder() icallback.AdminCallbackBuilderService
}
