package provider

import (
	ikb "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/keyboard"
)

type KeyboardProvider interface {
	AdminKeyboard() ikb.AdminKeyboardService
	UserKeyboard() ikb.UserKeyboardService
}
