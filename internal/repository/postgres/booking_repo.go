package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type BookingRepo interface {
	FindActiveByChatID(telegramID int64) (*model.Booking, error)
	ExistsByChatID(id int64) bool
	UpdateRimRadius(chatID int64, rimRadius string) error
	UpdateStatus(chatID int64, status model.BookingStatus) error
	UpdatePrice(chatID int64, price int64) error
	Save(booking *model.Booking) (*model.Booking, error)
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
	AutoConfirm(chatID int64) error
	Cancel(chatID int64) error
}

type bookingRepo struct {
	db *sql.DB
}

func (b *bookingRepo) FindActiveByChatID(chatID int64) (*model.Booking, error) {
	var booking model.Booking
	if err := b.db.QueryRow(FindActiveByChatID, chatID).Scan(
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[find_booking_by_chat_id] failed for %d: %v", chatID, err)
	}
	return &booking, nil
}

func (b *bookingRepo) ExistsByChatID(chatID int64) bool {
	var exists bool
	if err := b.db.QueryRow(ExistsByChatID, chatID).Scan(&exists); err != nil {
		return false
	}
	return exists
}

func (b *bookingRepo) UpdateRimRadius(chatID int64, rimRadius string) error {
	_, err := b.db.Exec(UpdateRimRadiusByChatID, rimRadius, chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (b *bookingRepo) UpdateStatus(chatID int64, status model.BookingStatus) error {
	_, err := b.db.Exec(UpdateStatusByChatID, status, chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (b *bookingRepo) UpdatePrice(chatID int64, price int64) error {
	_, err := b.db.Exec(UpdatePriceByChatID, price, chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	return nil
}

func (b *bookingRepo) Save(booking *model.Booking) (*model.Booking, error) {
	now := time.Now().UTC().Add(3 * time.Hour)
	err := b.db.QueryRow(SaveBooking,
		booking.ChatID,
		booking.Date,
		booking.Time,
		booking.Service,
		booking.RimRadius,
		booking.TotalPrice,
		true,
		model.NotConfirmed,
		now,
		now,
	).Scan(&booking.ID)
	if err != nil {
		return nil, utils.WrapError(err)
	}
	return booking, nil
}

func (b *bookingRepo) SetPhone(phone string, chatID int64) error {
	if _, err := b.db.Exec(SetPhoneByChatID, phone, chatID); err != nil {
		return fmt.Errorf("[set_phone_by_chat_id] query error: %w", err)
	}
	return nil
}

func (b *bookingRepo) Confirm(chatID int64) error {
	if _, err := b.db.Exec(ConfirmBooking, model.Confirmed, "an_user", chatID); err != nil {
		return fmt.Errorf("[confirm_booking] error: %w", err)
	}
	return nil
}

func (b *bookingRepo) AutoConfirm(chatID int64) error {
	if _, err := b.db.Exec(ConfirmBooking, model.Confirmed, "system", chatID); err != nil {
		return fmt.Errorf("[confirm_booking] error: %w", err)
	}
	return nil
}

func (b *bookingRepo) Cancel(chatID int64) error {
	if _, err := b.db.Exec(CancelBooking, model.Cancelled, chatID); err != nil {
		return fmt.Errorf("[cancel_booking] error: %w", err)
	}
	return nil
}
