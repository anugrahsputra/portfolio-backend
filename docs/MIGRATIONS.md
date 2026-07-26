# Database Migrations

How to add versioned schema migrations to this project, and how to use them
afterwards.

## Why

Current state: `sql/schema/schema.sql` is re-run wholesale when `INIT_SCHEMA=true`.
Every statement is `create ... if not exists`, so it works on an empty database and
does nothing on an existing one. Add a column and that file never applies it. No
history, no rollback, no way to tell which database is at which version.

This guide uses [goose](https://github.com/pressly/goose): plain SQL files, one
binary, and sqlc understands its annotations.

Lazier alternative: numbered `.sql` files applied by hand with `psql` — works until
you forget which ones ran. Heavier alternatives: golang-migrate, atlas.

## Concept first

A migration is one **ordered, immutable** schema change, applied exactly once. The
tool keeps a table (`goose_db_version`) listing applied versions. On `up` it diffs
the files against that table and runs the missing ones in order.

Three rules; everything else is detail:

1. **Never edit an applied migration.** Fix forward with a new one.
2. **Each file is one logical change**, with an Up and a Down.
3. **The schema is the sum of all migrations.** No separate `schema.sql` truth.

---

## Step 1 — Install goose

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Verify with `goose -version`. Requires `$(go env GOPATH)/bin` on `PATH`.

## Step 2 — Create the migrations directory

```bash
mkdir -p sql/migrations
```

Why a new directory instead of `sql/schema/`: sqlc reads every `.sql` file in its
schema directory, and `seed.sql` lives there. A dedicated directory stays pure DDL.

## Step 3 — Write the baseline migration

The first migration is the current schema, so that a fresh database and the existing
one converge on the same state.

Create `sql/migrations/00001_init.sql` — this is `sql/schema/schema.sql` verbatim,
wrapped in goose annotations:

```sql
-- +goose Up
-- ...paste all of sql/schema/schema.sql here, unchanged...

-- +goose Down
drop table if exists projects;
drop table if exists skills;
drop table if exists languages;
drop table if exists experiences;
drop table if exists educations;
drop table if exists profile_urls;
drop table if exists profiles;
drop type if exists proficiency_level;
```

Two things to understand here:

- **Write plain SQL, exactly as it would run on an empty database.** Do not try to
  make the baseline idempotent — in particular, do not wrap `create type` in a
  `do $$ ... $$` guard to survive a re-run. sqlc's parser treats such a block as an
  opaque string, never sees the enum, and silently degrades
  `Language.Proficiency` from `ProficiencyLevel` to `interface{}`. Idempotency is
  not the migration's job; the version table is.
- **Down order** — the reverse of Up. Children before parents, because of the foreign
  keys.

A file containing `do $$`, a function body, or a trigger genuinely needs
`-- +goose StatementBegin` / `-- +goose StatementEnd` around it, because goose splits
on `;` and those blocks contain their own semicolons. The baseline needs neither.

## Step 4 — Apply it

On an **empty** database, just run it:

```bash
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
goose -dir sql/migrations postgres "$DATABASE_URL" up
goose -dir sql/migrations postgres "$DATABASE_URL" status
```

Expect `00001_init.sql` to show as applied, and confirm the bookkeeping table exists:

```bash
psql "$DATABASE_URL" -c '\d goose_db_version'
```

### Baselining a database that already has the schema

Anything previously created by `INIT_SCHEMA=true` — production, a teammate's local
database — already has these tables and the `proficiency_level` type. Running the
baseline there fails with `type "proficiency_level" already exists (SQLSTATE 42710)`.
That is correct behaviour: the migration must not run, it must be *recorded as
already run*. Once per such database:

```bash
# 1. creates the (empty) goose_db_version table and reports 00001 as Pending
goose -dir sql/migrations postgres "$DATABASE_URL" status

# 2. record the baseline as applied, without executing it
psql "$DATABASE_URL" -c "insert into goose_db_version (version_id, is_applied) values (1, true);"

# 3. now reports 00001 as Applied; migration 00002 onward will run normally
goose -dir sql/migrations postgres "$DATABASE_URL" status
```

This is a one-time transition step. Every migration after the baseline runs the
ordinary way, everywhere.

## Step 5 — Point sqlc at the migrations

In `sqlc.yaml`:

```yaml
    schema: "sql/migrations"     # was: "sql/schema"
```

sqlc parses goose annotations and ignores `-- +goose Down`, so the types it generates
always match the migrated database.

```bash
sqlc generate && go build ./... && go test ./...
```

Expect zero diff in `internal/db/`, since the schema itself has not changed. If files
do change, the baseline has drifted from `schema.sql` — fix the baseline, not the
generated code. Watch specifically for a field turning into `interface{}`: that means
sqlc could not parse the type's declaration.

## Step 6 — Run migrations on startup (optional, recommended for the Docker deploy)

The container has no `goose` binary and no `psql`. Embed the migration files instead.

`sql/migrations/embed.go`:

```go
package migrations

import "embed"

//go:embed *.sql
var FS embed.FS
```

In `config/database.go`, replacing `InitSchema`:

```go
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
```

In `cmd/api/main.go`, replacing the `INIT_SCHEMA` block:

```go
	if err := db.Migrate(); err != nil {
		slog.Error("Failed to run migrations", "error", err)
		os.Exit(1)
	}
```

Note the `os.Exit(1)`: if the schema is wrong the app must not serve. The old code
only logged and carried on.

Then clean up: drop `COPY --from=builder /app/sql/schema/schema.sql ...` from the
Dockerfile, drop `INIT_SCHEMA` from `.env.example` and the README, delete
`InitSchema`, and delete `sql/schema/schema.sql`. Keep `seed.sql`.

Dependencies: `go get github.com/pressly/goose/v3` (`jackc/pgx/v5/stdlib` is already
in the module).

## Step 7 — Add justfile recipes

```make
# Create a new migration
migrate-new name:
    goose -dir sql/migrations -s create {{name}} sql

# Apply all pending migrations
migrate-up:
    goose -dir sql/migrations postgres "$DATABASE_URL" up

# Roll back the last migration
migrate-down:
    goose -dir sql/migrations postgres "$DATABASE_URL" down

# Show migration status
migrate-status:
    goose -dir sql/migrations postgres "$DATABASE_URL" status
```

`set dotenv-load` is already in the justfile, so `$DATABASE_URL` comes from `.env`.

---

## The loop you use from now on

Say you add `profiles.headline`:

```bash
just migrate-new add_profile_headline
```

That writes `sql/migrations/00002_add_profile_headline.sql`. The `-s` flag in the
recipe selects sequential numbering; goose defaults to timestamp prefixes. This
project uses sequential — do not mix the two.

```sql
-- +goose Up
alter table profiles add column headline text;

-- +goose Down
alter table profiles drop column headline;
```

Then:

```bash
just migrate-up          # verify: just migrate-status shows it applied
just migrate-down        # verify: rollback works. do this once, now, not in prod
just migrate-up
# edit sql/queries/profiles.sql to select/insert the new column
just generate            # sqlc regenerates internal/db from the migrations
just test
```

Repo-specific: a new column also means touching the `internal/domain` entity, its
`Merge` method, the repository's `toDomain()`, and the DTO. The migration is only the
database half.

## Traps

- **An applied migration is frozen.** Once it has run anywhere but your laptop, write
  a new migration to fix it.
- **`NOT NULL` on a populated table takes three steps**: add the column nullable,
  backfill it with an `update`, then `set not null`. One migration with three
  statements is fine.
- **Down migrations lie about data.** `drop column` rolls back the schema, not the
  deleted values. The rollback plan for production is a backup restore; `down` is a
  development convenience.
- **Postgres DDL is transactional**, so a failed migration rolls back as a whole and
  leaves nothing half-applied. The exception is `create index concurrently`, which
  needs `-- +goose NO TRANSACTION`.
- **sqlc reads the migrations, so keep them parseable.** Anything hidden inside a
  `do $$ ... $$` block or a function body is invisible to sqlc, and types declared
  there disappear from `internal/db/` without any error — the generated field just
  becomes `interface{}`. Always check `git diff internal/db/` after `sqlc generate`.
- **Two developers, two new migrations** produce conflicting version numbers.
  Sequential numbering collides loudly (good); timestamps merge silently but may then
  apply in an unintended order.

## Not covered

Down-migration tests, a CI check for pending migrations, and a production runbook.
Add them once there is more than one environment.
