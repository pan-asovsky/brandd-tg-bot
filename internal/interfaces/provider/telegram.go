package provider

import (
	itelegram "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
)

type TelegramProvider interface {
	Admin() itelegram.TelegramAdminService
	User() itelegram.TelegramUserService
	Common() itelegram.TelegramCommonService
}
