package telegram

import (
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type adminTelegramService struct {
	tgCommon itg.TelegramCommonService
	kb       isvc.AdminKeyboardService
}

func NewTelegramAdminService(tgCommon itg.TelegramCommonService, kb isvc.AdminKeyboardService) itg.TelegramAdminService {
	return &adminTelegramService{tgCommon: tgCommon, kb: kb}
}

func (tas *adminTelegramService) StartMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tas.tgCommon.SendMessage(chatID, "Тут потом всякое будет, ух! А пока...")
	})
}

func (tas *adminTelegramService) ChoiceMenu(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return tas.tgCommon.SendKeyboardMessage(chatID, admflow.ChoiceContinueFlow, tas.kb.AdminGreetingKeyboard())
	})
}
