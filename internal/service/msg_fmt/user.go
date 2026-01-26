package msg_fmt

import (
	"database/sql"
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constant"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
	ifmt "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/fmt"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

// userMessageFormatterService todo: подумать надо ли разделить форматирование на Заявки, Клиента и Админа.
type userMessageFormatterService struct {
	dateTime isvc.DateTimeService
}

func NewUserMessageFormatterService(dateTime isvc.DateTimeService) ifmt.UserMessageFormatterService {
	return &userMessageFormatterService{dateTime: dateTime}
}

func (umfs *userMessageFormatterService) Confirm(date, startTime string) string {
	return fmt.Sprintf(usflow.ConfirmMsg, date, startTime)
}

func (umfs *userMessageFormatterService) PreConfirm(booking *entity.Booking) (string, error) {
	date, err := umfs.dateTime.FormatDate(booking.Date, "2006-01-02", "02.01.2006")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(
		usflow.PreConfirmMsg,
		date,
		consts.Time[booking.Time],
		booking.Time,
		consts.ServiceNames[booking.Service],
		booking.RimRadius,
		SQLNullIntToInt64(booking.TotalPrice),
	), nil
}

func (umfs *userMessageFormatterService) My(booking *entity.Booking) (string, error) {
	view, err := umfs.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		usflow.ActiveBooking,
		view,
		consts.ServiceNames[booking.Service],
		SQLNullIntToInt64(booking.TotalPrice),
	), nil
}

func (umfs *userMessageFormatterService) Restriction(booking *entity.Booking) (string, error) {
	view, err := umfs.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(usflow.BookingRestriction, view), nil
}

func (umfs *userMessageFormatterService) PreCancel(date, time string) (string, error) {
	view, err := umfs.dateTime.FormatDateTimeToShortView(date, time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(usflow.BookingPreCancellation, view), nil
}

func (umfs *userMessageFormatterService) BookingPreview(booking *entity.Booking) (string, error) {
	view, err := umfs.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(
		admflow.BookingInfo,
		view,
		booking.RimRadius,
		consts.ServiceNames[booking.Service],
		SQLNullIntToInt64(booking.TotalPrice),
		SQLNullString(booking.UserPhone),
	), nil
}

func SQLNullIntToInt64(nullable sql.NullInt64) int64 {
	if nullable.Valid {
		return nullable.Int64
	}
	return 0
}

func SQLNullString(nullable sql.NullString) string {
	if nullable.Valid {
		return nullable.String
	}
	return ""
}
