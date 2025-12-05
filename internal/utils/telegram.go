package utils

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendKeyboardMsg(chatID int64, text string, kb api.InlineKeyboardMarkup, bot *api.BotAPI) {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := bot.Send(msg); err != nil {
		log.Printf("[send_keyboard_message] error sending message: %s", err)
	}
}

func SendMsg(chatID int64, text string, bot *api.BotAPI) {
	msg := api.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("[send_message] error sending message: %s", err)
	}
}

func SendRequestPhoneMsg(chatID int64, text string, kb api.ReplyKeyboardMarkup, bot *api.BotAPI) {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := bot.Send(msg); err != nil {
		log.Printf("[send_request_phone_message] error sending message: %s", err)
	}
}
