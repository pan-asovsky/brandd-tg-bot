package provider

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/telegram"
)

type telegramProvider struct {
	botAPI         *tgapi.BotAPI
	svcProvider    iprovider.ServiceProvider
	msgFmtProvider iprovider.MessageFormattingProvider
}

func NewTelegramProvider(botAPI *tgapi.BotAPI, svcProvider iprovider.ServiceProvider, msgFmtProvider iprovider.MessageFormattingProvider) iprovider.TelegramProvider {
	return &telegramProvider{botAPI: botAPI, svcProvider: svcProvider, msgFmtProvider: msgFmtProvider}
}

func (tp *telegramProvider) User() itg.TelegramUserService {
	return telegram.NewTelegramUserService(tp.svcProvider.UserKeyboard(), tp.svcProvider.DateTime(), tp.msgFmtProvider, tp.Common())
}

func (tp *telegramProvider) Admin() itg.TelegramAdminService {
	return telegram.NewTelegramAdminService(tp.Common(), tp.svcProvider.AdminKeyboard())
}

func (tp *telegramProvider) Common() itg.TelegramCommonService {
	return telegram.NewTelegramCommonService(tp.botAPI)
}
