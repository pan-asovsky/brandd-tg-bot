package message_formatting

import (
	"database/sql"
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type bookingMessageFormattingService struct {
	dateTime i.DateTimeService
}

func (b *bookingMessageFormattingService) Confirm(date, startTime string) string {
	return fmt.Sprintf(usflow.ConfirmMsg, date, startTime)
}

func (b *bookingMessageFormattingService) PreConfirm(booking *entity.Booking) (string, error) {
	date, err := b.dateTime.FormatDate(booking.Date, "2006-01-02", "02.01.2006")
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

func (b *bookingMessageFormattingService) My(booking *entity.Booking) (string, error) {
	view, err := b.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
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

func (b *bookingMessageFormattingService) Restriction(booking *entity.Booking) (string, error) {
	view, err := b.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(usflow.BookingRestriction, view), nil
}

func (b *bookingMessageFormattingService) PreCancel(date, time string) (string, error) {
	view, err := b.dateTime.FormatDateTimeToShortView(date, time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(usflow.BookingPreCancellation, view), nil
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
