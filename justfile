# portfolio-backend justfile
# https://github.com/casey/just

set dotenv-load

default:
    @just --list

# ── Dev ───────────────────────────────────────────────────────────────────────

# Run the API server
run:
    go run ./cmd/api

# Run with hot-reload (requires air)
dev:
    air

# ── Build ─────────────────────────────────────────────────────────────────────

# Build the binary
build:
    go build -o portfolio-api ./cmd/api

# Clean build artifacts
clean:
    rm -f portfolio-api
    rm -rf tmp/

# ── Test ──────────────────────────────────────────────────────────────────────

# Run all tests
test:
    go test ./...

# Run tests for a specific layer (usecase | handler | mapper)
test-layer layer:
    go test ./test/{{layer}}/...

# Run a specific test by name
test-run name:
    go test ./... -run {{name}}

# ── Database ──────────────────────────────────────────────────────────────────

# Generate sqlc code from SQL queries
generate:
    sqlc generate

# Seed the database
seed:
    psql $DATABASE_URL -f sql/schema/seed.sql

# ── Docker ────────────────────────────────────────────────────────────────────

# Build and start the container
docker-up:
    docker compose up --build -d

# Stop and remove the container
docker-down:
    docker compose down

# Tail container logs
docker-logs:
    docker compose logs -f app
