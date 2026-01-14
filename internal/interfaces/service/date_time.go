package service

import "time"

type DateTimeService interface {
	FormatDateTimeToShortView(d, t, inDateFormat string) (string, error)
	FormatDate(date, inLayout, outLayout string) (string, error)
	ParseDate(date, inLayout string) (time.Time, error)
	ParseToStartEndTime(time string) (start, end string)
}
