package provider

import (
	"database/sql"

	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interface/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/repository"
)

type repoProvider struct {
	db *sql.DB
}

func NewRepoProvider(db *sql.DB) iprovider.RepoProvider {
	return &repoProvider{db: db}
}

func (p *repoProvider) Service() irepo.ServiceRepo {
	return repository.NewServiceRepo(p.db)
}

func (p *repoProvider) Price() irepo.PriceRepo {
	return repository.NewPriceRepo(p.db)
}

func (p *repoProvider) Config() irepo.ConfigRepo {
	return repository.NewConfigRepo(p.db)
}

func (p *repoProvider) Slot() irepo.SlotRepo {
	return repository.NewSlotRepo(p.db)
}

func (p *repoProvider) Booking() irepo.BookingRepo {
	return repository.NewBookingRepo(p.db)
}

func (p *repoProvider) User() irepo.UserRepo {
	return repository.NewUserRepo(p.db)
}
