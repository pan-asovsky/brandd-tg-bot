package stat

import (
	"fmt"
	"time"
)

type Period struct {
	From, To time.Time
	Label    Label
}

type Label string

const (
	Today     Label = "T"
	Yesterday Label = "Y"
	Week      Label = "W"
	Month     Label = "M"
)

func (p Period) Format() (view string) {
	l := "02.01"

	switch p.Label {
	case Today, Yesterday:
		return p.From.Format(l)
	default:
		return fmt.Sprintf("%s-%s", p.From.Format(l), p.To.Format(l))
	}
}

type Stats struct {
	PendingCount   int
	ActiveCount    int
	CompletedCount int
	CanceledCount  int
	NoShowCount    int
}
