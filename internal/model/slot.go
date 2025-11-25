package model

import "time"

type Slot struct {
	ID          int64     `db:"id"`
	Date        time.Time `db:"date"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
	IsAvailable bool      `db:"is_available"`
	CreatedAt   time.Time `db:"created_at"`
}

type Timeslot struct {
	Start time.Time
	End   time.Time
}

type Zone map[string][]Timeslot

type ZoneDefinition struct {
	Name  string
	Start string
	End   string
}

var ZonesDefinition = []ZoneDefinition{
	{Name: "09:00-12:00", Start: "09:00", End: "12:00"},
	{Name: "12:00-15:00", Start: "12:00", End: "15:00"},
	{Name: "15:00-18:00", Start: "15:00", End: "18:00"},
	{Name: "18:00-21:00", Start: "18:00", End: "21:00"},
}
