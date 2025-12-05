package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type SlotRepo interface {
	IsTodayAvailable() bool
	GetAvailableSlots(date string) ([]model.Slot, error)
	FindByDateAndTime(date, start, end string) (*model.Slot, error)
	MarkUnavailable(date, start, end string) error
}

type slotRepo struct {
	db *sql.DB
}

func (s *slotRepo) IsTodayAvailable() bool {
	var available bool
	err := s.db.QueryRow(IsTodayAvailable).Scan(&available)
	return err == nil && available
}

func (s *slotRepo) GetAvailableSlots(date string) ([]model.Slot, error) {
	rows, err := s.db.Query(GetZonesByDate, date)
	if err != nil {
		return nil, fmt.Errorf("[get_available_slots] failed: %v", err)
	}
	defer rows.Close()

	var slots []model.Slot
	for rows.Next() {
		var slot model.Slot
		var (
			sqlDate time.Time
			start   time.Time
			end     time.Time
			created time.Time
		)
		if err := rows.Scan(
			&slot.ID,
			&sqlDate,
			&start,
			&end,
			&slot.IsAvailable,
			&created,
		); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}

		slot.Date = sqlDate.Format("2006-01-02")
		slot.StartTime = start.Format("15:04")
		slot.EndTime = end.Format("15:04")
		slot.CreatedAt = created.Format("2006-01-02 15:04")

		//log.Printf("[get_available_slots] mapped slot: %+v", slot)
		slots = append(slots, slot)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	//log.Printf("[get_available_slots] founded %d slots for date: %s", len(slots), date)
	return slots, nil
}

func (s *slotRepo) FindByDateAndTime(date, start, end string) (*model.Slot, error) {
	var (
		sqlDate   time.Time
		startTime time.Time
		endTime   time.Time
		created   time.Time
		slot      model.Slot
	)
	if err := s.db.QueryRow(GetSlotByDateAndTime, date, start, end).Scan(
		&slot.ID,
		&sqlDate,
		&startTime,
		&endTime,
		&slot.IsAvailable,
		&created,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[find_by_date_and_time] slot not founded: %v", err)
		}
		return nil, fmt.Errorf("[find_slot_by_date_time] failed: %v", err)
	}

	slot.Date = sqlDate.Format("02-01-2006")
	slot.StartTime = startTime.Format("15:04")
	slot.EndTime = endTime.Format("15:04")
	slot.CreatedAt = created.Format("15:04 02-01-2006")

	return &slot, nil
}

func (s *slotRepo) MarkUnavailable(date, start, end string) error {
	if err := s.db.QueryRow(MarkSlotUnavailable, date, start, end).Scan(); err != nil {
		return fmt.Errorf("[set_unavailable] update error: %v", err)
	}
	return nil
}
