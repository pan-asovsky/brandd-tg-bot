package service

import (
	"fmt"
	"strings"
	"time"

	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type dateTimeService struct{}

func NewDateTimeService() isvc.DateTimeService {
	return &dateTimeService{}
}

func (dts *dateTimeService) FormatDateTimeToShortView(d, t, inLayout string) (string, error) {
	date, err := time.Parse(inLayout, d)
	if err != nil {
		return "", utils.WrapError(err)
	}

	day := fmt.Sprintf("%02d", date.Day())
	monthNumber := fmt.Sprintf("%02d", int(date.Month()))
	weekDayRu := daysOfWeekShortRu[date.Weekday().String()]

	return fmt.Sprintf("%s.%s (%s), %s", day, monthNumber, weekDayRu, t), nil
}

func (dts *dateTimeService) FormatDate(date, inLayout, outLayout string) (string, error) {
	t, err := time.Parse(inLayout, date)
	if err != nil {
		return date, utils.WrapError(err)
	}

	formatted := t.Format(outLayout)
	return formatted, nil
}

func (dts *dateTimeService) ParseDate(date, inLayout string) (time.Time, error) {
	parsed, err := time.Parse(inLayout, date)
	if err != nil {
		return time.Now(), utils.WrapError(err)
	}
	return parsed, nil
}

func (dts *dateTimeService) ParseToStartEndTime(time string) (start, end string) {
	split := strings.Split(time, "-")
	start = split[0]
	end = split[1]
	return start, end
}

var daysOfWeekShortRu = map[string]string{
	"Monday":    "пн",
	"Tuesday":   "вт",
	"Wednesday": "ср",
	"Thursday":  "чт",
	"Friday":    "пт",
	"Saturday":  "сб",
	"Sunday":    "вс",
}
