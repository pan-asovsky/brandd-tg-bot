package handler

import (
	"errors"
	"log"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constant"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/admin"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/user"
	ihandler "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/handler"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/provider"
)

type updateHandler struct {
	telegramProvider                   iprovider.TelegramProvider
	userMessage, adminMessage, command ihandler.MessageHandler
	userCallback, adminCallback        ihandler.CallbackHandler
}

func NewUpdateHandler(container provider.Container) ihandler.UpdateHandler {
	return &updateHandler{
		telegramProvider: container.Telegram,
		command:          NewCommandHandler(container.Telegram, container.Service),
		userCallback:     user.NewUserCallbackHandler(container),
		adminCallback:    admin.NewAdminCallbackHandler(container),
		userMessage:      user.NewUserMessageHandler(container),
		adminMessage:     admin.NewAdminMessageHandler(container.Telegram, container.Service),
	}
}

func (uh *updateHandler) Handle(update *tgapi.Update) error {
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

func (uh *updateHandler) handleCallback(callback *tgapi.CallbackQuery) error {
	uh.telegramProvider.Common().AfterCallbackCleanup(callback)
	log.Printf("[handle_callback] callback received: %s", callback.Data)

	data := callback.Data
	switch {
	case strings.HasPrefix(data, consts.Flow):
		return uh.handleFlow(callback)

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

func (uh *updateHandler) handleFlow(callback *tgapi.CallbackQuery) error {
	cut, ok := strings.CutPrefix(callback.Data, consts.Flow)
	if !ok {
		return errors.New("[handle_flow] cut prefix failed for: " + callback.Data)
	}
	switch cut {
	case consts.ADMIN:
		return uh.telegramProvider.Admin().StartMenu(callback.Message.Chat.ID)
	case consts.USER:
		return uh.telegramProvider.User().StartMenu(callback.Message.Chat.ID)
	}

	return nil
}
