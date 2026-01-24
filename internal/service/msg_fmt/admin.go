package msg_fmt

import (
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	ifmt "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/fmt"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type adminMessageFormattingService struct {
	dateTime isvc.DateTimeService
}

func NewAdminMessageFormattingService(dateTime isvc.DateTimeService) ifmt.AdminMessageFormatterService {
	return &adminMessageFormattingService{dateTime: dateTime}
}

func (amfs *adminMessageFormattingService) NewBookingNotify(booking *entity.Booking) (string, error) {
	view, err := amfs.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
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
