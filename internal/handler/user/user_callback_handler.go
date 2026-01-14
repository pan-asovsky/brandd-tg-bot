package user

import (
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
)

type userCallbackHandler struct {
	botAPI           *tgapi.BotAPI
	serviceProvider  iprovider.ServiceProvider
	repoProvider     iprovider.RepoProvider
	cacheProvider    iprovider.CacheProvider
	telegramProvider iprovider.TelegramProvider
	callbackProvider iprovider.CallbackProvider
	handlers         map[string]CallbackFunc
}

func NewUserCallbackHandler(
	botAPI *tgapi.BotAPI,
	serviceProvider iprovider.ServiceProvider,
	repoProvider iprovider.RepoProvider,
	cacheProvider iprovider.CacheProvider,
	telegramProvider iprovider.TelegramProvider,
	callbackProvider iprovider.CallbackProvider,
) ihandler.CallbackHandler {
	uch := &userCallbackHandler{
		botAPI:           botAPI,
		serviceProvider:  serviceProvider,
		repoProvider:     repoProvider,
		cacheProvider:    cacheProvider,
		telegramProvider: telegramProvider,
		callbackProvider: callbackProvider,
		handlers:         map[string]CallbackFunc{},
	}

	uch.register(usflow.PrefixDate, uch.handleDate)
	uch.register(usflow.PrefixZone, uch.handleZone)
	uch.register(usflow.PrefixTime, uch.handleTime)
	uch.register(usflow.PrefixServiceSelect, uch.handleServiceSelect)
	uch.register(usflow.PrefixServiceConfirm, uch.handleServiceConfirm)
	uch.register(usflow.PrefixRim, uch.handleRim)
	uch.register(usflow.PrefixConfirm, uch.handleConfirm)
	uch.register(usflow.PrefixBooking, uch.handleBooking)

	uch.register(usflow.PrefixBack, uch.handleBack)

	return uch
}

type CallbackFunc func(cb *tgapi.CallbackQuery) error

func (uch *userCallbackHandler) register(prefix string, handler CallbackFunc) {
	uch.handlers[prefix] = handler
}

func (uch *userCallbackHandler) Handle(query *tgapi.CallbackQuery) error {
	for prefix, handler := range uch.handlers {
		if strings.HasPrefix(query.Data, prefix) {
			return handler(query)
		}
	}

	return nil
}
