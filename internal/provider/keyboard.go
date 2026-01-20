package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	ikb "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/keyboard"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/keyboard"
)

type keyboardProvider struct {
	dateTime         isvc.DateTimeService
	callbackProvider iprovider.CallbackProvider
}

func NewKeyboardProvider(dateTime isvc.DateTimeService, callbackProvider iprovider.CallbackProvider) iprovider.KeyboardProvider {
	return &keyboardProvider{dateTime: dateTime, callbackProvider: callbackProvider}
}

func (kp *keyboardProvider) AdminKeyboard() ikb.AdminKeyboardService {
	return keyboard.NewAdminKeyboardService(kp.callbackProvider.AdminCallbackBuilder(), kp.dateTime)
}

func (kp *keyboardProvider) UserKeyboard() ikb.UserKeyboardService {
	return keyboard.NewUserKeyboardService(kp.callbackProvider.UserCallbackBuilder(), kp.dateTime)
}
