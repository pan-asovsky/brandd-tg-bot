package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type slotRepo struct {
	db *sql.DB
}

func NewSlotRepo(db *sql.DB) SlotRepository {
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
		if err := rows.Scan(
			&slot.ID,
			&slot.Date,
			&slot.StartTime,
			&slot.EndTime,
			&slot.IsAvailable,
			&slot.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("row scan error: %w", err)
		}
		slots = append(slots, slot)
	}

	return slots, nil
}

func (s *slotRepo) FindByZone(zone model.Zone) []model.Slot {
	return nil
}

func (s *slotRepo) FindByTimeslot(timeslot model.Timeslot) []model.Slot {
	return nil
}
