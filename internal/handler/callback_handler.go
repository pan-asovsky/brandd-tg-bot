package handler

import (
	"log"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	rd "github.com/pan-asovsky/brandd-tg-bot/internal/repository/redis"
	svc "github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type CallbackHandler interface {
	Handle(query *api.CallbackQuery) error
}

type callbackHandler struct {
	api              *api.BotAPI
	svcProvider      *svc.Provider
	pgProvider       *pg.Provider
	serviceTypeCache rd.ServiceTypeCacheService
	handlers         map[string]CallbackFunc
}

func NewCallbackHandler(api *api.BotAPI, svcProvider *svc.Provider, pgProvider *pg.Provider, serviceTypeCache rd.ServiceTypeCacheService) CallbackHandler {
	ch := &callbackHandler{
		api:              api,
		svcProvider:      svcProvider,
		pgProvider:       pgProvider,
		serviceTypeCache: serviceTypeCache,
		handlers:         map[string]CallbackFunc{},
	}

	ch.register(consts.PrefixMenu, ch.handleMenu)
	ch.register(consts.PrefixDate, ch.handleDate)
	ch.register(consts.PrefixZone, ch.handleZone)
	ch.register(consts.PrefixTime, ch.handleTime)
	ch.register(consts.PrefixServiceSelect, ch.handleServiceSelect)
	ch.register(consts.PrefixServiceConfirm, ch.handleServiceConfirm)
	ch.register(consts.PrefixRim, ch.handleRim)
	ch.register(consts.PrefixConfirm, ch.handleConfirm)
	ch.register(consts.PrefixBooking, ch.handleBooking)

	ch.register(consts.PrefixBack, ch.handleBack)

	return ch
}

type CallbackFunc func(cb *api.CallbackQuery) error

func (c *callbackHandler) register(prefix string, handler CallbackFunc) {
	c.handlers[prefix] = handler
}

func (c *callbackHandler) Handle(query *api.CallbackQuery) error {
	c.cleanup(query)

	for prefix, handler := range c.handlers {
		if strings.HasPrefix(query.Data, prefix) {
			//cd := strings.TrimPrefix(query.Data, prefix)
			return handler(query)
		}
		//return handler(query)
	}

	return nil
}

func (c *callbackHandler) cleanup(cb *api.CallbackQuery) {
	c.answerCallback(cb.ID)

	if cb.Message != nil {
		c.deletePreviousMsg(cb.Message.Chat.ID, cb.Message.MessageID)
	}
}

func (c *callbackHandler) answerCallback(callbackID string) {
	if _, err := c.api.Request(api.NewCallback(callbackID, "")); err != nil {
		log.Printf("error answer to callback %s: %v", callbackID, err)
	}
}

func (c *callbackHandler) deletePreviousMsg(chatID int64, messageID int) {
	if _, err := c.api.Request(api.NewDeleteMessage(chatID, messageID)); err != nil {
		log.Printf("error delete previous message %d: %v", messageID, err)
	}
}
