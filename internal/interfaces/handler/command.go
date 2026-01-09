package interfaces

import api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type CommandHandler interface {
	Handle(msg *api.Message) error
}
