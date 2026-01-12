package utils

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PrintKeyboard(keyboard [][]api.InlineKeyboardButton) {
	for x, row := range keyboard {
		fmt.Printf("Строка %d:\n", x+1)
		for j, btn := range row {
			callbackData := "nil"
			if btn.CallbackData != nil {
				callbackData = *btn.CallbackData
			}
			fmt.Printf("Button %d: \"%s\" → Callback: \"%s\"\n",
				j+1, btn.Text, callbackData)
		}
	}
}
