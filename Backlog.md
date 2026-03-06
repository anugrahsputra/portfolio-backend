# Portfolio Backend Backlog

## Done
- **Project Infrastructure**
    - [x] Initial Clean/Hexagonal Architecture scaffolding.
    - [x] Environment-based configuration management (config/config.go).
    - [x] Database connection pooling and initialization logic (config/database.go).
    - [x] Centralized logging setup (pkg/logger/logger.go).
- **Database & Data Access**
    - [x] PostgreSQL Schema definition (sql/schema/schema.sql).
    - [x] SQL queries for all entities (sql/queries/*.sql).
    - [x] SQLC code generation setup (internal/db/).
- **Domain & Core Logic**
    - [x] Domain entity definitions for Profile, Education, Experience, Language, and Skill.
    - [x] Repository interface definitions in the domain layer.
- **Core Repository Implementations**
    - [x] Profile Repository and Mapper (internal/repository/profile_repository.go).
    - [x] Profile URL Repository and Mapper.
    - [x] Education Repository and Mapper.
    - [x] Experience Repository and Mapper.

## On Progress
- **Repository Implementation (Missing Entities)**
    - [ ] Skill Repository implementation.
    - [ ] Language Repository implementation.
- **Mapper Layer (Missing Entities)**
    - [ ] Skill Mapper (internal/mapper/).
    - [ ] Language Mapper (internal/mapper/).

## Todo
- **API & Routing**
    - [ ] Select and setup HTTP router (Chi, Gin, or standard library).
    - [ ] Implement API handlers/controllers for all entities.
    - [ ] Refine request/response structures and error handling.
- **Testing & Quality**
    - [ ] Unit tests for repositories and mappers.
    - [ ] Integration tests for API endpoints.
    - [ ] Input validation for all incoming requests.
- **DevOps & Documentation**
    - [ ] Database migration tool integration (e.g., golang-migrate).
    - [ ] Dockerize the application.
    - [ ] Generate API documentation (Swagger/OpenAPI).
- **Advanced Features**
    - [ ] Add support for image/asset uploads for profile/experience.
