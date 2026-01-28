package provider

import (
	"database/sql"

	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/repo"
)

type repoProvider struct {
	db *sql.DB
}

func NewRepoProvider(db *sql.DB) iprovider.RepoProvider {
	return &repoProvider{db: db}
}

func (p *repoProvider) Service() irepo.ServiceRepo {
	return repo.NewServiceRepo(p.db)
}

func (p *repoProvider) Price() irepo.PriceRepo {
	return repo.NewPriceRepo(p.db)
}

func (p *repoProvider) Config() irepo.ConfigRepo {
	return repo.NewConfigRepo(p.db)
}

func (p *repoProvider) Slot() irepo.SlotRepo {
	return repo.NewSlotRepo(p.db)
}

func (p *repoProvider) Booking() irepo.BookingRepo {
	return repo.NewBookingRepo(p.db)
}

func (p *repoProvider) User() irepo.UserRepo {
	return repo.NewUserRepo(p.db)
}
