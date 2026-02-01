package service

import "time"

type DateTimeService interface {
	FormatDateTimeToShortView(date time.Time, time string) (string, error)
	FormatDate(date time.Time, outLayout string) string
	ParseDate(date, inLayout string) (time.Time, error)
	ParseToStartEndTime(time string) (start, end string)
}
