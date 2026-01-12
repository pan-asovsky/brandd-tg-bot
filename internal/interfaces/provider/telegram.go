package provider

import (
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
)

type TelegramProvider interface {
	Admin() i.TelegramAdminService
	User() i.TelegramUserService
	Common() i.TelegramCommonService
}
