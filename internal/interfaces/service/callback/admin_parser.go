package callback

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type AdminCallbackParserService interface {
	Parse(query *tgapi.CallbackQuery) (string, error)
	ParseNoShow(query *tgapi.CallbackQuery) (*model.BookingInfo, error)
	ParseComplete(query *tgapi.CallbackQuery) (*model.BookingInfo, error)
}
