package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pan-asovsky/brandd-tg-bot/internal/config"
	"github.com/pressly/goose/v3"
)

func NewDBConn(cfg *config.Config, ctx context.Context) (*sql.DB, error) {
	dsn := cfg.DBDsn()

	//todo: conn, err := pgx.Connect(ctx, dsn)
	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to sql db: %w", err)
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(15 * time.Minute)

	ping, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err = sqlDB.PingContext(ping); err != nil {
		return nil, fmt.Errorf("failed to ping sql fb: %w", err)
	}

	goose.SetBaseFS(nil)
	//todo: as cfg field
	if err = goose.Up(sqlDB, "migrations"); err != nil {
		return nil, fmt.Errorf("failed to migrate sql db: %w", err)
	}
	log.Println("sql db migrations applied successfully")

	return sqlDB, nil
}
