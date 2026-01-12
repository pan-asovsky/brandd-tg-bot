package tg_svc

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
)

type telegramProvider struct {
	botAPI         *tgapi.BotAPI
	svcProvider    p.ServiceProvider
	msgFmtProvider p.MessageFormattingProvider
}

func NewTelegramProvider(botAPI *tgapi.BotAPI, svcProvider p.ServiceProvider, msgFmtProvider p.MessageFormattingProvider) p.TelegramProvider {
	return &telegramProvider{botAPI: botAPI, svcProvider: svcProvider, msgFmtProvider: msgFmtProvider}
}

func (tp *telegramProvider) User() i.TelegramUserService {
	return &telegramUserService{kb: tp.svcProvider.Keyboard(), dateTime: tp.svcProvider.DateTime(), msgFmtProvider: tp.msgFmtProvider, botAPI: tp.botAPI}
}

func (tp *telegramProvider) Admin() i.TelegramAdminService {
	return &adminService{botAPI: tp.botAPI}
}

func (tp *telegramProvider) Common() i.TelegramCommonService {
	return &commonService{botAPI: tp.botAPI}
}
