package handler

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleConfirm(q *api.CallbackQuery, _ string) error {
	info := &types.UserSessionInfo{ChatID: q.Message.Chat.ID}
	return utils.WrapFunctionError(func() error {
		return c.svcProvider.Telegram().RequestUserPhone(info)
	})
}
