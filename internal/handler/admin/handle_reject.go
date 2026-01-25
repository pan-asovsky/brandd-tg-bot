package admin

import (
	"errors"
	"log"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleReject(query *tgapi.CallbackQuery) error {
	log.Printf("[handle_reject] callback: %s", query.Data)

	action, ok := strings.CutPrefix(query.Data, admflow.PrefixReject)
	if !ok {
		return errors.New("[handle_reject] invalid action " + query.Data)
	}

	var backDirection string
	switch model.Status(action) {
	case model.PreNoShow, model.PreCompleted:
		backDirection = admflow.MenuPrefix + admflow.Bookings
	default:
		backDirection = admflow.Menu
	}

	return utils.WrapFunctionError(func() error {
		return ach.tgProvider.Admin().RejectAction(query.Message.Chat.ID, backDirection)
	})
}
