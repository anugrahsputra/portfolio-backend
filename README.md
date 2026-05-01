# Portfolio Backend

A Go-based REST API backend for managing personal portfolio data, including profiles, education, work experience, languages, skills, and projects.

## Features

- Profile management (CRUD operations)
- Education history tracking
- Work experience management
- Language proficiency tracking
- Technical skills management
- Profile URLs management
- Project portfolio management
- Full resume retrieval (aggregated data)
- Clean/Hexagonal architecture
- PostgreSQL database with SQLC code generation

## Tech Stack

- **Language:** Go (v1.25.0+)
- **Database:** PostgreSQL
- **DB Driver:** pgx/v5
- **Code Generation:** sqlc
- **Logging:** github.com/op/go-logging

## Project Structure

```
portfolio-backend/
├── cmd/api/              # Application entry point
├── config/               # Configuration management
├── internal/
│   ├── db/              # SQLC-generated database code
│   ├── domain/          # Core business entities and interfaces
│   ├── mapper/          # Database to domain model conversion
│   ├── repository/      # Data access implementation
│   ├── delivery/        # HTTP handlers, DTOs, routes
│   └── usecase/         # Business logic
├── sql/
│   ├── schema/          # PostgreSQL schema definitions
│   └── queries/         # SQL queries for sqlc
├── pkg/                 # Shared packages (middleware, utilities)
└── test/                # Test files
```

## Prerequisites

- Go 1.25.0 or higher
- PostgreSQL instance
- sqlc (for code generation)

## Quick Start

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd portfolio-backend
   ```

2. Set up environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your database credentials
   ```

3. Initialize the database schema:

   ```bash
   export INIT_SCHEMA=true
   ```

4. Generate database code:

   ```bash
   sqlc generate
   ```

5. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

## Configuration

The application can be configured via environment variables:

| Variable     | Description                           | Default                                                              |
| ------------ | ------------------------------------- | -------------------------------------------------------------------- |
| PORT         | Server port                           | 8082                                                                 |
| DATABASE_URL | PostgreSQL connection string          | postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable |
| ENV          | Environment                           | development                                                          |
| INIT_SCHEMA  | Initialize database schema on startup | false                                                                |

## API Endpoints

All endpoints are prefixed with `/api/v1`.

### Profiles

- `POST /api/v1/profile` - Create a new profile
- `GET /api/v1/profile/:id` - Get a profile by ID
- `PUT /api/v1/profile/:id` - Update a profile
- `DELETE /api/v1/profile/:id` - Delete a profile

### Education

- `POST /api/v1/education` - Create education entry
- `GET /api/v1/education/:profile_id` - Get education by profile
- `PUT /api/v1/education/:education_id` - Update education
- `DELETE /api/v1/education/:education_id` - Delete education

### Experience

- `POST /api/v1/experience` - Create experience entry
- `GET /api/v1/experience/:profile_id` - Get experiences by profile
- `PUT /api/v1/experience/:experience_id` - Update experience
- `DELETE /api/v1/experience/:experience_id` - Delete experience

### Languages

- `POST /api/v1/language` - Create language entry
- `GET /api/v1/language/:profile_id` - Get languages by profile
- `PUT /api/v1/language/:language_id` - Update language
- `DELETE /api/v1/language/:language_id` - Delete language

### Profile URLs

- `POST /api/v1/profile-url` - Create profile URL
- `GET /api/v1/profile-url/:profile_url_id` - Get profile URL
- `PUT /api/v1/profile-url/:profile_url_id` - Update profile URL
- `DELETE /api/v1/profile-url/:profile_url_id` - Delete profile URL

### Skills

- `POST /api/v1/skill` - Create skill entry
- `GET /api/v1/skill/:profile_id` - Get skills by profile
- `PUT /api/v1/skill/:skill_id` - Update skill
- `DELETE /api/v1/skill/:skill_id` - Delete skill

### Projects

- `POST /api/v1/project` - Create project entry
- `GET /api/v1/project/:profile_id` - Get projects by profile
- `PUT /api/v1/project/:project_id` - Update project
- `DELETE /api/v1/project/:project_id` - Delete project

### Resumes

- `GET /api/v1/resume/:profile_id` - Get full resume by profile ID

## Building

Build the application:

```bash
go build -o portfolio-api cmd/api/main.go
```

## Testing

Run tests:

```bash
go test ./...
```

## License

MIT

