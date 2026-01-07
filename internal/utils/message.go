package utils

import (
	"database/sql"
	"fmt"
	"time"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

func FmtConfirmMsg(booking *model.Booking) (string, error) {
	date, err := FormatDate(booking.Date)
	if err != nil {
		return "", WrapError(err)
	}

	return fmt.Sprintf(
		consts.PreConfirmMsg,
		date,
		consts.Time[booking.Time],
		booking.Time,
		consts.ServiceNames[booking.Service],
		booking.RimRadius,
		SQLNullIntToInt64(booking.TotalPrice),
	), nil
}

func FmtMyBookingMsg(booking *model.Booking) (string, error) {
	date, err := FormatDate(booking.Date)
	if err != nil {
		return "", WrapError(err)
	}

	return fmt.Sprintf(
		consts.ActiveBooking,
		booking.Time,
		date,
		consts.ServiceNames[booking.Service],
		SQLNullIntToInt64(booking.TotalPrice),
	), nil
}

func FmtBookingRestrictionMsg(booking *model.Booking) (string, error) {
	date, err := FormatDate(booking.Date)
	if err != nil {
		return "", WrapError(err)
	}

	return fmt.Sprintf(consts.BookingRestriction, booking.Time, date), nil
}

func FmtBookingPreCancelMsg(date, time string) (string, error) {
	date, err := FormatDate(date)
	if err != nil {
		return "", WrapError(err)
	}

	return fmt.Sprintf(consts.BookingPreCancellation, time, date), nil
}

func FormatDate(date string) (string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date, WrapError(err)
	}

	formatted := t.Format("02.01.2006")
	//log.Printf("[format_date] in: %s, out: %s", date, formatted)
	return formatted, nil
}

func SQLNullIntToInt64(nullable sql.NullInt64) int64 {
	if nullable.Valid {
		return nullable.Int64
	}
	return 0
}
