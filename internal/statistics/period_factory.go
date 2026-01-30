package statistics

import (
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
)

type PeriodFactory struct {
	loc *time.Location
}

func NewPeriodFactory() *PeriodFactory {
	loc, _ := time.LoadLocation("Europe/Moscow")
	return &PeriodFactory{loc: loc}
}

func (pf *PeriodFactory) FromLabel(l stat.Label) stat.Period {
	switch l {
	case stat.Yesterday:
		return pf.Yesterday(time.Now())
	case stat.Week:
		return pf.Week(time.Now())
	case stat.Month:
		return pf.Month(time.Now())
	default:
		return pf.Today(time.Now())
	}
}

func (pf *PeriodFactory) Today(t time.Time) stat.Period {
	return stat.Period{From: pf.today(t), To: t, Label: stat.Today}
}

func (pf *PeriodFactory) Yesterday(t time.Time) stat.Period {
	today := pf.today(t)
	to := time.Date(today.Year(), today.Month(), today.Day()-1, 23, 59, 59, 0, pf.loc)
	from := time.Date(today.Year(), today.Month(), today.Day()-1, 0, 0, 0, 0, pf.loc)
	return stat.Period{From: from, To: to, Label: stat.Yesterday}
}

func (pf *PeriodFactory) Week(t time.Time) stat.Period {
	today := pf.today(t)
	weekday := int(today.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	sunday := today.AddDate(0, 0, -(weekday))
	monday := sunday.AddDate(0, 0, -6)
	sunday = time.Date(sunday.Year(), sunday.Month(), sunday.Day(), 23, 59, 59, 0, pf.loc)

	return stat.Period{From: monday, To: sunday, Label: stat.Week}
}

func (pf *PeriodFactory) Month(t time.Time) stat.Period {
	today := pf.today(t)
	y, m, _ := today.Date()

	endMonth := time.Date(y, m-1, 31, 23, 59, 0, 0, pf.loc)
	startMonth := time.Date(y, m-1, 1, 0, 0, 0, 0, pf.loc)
	return stat.Period{From: startMonth, To: endMonth, Label: stat.Month}
}

func (pf *PeriodFactory) today(t time.Time) time.Time {
	local := t.In(pf.loc)
	y, m, d := local.Date()
	from := time.Date(y, m, d, 0, 0, 0, 0, pf.loc)
	return from
}
