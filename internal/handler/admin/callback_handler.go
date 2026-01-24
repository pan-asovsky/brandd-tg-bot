package admin

import (
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/provider"
)

type adminCallbackHandler struct {
	serviceProvider  iprovider.ServiceProvider
	repoProvider     iprovider.RepoProvider
	tgProvider       iprovider.TelegramProvider
	callbackProvider iprovider.CallbackProvider
	keyboardProvider iprovider.KeyboardProvider
	handlers         map[string]CallbackFunc
}

func NewAdminCallbackHandler(container provider.Container) ihandler.CallbackHandler {
	ach := &adminCallbackHandler{
		serviceProvider:  container.ServiceProvider,
		repoProvider:     container.RepoProvider,
		tgProvider:       container.TelegramProvider,
		callbackProvider: container.CallbackProvider,
		keyboardProvider: container.KeyboardProvider,
		handlers:         map[string]CallbackFunc{},
	}

	ach.register(admflow.MenuPrefix, ach.handleMenu)
	ach.register(admflow.PrefixBooking, ach.handleBookings)
	ach.register(admflow.PrefixStatistics, ach.handleStatistics)
	ach.register(admflow.PrefixSettings, ach.handleSettings)
	ach.register(admflow.PrefixBack, ach.handleBack)
	ach.register(admflow.PrefixComplete, ach.handleComplete)
	ach.register(admflow.PrefixNoShow, ach.handleNoShow)
	ach.register(admflow.PrefixReject, ach.handleReject)

	return ach
}

type CallbackFunc func(cb *tgapi.CallbackQuery) error

func (ach *adminCallbackHandler) register(prefix string, handler CallbackFunc) {
	ach.handlers[prefix] = handler
}

func (ach *adminCallbackHandler) Handle(query *tgapi.CallbackQuery) error {
	for prefix, handler := range ach.handlers {
		if strings.HasPrefix(query.Data, prefix) {
			return handler(query)
		}
	}

	return nil
}
