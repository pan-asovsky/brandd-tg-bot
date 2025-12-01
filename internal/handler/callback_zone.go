package handler

import (
	"fmt"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (c *callbackHandler) handleZone(q *api.CallbackQuery, cd string) error {
	//log.Printf("[handle_zone] callback: %s", cd)
	z, d := utils.ParseZoneCallback(cd)

	zones, err := c.slot.GetAvailableZones(d)
	if err != nil {
		return fmt.Errorf("[handle_zone] error getting zones: %w", err)
	}

	timeslots := zones[z]
	kb := c.kb.TimeKeyboard(timeslots, z, d)
	utils.SendKeyboardMessage(q.Message.Chat.ID, consts.TimeMsg, kb, c.api)

	return nil
}
