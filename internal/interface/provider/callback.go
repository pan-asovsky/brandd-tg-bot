package provider

import icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/callback"

type CallbackProvider interface {
	UserCallbackParser() icallback.UserCallbackParserService
	AdminCallbackParser() icallback.AdminCallbackParserService
	UserCallbackBuilder() icallback.UserCallbackBuilderService
	AdminCallbackBuilder() icallback.AdminCallbackBuilderService
}
