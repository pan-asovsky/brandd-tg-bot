package msg_fmt

import (
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	msg_fmt2 "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/msg_fmt"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type adminMessageFormattingService struct {
	dateTime isvc.DateTimeService
}

func NewAdminMessageFormattingService(dateTime isvc.DateTimeService) msg_fmt2.AdminMessageFormatterService {
	return &adminMessageFormattingService{dateTime: dateTime}
}

func (a *adminMessageFormattingService) NewBookingNotify(booking *entity.Booking) (string, error) {
	view, err := a.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(
		admflow.NewBookingNotification,
		view,
		booking.RimRadius,
		consts.ServiceNames[booking.Service],
		SQLNullIntToInt64(booking.TotalPrice),
		SQLNullString(booking.UserPhone),
	), nil
}
