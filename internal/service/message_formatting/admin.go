package message_formatting

import (
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type adminMessageFormattingService struct {
	dateTime i.DateTimeService
}

func (a *adminMessageFormattingService) NewBookingNotify(booking *model.Booking) (string, error) {
	view, err := a.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		consts.NewBookingNotification,
		view,
		booking.RimRadius,
		consts.ServiceNames[booking.Service],
		SQLNullIntToInt64(booking.TotalPrice),
		SQLNullString(booking.UserPhone),
	), nil
}
