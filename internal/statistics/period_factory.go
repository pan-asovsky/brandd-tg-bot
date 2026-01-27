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
	local := t.In(pf.loc)
	y, m, d := local.Date()
	from := time.Date(y, m, d, 0, 0, 0, 0, pf.loc)
	to := from.AddDate(0, 0, 1)

	return stat.Period{From: from, To: to}
}
