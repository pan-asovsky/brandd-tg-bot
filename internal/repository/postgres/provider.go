package postgres

import "database/sql"

type PgProvider struct {
	db *sql.DB
}

func NewPgProvider(db *sql.DB) *PgProvider {
	return &PgProvider{db: db}
}

func (p *PgProvider) Service() ServiceRepo {
	return &serviceRepo{db: p.db}
}

func (p *PgProvider) Price() PriceRepo {
	return &priceRepo{db: p.db}
}

func (p *PgProvider) Config() ConfigRepo {
	return &configRepo{db: p.db}
}

func (p *PgProvider) Slot() SlotRepo {
	return &slotRepo{db: p.db}
}

func (p *PgProvider) Booking() BookingRepo {
	return &bookingRepo{db: p.db}
}
