package config

import (
	"context"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *Config) *redis.Client {
	opts, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		slog.Error("Failed to parse Redis URL, using client with default options", "error", err)
		opts = &redis.Options{Addr: "localhost:6379"}
	}

	// ponytail: fail-open only helps if it also fails fast — cap retries/timeouts so an
	// unreachable Redis costs milliseconds per request, not seconds.
	opts.MaxRetries = 1
	opts.DialTimeout = 200 * time.Millisecond
	opts.ReadTimeout = 200 * time.Millisecond
	opts.WriteTimeout = 200 * time.Millisecond

	client := redis.NewClient(opts)

	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Warn("Redis ping failed, continuing without cache/rate-limit guarantees", "error", err)
	} else {
		slog.Info("Redis connected successfully")
	}

	return client
}
