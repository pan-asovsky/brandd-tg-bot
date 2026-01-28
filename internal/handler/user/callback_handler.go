package user

import (
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/user_flow"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/provider"
)

type userCallbackHandler struct {
	service      iprovider.ServiceProvider
	repo         iprovider.RepoProvider
	cache        iprovider.CacheProvider
	telegram     iprovider.TelegramProvider
	callback     iprovider.CallbackProvider
	notification iprovider.NotificationProvider
	handlers     map[string]CallbackFunc
}

func NewUserCallbackHandler(container provider.Container) ihandler.CallbackHandler {
	uch := &userCallbackHandler{
		service:      container.Service,
		repo:         container.Repo,
		cache:        container.Cache,
		telegram:     container.Telegram,
		callback:     container.Callback,
		notification: container.Notification,
		handlers:     map[string]CallbackFunc{},
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
