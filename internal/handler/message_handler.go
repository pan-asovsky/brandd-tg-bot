package handler

import (
	"context"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler interface {
	Handle(ctx context.Context, msg *tgbot.Message) error
}
