package handler

import (
	"fmt"
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleDate(q *api.CallbackQuery, cd string) error {
	log.Printf("[handle_date] callback: %s", cd)

	zones, err := c.slot.GetAvailableZones(cd)
	if err != nil {
		return fmt.Errorf("error getting zones: %s", err)
	}

	if err = c.cache.CacheZones(cd, zones); err != nil {
		return fmt.Errorf("error caching zones: %s", err)
	}

	kb := c.kb.ZoneKeyboard(zones, cd)
	utils.SendKeyboardMessage(q.Message.Chat.ID, consts.ZoneChoosingMsg, kb, c.api)

	return nil
}
