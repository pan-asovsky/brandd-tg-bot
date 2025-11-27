package handler

import (
	"context"
	"log"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	rd "github.com/pan-asovsky/brandd-tg-bot/internal/repository/redis"
	kb "github.com/pan-asovsky/brandd-tg-bot/internal/service/keyboard"
	slot "github.com/pan-asovsky/brandd-tg-bot/internal/service/slot"
)

type callbackHandler struct {
	api      *api.BotAPI
	kb       kb.KeyboardService
	slot     slot.SlotService
	cache    rd.ZoneCache
	handlers map[string]CallbackFunc
}

func NewCallbackHandler(
	api *api.BotAPI,
	kb kb.KeyboardService,
	slot slot.SlotService,
	cache rd.ZoneCache,
) CallbackHandler {
	ch := &callbackHandler{
		api:      api,
		kb:       kb,
		slot:     slot,
		cache:    cache,
		handlers: map[string]CallbackFunc{},
	}

	ch.register(consts.PrefixMenu, ch.handleMenu)
	ch.register(consts.PrefixDate, ch.handleDate)
	ch.register(consts.PrefixZone, ch.handleZone)
	ch.register(consts.PrefixTime, ch.handleTime)
	ch.register(consts.PrefixService, ch.handleService)
	ch.register(consts.PrefixRim, ch.handleRim)
	ch.register(consts.PrefixConfirm, ch.handleConfirm)

	return ch
}

type CallbackFunc func(cb *api.CallbackQuery, data string) error

func (c *callbackHandler) register(prefix string, handler CallbackFunc) {
	c.handlers[prefix] = handler
}

func (c *callbackHandler) Handle(ctx context.Context, query *api.CallbackQuery) error {
	c.cleanup(query)

	for prefix, handler := range c.handlers {
		if strings.HasPrefix(query.Data, prefix) {
			cb := strings.TrimPrefix(query.Data, prefix)
			return handler(query, cb)
		}
	}

	log.Printf("callback handler not found for %s", query.Data)
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
