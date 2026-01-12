package postgres

import (
	"database/sql"

	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
)

type provider struct {
	db *sql.DB
}

func NewRepoProvider(db *sql.DB) p.RepoProvider {
	return &provider{db: db}
}

func (p *provider) Service() i.ServiceRepo {
	return &serviceRepo{db: p.db}
}

func (p *provider) Price() i.PriceRepo {
	return &priceRepo{db: p.db}
}

func (p *provider) Config() i.ConfigRepo {
	return &configRepo{db: p.db}
}

func (p *provider) Slot() i.SlotRepo {
	return &slotRepo{db: p.db}
}

func (p *provider) Booking() i.BookingRepo {
	return &bookingRepo{db: p.db}
}

func (p *provider) User() i.UserRepo {
	return &userRepo{db: p.db}
}
