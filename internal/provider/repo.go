package provider

import (
	"database/sql"

	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
)

type repoProvider struct {
	db *sql.DB
}

func NewRepoProvider(db *sql.DB) iprovider.RepoProvider {
	return &repoProvider{db: db}
}

func (p *repoProvider) Service() irepo.ServiceRepo {
	return postgres.NewServiceRepo(p.db)
}

func (p *repoProvider) Price() irepo.PriceRepo {
	return postgres.NewPriceRepo(p.db)
}

func (p *repoProvider) Config() irepo.ConfigRepo {
	return postgres.NewConfigRepo(p.db)
}

func (p *repoProvider) Slot() irepo.SlotRepo {
	return postgres.NewSlotRepo(p.db)
}

func (p *repoProvider) Booking() irepo.BookingRepo {
	return postgres.NewBookingRepo(p.db)
}

func (p *repoProvider) User() irepo.UserRepo {
	return postgres.NewUserRepo(p.db)
}
