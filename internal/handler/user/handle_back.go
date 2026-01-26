package user

import (
	"errors"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constant/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (uch *userCallbackHandler) handleBack(query *tgapi.CallbackQuery) error {
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return errors.New("[handle_back] invalid callback" + query.Data)
	}

	switch payload {
	case consts.Menu:
		return utils.WrapFunctionError(func() error {
			return uch.telegram.User().StartMenu(query.Message.Chat.ID)
		})
	default:
		return nil
	}
}
