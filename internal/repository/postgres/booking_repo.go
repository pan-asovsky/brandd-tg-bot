package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type BookingRepo interface {
	FindActiveByChatID(telegramID int64) (*model.Booking, error)
	UpdateStatus(id int64, status model.BookingStatus) error
	Save(booking *model.Booking) error
	SetPhone(phone string, chatID int64) error
}

type bookingRepo struct {
	db *sql.DB
}

func (b *bookingRepo) FindActiveByChatID(chatID int64) (*model.Booking, error) {
	var booking *model.Booking
	if err := b.db.QueryRow(FindActiveByChatID, chatID).Scan(&booking); err != nil {
		return nil, fmt.Errorf("[find_active_booking_by_chat_id] failed for chat %v: %w", chatID, err)
	}
	return booking, nil
}

func (b *bookingRepo) UpdateStatus(id int64, status model.BookingStatus) error {
	return nil
}

func (b *bookingRepo) Save(booking *model.Booking) error {
	err := b.db.QueryRow(SaveBooking,
		booking.ChatID,
		booking.SlotID,
		booking.ServiceTypeID,
		booking.RimRadius,
		true,
		model.NotConfirmed,
		time.Now(),
		time.Now()).Scan(&booking.ID)
	if err := err; err != nil {
		return fmt.Errorf("[save_booking] error: %w", err)
	}
	//log.Printf("[save_booking] booking %v saved", booking.ID)
	return nil
}

func (b *bookingRepo) SetPhone(phone string, chatID int64) error {
	_, err := b.db.Exec(SetPhoneByChatID, phone, chatID)
	if err != nil {
		return fmt.Errorf("[set_phone_by_chat_id] error: %w", err)
	}
	return nil
}
