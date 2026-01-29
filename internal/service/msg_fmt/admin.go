package msg_fmt

import (
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constant"
	"github.com/pan-asovsky/brandd-tg-bot/internal/constant/notification"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	ifmt "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/fmt"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
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
		notification.NewBooking,
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
		notification.CancelBooking,
		view,
		SQLNullString(booking.UserPhone),
	), nil
}

func (amfs *adminMessageFormatterService) BookingCompleted(_ *entity.Booking) (string, error) {
	return notification.CompleteBooking, nil
}

func (amfs *adminMessageFormatterService) Statistics(s stat.Stats, p stat.Period) string {
	return fmt.Sprintf(
		consts.StatShortMessage,
		consts.PeriodLabels[p.Label], p.Format(),
		s.CompletedCount,
		s.NoShowCount,
		s.CanceledCount,
	)
}
