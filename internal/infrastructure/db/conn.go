package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	pgxw "github.com/jackc/pgx/v5/stdlib"
	"github.com/pan-asovsky/brandd-tg-bot/internal/config"
	"github.com/pressly/goose/v3"
)

func NewPGXPool(cfg *config.Config, ctx context.Context) (*pgxpool.Pool, error) {
	dsn := cfg.DBDsn()

	pgxCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse conn: %w", err)
	}
	pgxCfg.MaxConns = 10
	pgxCfg.MaxConnLifetime = 10 * time.Minute
	pgxCfg.MaxConnIdleTime = 15 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, pgxCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping sql fb: %w", err)
	}

	sqlDB := pgxw.OpenDBFromPool(pool)
	if err = goose.Up(sqlDB, "migrations"); err != nil {
		return nil, fmt.Errorf("failed to migrate sql db: %w", err)
	}
	log.Println("sql migrations applied successfully")

	return pool, nil
}
