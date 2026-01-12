package tg_svc

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type telegramCommonService struct {
	botAPI *tgapi.BotAPI
}

func (tcs *telegramCommonService) RemoveReplyKeyboard(chatID int64, message string) error {
	msg := tgapi.NewMessage(chatID, message)
	msg.ReplyMarkup = tgapi.ReplyKeyboardRemove{RemoveKeyboard: true}
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *telegramCommonService) AfterCallbackCleanup(cb *tgapi.CallbackQuery) {
	tcs.answerCallback(cb.ID)

	if cb.Message != nil {
		tcs.deletePreviousMsg(cb.Message.Chat.ID, cb.Message.MessageID)
	}
}

func (tcs *telegramCommonService) SendKeyboardMessage(chatID int64, text string, kb tgapi.InlineKeyboardMarkup) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *telegramCommonService) SendEditedKeyboard(chatID int64, messageID int, kb tgapi.InlineKeyboardMarkup) error {
	msg := tgapi.NewEditMessageReplyMarkup(chatID, messageID, kb)
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *telegramCommonService) SendMessage(chatID int64, text string) error {
	msg := tgapi.NewMessage(chatID, text)
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *telegramCommonService) SendMessageHTMLMode(chatID int64, text string) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ParseMode = tgapi.ModeHTML
	msg.DisableWebPagePreview = true
	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *telegramCommonService) SendKeyboardMessageHTMLMode(chatID int64, text string, kb tgapi.InlineKeyboardMarkup) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ParseMode = tgapi.ModeHTML
	msg.DisableWebPagePreview = true
	msg.ReplyMarkup = kb

	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *telegramCommonService) SendRequestPhoneMessage(chatID int64, text string, kb tgapi.ReplyKeyboardMarkup) error {
	msg := tgapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	if _, err := tcs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (tcs *telegramCommonService) answerCallback(callbackID string) {
	if _, err := tcs.botAPI.Request(tgapi.NewCallback(callbackID, "")); err != nil {
		log.Printf("error answer to callback %s: %v", callbackID, err)
	}
}

func (tcs *telegramCommonService) deletePreviousMsg(chatID int64, messageID int) {
	if _, err := tcs.botAPI.Request(tgapi.NewDeleteMessage(chatID, messageID)); err != nil {
		log.Printf("error delete previous message %d: %v", messageID, err)
	}
}
