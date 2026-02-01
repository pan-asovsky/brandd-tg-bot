package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type slotRepo struct {
	pool *pgxpool.Pool
}

const (
	timeLay     = "15:04"
	dateTimeLay = "2006-01-02 15:04:05"
)

func NewSlotRepo(p *pgxpool.Pool) irepo.SlotRepo {
	return &slotRepo{pool: p}
}

func (sr *slotRepo) IsTodayAvailable() bool {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	var available bool
	err := sr.pool.QueryRow(ctx, IsTodayAvailable).Scan(&available)

	return err == nil && available
}

func (sr *slotRepo) GetAvailableSlots(date string) ([]entity.Slot, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()
	rows, err := sr.pool.Query(ctx, GetZonesByDate, date)

	if err != nil {
		return nil, fmt.Errorf("[get_available_slots] query error: %w", err)
	}
	defer rows.Close()

	var slots []entity.Slot
	for rows.Next() {
		var slot entity.Slot
		var (
			//sqlDate time.Time
			start   time.Time
			end     time.Time
			created time.Time
		)
		if err = rows.Scan(
			&slot.ID,
			&slot.Date,
			&start,
			&end,
			&slot.IsAvailable,
			&created,
		); err != nil {
			return nil, fmt.Errorf("[get_available_slots] row scan error: %w", err)
		}

		//slot.Date = sqlDate.Format(dateLay)
		slot.StartTime = start.Format(timeLay)
		slot.EndTime = end.Format(timeLay)
		slot.CreatedAt = created.Format(dateTimeLay)

		slots = append(slots, slot)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[get_available_slots] rows error: %w", err)
	}
	return slots, nil
}

func (sr *slotRepo) FindByDateAndTime(date time.Time, start string) (*entity.Slot, error) {
	var (
		startTime time.Time
		endTime   time.Time
		created   time.Time
		slot      entity.Slot
	)

	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	if err := sr.pool.QueryRow(ctx, GetZonesByDate, date).Scan(
		&slot.ID,
		&slot.Date,
		&startTime,
		&endTime,
		&slot.IsAvailable,
		&created,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[find_slot_by_date_time] not founded for %sr %sr %v", date, start, err)
		}
		return nil, fmt.Errorf("[find_slot_by_date_time] failed: %w", err)
	}

	slot.StartTime = startTime.Format(timeLay)
	slot.EndTime = endTime.Format(timeLay)
	slot.CreatedAt = created.Format(dateTimeLay)

	return &slot, nil
}

func (sr *slotRepo) MarkUnavailable(date time.Time, startTime string) error {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	if _, err := sr.pool.Exec(ctx, MarkSlotUnavailable, date, startTime); err != nil {
		return utils.WrapError(err)
	}
	return nil
}

func (sr *slotRepo) FreeUp(date time.Time, startTime string) error {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	result, err := sr.pool.Exec(ctx, FreeUpSlot, date, startTime)
	if err != nil {
		return utils.WrapError(err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("[free_up_slot] no free up slot: %s - %s", date, startTime)
	}

	return nil
}
