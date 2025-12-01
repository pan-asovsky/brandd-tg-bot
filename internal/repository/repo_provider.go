package repository

import (
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
)

type RepoProvider interface {
	Service() pg.ServiceRepo
	Price() pg.PriceRepo
	Config() pg.ConfigRepo
	Slot() pg.SlotRepo
	Booking() pg.BookingRepo
}
