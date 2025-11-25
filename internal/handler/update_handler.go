package handler

import (
	"context"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	rd "github.com/pan-asovsky/brandd-tg-bot/internal/repository/redis"
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
	cache rd.ZoneCache,
) *UpdateHandler {
	return &UpdateHandler{
		command:  NewCommandHandler(api, kb),
		callback: NewCallbackHandler(api, kb, slot, cache),
		message:  NewMessageHandler(api),
	}
}

func (h *UpdateHandler) Handle(ctx context.Context, update *tg.Update) error {
	switch {
	case update.Message != nil && update.Message.IsCommand():
		return h.command.Handle(ctx, update.Message)
	case update.CallbackQuery != nil:
		return h.callback.Handle(ctx, update.CallbackQuery)
	case update.Message != nil:
		return h.message.Handle(ctx, update.Message)
	}
	return nil
}
