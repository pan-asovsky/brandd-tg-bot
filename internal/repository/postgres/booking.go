package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type bookingRepo struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) irepo.BookingRepo {
	return &bookingRepo{db: db}
}

func (br *bookingRepo) FindActiveNotPending(chatID int64) (*entity.Booking, error) {
	booking, err := scanBooking(br.db.QueryRow(FindActiveNotPending, chatID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[find_booking_by_chat_id] failed for %d: %v", chatID, err)
	}
	return booking, nil
}

func (br *bookingRepo) FindPending(chatID int64) (*entity.Booking, error) {
	booking, err := scanBooking(br.db.QueryRow(FindActivePending, chatID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[find_booking_by_chat_id] failed for %d: %v", chatID, err)
	}
	return booking, nil
}

func (br *bookingRepo) Exists(chatID int64) bool {
	var exists bool
	if err := br.db.QueryRow(BookingExists, chatID).Scan(&exists); err != nil {
		return false
	}
	return exists
}

func (br *bookingRepo) UpdateRimRadius(chatID int64, rimRadius string) error {
	_, err := br.db.Exec(UpdateRimRadius, rimRadius, chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (br *bookingRepo) UpdateStatus(chatID int64, status entity.BookingStatus) error {
	_, err := br.db.Exec(UpdateStatus, status, chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (br *bookingRepo) UpdatePrice(chatID int64, price int64) error {
	_, err := br.db.Exec(UpdatePrice, price, chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (br *bookingRepo) UpdateService(chatID int64, service string) error {
	_, err := br.db.Exec(UpdateService, service, chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (br *bookingRepo) Save(booking *entity.Booking) (*entity.Booking, error) {
	now := time.Now().UTC().Add(3 * time.Hour)
	err := br.db.QueryRow(SaveBooking,
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
	).Scan(&booking.ID)
	if err != nil {
		return nil, utils.WrapError(err)
	}
	return booking, nil
}

func (br *bookingRepo) SetPhone(phone string, chatID int64) error {
	if _, err := br.db.Exec(SetPhone, phone, chatID); err != nil {
		return fmt.Errorf("[set_phone_by_chat_id] query error: %w", err)
	}
	return nil
}

func (br *bookingRepo) Confirm(chatID int64) error {
	if _, err := br.db.Exec(ConfirmBooking, entity.Confirmed, "an_user", chatID); err != nil {
		return fmt.Errorf("[confirm_booking] error: %w", err)
	}
	return nil
}

func (br *bookingRepo) AutoConfirm(chatID int64) error {
	if _, err := br.db.Exec(ConfirmBooking, entity.Confirmed, "system", chatID); err != nil {
		return fmt.Errorf("[confirm_booking] error: %w", err)
	}
	return nil
}

func (br *bookingRepo) Cancel(chatID int64) error {
	if _, err := br.db.Exec(CancelBooking, entity.Cancelled, chatID); err != nil {
		return fmt.Errorf("[cancel_booking] error: %w", err)
	}
	return nil
}

func (br *bookingRepo) FindAllActive() ([]entity.Booking, error) {
	rows, err := br.db.Query(FindAllActive)
	if err != nil {
		return nil, fmt.Errorf("[find_all_active_bookings] query error: %w", err)
	}
	defer rows.Close()

	var bookings []entity.Booking
	for rows.Next() {
		booking, err := scanBooking(rows)
		if err != nil {
			return nil, fmt.Errorf("[find_all_active_bookings] rows scan error: %w", err)
		}
		bookings = append(bookings, *booking)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[find_all_active_bookings] rows error: %w", err)
	}

	return bookings, nil
}

func (br *bookingRepo) Find(bookingID int64) (*entity.Booking, error) {
	booking, err := scanBooking(br.db.QueryRow(FindByID, bookingID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[find_booking] booking %d not founded", bookingID)
		}
		return nil, utils.WrapError(err)
	}

	return booking, nil
}

func (br *bookingRepo) Close(info *model.BookingInfo) (*entity.Booking, error) {
	booking, err := scanBooking(br.db.QueryRow(Close, info.Status, info.UserChatID, info.BookingID))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[close_booking] booking %d not founded", info.BookingID)
		}
		return nil, utils.WrapError(err)
	}

	return booking, nil
}

func scanBooking(scanner interface {
	Scan(dest ...any) error
}) (*entity.Booking, error) {
	var booking entity.Booking

	if err := scanner.Scan(
		&booking.ID,
		&booking.ChatID,
		&booking.UserPhone,
		&booking.Date,
		&booking.Time,
		&booking.Service,
		&booking.RimRadius,
		&booking.TotalPrice,
		&booking.Status,
		&booking.IsActive,
		&booking.CreatedAt,
		&booking.UpdatedAt,
		&booking.ConfirmedBy,
		&booking.CancelledBy,
		&booking.Notes,
	); err != nil {
		return nil, err
	}

	return &booking, nil
}
