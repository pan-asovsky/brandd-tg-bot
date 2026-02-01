package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type bookingRepo struct {
	pool *pgxpool.Pool
}

const TwoSec = 2 * time.Second

func NewBookingRepo(p *pgxpool.Pool) irepo.BookingRepo {
	return &bookingRepo{pool: p}
}

func (br *bookingRepo) FindActiveNotPending(chatID int64) (*entity.Booking, error) {
	return br.findOne(FindActiveNotPending, "find_active_not_pending", chatID)
}

func (br *bookingRepo) FindActivePending(chatID int64) (*entity.Booking, error) {
	return br.findOne(FindActivePending, "find_active_pending", chatID)
}

func (br *bookingRepo) FindAnyActive(chatID int64) (*entity.Booking, error) {
	return br.findOne(FindAnyActive, "find_any_active", chatID)
}

func (br *bookingRepo) Exists(chatID int64) bool {
	ok, err := br.exists(BookingExists, chatID)
	return err == nil && ok
}

func (br *bookingRepo) UpdateRimRadius(chatID int64, rimRadius string) error {
	return br.exec(UpdateRimRadius, "update_rim_radius", rimRadius, chatID)
}

func (br *bookingRepo) UpdateStatus(chatID int64, status entity.BookingStatus) error {
	return br.exec(UpdateStatus, "update_status", status, chatID)
}

func (br *bookingRepo) UpdatePrice(chatID int64, price int64) error {
	return br.exec(UpdatePrice, "update_price", price, chatID)
}

func (br *bookingRepo) UpdateService(chatID int64, service string) error {
	return br.exec(UpdateService, "update_service", service, chatID)
}

func (br *bookingRepo) Save(booking *entity.Booking) (*entity.Booking, error) {
	now := time.Now().UTC().Add(3 * time.Hour)
	ctx, cancel := context.WithTimeout(context.Background(), TwoSec)
	defer cancel()
	if err := br.pool.QueryRow(ctx, SaveBooking,
		booking.ChatID,
		booking.Date,
		booking.Time,
		booking.Service,
		booking.RimRadius,
		booking.TotalPrice,
		true,
		booking.Status,
		now,
		now,
	).Scan(&booking.ID); err != nil {
		return nil, utils.WrapError(err)
	}

	return booking, nil
}

func (br *bookingRepo) SetPhone(phone string, chatID int64) error {
	return br.exec(SetPhone, "set_phone", phone, chatID)
}

func (br *bookingRepo) Confirm(chatID int64) error {
	return br.confirm(chatID, "user")
}

func (br *bookingRepo) AutoConfirm(chatID int64) error {
	return br.confirm(chatID, "system")
}

func (br *bookingRepo) Cancel(chatID int64) error {
	return br.exec(CancelBooking, "cancel", entity.Cancelled, chatID)
}

func (br *bookingRepo) FindAllActive() ([]entity.Booking, error) {
	return br.findMany(FindAllActive, "find_all_active")
}

func (br *bookingRepo) Find(bookingID int64) (*entity.Booking, error) {
	booking, err := br.findOne(FindByID, "find_by_id", bookingID)
	if booking == nil && err != nil {
		return nil, fmt.Errorf("[find_booking] booking %d not founded", bookingID)
	}
	return booking, nil
}

func (br *bookingRepo) Close(info *model.BookingInfo) (*entity.Booking, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	booking, err := scan(br.pool.QueryRow(ctx, Close, entity.BookingStatus(info.Status), info.UserChatID, info.BookingID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[close_booking] booking %d not founded", info.BookingID)
		}
		return nil, utils.WrapError(err)
	}

	return booking, nil
}

func (br *bookingRepo) StatusesByPeriod(period stat.Period) (stat.Stats, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	var stats stat.Stats
	row := br.pool.QueryRow(ctx, StatusesByPeriod, period.From, period.To)

	if err := row.Scan(
		&stats.ActiveCount,
		&stats.CanceledCount,
		&stats.CompletedCount,
		&stats.NoShowCount,
		&stats.PendingCount,
	); err != nil {
		return stats, utils.WrapError(err)
	}

	return stats, nil
}

func (br *bookingRepo) findOne(query, tag string, args ...any) (*entity.Booking, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	booking, err := scan(br.pool.QueryRow(ctx, query, args...))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[%s] query error: %w", tag, err)
	}

	return booking, nil
}

func (br *bookingRepo) findMany(query, tag string, args ...any) ([]entity.Booking, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	rows, err := br.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("[%s] query error: %w", tag, err)
	}
	defer rows.Close()

	var bookings []entity.Booking
	for rows.Next() {
		booking, err := scan(rows)
		if err != nil {
			return nil, fmt.Errorf("[%s] rows scan error: %w", tag, err)
		}

		bookings = append(bookings, *booking)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[%s] rows error: %w", tag, err)
	}

	return bookings, nil
}

func scan(scanner interface {
	Scan(dest ...any) error
}) (*entity.Booking, error) {
	var booking entity.Booking
	var t time.Time
	if err := scanner.Scan(
		&booking.ID,
		&booking.ChatID,
		&booking.UserPhone,
		&booking.Date,
		&t,
		&booking.Service,
		&booking.RimRadius,
		&booking.TotalPrice,
		&booking.Status,
		&booking.Active,
		&booking.CreatedAt,
		&booking.UpdatedAt,
		&booking.ConfirmedBy,
		&booking.CancelledBy,
		&booking.Notes,
	); err != nil {
		return nil, err
	}

	booking.Time = t.Format(timeLay)

	return &booking, nil
}

func (br *bookingRepo) confirm(chatID int64, confirmedBy string) error {
	return br.exec(ConfirmBooking, "confirm_booking", entity.Confirmed, confirmedBy, chatID)
}

func (br *bookingRepo) exec(query, tag string, args ...any) error {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	if _, err := br.pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("[%s] exec error: %w", tag, err)
	}

	return nil
}

func (br *bookingRepo) exists(query string, args ...any) (bool, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	var exists bool
	err := br.pool.QueryRow(ctx, query, args...).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func CtxWithTimeout(t time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), t)
}
