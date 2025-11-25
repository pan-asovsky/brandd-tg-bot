package postgres

import (
	"database/sql"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type BookingRepository interface {
	Create(booking *model.Booking) error
	FindActiveByTelegramID(telegramID int64) (*model.Booking, error)
	UpdateStatus(id int64, status model.BookingStatus) error
}

type bookingRepo struct {
	db *sql.DB
}

func (b *bookingRepo) Create(booking *model.Booking) error {
	//TODO implement me
	panic("implement me")
}

func (b *bookingRepo) FindActiveByTelegramID(telegramID int64) (*model.Booking, error) {
	//TODO implement me
	panic("implement me")
}

func (b *bookingRepo) UpdateStatus(id int64, status model.BookingStatus) error {
	//TODO implement me
	panic("implement me")
}

func NewBookingRepo(db *sql.DB) BookingRepository {
	return &bookingRepo{db: db}
}
