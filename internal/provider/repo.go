package provider

import (
	"github.com/jackc/pgx/v5/pgxpool"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/repo"
)

type repoProvider struct {
	pool *pgxpool.Pool
}

func NewRepoProvider(p *pgxpool.Pool) iprovider.RepoProvider {
	return &repoProvider{pool: p}
}

func (p *repoProvider) Service() irepo.ServiceRepo {
	return repo.NewServiceRepo(p.pool)
}

func (p *repoProvider) Price() irepo.PriceRepo {
	return repo.NewPriceRepo(p.pool)
}

func (p *repoProvider) Config() irepo.ConfigRepo {
	return repo.NewConfigRepo(p.pool)
}

func (p *repoProvider) Slot() irepo.SlotRepo {
	return repo.NewSlotRepo(p.pool)
}

func (p *repoProvider) Booking() irepo.BookingRepo {
	return repo.NewBookingRepo(p.pool)
}

func (p *repoProvider) User() irepo.UserRepo {
	return repo.NewUserRepo(p.pool)
}
