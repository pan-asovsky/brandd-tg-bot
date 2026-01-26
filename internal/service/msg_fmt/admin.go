package msg_fmt

import (
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constant"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
	ifmt "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/fmt"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type adminMessageFormatterService struct {
	dateTime isvc.DateTimeService
}

func NewAdminMessageFormatterService(dateTime isvc.DateTimeService) ifmt.AdminMessageFormatterService {
	return &adminMessageFormatterService{dateTime: dateTime}
}

func (amfs *adminMessageFormatterService) BookingCreated(booking *entity.Booking) (string, error) {
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

func (amfs *adminMessageFormatterService) BookingCancelled(booking *entity.Booking) (string, error) {
	view, err := amfs.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(
		admflow.CancelBookingNotification,
		view,
		SQLNullString(booking.UserPhone),
	), nil
}

func (amfs *adminMessageFormatterService) BookingCompleted(_ *entity.Booking) (string, error) {
	return admflow.CompleteBookingNotification, nil
}
