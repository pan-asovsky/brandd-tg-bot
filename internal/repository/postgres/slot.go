package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type slotRepo struct {
	db *sql.DB
}

func NewSlotRepo(db *sql.DB) irepo.SlotRepo {
	return &slotRepo{db: db}
}

func (s *slotRepo) IsTodayAvailable() bool {
	var available bool
	err := s.db.QueryRow(IsTodayAvailable).Scan(&available)
	return err == nil && available
}

func (s *slotRepo) GetAvailableSlots(date string) ([]entity.Slot, error) {
	rows, err := s.db.Query(GetZonesByDate, date)
	if err != nil {
		return nil, fmt.Errorf("[get_available_slots] query error: %w", err)
	}
	defer rows.Close()

	var slots []entity.Slot
	for rows.Next() {
		var slot entity.Slot
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
			return nil, fmt.Errorf("[get_available_slots] row scan error: %w", err)
		}

		slot.Date = sqlDate.Format("2006-01-02")
		slot.StartTime = start.Format("15:04")
		slot.EndTime = end.Format("15:04")
		slot.CreatedAt = created.Format("2006-01-02 15:04")

		slots = append(slots, slot)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[get_available_slots] rows error: %w", err)
	}

	return slots, nil
}

func (s *slotRepo) FindByDateAndTime(date, start string) (*entity.Slot, error) {
	var (
		sqlDate   time.Time
		startTime time.Time
		endTime   time.Time
		created   time.Time
		slot      entity.Slot
	)

	if err := s.db.QueryRow(GetSlotByDateAndTime, date, start).Scan(
		&slot.ID,
		&sqlDate,
		&startTime,
		&endTime,
		&slot.IsAvailable,
		&created,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[find_slot_by_date_time] not founded for %s %s %v", date, start, err)
		}
		return nil, fmt.Errorf("[find_slot_by_date_time] failed: %w", err)
	}

	slot.Date = sqlDate.Format("2006-01-02")
	slot.StartTime = startTime.Format("15:04")
	slot.EndTime = endTime.Format("15:04")
	slot.CreatedAt = created.Format("15:04 02-01-2006")

	return &slot, nil
}

func (s *slotRepo) MarkUnavailable(date, startTime string) error {
	if _, err := s.db.Exec(MarkSlotUnavailable, date, startTime); err != nil {
		return utils.WrapError(err)
	}
	return nil
}

func (s *slotRepo) FreeUp(date, startTime string) error {
	result, err := s.db.Exec(FreeUpSlot, date, startTime)
	affected, err := result.RowsAffected()
	if err != nil {
		return utils.WrapError(err)
	}
	if affected == 0 {
		return errors.New("[free_up_slot] no free up slot: " + date + " " + startTime)
	}

	return nil
}
