package utils

import (
	"database/sql"
	"fmt"
	"time"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

func FmtConfirmMsg(booking *model.Booking) (string, error) {
	date, err := formatDate(booking.Date)
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
		toInt64(booking.TotalPrice),
	), nil
}

func formatDate(date string) (string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return date, WrapError(err)
	}

	return t.Format("02.01.2006"), nil
}

func toInt64(nullable sql.NullInt64) int64 {
	if nullable.Valid {
		return nullable.Int64
	}
	return 0
}
