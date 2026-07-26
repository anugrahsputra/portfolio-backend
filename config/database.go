package config

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/anugrahsputra/portfolio-backend/sql/migrations"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(cfg *Config) (*Database, error) {
	config, err := pgxpool.ParseConfig(cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	slog.Info("Database connected successfully")
	return &Database{Pool: pool}, nil
}

func (db *Database) Migrate() error {
	goose.SetBaseFS(migrations.FS)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}
	sqlDB := stdlib.OpenDBFromPool(db.Pool)
	defer sqlDB.Close()
	if err := goose.Up(sqlDB, "."); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	slog.Info("Database migrations applied")
	return nil
}

func (db *Database) Close() {
	db.Pool.Close()
}
