package handler

import (
	"errors"
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/admin"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/user"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type updateHandler struct {
	tgapi                                     *tg.BotAPI
	message, command                          i.MessageHandler
	userCallbackHandler, adminCallbackHandler i.CallbackHandler
}

func NewUpdateHandler(tgapi *tg.BotAPI, svcProvider *service.Provider, pgProvider *pg.Provider, cacheProvider *cache.Provider) i.UpdateHandler {
	return &updateHandler{
		tgapi:                tgapi,
		command:              NewCommandHandler(tgapi, svcProvider),
		userCallbackHandler:  user.NewUserCallbackHandler(tgapi, svcProvider, pgProvider, cacheProvider),
		adminCallbackHandler: admin.NewAdminCallbackHandler(),
		message:              NewMessageHandler(tgapi, svcProvider, cacheProvider),
	}
}

func (h *updateHandler) Handle(update *tg.Update) error {
	if update == nil {
		return errors.New("update is nil")
	}

	if msg := update.Message; msg != nil && msg.IsCommand() {
		return h.command.Handle(msg)
	}

	if callback := update.CallbackQuery; callback != nil {
		return h.handleCallback(callback)
	}

	if msg := update.Message; msg != nil {
		return h.message.Handle(msg)
	}

	return nil
}

func (h *updateHandler) handleCallback(callback *tg.CallbackQuery) error {
	h.cleanup(callback)

	log.Printf("[handle_callback] callback received: %s", callback.Data)

	data := callback.Data
	switch {
	case strings.HasPrefix(data, admflow.AdminPrefix):
		cut, ok := strings.CutPrefix(data, admflow.AdminPrefix)
		if !ok {
			return errors.New("[handle_callback] invalid prefix: " + data)
		}
		callback.Data = cut
		log.Printf("[handle_admin] data: %s", callback.Data)
		return h.adminCallbackHandler.Handle(callback)

	case strings.HasPrefix(data, usflow.UserPrefix):
		cut, ok := strings.CutPrefix(data, usflow.UserPrefix)
		if !ok {
			return errors.New("[handle_callback] invalid prefix: " + data)
		}
		callback.Data = cut
		log.Printf("[handle_user] data: %s", callback.Data)
		return h.userCallbackHandler.Handle(callback)

	default:
		log.Printf("[handle_callback] unexpected prefix: %s", data)
	}

	return nil
}

func (h *updateHandler) cleanup(cb *tg.CallbackQuery) {
	h.answerCallback(cb.ID)

	if cb.Message != nil {
		h.deletePreviousMsg(cb.Message.Chat.ID, cb.Message.MessageID)
	}
}

func (h *updateHandler) answerCallback(callbackID string) {
	if _, err := h.tgapi.Request(tg.NewCallback(callbackID, "")); err != nil {
		log.Printf("error answer to callback %s: %v", callbackID, err)
	}
}

func (h *updateHandler) deletePreviousMsg(chatID int64, messageID int) {
	if _, err := h.tgapi.Request(tg.NewDeleteMessage(chatID, messageID)); err != nil {
		log.Printf("error delete previous message %d: %v", messageID, err)
	}
}
