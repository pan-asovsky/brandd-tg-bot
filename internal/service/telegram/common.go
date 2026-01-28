package telegram

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type commonTelegramService struct {
	botAPI *tgapi.BotAPI
}

func NewTelegramCommonService(botAPI *tgapi.BotAPI) itg.TelegramCommonService {
	return &commonTelegramService{botAPI: botAPI}
}

func (tcs *commonTelegramService) RemoveReplyKeyboard(chatID int64, message string) error {
	msg := tgapi.NewMessage(chatID, message)
	msg.ReplyMarkup = tgapi.ReplyKeyboardRemove{RemoveKeyboard: true}
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *commonTelegramService) AfterCallbackCleanup(cb *tgapi.CallbackQuery) {
	tcs.answerCallback(cb.ID)

	if cb.Message != nil {
		tcs.deletePreviousMsg(cb.Message.Chat.ID, cb.Message.MessageID)
	}
}

func (tcs *commonTelegramService) SendKeyboardMessage(chatID int64, text string, kb tgapi.InlineKeyboardMarkup) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *commonTelegramService) SendEditedKeyboard(chatID int64, messageID int, kb tgapi.InlineKeyboardMarkup) error {
	msg := tgapi.NewEditMessageReplyMarkup(chatID, messageID, kb)
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *commonTelegramService) SendMessage(chatID int64, text string) error {
	msg := tgapi.NewMessage(chatID, text)
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *commonTelegramService) SendMessageHTMLMode(chatID int64, text string) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ParseMode = tgapi.ModeHTML
	msg.DisableWebPagePreview = true
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *commonTelegramService) SendKeyboardMessageHTMLMode(chatID int64, text string, kb tgapi.InlineKeyboardMarkup) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ParseMode = tgapi.ModeHTML
	msg.DisableWebPagePreview = true
	msg.ReplyMarkup = kb

	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *commonTelegramService) SendRequestPhoneMessage(chatID int64, text string, kb tgapi.ReplyKeyboardMarkup) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ParseMode = tgapi.ModeHTML
	msg.ReplyMarkup = kb

	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *commonTelegramService) answerCallback(callbackID string) {
	if _, err := tcs.botAPI.Request(tgapi.NewCallback(callbackID, "")); err != nil {
		log.Printf("error answer to callback %s: %v", callbackID, err)
	}
}

func (tcs *commonTelegramService) deletePreviousMsg(chatID int64, messageID int) {
	if _, err := tcs.botAPI.Request(tgapi.NewDeleteMessage(chatID, messageID)); err != nil {
		log.Printf("error delete previous message %d: %v", messageID, err)
	}
}
