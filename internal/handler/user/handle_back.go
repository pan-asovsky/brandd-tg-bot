package user

import (
	"errors"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleBack(query *tg.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_back] invalid callback" + query.Data)
	}

	switch payload {
	case consts.Menu:
		return utils.WrapFunctionError(func() error {
			return uch.tgProvider.User().StartMenu(query.Message.Chat.ID)
		})
	default:
		return nil
	}
}
