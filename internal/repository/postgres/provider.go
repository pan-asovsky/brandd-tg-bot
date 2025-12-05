package postgres

import "database/sql"

type Provider struct {
	db *sql.DB
}

func NewPgProvider(db *sql.DB) *Provider {
	return &Provider{db: db}
}

func (p *Provider) Service() ServiceRepo {
	return &serviceRepo{db: p.db}
}

func (p *Provider) Price() PriceRepo {
	return &priceRepo{db: p.db}
}

func (p *Provider) Config() ConfigRepo {
	return &configRepo{db: p.db}
}

func (p *Provider) Slot() SlotRepo {
	return &slotRepo{db: p.db}
}

func (p *Provider) Booking() BookingRepo {
	return &bookingRepo{db: p.db}
}
