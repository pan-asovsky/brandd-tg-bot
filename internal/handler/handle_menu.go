package handler

//import (
//	"errors"
//	"strings"
//
//	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
//	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
//	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
//)
//
//func (c *callbackHandler) handleMenu(query *api.CallbackQuery) error {
//	_, payload, ok := strings.Cut(query.Data, "::")
//	if !ok {
//		return errors.New("[handle_menu]: invalid callback: " + query.Data)
//	}
//
//	switch payload {
//	case consts.Help:
//		return c.handleHelp(query)
//	}
//	return nil
//}
//
//func (c *callbackHandler) handleHelp(q *api.CallbackQuery) error {
//	return utils.WrapFunctionError(func() error { return c.svcProvider.Telegram().SendHelpMessage(q.Message.Chat.ID) })
//}
