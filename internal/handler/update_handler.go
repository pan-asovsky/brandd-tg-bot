package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	repo "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type UpdateHandler struct {
	command  CommandHandler
	callback CallbackHandler
	message  MessageHandler
}

func NewUpdateHandler(
	api *tg.BotAPI,
	svcProvider *service.Provider,
	repoProvider *repo.Provider,
) *UpdateHandler {
	return &UpdateHandler{
		command:  NewCommandHandler(api, svcProvider.Keyboard()),
		callback: NewCallbackHandler(api, svcProvider, repoProvider),
		message:  NewMessageHandler(api, svcProvider.Booking(), svcProvider.Telegram()),
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
