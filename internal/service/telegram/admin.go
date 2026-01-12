package tg_svc

import (
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type telegramAdminService struct {
	tgCommon itg.TelegramCommonService
	kb       i.KeyboardService
}

func (tas *telegramAdminService) StartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tas.tgCommon.SendMessage(chatID, "Тут потом всякое будет, ух! А пока...")
	})
}

func (tas *telegramAdminService) ChoiceMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tas.tgCommon.SendKeyboardMessage(chatID, admflow.ChoiceContinueFlow, tas.kb.AdminGreetingKeyboard())
	})
}
