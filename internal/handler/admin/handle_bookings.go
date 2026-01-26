package admin

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

func (ach *adminCallbackHandler) handleBookings(query *tgapi.CallbackQuery) error {
	log.Printf("[handle_bookings] callback: %s", query.Data)

	data, ok := strings.CutPrefix(query.Data, admflow.PrefixBooking)
	if !ok {
		return fmt.Errorf("[handle_bookings] error cuting prefix: %s", query.Data)
	}

	bookingID, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return utils.WrapError(err)
	}

	booking, err := ach.service.Booking().Find(bookingID)
	if err != nil {
		return utils.WrapError(err)
	}

	return utils.WrapFunctionError(func() error {
		return ach.telegram.Admin().BookingPreview(query.Message.Chat.ID, booking)
	})
}
