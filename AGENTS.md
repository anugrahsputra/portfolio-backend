# Portfolio Backend - Project Context

Go-based backend for a personal portfolio system using Gin, PostgreSQL, and SQLC.

## Commands

```bash
# Run all tests
go test ./test/...

# Run a single test file
go test ./test/usecase/profile_usecase_test.go

# Run a single test function
go test ./test/usecase/ -run TestCreateProfile -v

# Run tests in a specific package
go test ./test/handler/ -v

# Generate DB code from SQL
sqlc generate

# Run the application
go run cmd/api/main.go

# Build
go build -o portfolio-api cmd/api/main.go

# Format code
gofmt -w .

# Vet code
go vet ./...
```

## Architecture

Clean/Hexagonal architecture with layered responsibilities:

- `internal/domain/` - Entities and repository interfaces
- `internal/repository/` - Concrete repository implementations
- `internal/usecase/` - Business logic (use cases)
- `internal/delivery/handler/` - HTTP handlers (Gin)
- `internal/delivery/dto/` - Request/response DTOs
- `internal/delivery/route/` - Route definitions
- `internal/mapper/` - DB model <-> domain entity conversion
- `internal/db/` - SQLC-generated code (DO NOT EDIT)
- `sql/schema/` - PostgreSQL schema definitions
- `sql/queries/` - SQL queries for SQLC
- `config/` - DB connection and environment config
- `test/` - All tests (handler/, usecase/, mapper/)

## Code Style

### Imports

- Group imports: stdlib, third-party, internal (blank line between groups)
- Use full module path: `github.com/anugrahsputra/portfolio-backend/internal/...`
- Remove unused imports immediately

### Naming Conventions

- **Packages:** lowercase, single word (`domain`, `mapper`, `handler`)
- **Interfaces:** `EntityRepository`, `EntityUsecase` (e.g., `ProfileRepository`)
- **Structs:** `EntityHandler`, `EntityMapper`, `EntityReq`, `EntityRes`
- **Constructor functions:** `NewProfileHandler`, `NewProfileUsecase`
- **DTOs:** `ProfileReq`, `ProfileRes`, `NoDataResponse`
- **Domain inputs:** `ProfileInput`, `ProfileUpdateInput` (pointer fields for partial updates)

### Types

- Use `string` for UUIDs in domain layer, `uuid.UUID` in DB layer
- Use pointers (`*string`) for optional/partial update fields
- Domain entities use plain types; mappers handle DB conversion
- JSON tags on DTOs only, not on domain structs

### Error Handling

- Wrap errors with context: `fmt.Errorf("operation failed: %w", err)`
- Return `domain.ErrNotFound` for missing resources
- Handlers log via `c.Error(err)` and return user-friendly messages
- Never expose internal error details to clients

### HTTP Handlers

- Extract context: `ctx := c.Request.Context()`
- Bind JSON: `c.ShouldBindJSON(&req)`
- Use `dto.Response{}` for success with data, `dto.NoDataResponse{}` for no-data responses
- Set appropriate HTTP status codes (201 created, 200 ok, 400 bad request, 404 not found, 500 server error)

### Testing

- Test files live in `test/` mirroring internal structure
- Package suffix `_test` (e.g., `usecase_test`)
- Use `testify` for assertions (`assert.NoError`, `assert.Equal`) and mocks (`mock.Mock`)
- Mock repositories at the domain interface level, not DB layer
- Table-driven tests with `t.Run("success", ...)` and `t.Run("error", ...)`
- Always call `mockRepo.AssertExpectations(t)`

## Development Conventions

- **SQL-First:** Update `sql/schema/` or `sql/queries/`, then run `sqlc generate` before modifying repositories
- **Domain Isolation:** Repositories return domain entities, never SQLC models. Use `internal/mapper` for conversions
- **Never edit `internal/db/`** - it is SQLC-generated
- **Env vars:** `DATABASE_URL` and `PORT` with fallbacks in `config/config.go`
- **Logging:** `github.com/op/go-logging`

## Backlog Workflow

This project uses Backlog.md MCP for task management. Read `backlog://workflow/overview` before creating tasks.
