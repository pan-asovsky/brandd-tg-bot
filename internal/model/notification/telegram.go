package notification

import (
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type TelegramChannel struct {
	tgCommon itg.TelegramCommonService
}

func NewTelegramChannel(tgCommon itg.TelegramCommonService) *TelegramChannel {
	return &TelegramChannel{tgCommon: tgCommon}
}

func (tc *TelegramChannel) Send(chatID int64, msg string) error {
	return utils.WrapFunctionError(func() error {
		return tc.tgCommon.SendMessageHTMLMode(chatID, msg)
	})
}
