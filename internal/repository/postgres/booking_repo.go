package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type BookingRepo interface {
	FindActiveByChatID(telegramID int64) (*model.Booking, error)
	UpdateStatus(id int64, status model.BookingStatus) error
	Save(booking *model.Booking) error
	SetPhone(phone string, chatID int64) error
	Confirm(chatID int64) error
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
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[find_booking_by_chat_id] not founded for %d: %v", chatID, err)
		}
		return nil, fmt.Errorf("[find_booking_by_chat_id] failed for %d: %v", chatID, err)
	}
	return &booking, nil
}

func (b *bookingRepo) UpdateStatus(id int64, status model.BookingStatus) error {
	return nil
}

func (b *bookingRepo) Save(booking *model.Booking) error {
	err := b.db.QueryRow(SaveBooking,
		booking.ChatID,
		booking.Date,
		booking.Time,
		booking.Service,
		booking.RimRadius,
		true,
		model.NotConfirmed,
		time.Now(),
		time.Now()).Scan(&booking.ID)
	if err := err; err != nil {
		return fmt.Errorf("[save_booking] error: %w", err)
	}
	return nil
}

func (b *bookingRepo) SetPhone(phone string, chatID int64) error {
	if _, err := b.db.Exec(SetPhoneByChatID, phone, chatID); err != nil {
		return fmt.Errorf("[set_phone_by_chat_id] query error: %w", err)
	}
	return nil
}

func (b *bookingRepo) Confirm(chatID int64) error {
	if _, err := b.db.Exec(ConfirmBooking, model.Confirmed, chatID); err != nil {
		return fmt.Errorf("[confirm_booking] error: %w", err)
	}
	return nil
}
