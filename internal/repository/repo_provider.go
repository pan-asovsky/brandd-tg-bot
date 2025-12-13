package repository

import (
	"database/sql"

	"github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/redis/go-redis/v9"
)

type RepoProvider struct {
	db    *sql.DB
	cache *redis.Client
}

func NewRepoProvider(db *sql.DB, cache *redis.Client) *RepoProvider {
	return &RepoProvider{db: db, cache: cache}
}

func (rp *RepoProvider) PgProvider() *postgres.Provider {
	return postgres.NewPgProvider(rp.db)
}
