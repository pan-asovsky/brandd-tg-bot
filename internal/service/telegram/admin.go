package tg_svc

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type adminService struct {
	botAPI *tgapi.BotAPI
}

func (as *adminService) StartMenu(chatID int64) error {
	msg := tgapi.NewMessage(chatID, "Тут будет админская менюшечка")
	if _, err := as.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}
	return nil
}
