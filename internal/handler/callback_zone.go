package handler

import (
	"fmt"
	"log"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleZone(q *api.CallbackQuery, cd string) error {
	log.Printf("[handle_zone] callback: %s", cd)

	z, d, err := utils.ParseCallback(cd)
	if err != nil {
		return fmt.Errorf("[handle_zone] error parsing %s: %w", cd, err)
	}

	zones, err := c.cache.GetZones(d)
	if err != nil {
		return fmt.Errorf("[handle_zone] cache error: %w", err)
	}

	timeslots := zones[z]
	log.Printf("[handle_zone] timeslots: %s", timeslots)

	kb := c.kb.TimeKeyboard(timeslots)
	utils.SendKeyboardMessage(q.Message.Chat.ID, consts.TimeChoosingMsg, kb, c.api)

	return nil
}
