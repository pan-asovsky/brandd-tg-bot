package handler

import (
	"log"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	svc "github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type callbackHandler struct {
	api         *api.BotAPI
	kb          svc.KeyboardService
	slot        svc.SlotService
	lockSvc     svc.LockService
	bookingSvc  svc.BookingService
	telegramSvc svc.TelegramService
	svcRepo     pg.ServiceRepo
	priceRepo   pg.PriceRepo
	cfgRepo     pg.ConfigRepo
	handlers    map[string]CallbackFunc
}

func NewCallbackHandler(
	api *api.BotAPI,
	kb svc.KeyboardService,
	slot svc.SlotService,
	lockSvc svc.LockService,
	bookingSvc svc.BookingService,
	telegramSvc svc.TelegramService,
	svcRepo pg.ServiceRepo,
	priceRepo pg.PriceRepo,
	cfgRepo pg.ConfigRepo,
) CallbackHandler {
	ch := &callbackHandler{
		api:         api,
		kb:          kb,
		slot:        slot,
		lockSvc:     lockSvc,
		svcRepo:     svcRepo,
		bookingSvc:  bookingSvc,
		telegramSvc: telegramSvc,
		priceRepo:   priceRepo,
		cfgRepo:     cfgRepo,
		handlers:    map[string]CallbackFunc{},
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

func (c *callbackHandler) Handle(query *api.CallbackQuery) error {
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
