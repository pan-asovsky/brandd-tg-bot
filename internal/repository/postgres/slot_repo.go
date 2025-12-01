package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type SlotRepo interface {
	IsTodayAvailable() bool
	GetAvailableSlots(date string) ([]model.Slot, error)
}

type slotRepo struct {
	db *sql.DB
}

func NewSlotRepo(db *sql.DB) SlotRepo {
	return &slotRepo{db: db}
}

func (s *slotRepo) IsTodayAvailable() bool {
	var available bool
	err := s.db.QueryRow(IsTodayAvailable).Scan(&available)
	return err == nil && available
}

func (s *slotRepo) GetAvailableSlots(date string) ([]model.Slot, error) {
	rows, err := s.db.Query(GetZonesByDate, date)
	if err != nil {
		return nil, fmt.Errorf("get available slots request failed: %v", err)
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
