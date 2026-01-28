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

func (pf *PeriodFactory) Day(t time.Time) stat.Period {
	to := pf.today(t)
	from := to.AddDate(0, 0, -1)
	return stat.Period{From: from, To: to}
}

func (pf *PeriodFactory) Week(t time.Time) stat.Period {
	today := pf.today(t)
	weekday := int(today.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	thisMonday := today.AddDate(0, 0, -(weekday - 1))
	lastMonday := thisMonday.AddDate(0, 0, -7)
	return stat.Period{From: lastMonday, To: thisMonday}
}

func (pf *PeriodFactory) Month(t time.Time) stat.Period {
	today := pf.today(t)
	y, m, _ := today.Date()

	monthStart := time.Date(y, m, 1, 0, 0, 0, 0, pf.loc)
	lastMonth := monthStart.AddDate(0, -1, 0)
	return stat.Period{From: lastMonth, To: monthStart}
}

func (pf *PeriodFactory) today(t time.Time) time.Time {
	local := t.In(pf.loc)
	y, m, d := local.Date()
	from := time.Date(y, m, d, 0, 0, 0, 0, pf.loc)
	return from
}
