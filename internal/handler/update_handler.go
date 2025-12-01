package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	kb "github.com/pan-asovsky/brandd-tg-bot/internal/service/keyboard"
	slot "github.com/pan-asovsky/brandd-tg-bot/internal/service/slot"
)

type UpdateHandler struct {
	command  CommandHandler
	callback CallbackHandler
	message  MessageHandler
}

func NewUpdateHandler(
	api *tg.BotAPI,
	kb kb.KeyboardService,
	slot slot.SlotService,
	lockSvc service.LockService,
	svcRepo pg.ServiceRepo,
	priceRepo pg.PriceRepo,
) *UpdateHandler {
	return &UpdateHandler{
		command:  NewCommandHandler(api, kb),
		callback: NewCallbackHandler(api, kb, slot, lockSvc, svcRepo, priceRepo),
		message:  NewMessageHandler(api),
	}
}

func (h *UpdateHandler) Handle(update *tg.Update) error {
	switch {
	case update.Message != nil && update.Message.IsCommand():
		return h.command.Handle(update.Message)
	case update.CallbackQuery != nil:
		return h.callback.Handle(update.CallbackQuery)
	case update.Message != nil:
		return h.message.Handle(update.Message)
	}
	return nil
}
