package admin

import (
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interface/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/provider"
)

type adminCallbackHandler struct {
	service      iprovider.ServiceProvider
	repo         iprovider.RepoProvider
	telegram     iprovider.TelegramProvider
	callback     iprovider.CallbackProvider
	keyboard     iprovider.KeyboardProvider
	notification iprovider.NotificationProvider
	handlers     map[string]CallbackFunc
}

func NewAdminCallbackHandler(container provider.Container) ihandler.CallbackHandler {
	ach := &adminCallbackHandler{
		service:      container.ServiceProvider,
		repo:         container.RepoProvider,
		telegram:     container.TelegramProvider,
		callback:     container.CallbackProvider,
		keyboard:     container.KeyboardProvider,
		notification: container.NotificationProvider,
		handlers:     map[string]CallbackFunc{},
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
