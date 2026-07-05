# Redis Cache + Email Rate Limit Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add Redis-backed response caching for public GET endpoints and IP rate limiting for the email contact form, both fail-open if Redis is unreachable.

**Architecture:** A `config/redis.go` connector builds a `*redis.Client` from `REDIS_URL`. Two new Gin middlewares (`pkg/middleware/cache.go`, `pkg/middleware/ratelimit.go`) are wired into `setup_router.go` — cache globally (self-guards on GET), rate limit only on `/send-email`. Every Redis call is fail-open: an error is treated as "cache miss" / "not rate limited", never a request failure.

**Tech Stack:** `github.com/redis/go-redis/v9` (runtime), `github.com/alicebob/miniredis/v2` (test-only, in-process fake Redis).

Spec: `docs/superpowers/specs/2026-07-05-redis-cache-and-ratelimit-design.md`

---

## File Structure

- Create: `config/redis.go` — `NewRedis(cfg *Config) *redis.Client`
- Modify: `config/config.go` — add `RedisURL` field
- Create: `pkg/middleware/cache.go` — `CacheMiddleware(rdb, ttl)`
- Create: `pkg/middleware/ratelimit.go` — `RateLimitMiddleware(rdb, limit, window)`
- Create: `test/middleware/cache_middleware_test.go`
- Create: `test/middleware/ratelimit_middleware_test.go`
- Modify: `internal/delivery/route/setup_router.go` — accept `rdb`, wire both middlewares
- Modify: `cmd/api/main.go` — construct Redis client, pass to `SetupRouter`
- Modify: `.env.example` — document `REDIS_URL`
- Modify: `go.mod` / `go.sum` — add `go-redis/v9` and `miniredis/v2`

---

### Task 1: Add dependencies

**Files:**
- Modify: `go.mod`, `go.sum`

- [ ] **Step 1: Add runtime and test dependencies**

Run:
```bash
go get github.com/redis/go-redis/v9
go get github.com/alicebob/miniredis/v2
```

Expected: `go.mod` gains both requires, `go.sum` updated, no errors.

- [ ] **Step 2: Verify module resolves**

Run: `go build ./...`
Expected: builds successfully (no source uses the new deps yet, so this just confirms the module graph resolves).

- [ ] **Step 3: Commit**

```bash
git add go.mod go.sum
git commit -m "chore: add go-redis and miniredis dependencies"
```

---

### Task 2: Redis config + connector

**Files:**
- Modify: `config/config.go`
- Create: `config/redis.go`

- [ ] **Step 1: Add `RedisURL` to `Config`**

In `config/config.go`, add the field and env lookup:

```go
type Config struct {
	Port        string
	DatabaseURL string
	Env         string
	GmailUser   string
	GmailPass   string
	RedisURL    string
}

func Load() *Config {
	env := getEnv("ENV", "development")

	return &Config{
		Port:        getEnv("PORT", "8082"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"),
		Env:         env,
		GmailUser:   getEnv("GMAIL_USER", ""),
		GmailPass:   getEnv("GMAIL_PASS", ""),
		RedisURL:    getEnv("REDIS_URL", "redis://localhost:6379/0"),
	}
}
```

- [ ] **Step 2: Create `config/redis.go`**

```go
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
```

- [ ] **Step 3: Verify it builds**

Run: `go build ./...`
Expected: builds successfully.

- [ ] **Step 4: Commit**

```bash
git add config/config.go config/redis.go
git commit -m "feat(config): add Redis connector with fail-open startup"
```

---

### Task 3: Cache middleware (TDD)

**Files:**
- Create: `pkg/middleware/cache.go`
- Test: `test/middleware/cache_middleware_test.go`

- [ ] **Step 1: Write the failing tests**

Create `test/middleware/cache_middleware_test.go`:

```go
package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func newTestRedis(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
	mr, err := miniredis.Run()
	assert.NoError(t, err)
	t.Cleanup(mr.Close)

	client := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	return client, mr
}

func newTestRouter(rdb *redis.Client, hitCount *int) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.CacheMiddleware(rdb, time.Minute))
	r.GET("/thing", func(c *gin.Context) {
		*hitCount++
		c.JSON(http.StatusOK, gin.H{"count": *hitCount})
	})
	r.POST("/thing", func(c *gin.Context) {
		*hitCount++
		c.JSON(http.StatusOK, gin.H{"count": *hitCount})
	})
	return r
}

func TestCacheMiddleware_SecondGetIsServedFromCache(t *testing.T) {
	rdb, _ := newTestRedis(t)
	hits := 0
	r := newTestRouter(rdb, &hits)

	w1 := httptest.NewRecorder()
	req1 := httptest.NewRequest(http.MethodGet, "/thing", nil)
	r.ServeHTTP(w1, req1)
	assert.Equal(t, http.StatusOK, w1.Code)
	assert.JSONEq(t, `{"count":1}`, w1.Body.String())

	w2 := httptest.NewRecorder()
	req2 := httptest.NewRequest(http.MethodGet, "/thing", nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, http.StatusOK, w2.Code)
	assert.JSONEq(t, `{"count":1}`, w2.Body.String())
	assert.Equal(t, 1, hits, "handler should only run once; second GET served from cache")
}

func TestCacheMiddleware_NonGetBypassesCache(t *testing.T) {
	rdb, _ := newTestRedis(t)
	hits := 0
	r := newTestRouter(rdb, &hits)

	for i := 0; i < 2; i++ {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/thing", nil)
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	}

	assert.Equal(t, 2, hits, "POST requests must never be cached")
}

func TestCacheMiddleware_FailsOpenWhenRedisDown(t *testing.T) {
	rdb, mr := newTestRedis(t)
	mr.Close()

	hits := 0
	r := newTestRouter(rdb, &hits)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/thing", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, 1, hits, "request must still succeed via the handler when Redis is unreachable")
}
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `go test ./test/middleware/... -run TestCacheMiddleware -v`
Expected: FAIL — `middleware.CacheMiddleware` undefined.

- [ ] **Step 3: Implement `pkg/middleware/cache.go`**

```go
package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type cachedResponse struct {
	Status      int    `json:"status"`
	ContentType string `json:"content_type"`
	Body        []byte `json:"body"`
}

type bodyCapturingWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyCapturingWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func CacheMiddleware(rdb *redis.Client, ttl time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodGet {
			c.Next()
			return
		}

		ctx := context.Background()
		key := "cache:" + c.Request.URL.RequestURI()

		cached, err := rdb.Get(ctx, key).Bytes()
		if err == nil {
			var resp cachedResponse
			if json.Unmarshal(cached, &resp) == nil {
				c.Header("Content-Type", resp.ContentType)
				c.Data(resp.Status, resp.ContentType, resp.Body)
				c.Abort()
				return
			}
		}

		writer := &bodyCapturingWriter{ResponseWriter: c.Writer, body: &bytes.Buffer{}}
		c.Writer = writer

		c.Next()

		if writer.Status() != http.StatusOK {
			return
		}

		encoded, err := json.Marshal(cachedResponse{
			Status:      writer.Status(),
			ContentType: writer.Header().Get("Content-Type"),
			Body:        writer.body.Bytes(),
		})
		if err != nil {
			return
		}

		rdb.Set(ctx, key, encoded, ttl)
	}
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `go test ./test/middleware/... -run TestCacheMiddleware -v`
Expected: PASS (all three tests).

- [ ] **Step 5: Commit**

```bash
git add pkg/middleware/cache.go test/middleware/cache_middleware_test.go
git commit -m "feat(middleware): add fail-open Redis cache middleware for GET requests"
```

---

### Task 4: Rate limit middleware (TDD)

**Files:**
- Create: `pkg/middleware/ratelimit.go`
- Test: `test/middleware/ratelimit_middleware_test.go`

- [ ] **Step 1: Write the failing tests**

Create `test/middleware/ratelimit_middleware_test.go`:

```go
package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func newRateLimitRouter(rdb *redis.Client, limit int, window time.Duration) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/send-email", middleware.RateLimitMiddleware(rdb, limit, window), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"sent": true})
	})
	return r
}

func doPost(r *gin.Engine, ip string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/send-email", nil)
	req.RemoteAddr = ip + ":12345"
	r.ServeHTTP(w, req)
	return w
}

func TestRateLimitMiddleware_AllowsUpToLimitThenBlocks(t *testing.T) {
	rdb, _ := newTestRedis(t)
	r := newRateLimitRouter(rdb, 3, time.Minute)

	for i := 0; i < 3; i++ {
		w := doPost(r, "1.2.3.4")
		assert.Equal(t, http.StatusOK, w.Code, "request %d should be allowed", i+1)
	}

	w := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusTooManyRequests, w.Code, "4th request in window should be blocked")
}

func TestRateLimitMiddleware_DifferentIPsTrackedSeparately(t *testing.T) {
	rdb, _ := newTestRedis(t)
	r := newRateLimitRouter(rdb, 1, time.Minute)

	w1 := doPost(r, "1.1.1.1")
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := doPost(r, "2.2.2.2")
	assert.Equal(t, http.StatusOK, w2.Code, "different IP must have its own counter")
}

func TestRateLimitMiddleware_ResetsAfterWindowExpires(t *testing.T) {
	rdb, mr := newTestRedis(t)
	r := newRateLimitRouter(rdb, 1, time.Minute)

	w1 := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusOK, w1.Code)

	w2 := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusTooManyRequests, w2.Code)

	mr.FastForward(time.Minute + time.Second)

	w3 := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusOK, w3.Code, "counter must reset once the window elapses")
}

func TestRateLimitMiddleware_FailsOpenWhenRedisDown(t *testing.T) {
	rdb, mr := newTestRedis(t)
	mr.Close()

	r := newRateLimitRouter(rdb, 1, time.Minute)

	w := doPost(r, "1.2.3.4")
	assert.Equal(t, http.StatusOK, w.Code, "request must succeed when Redis is unreachable")
}
```

- [ ] **Step 2: Run tests to verify they fail**

Run: `go test ./test/middleware/... -run TestRateLimitMiddleware -v`
Expected: FAIL — `middleware.RateLimitMiddleware` undefined.

- [ ] **Step 3: Implement `pkg/middleware/ratelimit.go`**

```go
package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RateLimitMiddleware(rdb *redis.Client, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		key := "ratelimit:sendemail:" + c.ClientIP()

		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			c.Next()
			return
		}

		if count == 1 {
			rdb.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.JSON(http.StatusTooManyRequests, dto.NoDataResponse{
				Status:  http.StatusTooManyRequests,
				Message: "Too many requests, please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
```

- [ ] **Step 4: Run tests to verify they pass**

Run: `go test ./test/middleware/... -run TestRateLimitMiddleware -v`
Expected: PASS (all four tests).

- [ ] **Step 5: Commit**

```bash
git add pkg/middleware/ratelimit.go test/middleware/ratelimit_middleware_test.go
git commit -m "feat(middleware): add fail-open Redis rate limit middleware"
```

---

### Task 5: Wire middlewares into the router and main

**Files:**
- Modify: `internal/delivery/route/setup_router.go`
- Modify: `cmd/api/main.go`
- Modify: `.env.example`

- [ ] **Step 1: Update `setup_router.go` signature and imports**

In `internal/delivery/route/setup_router.go`, change the import block to add `"time"` and `"github.com/redis/go-redis/v9"`:

```go
import (
	"os"
	"time"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/repository"
	"github.com/anugrahsputra/portfolio-backend/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)
```

Change the function signature:

```go
func SetupRouter(db *config.Database, mail *config.Mail, cfg *config.Config, rdb *redis.Client) *gin.Engine {
```

- [ ] **Step 2: Register the cache middleware globally**

Immediately after the existing `route.Use(cors.New(...))` block (still before `apiKey := os.Getenv("API_KEY")`), add:

```go
	route.Use(middleware.CacheMiddleware(rdb, 60*time.Second))
```

- [ ] **Step 3: Register the rate limit middleware on send-email**

Find:

```go
		se := api.Group("/send-email")
		se.Use(middleware.AuthMiddleware(apiKey))
		se.POST("/", contactForm.SendMail)
```

Replace with:

```go
		se := api.Group("/send-email")
		se.Use(middleware.AuthMiddleware(apiKey))
		se.Use(middleware.RateLimitMiddleware(rdb, 3, time.Minute))
		se.POST("/", contactForm.SendMail)
```

- [ ] **Step 4: Update `cmd/api/main.go` to construct and pass the Redis client**

Add the call after `mail, err := config.NewMail(cfg)` block and before `r := route.SetupRouter(...)`:

```go
	rdb := config.NewRedis(cfg)

	r := route.SetupRouter(db, mail, cfg, rdb)
```

- [ ] **Step 5: Document `REDIS_URL` in `.env.example`**

Append to `.env.example`:

```
# Redis Configuration
REDIS_URL=redis://localhost:6379/0
```

- [ ] **Step 6: Build and run full test suite**

Run: `go build ./... && go test ./...`
Expected: builds successfully, all tests pass (including the new middleware tests and all pre-existing handler/usecase/mapper tests).

- [ ] **Step 7: Commit**

```bash
git add internal/delivery/route/setup_router.go cmd/api/main.go .env.example
git commit -m "feat(route): wire Redis cache and rate-limit middlewares into the router"
```

---

## Self-Review Notes

- **Spec coverage:** cache (Task 3), rate limit (Task 4), config/connector + fail-open startup (Task 2), wiring + REDIS_URL doc (Task 5), miniredis-based tests with no real Redis required (Tasks 3 & 4) — all spec sections covered. No docker-compose changes, per spec's non-goals.
- **Placeholder scan:** no TBD/TODO; every step has runnable code and exact commands.
- **Type consistency:** `CacheMiddleware(rdb *redis.Client, ttl time.Duration)` and `RateLimitMiddleware(rdb *redis.Client, limit int, window time.Duration)` signatures match between test files (Tasks 3–4) and wiring (Task 5). `cachedResponse` struct is internal to `cache.go`, not referenced elsewhere.
