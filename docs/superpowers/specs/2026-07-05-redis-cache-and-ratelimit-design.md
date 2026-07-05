# Redis: response cache + email rate limit

Date: 2026-07-05

## Goal

Add Redis to the portfolio backend for two purposes:
1. Cache public GET responses (profile, profile-url, experience, education, skill, language, project, resume) — these are read-heavy and rarely written (admin-only, gated by API key).
2. Rate-limit `POST /send-email` to prevent spam/abuse.

## Non-goals

- No session storage (auth is a static API key, no sessions/JWT).
- No active cache invalidation on writes.
- No rate limiting on any endpoint other than `/send-email`.
- No Redis service added to `docker-compose.yaml` — Redis is assumed external, same pattern as Postgres (`DATABASE_URL` today).

## Architecture

- `config/redis.go` (new): `NewRedis(cfg *Config) *redis.Client`. Builds client via `redis.ParseURL(cfg.RedisURL)` + `redis.NewClient`. Pings on startup; on failure, logs a warning via `slog` and returns the client anyway — startup must not fail if Redis is down (fail-open applies at startup too).
- `config/config.go`: add `RedisURL string` field, env var `REDIS_URL`, default `redis://localhost:6379/0`.
- `cmd/api/main.go`: construct the Redis client after config load, pass it into `route.SetupRouter`.
- `go.mod`: add `github.com/redis/go-redis/v9` (runtime) and `github.com/alicebob/miniredis/v2` (test-only).

## Components

### 1. Cache middleware — `pkg/middleware/cache.go`

`CacheMiddleware(rdb *redis.Client, ttl time.Duration) gin.HandlerFunc`

- Runs on every request (registered once via `route.Use(...)` in `setup_router.go`, after CORS).
- Guards internally: if `c.Request.Method != http.MethodGet`, call `c.Next()` and return immediately — no Redis calls for non-GET.
- Cache key: `"cache:" + c.Request.URL.RequestURI()` (full path + query string).
- On GET:
  - `rdb.Get(ctx, key)` — on hit, JSON-decode a small struct `{Status int, ContentType string, Body []byte}`, write it directly to the response, `c.Abort()`.
  - On miss or any Redis error, wrap the `gin.ResponseWriter` to capture status/body, call `c.Next()`.
  - After the handler runs: if captured status is `200`, `rdb.Set(ctx, key, encoded, ttl)`. Ignore/log errors from `Set` — never fail the request over a cache-write failure.
- TTL: 60 seconds, passed in from `setup_router.go`.

### 2. Rate limit middleware — `pkg/middleware/ratelimit.go`

`RateLimitMiddleware(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc`

- Applied only to the `send-email` route group, alongside the existing `AuthMiddleware`.
- Key: `"ratelimit:sendemail:" + c.ClientIP()`.
- Logic: `INCR` the key. If the returned count is `1` (first hit in this window), set `EXPIRE key window`. If count `> limit`, respond `429` with the existing `dto.NoDataResponse` shape and `c.Abort()`.
- Limit: 3 requests per 60-second window.
- Any Redis error during `INCR`/`EXPIRE` → fail-open: log a warning, let the request through.

### 3. Wiring — `internal/delivery/route/setup_router.go`

- `SetupRouter` gains a `rdb *redis.Client` parameter.
- `route.Use(middleware.CacheMiddleware(rdb, 60*time.Second))` added after the CORS middleware, before route groups are registered.
- `se.Use(middleware.RateLimitMiddleware(rdb, 3, time.Minute))` added to the existing `se := api.Group("/send-email")` group, alongside `se.Use(middleware.AuthMiddleware(apiKey))`.

## Data flow

`GET /api/v1/project/:profile_id`:
1. CacheMiddleware checks Redis for `cache:/api/v1/project/<id>`.
2. Miss → handler → usecase → repository → Postgres. Response captured, cached for 60s.
3. Any request to the same URL within 60s is served directly from Redis; Postgres is not touched.

`POST /api/v1/send-email`:
1. AuthMiddleware checks API key (unchanged).
2. RateLimitMiddleware increments `ratelimit:sendemail:<ip>`; blocks with 429 past 3/min.
3. Passes through to the existing contact-form handler.

## Error handling

Redis is never a hard dependency for request handling:
- Cache get/set errors → treated as cache miss / skip cache write. Request still succeeds via Postgres.
- Rate-limit incr/expire errors → treated as "not rate limited". Request still succeeds.
- Startup Redis ping failure → logged, server still starts.

This is fail-open throughout, per explicit decision — availability of the API is prioritized over caching/rate-limiting correctness.

## Testing

Using `github.com/alicebob/miniredis/v2` (in-process fake Redis, test dependency only):
- Cache middleware: first GET misses and stores; second identical GET is served from cache (handler not invoked again — assert via a call counter); non-GET requests bypass Redis entirely; a `POST`/`PUT`/`DELETE` never gets cached; simulate a closed miniredis connection to assert fail-open (request still succeeds).
- Rate limit middleware: first `limit` requests within the window succeed; the next one gets 429; after the window elapses (or key manually expired in the test), requests succeed again; simulate a closed connection to assert fail-open.

No real Redis instance required for `go test ./...`.
