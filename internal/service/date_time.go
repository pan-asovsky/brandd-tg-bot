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

func (dts *dateTimeService) FormatDateTimeToShortView(date time.Time, time string) (string, error) {
	day := fmt.Sprintf("%02d", date.Day())
	monthNumber := fmt.Sprintf("%02d", int(date.Month()))
	weekDayRu := daysOfWeekShortRu[date.Weekday().String()]

	return fmt.Sprintf("%s.%s (%s), %s", day, monthNumber, weekDayRu, time), nil
}

func (dts *dateTimeService) formatDateTimeToExtendedView(date time.Time, time string) (string, error) {
	day := fmt.Sprintf("%02d", date.Day())
	monthRuName := monthNamesRu[int(date.Month())]
	weekDayRu := daysOfWeekFullRu[date.Weekday().String()]

	return fmt.Sprintf("%s.%s (%s), %s", day, monthRuName, weekDayRu, time), nil
}

func (dts *dateTimeService) FormatDate(date time.Time, outLayout string) string {
	return date.Format(outLayout)
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

var daysOfWeekFullRu = map[string]string{
	"Monday":    "понедельник",
	"Tuesday":   "вторник",
	"Wednesday": "среда",
	"Thursday":  "четверг",
	"Friday":    "пятница",
	"Saturday":  "суббота",
	"Sunday":    "воскресенье",
}

var monthNamesRu = map[int]string{
	1:  "Январь",
	2:  "Февраль",
	3:  "Март",
	4:  "Апрель",
	5:  "Май",
	6:  "Июнь",
	7:  "Июль",
	8:  "Август",
	9:  "Сентябрь",
	10: "Октябрь",
	11: "Ноябрь",
	12: "Декабрь",
}
