package service

import (
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type telegramService struct {
	kb    KeyboardService
	tgapi *api.BotAPI
}

func (t *telegramService) ProcessMenu(bookings []AvailableBooking, chatID int64) {
	log.Printf("[process_menu] bookings %v, chatID %d", bookings, chatID)
	kb := t.kb.DateKeyboard(bookings)
	t.sendKeyboardMsg(chatID, consts.DateMsg, kb)
}

func (t *telegramService) ProcessDate(zone model.Zone, info *types.UserSessionInfo) {
	log.Printf("[process_date] zone %v, chatID %d", zone, info.ChatID)
	kb := t.kb.ZoneKeyboard(zone, info.Date)
	t.sendKeyboardMsg(info.ChatID, consts.ZoneMsg, kb)
}

func (t *telegramService) ProcessZone(timeslots []model.Timeslot, info *types.UserSessionInfo) {
	log.Printf("[process_zone] timeslots %v, chatID %d", timeslots, info.ChatID)
	kb := t.kb.TimeKeyboard(timeslots, info)
	t.sendKeyboardMsg(info.ChatID, consts.TimeMsg, kb)
}

func (t *telegramService) ProcessTime() {
}

func (t *telegramService) ProcessServiceType() {
}

func (t *telegramService) ProcessRimRadius() {
}

func (t *telegramService) ProcessPhone() {
}

func (t *telegramService) sendKeyboardMsg(chatID int64, text string, kb api.InlineKeyboardMarkup) {
	log.Printf("[send_keyboard_msg] chatID: %d, text: %s, keyboard: %v", chatID, text, kb)
	msg := api.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := t.tgapi.Send(msg); err != nil {
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
