package interfaces

import api "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type CallbackHandler interface {
	Handle(query *api.CallbackQuery) error
}
