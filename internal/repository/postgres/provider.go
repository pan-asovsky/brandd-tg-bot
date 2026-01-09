package postgres

import (
	"database/sql"

	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
)

type Provider struct {
	db *sql.DB
}

func NewPgProvider(db *sql.DB) *Provider {
	return &Provider{db: db}
}

func (p *Provider) Service() i.ServiceRepo {
	return &serviceRepo{db: p.db}
}

func (p *Provider) Price() i.PriceRepo {
	return &priceRepo{db: p.db}
}

func (p *Provider) Config() i.ConfigRepo {
	return &configRepo{db: p.db}
}

func (p *Provider) Slot() i.SlotRepo {
	return &slotRepo{db: p.db}
}

func (p *Provider) Booking() i.BookingRepo {
	return &bookingRepo{db: p.db}
}

func (p *Provider) User() i.UserRepo {
	return &userRepo{db: p.db}
}
