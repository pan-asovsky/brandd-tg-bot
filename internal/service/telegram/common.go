package tg_svc

import (
	"log"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type commonService struct {
	botAPI *tgapi.BotAPI
}

func (cs *commonService) RemoveReplyKeyboard(chatID int64, message string) error {
	msg := tgapi.NewMessage(chatID, message)
	msg.ReplyMarkup = tgapi.ReplyKeyboardRemove{RemoveKeyboard: true}
	if _, err := cs.botAPI.Send(msg); err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (cs *commonService) AfterCallbackCleanup(cb *tgapi.CallbackQuery) {
	cs.answerCallback(cb.ID)

	if cb.Message != nil {
		cs.deletePreviousMsg(cb.Message.Chat.ID, cb.Message.MessageID)
	}
}

func (cs *commonService) answerCallback(callbackID string) {
	if _, err := cs.botAPI.Request(tgapi.NewCallback(callbackID, "")); err != nil {
		log.Printf("error answer to callback %s: %v", callbackID, err)
	}
}

func (cs *commonService) deletePreviousMsg(chatID int64, messageID int) {
	if _, err := cs.botAPI.Request(tgapi.NewDeleteMessage(chatID, messageID)); err != nil {
		log.Printf("error delete previous message %d: %v", messageID, err)
	}
}
