package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type CommandHandler interface {
	Handle(ctx context.Context, msg *tgbotapi.Message) error
}
