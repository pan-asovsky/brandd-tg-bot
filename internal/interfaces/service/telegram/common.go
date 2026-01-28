package tg

import (
	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramCommonService interface {
	RemoveReplyKeyboard(chatID int64, message string) error
	AfterCallbackCleanup(cb *tgapi.CallbackQuery)

	SendKeyboardMessage(chatID int64, text string, kb tgapi.InlineKeyboardMarkup) error
	SendEditedKeyboard(chatID int64, messageID int, kb tgapi.InlineKeyboardMarkup) error
	SendMessage(chatID int64, text string) error
	SendMessageHTMLMode(chatID int64, text string) error
	SendKeyboardMessageHTMLMode(chatID int64, text string, kb tgapi.InlineKeyboardMarkup) error
	SendRequestPhoneMessage(chatID int64, text string, kb tgapi.ReplyKeyboardMarkup) error
}
