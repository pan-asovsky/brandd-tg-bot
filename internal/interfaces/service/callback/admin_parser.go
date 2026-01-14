package callback

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type AdminCallbackParserService interface {
	Parse(query *tgapi.CallbackQuery) (string, error)
}
