package postgres

import (
	"database/sql"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type BookingRepo interface {
	Create(booking *model.Booking) error
	FindActiveByTelegramID(telegramID int64) (*model.Booking, error)
	UpdateStatus(id int64, status model.BookingStatus) error
}

type bookingRepo struct {
	db *sql.DB
}

func (b *bookingRepo) Create(booking *model.Booking) error {
	return nil
}

func (b *bookingRepo) FindActiveByTelegramID(telegramID int64) (*model.Booking, error) {
	//TODO implement me
	return nil, nil
}

func (b *bookingRepo) UpdateStatus(id int64, status model.BookingStatus) error {
	return nil
}

func NewBookingRepo(db *sql.DB) BookingRepo {
	return &bookingRepo{db: db}
}
