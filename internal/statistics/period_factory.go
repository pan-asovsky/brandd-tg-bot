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
	to := pf.today(t)
	from := to.AddDate(0, 0, -1)
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
	return stat.Period{From: monday, To: sunday, Label: stat.Week}
}

func (pf *PeriodFactory) Month(t time.Time) stat.Period {
	today := pf.today(t)
	y, m, _ := today.Date()

	endMonth := time.Date(y, m-1, 31, 0, 0, 0, 0, pf.loc)
	startMonth := endMonth.AddDate(0, -1, 0)
	return stat.Period{From: startMonth, To: endMonth, Label: stat.Month}
}

func (pf *PeriodFactory) today(t time.Time) time.Time {
	local := t.In(pf.loc)
	y, m, d := local.Date()
	from := time.Date(y, m, d, 0, 0, 0, 0, pf.loc)
	return from
}
