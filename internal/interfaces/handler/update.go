package interfaces

import tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type UpdateHandler interface {
	Handle(update *tg.Update) error
}
