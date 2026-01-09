package message_formatting

import (
	"database/sql"
	"fmt"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type bookingMessageFormattingService struct {
	dateTime i.DateTimeService
}

func (b *bookingMessageFormattingService) Confirm(date, startTime string) string {
	return fmt.Sprintf(consts.ConfirmMsg, date, startTime)
}

func (b *bookingMessageFormattingService) PreConfirm(booking *model.Booking) (string, error) {
	date, err := b.dateTime.FormatDate(booking.Date, "2006-01-02", "01.02.2006")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(
		consts.PreConfirmMsg,
		date,
		consts.Time[booking.Time],
		booking.Time,
		consts.ServiceNames[booking.Service],
		booking.RimRadius,
		sqlNullIntToInt64(booking.TotalPrice),
	), nil
}

func (b *bookingMessageFormattingService) My(booking *model.Booking) (string, error) {
	view, err := b.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		consts.ActiveBooking,
		view,
		consts.ServiceNames[booking.Service],
		sqlNullIntToInt64(booking.TotalPrice),
	), nil
}

func (b *bookingMessageFormattingService) Restriction(booking *model.Booking) (string, error) {
	view, err := b.dateTime.FormatDateTimeToShortView(booking.Date, booking.Time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(consts.BookingRestriction, view), nil
}

func (b *bookingMessageFormattingService) PreCancel(date, time string) (string, error) {
	view, err := b.dateTime.FormatDateTimeToShortView(date, time, "2006-01-02")
	if err != nil {
		return "", utils.WrapError(err)
	}

	return fmt.Sprintf(consts.BookingPreCancellation, view), nil
}

func sqlNullIntToInt64(nullable sql.NullInt64) int64 {
	if nullable.Valid {
		return nullable.Int64
	}
	return 0
}
