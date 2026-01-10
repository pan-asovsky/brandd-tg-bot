package user

import (
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	svc "github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type userCallbackHandler struct {
	tgapi         *tg.BotAPI
	svcProvider   *svc.Provider
	pgProvider    *pg.Provider
	cacheProvider *cache.Provider
	handlers      map[string]CallbackFunc
}

func NewUserCallbackHandler(
	tgapi *tg.BotAPI,
	svcProvider *svc.Provider,
	pgProvider *pg.Provider,
	cacheProvider *cache.Provider,
) i.CallbackHandler {
	uch := &userCallbackHandler{
		tgapi:         tgapi,
		svcProvider:   svcProvider,
		pgProvider:    pgProvider,
		cacheProvider: cacheProvider,
		handlers:      map[string]CallbackFunc{},
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

type CallbackFunc func(cb *tg.CallbackQuery) error

func (c *userCallbackHandler) register(prefix string, handler CallbackFunc) {
	c.handlers[prefix] = handler
}

func (c *userCallbackHandler) Handle(query *tg.CallbackQuery) error {
	for prefix, handler := range c.handlers {
		if strings.HasPrefix(query.Data, prefix) {
			return handler(query)
		}
	}

	return nil
}
