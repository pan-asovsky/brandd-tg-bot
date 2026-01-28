package stat

import "time"

type Period struct {
	From, To time.Time
}

func (p Period) Format() (from, to string) {
	l := "02.01"
	return p.From.Format(l), p.To.Format(l)
}

type Stats struct {
	PendingCount   int
	ActiveCount    int
	CompletedCount int
	CanceledCount  int
	NoShowCount    int
}
