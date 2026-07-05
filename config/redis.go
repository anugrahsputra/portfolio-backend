package config

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *Config) *redis.Client {
	opts, err := redis.ParseURL(cfg.RedisURL)
	if err != nil {
		slog.Error("Failed to parse Redis URL, using client with default options", "error", err)
		opts = &redis.Options{Addr: "localhost:6379"}
	}

	client := redis.NewClient(opts)

	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Warn("Redis ping failed, continuing without cache/rate-limit guarantees", "error", err)
	} else {
		slog.Info("Redis connected successfully")
	}

	return client
}
