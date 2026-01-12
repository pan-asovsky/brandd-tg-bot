package handler

import (
	"errors"
	"log"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/admin"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/user"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
)

type updateHandler struct {
	tgapi                              *tg.BotAPI
	userMessage, adminMessage, command i.MessageHandler
	userCallback, adminCallback        i.CallbackHandler
}

func NewUpdateHandler(tgapi *tg.BotAPI, svcProvider p.ServiceProvider, repoProvider p.RepoProvider, cacheProvider p.CacheProvider, tgProvider p.TelegramProvider) i.UpdateHandler {
	return &updateHandler{
		tgapi:         tgapi,
		command:       NewCommandHandler(tgProvider, svcProvider),
		userCallback:  user.NewUserCallbackHandler(tgapi, svcProvider, repoProvider, cacheProvider, tgProvider),
		adminCallback: admin.NewAdminCallbackHandler(),
		userMessage:   user.NewUserMessageHandler(tgapi, svcProvider, cacheProvider, tgProvider),
		adminMessage:  admin.NewAdminMessageHandler(svcProvider),
	}
}

func (uh *updateHandler) Handle(update *tg.Update) error {
	if update == nil {
		return errors.New("update is nil")
	}

	if msg := update.Message; msg != nil && msg.IsCommand() {
		return uh.command.Handle(msg)
	}

	if callback := update.CallbackQuery; callback != nil {
		return uh.handleCallback(callback)
	}

	if msg := update.Message; msg != nil {
		return uh.userMessage.Handle(msg)
	}

	return nil
}

func (uh *updateHandler) handleCallback(callback *tg.CallbackQuery) error {
	uh.cleanup(callback)

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
		return uh.adminCallback.Handle(callback)

	case strings.HasPrefix(data, usflow.UserPrefix):
		cut, ok := strings.CutPrefix(data, usflow.UserPrefix)
		if !ok {
			return errors.New("[handle_callback] invalid prefix: " + data)
		}
		callback.Data = cut
		return uh.userCallback.Handle(callback)

	default:
		log.Printf("[handle_callback] unexpected prefix: %s", data)
	}

	return nil
}

// todo: вынести это в Telegram ?(Util)? Service
func (uh *updateHandler) cleanup(cb *tg.CallbackQuery) {
	uh.answerCallback(cb.ID)

	if cb.Message != nil {
		uh.deletePreviousMsg(cb.Message.Chat.ID, cb.Message.MessageID)
	}
}

func (uh *updateHandler) answerCallback(callbackID string) {
	if _, err := uh.tgapi.Request(tg.NewCallback(callbackID, "")); err != nil {
		log.Printf("error answer to callback %s: %v", callbackID, err)
	}
}

func (uh *updateHandler) deletePreviousMsg(chatID int64, messageID int) {
	if _, err := uh.tgapi.Request(tg.NewDeleteMessage(chatID, messageID)); err != nil {
		log.Printf("error delete previous userMessage, adminMessage %d: %v", messageID, err)
	}
}
