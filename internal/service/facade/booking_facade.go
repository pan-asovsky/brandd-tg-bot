package facade

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BookingFacade interface {
	MainMenu(ctx context.Context, msg *tgbotapi.Message) error
}
