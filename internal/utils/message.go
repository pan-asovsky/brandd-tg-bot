package utils

import (
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

func FmtRimMsg(booking *model.Booking) string {
	return fmt.Sprintf(
		consts.PreConfirmMsg,
		booking.Date,
		consts.Time[booking.Time],
		booking.Time,
		consts.ServiceNames[booking.Service],
		booking.RimRadius,
	)
}
