package provider

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/telegram"
)

type telegramProvider struct {
	botAPI           *tgapi.BotAPI
	svcProvider      iprovider.ServiceProvider
	keyboardProvider iprovider.KeyboardProvider
	msgFmtProvider   iprovider.MessageFormatterProvider
}

func NewTelegramProvider(
	botAPI *tgapi.BotAPI,
	svcProvider iprovider.ServiceProvider,
	keyboardProvider iprovider.KeyboardProvider,
	msgFmtProvider iprovider.MessageFormatterProvider,
) iprovider.TelegramProvider {
	return &telegramProvider{
		botAPI:           botAPI,
		svcProvider:      svcProvider,
		keyboardProvider: keyboardProvider,
		msgFmtProvider:   msgFmtProvider,
	}
}

func (tp *telegramProvider) User() itg.TelegramUserService {
	return telegram.NewTelegramUserService(
		tp.keyboardProvider.UserKeyboard(),
		tp.svcProvider.DateTime(),
		tp.msgFmtProvider,
		tp.Common(),
	)
}

func (tp *telegramProvider) Admin() itg.TelegramAdminService {
	return telegram.NewTelegramAdminService(
		tp.Common(),
		tp.keyboardProvider.AdminKeyboard(),
		tp.msgFmtProvider,
		tp.svcProvider.DateTime(),
	)
}

func (tp *telegramProvider) Common() itg.TelegramCommonService {
	return telegram.NewTelegramCommonService(tp.botAPI)
}
