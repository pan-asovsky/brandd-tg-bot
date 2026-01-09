package interfaces

import "time"

type DateTimeService interface {
	FormatDateTimeToShortView(d, t, inLayout string) (string, error)
	FormatDate(date, inLayout, outLayout string) (string, error)
	ParseDate(date, inLayout string) (time.Time, error)
}
