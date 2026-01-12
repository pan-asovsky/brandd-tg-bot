package telegram

import tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramCommonService interface {
	RemoveReplyKeyboard(chatID int64, message string) error
	AfterCallbackCleanup(cb *tgapi.CallbackQuery)
}
