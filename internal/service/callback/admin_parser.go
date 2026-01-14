package callback

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
)

type adminCallbackParserService struct{}

func NewAdminCallbackParserService() icallback.AdminCallbackParserService {
	return &adminCallbackParserService{}
}

func (acps *adminCallbackParserService) Parse(query *tgapi.CallbackQuery) (string, error) {
	return "", nil
}
