package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	repo "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type UpdateHandler struct {
	command  i.CommandHandler
	callback i.CallbackHandler
	message  i.MessageHandler
}

func NewUpdateHandler(api *tg.BotAPI, svcProvider *service.Provider, pgProvider *repo.Provider, serviceTypeCache cache.ServiceTypeCache) *UpdateHandler {
	return &UpdateHandler{
		command:  NewCommandHandler(api, svcProvider),
		callback: NewCallbackHandler(api, svcProvider, pgProvider, serviceTypeCache),
		message:  NewMessageHandler(api, svcProvider, serviceTypeCache),
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
