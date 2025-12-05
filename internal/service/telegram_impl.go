package service

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
)

type telegramService struct {
	kb    KeyboardService
	tgapi *api.BotAPI
}

func (t *telegramService) ProcessMenu(bookings []AvailableBooking, chatID int64) {
	kb := t.kb.DateKeyboard(bookings)
	sendKeyboardMsg(chatID, consts.DateMsg, kb, t.tgapi)
}

func (t *telegramService) ProcessDate() {
}

func (t *telegramService) ProcessZone() {
}

func (t *telegramService) ProcessTime() {
}

func (t *telegramService) ProcessServiceType() {
}

func (t *telegramService) ProcessRimRadius() {
}

func (t *telegramService) ProcessPhone() {
}

func sendKeyboardMsg(chatID int64, text string, kb api.InlineKeyboardMarkup, bot *api.BotAPI) {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := bot.Send(msg); err != nil {
		log.Printf("[send_keyboard_message] error sending message: %s", err)
	}
}

func sendMsg(chatID int64, text string, bot *api.BotAPI) {
	msg := api.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("[send_message] error sending message: %s", err)
	}
}

func sendRequestPhoneMsg(chatID int64, text string, kb api.ReplyKeyboardMarkup, bot *api.BotAPI) {
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := bot.Send(msg); err != nil {
		log.Printf("[send_request_phone_message] error sending message: %s", err)
	}
}
