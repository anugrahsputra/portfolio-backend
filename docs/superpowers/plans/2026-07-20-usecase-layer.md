# Usecase Layer Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Introduce `internal/usecase/` owning fetch→merge→write orchestration, reducing repositories to dumb CRUD that never sees an `UpdateInput`.

**Architecture:** Handlers depend on `domain.XUsecase` ports instead of `domain.XRepository`. The usecase interfaces inherit today's repository signatures verbatim, so handler call sites do not change. Repository interfaces change underneath: `Update` takes a fully-merged domain entity. The usecase does the read-modify-write: fetch current → `current.Merge(in)` → write. `Merge` methods already exist on the entities (added 2026-07-20) and stay where they are — merging is domain logic, orchestrating the fetch around it is usecase logic.

**Tech Stack:** Go 1.x, Gin, sqlc + pgx/v5, testify/mock, gomail.v2.

## Global Constraints

- Module path: `github.com/anugrahsputra/portfolio-backend`
- Repository interfaces live in `internal/domain/<entity>.go`. Usecase interfaces go in the same file, directly above the repository interface.
- Usecase implementations: `internal/usecase/<entity>_usecase.go`, unexported struct `xUsecase`, constructor `NewXUsecase(...) domain.XUsecase`.
- `internal/db/` is sqlc-generated. DO NOT EDIT. No SQL or schema changes in this plan.
- Every entity's `Merge` method stays in `internal/domain/`. Do not move merge logic into usecases.
- `ProfileID` is never mergeable and never reassignable.
- After every task: `go build ./... && go vet ./... && go test ./... -count=1` must pass.
- Commit after each task. Conventional Commits, subject ≤50 chars.
- **The build breaks mid-task and that is expected.** Changing a repository interface
  (step 3) orphans its handler until the handler is repointed (step 6). Do not try to
  keep `go build` green between steps within a task — steps 3 through 8 are one atomic
  change. The build must be green at each task's final verification step.

## Scope Note

Nine domains. Five have merge orchestration (project, education, experience, profile, profile_url) — these are the substance of the plan. Two overwrite wholesale with non-pointer `UpdateInput` fields (skill, language) — their usecases forward, and that is correct, not a shortcut. One is read-only (resume). One is already a usecase misfiled as a repository (contact form) — Task 9 fixes that, and it is the most valuable task here.

---

### Task 1: Project usecase (the template)

Do this task first. Tasks 2–5 follow its exact shape.

**Files:**
- Modify: `internal/domain/project.go` (add `ProjectUsecase`, change `ProjectRepository`)
- Create: `internal/usecase/project_usecase.go`
- Modify: `internal/repository/project_repository.go` (Update becomes dumb)
- Modify: `internal/delivery/handler/project_handler.go` (depend on usecase)
- Create: `test/usecase/project_usecase_test.go`
- Modify: `test/handler/project_handler_test.go` (rename mock type)

**Interfaces:**
- Consumes: `domain.Project.Merge(ProjectUpdateInput) Project` (already exists, `internal/domain/project.go`)
- Produces: `domain.ProjectUsecase`, `usecase.NewProjectUsecase(domain.ProjectRepository) domain.ProjectUsecase`

- [ ] **Step 1: Write the failing usecase test**

Create `test/usecase/project_usecase_test.go`:

```go
package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProjectRepository struct {
	mock.Mock
}

func (m *MockProjectRepository) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	args := m.Called(ctx, pr)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectRepository) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Project), args.Error(1)
}

func (m *MockProjectRepository) GetProjectByID(ctx context.Context, id string) (domain.Project, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectRepository) UpdateProject(ctx context.Context, pr domain.Project) (domain.Project, error) {
	args := m.Called(ctx, pr)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectRepository) DeleteProject(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func strPtr(s string) *string { return &s }

func TestProjectUsecaseUpdate(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	current := domain.Project{
		ID:        "p1",
		ProfileID: "owner",
		Title:     "old",
		Location:  "Jakarta",
		StartDate: start,
	}

	t.Run("fetches current, merges, writes the merged entity", func(t *testing.T) {
		repo := new(MockProjectRepository)
		uc := usecase.NewProjectUsecase(repo)

		expected := current
		expected.Title = "new"

		repo.On("GetProjectByID", mock.Anything, "p1").Return(current, nil)
		repo.On("UpdateProject", mock.Anything, expected).Return(expected, nil)

		got, err := uc.UpdateProject(context.Background(), "p1", domain.ProjectUpdateInput{
			Title: strPtr("new"),
		})

		assert.NoError(t, err)
		assert.Equal(t, "new", got.Title)
		// The repository must receive a fully merged entity, never an UpdateInput.
		repo.AssertExpectations(t)
	})

	t.Run("does not write when the fetch fails", func(t *testing.T) {
		repo := new(MockProjectRepository)
		uc := usecase.NewProjectUsecase(repo)

		repo.On("GetProjectByID", mock.Anything, "missing").
			Return(domain.Project{}, errors.New("not found"))

		_, err := uc.UpdateProject(context.Background(), "missing", domain.ProjectUpdateInput{
			Title: strPtr("new"),
		})

		assert.Error(t, err)
		repo.AssertNotCalled(t, "UpdateProject", mock.Anything, mock.Anything)
	})

	t.Run("passes through non-orchestrated calls", func(t *testing.T) {
		repo := new(MockProjectRepository)
		uc := usecase.NewProjectUsecase(repo)

		repo.On("DeleteProject", mock.Anything, "p1").Return(nil)

		assert.NoError(t, uc.DeleteProject(context.Background(), "p1"))
		repo.AssertExpectations(t)
	})
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./test/usecase/ -count=1`
Expected: FAIL — `internal/usecase` package does not exist, `undefined: usecase.NewProjectUsecase`.

- [ ] **Step 3: Split the domain interfaces**

In `internal/domain/project.go`, replace the `ProjectRepository` interface block with both interfaces:

```go
// ProjectUsecase is the port the HTTP layer depends on. It accepts partial
// update input and owns the read-modify-write cycle.
type ProjectUsecase interface {
	CreateProject(ctx context.Context, pr ProjectInput) (Project, error)
	GetProjects(ctx context.Context, profileID string) ([]Project, error)
	UpdateProject(ctx context.Context, id string, pr ProjectUpdateInput) (Project, error)
	DeleteProject(ctx context.Context, id string) error
}

// ProjectRepository is storage only. UpdateProject writes an already-merged
// entity; it never sees a ProjectUpdateInput.
type ProjectRepository interface {
	CreateProject(ctx context.Context, pr ProjectInput) (Project, error)
	GetProjects(ctx context.Context, profileID string) ([]Project, error)
	GetProjectByID(ctx context.Context, id string) (Project, error)
	UpdateProject(ctx context.Context, pr Project) (Project, error)
	DeleteProject(ctx context.Context, id string) error
}
```

Note: `GetProjectByID` is newly promoted into the interface. The implementation already exists in `internal/repository/project_repository.go`.

- [ ] **Step 4: Write the usecase**

Create `internal/usecase/project_usecase.go`:

```go
package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type projectUsecase struct {
	repo domain.ProjectRepository
}

func NewProjectUsecase(repo domain.ProjectRepository) domain.ProjectUsecase {
	return &projectUsecase{repo: repo}
}

func (u *projectUsecase) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	return u.repo.CreateProject(ctx, pr)
}

func (u *projectUsecase) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	return u.repo.GetProjects(ctx, profileID)
}

func (u *projectUsecase) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	current, err := u.repo.GetProjectByID(ctx, id)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to load project for update: %w", err)
	}

	return u.repo.UpdateProject(ctx, current.Merge(pr))
}

func (u *projectUsecase) DeleteProject(ctx context.Context, id string) error {
	return u.repo.DeleteProject(ctx, id)
}
```

- [ ] **Step 5: Make the repository dumb**

In `internal/repository/project_repository.go`, replace the whole `UpdateProject` method with:

```go
func (r *projectRepository) UpdateProject(ctx context.Context, pr domain.Project) (domain.Project, error) {
	idStr, err := uuid.Parse(pr.ID)
	if err != nil {
		return domain.Project{}, fmt.Errorf("invalid project id: %w", err)
	}

	var ed pgtype.Date
	if pr.EndDate != nil {
		ed = pgtype.Date{Time: *pr.EndDate, Valid: !pr.EndDate.IsZero()}
	}

	param := db.UpdateProjectParams{
		ID:            idStr,
		Title:         pr.Title,
		Description:   pr.Description,
		TechStacks:    pr.TechStacks,
		LiveDemoUrl:   pr.LiveDemoUrl,
		GithubRepoUrl: pr.GithubRepoUrl,
		IsLive:        pr.IsLive,
		IsNda:         pr.IsNda,
		IsFeatured:    pr.IsFeatured,
		ImageUrl:      pr.ImageUrl,
		Company:       pr.Company,
		StartDate:     pgtype.Date{Time: pr.StartDate, Valid: true},
		EndDate:       ed,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}

	project, err := r.db.UpdateProject(ctx, param)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to update project: %w", err)
	}

	return r.toDomain(project), nil
}
```

The repository no longer fetches. `StartDate` keeps `Valid: true` — that is the NOT NULL bug fix from earlier, do not regress it.

- [ ] **Step 6: Point the handler at the usecase**

In `internal/delivery/handler/project_handler.go`, change the struct, constructor, and the four call sites:

```go
type ProjectHandler struct {
	uc domain.ProjectUsecase
}

func NewProjectHandler(u domain.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{uc: u}
}
```

Then replace every `h.repo.` with `h.uc.`. There are four: `CreateProject`, `GetProjects`, `UpdateProject`, `DeleteProject`. No argument lists change.

- [ ] **Step 7: Rename the handler test mock**

In `test/handler/project_handler_test.go`, rename the type `MockProjectRepository` → `MockProjectUsecase` (declaration, all method receivers, and all construction sites). Method signatures do NOT change — `domain.ProjectUsecase` has the same shape the old `domain.ProjectRepository` had.

Do not add `GetProjectByID` to this mock. The usecase port has no such method.

- [ ] **Step 8: Wire the router**

In `internal/delivery/route/setup_router.go`, add the usecase between repo and handler:

```go
	projectUC := usecase.NewProjectUsecase(projectRepo)
```

and change the handler construction:

```go
	project := handler.NewProjectHandler(projectUC)
```

Add the import `"github.com/anugrahsputra/portfolio-backend/internal/usecase"`.

- [ ] **Step 9: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS, including the new `test/usecase` package.

- [ ] **Step 10: Commit**

```bash
git add internal/domain/project.go internal/usecase/project_usecase.go \
        internal/repository/project_repository.go \
        internal/delivery/handler/project_handler.go \
        internal/delivery/route/setup_router.go \
        test/usecase/project_usecase_test.go test/handler/project_handler_test.go
git commit -m "feat(usecase): add project usecase, thin the repo"
```

---

### Task 2: Education usecase

`UpdateEducation` returns only `error`, so the usecase discards the written entity.

**Files:**
- Modify: `internal/domain/education.go`
- Create: `internal/usecase/education_usecase.go`
- Modify: `internal/repository/education_repository.go`
- Modify: `internal/delivery/handler/education_handler.go`
- Create: `test/usecase/education_usecase_test.go`
- Modify: `test/handler/education_handler_test.go`

**Interfaces:**
- Consumes: `domain.Education.Merge(EducationUpdateInput) Education`
- Produces: `domain.EducationUsecase`, `usecase.NewEducationUsecase(domain.EducationRepository) domain.EducationUsecase`

- [ ] **Step 1: Write the failing usecase test**

Create `test/usecase/education_usecase_test.go`. Note `strPtr` already exists in this package from Task 1 — do not redeclare it.

```go
package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEducationRepository struct {
	mock.Mock
}

func (m *MockEducationRepository) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *MockEducationRepository) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Education), args.Error(1)
}

func (m *MockEducationRepository) GetEducationByID(ctx context.Context, id string) (domain.Education, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Education), args.Error(1)
}

func (m *MockEducationRepository) UpdateEducation(ctx context.Context, e domain.Education) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *MockEducationRepository) DeleteEducation(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestEducationUsecaseUpdate(t *testing.T) {
	start := time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)
	current := domain.Education{ID: "e1", ProfileID: "owner", School: "MIT", Gpa: 3.9, StartDate: start}

	t.Run("fetches current, merges, writes the merged entity", func(t *testing.T) {
		repo := new(MockEducationRepository)
		uc := usecase.NewEducationUsecase(repo)

		expected := current
		expected.School = "Stanford"

		repo.On("GetEducationByID", mock.Anything, "e1").Return(current, nil)
		repo.On("UpdateEducation", mock.Anything, expected).Return(nil)

		err := uc.UpdateEducation(context.Background(), "e1", domain.EducationUpdateInput{
			School: strPtr("Stanford"),
		})

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("does not write when the fetch fails", func(t *testing.T) {
		repo := new(MockEducationRepository)
		uc := usecase.NewEducationUsecase(repo)

		repo.On("GetEducationByID", mock.Anything, "missing").
			Return(domain.Education{}, errors.New("not found"))

		err := uc.UpdateEducation(context.Background(), "missing", domain.EducationUpdateInput{
			School: strPtr("Stanford"),
		})

		assert.Error(t, err)
		repo.AssertNotCalled(t, "UpdateEducation", mock.Anything, mock.Anything)
	})
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./test/usecase/ -count=1 -run TestEducationUsecase`
Expected: FAIL — `undefined: usecase.NewEducationUsecase`.

- [ ] **Step 3: Split the domain interfaces**

In `internal/domain/education.go`, replace the `EducationRepository` block:

```go
// EducationUsecase is the port the HTTP layer depends on.
type EducationUsecase interface {
	CreateEducation(ctx context.Context, e EducationInput) error
	GetEducations(ctx context.Context, profileID string) ([]Education, error)
	GetEducationByID(ctx context.Context, id string) (Education, error)
	UpdateEducation(ctx context.Context, id string, e EducationUpdateInput) error
	DeleteEducation(ctx context.Context, id string) error
}

// EducationRepository is storage only. UpdateEducation writes an already-merged entity.
type EducationRepository interface {
	CreateEducation(ctx context.Context, e EducationInput) error
	GetEducations(ctx context.Context, profileID string) ([]Education, error)
	GetEducationByID(ctx context.Context, id string) (Education, error)
	UpdateEducation(ctx context.Context, e Education) error
	DeleteEducation(ctx context.Context, id string) error
}
```

- [ ] **Step 4: Write the usecase**

Create `internal/usecase/education_usecase.go`:

```go
package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type educationUsecase struct {
	repo domain.EducationRepository
}

func NewEducationUsecase(repo domain.EducationRepository) domain.EducationUsecase {
	return &educationUsecase{repo: repo}
}

func (u *educationUsecase) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	return u.repo.CreateEducation(ctx, e)
}

func (u *educationUsecase) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	return u.repo.GetEducations(ctx, profileID)
}

func (u *educationUsecase) GetEducationByID(ctx context.Context, id string) (domain.Education, error) {
	return u.repo.GetEducationByID(ctx, id)
}

func (u *educationUsecase) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	current, err := u.repo.GetEducationByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to load education for update: %w", err)
	}

	return u.repo.UpdateEducation(ctx, current.Merge(e))
}

func (u *educationUsecase) DeleteEducation(ctx context.Context, id string) error {
	return u.repo.DeleteEducation(ctx, id)
}
```

- [ ] **Step 5: Make the repository dumb**

In `internal/repository/education_repository.go`, replace the whole `UpdateEducation` method:

```go
func (r *educationRepository) UpdateEducation(ctx context.Context, e domain.Education) error {
	idStr, err := uuid.Parse(e.ID)
	if err != nil {
		return fmt.Errorf("invalid education id: %w", err)
	}

	var gd pgtype.Date
	if e.GraduationDate != nil {
		gd = pgtype.Date{Time: *e.GraduationDate, Valid: !e.GraduationDate.IsZero()}
	}

	param := db.UpdateEducationParams{
		ID:             idStr,
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      pgtype.Date{Time: e.StartDate, Valid: true},
		GraduationDate: gd,
		IsPresent:      e.IsPresent,
	}

	if _, err := r.db.UpdateEducation(ctx, param); err != nil {
		return fmt.Errorf("failed to update education: %w", err)
	}

	return nil
}
```

- [ ] **Step 6: Point the handler at the usecase**

In `internal/delivery/handler/education_handler.go`:

```go
type EducationHandler struct {
	uc domain.EducationUsecase
}

func NewEducationHandler(u domain.EducationUsecase) *EducationHandler {
	return &EducationHandler{uc: u}
}
```

Replace every `h.repo.` with `h.uc.`.

- [ ] **Step 7: Rename the handler test mock**

In `test/handler/education_handler_test.go`, rename `MockEducationRepository` → `MockEducationUsecase`. Signatures unchanged.

- [ ] **Step 8: Wire the router**

In `internal/delivery/route/setup_router.go`:

```go
	educationUC := usecase.NewEducationUsecase(educationRepo)
```

```go
	education := handler.NewEducationHandler(educationUC)
```

- [ ] **Step 9: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

- [ ] **Step 10: Commit**

```bash
git add internal/domain/education.go internal/usecase/education_usecase.go \
        internal/repository/education_repository.go \
        internal/delivery/handler/education_handler.go \
        internal/delivery/route/setup_router.go \
        test/usecase/education_usecase_test.go test/handler/education_handler_test.go
git commit -m "feat(usecase): add education usecase, thin the repo"
```

---

### Task 3: Experience usecase

`UpdateExperience` returns `(Experience, error)` — same shape as project.

**Files:**
- Modify: `internal/domain/experience.go`
- Create: `internal/usecase/experience_usecase.go`
- Modify: `internal/repository/experience_repository.go`
- Modify: `internal/delivery/handler/experience_handler.go`
- Create: `test/usecase/experience_usecase_test.go`
- Modify: `test/handler/experience_handler_test.go`

**Interfaces:**
- Consumes: `domain.Experience.Merge(ExperienceUpdateInput) Experience`
- Produces: `domain.ExperienceUsecase`, `usecase.NewExperienceUsecase(domain.ExperienceRepository) domain.ExperienceUsecase`

- [ ] **Step 1: Write the failing usecase test**

Create `test/usecase/experience_usecase_test.go`:

```go
package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockExperienceRepository struct {
	mock.Mock
}

func (m *MockExperienceRepository) CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error) {
	args := m.Called(ctx, ex)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) GetExperienceByID(ctx context.Context, id string) (domain.Experience, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) UpdateExperience(ctx context.Context, ex domain.Experience) (domain.Experience, error) {
	args := m.Called(ctx, ex)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) DeleteExperience(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestExperienceUsecaseUpdate(t *testing.T) {
	start := time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC)
	current := domain.Experience{
		ID:        "x1",
		ProfileID: "owner",
		Company:   "Acme",
		Position:  "Dev",
		StartDate: start,
	}

	t.Run("fetches current, merges, writes the merged entity", func(t *testing.T) {
		repo := new(MockExperienceRepository)
		uc := usecase.NewExperienceUsecase(repo)

		expected := current
		expected.Position = "Staff Dev"

		repo.On("GetExperienceByID", mock.Anything, "x1").Return(current, nil)
		repo.On("UpdateExperience", mock.Anything, expected).Return(expected, nil)

		got, err := uc.UpdateExperience(context.Background(), "x1", domain.ExperienceUpdateInput{
			Position: strPtr("Staff Dev"),
		})

		assert.NoError(t, err)
		assert.Equal(t, "Staff Dev", got.Position)
		assert.Equal(t, "Acme", got.Company)
		repo.AssertExpectations(t)
	})

	t.Run("does not write when the fetch fails", func(t *testing.T) {
		repo := new(MockExperienceRepository)
		uc := usecase.NewExperienceUsecase(repo)

		repo.On("GetExperienceByID", mock.Anything, "missing").
			Return(domain.Experience{}, errors.New("not found"))

		_, err := uc.UpdateExperience(context.Background(), "missing", domain.ExperienceUpdateInput{
			Position: strPtr("Staff Dev"),
		})

		assert.Error(t, err)
		repo.AssertNotCalled(t, "UpdateExperience", mock.Anything, mock.Anything)
	})
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./test/usecase/ -count=1 -run TestExperienceUsecase`
Expected: FAIL — `undefined: usecase.NewExperienceUsecase`.

- [ ] **Step 3: Split the domain interfaces**

In `internal/domain/experience.go`, replace the `ExperienceRepository` block:

```go
// ExperienceUsecase is the port the HTTP layer depends on.
type ExperienceUsecase interface {
	CreateExperience(ctx context.Context, ex ExperienceInput) (Experience, error)
	GetExperiences(ctx context.Context, profileID string) ([]Experience, error)
	GetExperienceByID(ctx context.Context, id string) (Experience, error)
	UpdateExperience(ctx context.Context, id string, ex ExperienceUpdateInput) (Experience, error)
	DeleteExperience(ctx context.Context, id string) error
}

// ExperienceRepository is storage only. UpdateExperience writes an already-merged entity.
type ExperienceRepository interface {
	CreateExperience(ctx context.Context, ex ExperienceInput) (Experience, error)
	GetExperiences(ctx context.Context, profileID string) ([]Experience, error)
	GetExperienceByID(ctx context.Context, id string) (Experience, error)
	UpdateExperience(ctx context.Context, ex Experience) (Experience, error)
	DeleteExperience(ctx context.Context, id string) error
}
```

- [ ] **Step 4: Write the usecase**

Create `internal/usecase/experience_usecase.go`:

```go
package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type experienceUsecase struct {
	repo domain.ExperienceRepository
}

func NewExperienceUsecase(repo domain.ExperienceRepository) domain.ExperienceUsecase {
	return &experienceUsecase{repo: repo}
}

func (u *experienceUsecase) CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error) {
	return u.repo.CreateExperience(ctx, ex)
}

func (u *experienceUsecase) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	return u.repo.GetExperiences(ctx, profileID)
}

func (u *experienceUsecase) GetExperienceByID(ctx context.Context, id string) (domain.Experience, error) {
	return u.repo.GetExperienceByID(ctx, id)
}

func (u *experienceUsecase) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	current, err := u.repo.GetExperienceByID(ctx, id)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("failed to load experience for update: %w", err)
	}

	return u.repo.UpdateExperience(ctx, current.Merge(ex))
}

func (u *experienceUsecase) DeleteExperience(ctx context.Context, id string) error {
	return u.repo.DeleteExperience(ctx, id)
}
```

- [ ] **Step 5: Make the repository dumb**

In `internal/repository/experience_repository.go`, replace the whole `UpdateExperience` method:

```go
func (r *experienceRepository) UpdateExperience(ctx context.Context, ex domain.Experience) (domain.Experience, error) {
	idStr, err := uuid.Parse(ex.ID)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("invalid experience id: %w", err)
	}

	var ed pgtype.Date
	if ex.EndDate != nil {
		ed = pgtype.Date{Time: *ex.EndDate, Valid: !ex.EndDate.IsZero()}
	}

	param := db.UpdateExperienceParams{
		ID:          idStr,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		Location:    ex.Location,
		StartDate:   pgtype.Date{Time: ex.StartDate, Valid: true},
		EndDate:     ed,
		IsPresent:   ex.IsPresent,
	}

	experience, err := r.db.UpdateExperience(ctx, param)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("failed to update experience: %w", err)
	}

	return r.toDomain(experience), nil
}
```

- [ ] **Step 6: Point the handler at the usecase**

In `internal/delivery/handler/experience_handler.go`:

```go
type ExperienceHandler struct {
	uc domain.ExperienceUsecase
}

func NewExperienceHandler(u domain.ExperienceUsecase) *ExperienceHandler {
	return &ExperienceHandler{uc: u}
}
```

Replace every `h.repo.` with `h.uc.`.

- [ ] **Step 7: Rename the handler test mock**

In `test/handler/experience_handler_test.go`, rename `MockExperienceRepository` → `MockExperienceUsecase`. Signatures unchanged.

- [ ] **Step 8: Wire the router**

```go
	experienceUC := usecase.NewExperienceUsecase(experienceRepo)
```

```go
	experience := handler.NewExperienceHandler(experienceUC)
```

- [ ] **Step 9: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

- [ ] **Step 10: Commit**

```bash
git add internal/domain/experience.go internal/usecase/experience_usecase.go \
        internal/repository/experience_repository.go \
        internal/delivery/handler/experience_handler.go \
        internal/delivery/route/setup_router.go \
        test/usecase/experience_usecase_test.go test/handler/experience_handler_test.go
git commit -m "feat(usecase): add experience usecase, thin the repo"
```

---

### Task 4: Profile usecase

`GetProfile` returns `*Profile`, so the usecase must nil-check before merging.

**Files:**
- Modify: `internal/domain/profile.go`
- Create: `internal/usecase/profile_usecase.go`
- Modify: `internal/repository/profile_repository.go`
- Modify: `internal/delivery/handler/profile_handler.go`
- Create: `test/usecase/profile_usecase_test.go`
- Modify: `test/handler/profile_handler_test.go`

**Interfaces:**
- Consumes: `domain.Profile.Merge(ProfileUpdateInput) Profile`
- Produces: `domain.ProfileUsecase`, `usecase.NewProfileUsecase(domain.ProfileRepository) domain.ProfileUsecase`

- [ ] **Step 1: Write the failing usecase test**

Create `test/usecase/profile_usecase_test.go`:

```go
package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProfileRepository struct {
	mock.Mock
}

func (m *MockProfileRepository) CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error) {
	args := m.Called(ctx, p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) UpdateProfile(ctx context.Context, p domain.Profile) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *MockProfileRepository) DeleteProfile(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProfileUsecaseUpdate(t *testing.T) {
	current := &domain.Profile{ID: "1", Name: "Ada", Title: "Engineer", Email: "a@b.c"}

	t.Run("fetches current, merges, writes the merged entity", func(t *testing.T) {
		repo := new(MockProfileRepository)
		uc := usecase.NewProfileUsecase(repo)

		expected := *current
		expected.Name = "Grace"

		repo.On("GetProfile", mock.Anything, "1").Return(current, nil)
		repo.On("UpdateProfile", mock.Anything, expected).Return(nil)

		err := uc.UpdateProfile(context.Background(), "1", domain.ProfileUpdateInput{
			Name: strPtr("Grace"),
		})

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("errors when the profile is missing instead of dereferencing nil", func(t *testing.T) {
		repo := new(MockProfileRepository)
		uc := usecase.NewProfileUsecase(repo)

		repo.On("GetProfile", mock.Anything, "gone").Return(nil, nil)

		err := uc.UpdateProfile(context.Background(), "gone", domain.ProfileUpdateInput{
			Name: strPtr("Grace"),
		})

		assert.Error(t, err)
		repo.AssertNotCalled(t, "UpdateProfile", mock.Anything, mock.Anything)
	})

	t.Run("does not write when the fetch fails", func(t *testing.T) {
		repo := new(MockProfileRepository)
		uc := usecase.NewProfileUsecase(repo)

		repo.On("GetProfile", mock.Anything, "boom").Return(nil, errors.New("db down"))

		err := uc.UpdateProfile(context.Background(), "boom", domain.ProfileUpdateInput{
			Name: strPtr("Grace"),
		})

		assert.Error(t, err)
		repo.AssertNotCalled(t, "UpdateProfile", mock.Anything, mock.Anything)
	})
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./test/usecase/ -count=1 -run TestProfileUsecase`
Expected: FAIL — `undefined: usecase.NewProfileUsecase`.

- [ ] **Step 3: Split the domain interfaces**

In `internal/domain/profile.go`, replace the `ProfileRepository` block:

```go
// ProfileUsecase is the port the HTTP layer depends on.
type ProfileUsecase interface {
	CreateProfile(ctx context.Context, p ProfileInput) (*Profile, error)
	GetProfiles(ctx context.Context) ([]Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	UpdateProfile(ctx context.Context, id string, p ProfileUpdateInput) error
	DeleteProfile(ctx context.Context, id string) error
}

// ProfileRepository is storage only. UpdateProfile writes an already-merged entity.
type ProfileRepository interface {
	CreateProfile(ctx context.Context, p ProfileInput) (*Profile, error)
	GetProfiles(ctx context.Context) ([]Profile, error)
	GetProfile(ctx context.Context, id string) (*Profile, error)
	UpdateProfile(ctx context.Context, p Profile) error
	DeleteProfile(ctx context.Context, id string) error
}
```

- [ ] **Step 4: Write the usecase**

Create `internal/usecase/profile_usecase.go`:

```go
package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type profileUsecase struct {
	repo domain.ProfileRepository
}

func NewProfileUsecase(repo domain.ProfileRepository) domain.ProfileUsecase {
	return &profileUsecase{repo: repo}
}

func (u *profileUsecase) CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error) {
	return u.repo.CreateProfile(ctx, p)
}

func (u *profileUsecase) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	return u.repo.GetProfiles(ctx)
}

func (u *profileUsecase) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	return u.repo.GetProfile(ctx, id)
}

func (u *profileUsecase) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {
	current, err := u.repo.GetProfile(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to load profile for update: %w", err)
	}
	if current == nil {
		return fmt.Errorf("profile %s not found", id)
	}

	return u.repo.UpdateProfile(ctx, current.Merge(p))
}

func (u *profileUsecase) DeleteProfile(ctx context.Context, id string) error {
	return u.repo.DeleteProfile(ctx, id)
}
```

- [ ] **Step 5: Make the repository dumb**

In `internal/repository/profile_repository.go`, replace the whole `UpdateProfile` method:

```go
func (r *profileRepository) UpdateProfile(ctx context.Context, p domain.Profile) error {
	idStr, err := uuid.Parse(p.ID)
	if err != nil {
		return fmt.Errorf("invalid profile id: %w", err)
	}

	param := db.UpdateProfileParams{
		ID:      idStr,
		Name:    p.Name,
		Title:   p.Title,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
	}

	if err := r.db.UpdateProfile(ctx, param); err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}
```

The `r.toDomain(existing)` call and its error branch are no longer needed here — the usecase does the fetch.

- [ ] **Step 6: Point the handler at the usecase**

In `internal/delivery/handler/profile_handler.go`:

```go
type ProfileHandler struct {
	uc domain.ProfileUsecase
}

func NewProfileHandler(u domain.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{uc: u}
}
```

Replace every `h.repo.` with `h.uc.`.

- [ ] **Step 7: Rename the handler test mock**

In `test/handler/profile_handler_test.go`, rename `MockProfileRepository` → `MockProfileUsecase`. Signatures unchanged.

- [ ] **Step 8: Wire the router**

```go
	profileUC := usecase.NewProfileUsecase(profileRepo)
```

```go
	profile := handler.NewProfileHandler(profileUC)
```

Leave `profileRepo` in scope — Task 9's contact form usecase still needs it.

- [ ] **Step 9: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

- [ ] **Step 10: Commit**

```bash
git add internal/domain/profile.go internal/usecase/profile_usecase.go \
        internal/repository/profile_repository.go \
        internal/delivery/handler/profile_handler.go \
        internal/delivery/route/setup_router.go \
        test/usecase/profile_usecase_test.go test/handler/profile_handler_test.go
git commit -m "feat(usecase): add profile usecase, thin the repo"
```

---

### Task 5: ProfileUrl usecase

**Files:**
- Modify: `internal/domain/profile_url.go`
- Create: `internal/usecase/profile_url_usecase.go`
- Modify: `internal/repository/profile_url_repository.go`
- Modify: `internal/delivery/handler/profile_url_handler.go`
- Create: `test/usecase/profile_url_usecase_test.go`
- Modify: `test/handler/profile_url_handler_test.go`

**Interfaces:**
- Consumes: `domain.ProfileUrl.Merge(ProfileUrlUpdateInput) ProfileUrl`
- Produces: `domain.ProfileUrlUsecase`, `usecase.NewProfileUrlUsecase(domain.ProfileUrlRepository) domain.ProfileUrlUsecase`

- [ ] **Step 1: Write the failing usecase test**

Create `test/usecase/profile_url_usecase_test.go`:

```go
package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProfileUrlRepository struct {
	mock.Mock
}

func (m *MockProfileUrlRepository) CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error) {
	args := m.Called(ctx, pu)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlRepository) GetProfileUrlByID(ctx context.Context, id string) (domain.ProfileUrl, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlRepository) GetProfileUrl(ctx context.Context, profileID string) ([]domain.ProfileUrl, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlRepository) UpdateProfileUrl(ctx context.Context, pu domain.ProfileUrl) error {
	args := m.Called(ctx, pu)
	return args.Error(0)
}

func (m *MockProfileUrlRepository) DeleteProfileUrl(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProfileUrlUsecaseUpdate(t *testing.T) {
	current := domain.ProfileUrl{ID: "u1", ProfileID: "owner", Label: "GitHub", Url: "https://gh"}

	t.Run("partial input merges only what is set", func(t *testing.T) {
		repo := new(MockProfileUrlRepository)
		uc := usecase.NewProfileUrlUsecase(repo)

		expected := current
		expected.Label = "Portfolio"

		repo.On("GetProfileUrlByID", mock.Anything, "u1").Return(current, nil)
		repo.On("UpdateProfileUrl", mock.Anything, expected).Return(nil)

		err := uc.UpdateProfileUrl(context.Background(), "u1", domain.ProfileUrlUpdateInput{
			Label: strPtr("Portfolio"),
		})

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("does not write when the fetch fails", func(t *testing.T) {
		repo := new(MockProfileUrlRepository)
		uc := usecase.NewProfileUrlUsecase(repo)

		repo.On("GetProfileUrlByID", mock.Anything, "missing").
			Return(domain.ProfileUrl{}, errors.New("not found"))

		err := uc.UpdateProfileUrl(context.Background(), "missing", domain.ProfileUrlUpdateInput{
			Label: strPtr("Portfolio"),
		})

		assert.Error(t, err)
		repo.AssertNotCalled(t, "UpdateProfileUrl", mock.Anything, mock.Anything)
	})
}
```

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./test/usecase/ -count=1 -run TestProfileUrlUsecase`
Expected: FAIL — `undefined: usecase.NewProfileUrlUsecase`.

- [ ] **Step 3: Split the domain interfaces**

In `internal/domain/profile_url.go`, replace the `ProfileUrlRepository` block:

```go
// ProfileUrlUsecase is the port the HTTP layer depends on.
type ProfileUrlUsecase interface {
	CreateProfileUrl(ctx context.Context, pu ProfileUrlInput) (*ProfileUrl, error)
	GetProfileUrlByID(ctx context.Context, id string) (ProfileUrl, error)
	GetProfileUrl(ctx context.Context, profileID string) ([]ProfileUrl, error)
	UpdateProfileUrl(ctx context.Context, id string, pu ProfileUrlUpdateInput) error
	DeleteProfileUrl(ctx context.Context, id string) error
}

// ProfileUrlRepository is storage only. UpdateProfileUrl writes an already-merged entity.
type ProfileUrlRepository interface {
	CreateProfileUrl(ctx context.Context, pu ProfileUrlInput) (*ProfileUrl, error)
	GetProfileUrlByID(ctx context.Context, id string) (ProfileUrl, error)
	GetProfileUrl(ctx context.Context, profileID string) ([]ProfileUrl, error)
	UpdateProfileUrl(ctx context.Context, pu ProfileUrl) error
	DeleteProfileUrl(ctx context.Context, id string) error
}
```

- [ ] **Step 4: Write the usecase**

Create `internal/usecase/profile_url_usecase.go`:

```go
package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type profileUrlUsecase struct {
	repo domain.ProfileUrlRepository
}

func NewProfileUrlUsecase(repo domain.ProfileUrlRepository) domain.ProfileUrlUsecase {
	return &profileUrlUsecase{repo: repo}
}

func (u *profileUrlUsecase) CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error) {
	return u.repo.CreateProfileUrl(ctx, pu)
}

func (u *profileUrlUsecase) GetProfileUrlByID(ctx context.Context, id string) (domain.ProfileUrl, error) {
	return u.repo.GetProfileUrlByID(ctx, id)
}

func (u *profileUrlUsecase) GetProfileUrl(ctx context.Context, profileID string) ([]domain.ProfileUrl, error) {
	return u.repo.GetProfileUrl(ctx, profileID)
}

func (u *profileUrlUsecase) UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error {
	current, err := u.repo.GetProfileUrlByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to load profile url for update: %w", err)
	}

	return u.repo.UpdateProfileUrl(ctx, current.Merge(pu))
}

func (u *profileUrlUsecase) DeleteProfileUrl(ctx context.Context, id string) error {
	return u.repo.DeleteProfileUrl(ctx, id)
}
```

- [ ] **Step 5: Make the repository dumb**

In `internal/repository/profile_url_repository.go`, replace the whole `UpdateProfileUrl` method:

```go
func (r *profileUrlRepository) UpdateProfileUrl(ctx context.Context, pu domain.ProfileUrl) error {
	idStr, err := uuid.Parse(pu.ID)
	if err != nil {
		return fmt.Errorf("invalid profile url id: %w", err)
	}

	param := db.UpdateProfileURLParams{
		ID:    idStr,
		Label: pu.Label,
		Url:   pu.Url,
	}

	if _, err := r.db.UpdateProfileURL(ctx, param); err != nil {
		return fmt.Errorf("failed to update profile url: %w", err)
	}

	return nil
}
```

- [ ] **Step 6: Point the handler at the usecase**

In `internal/delivery/handler/profile_url_handler.go`:

```go
type ProfileUrlHandler struct {
	uc domain.ProfileUrlUsecase
}

func NewProfileUrlHandler(u domain.ProfileUrlUsecase) *ProfileUrlHandler {
	return &ProfileUrlHandler{uc: u}
}
```

Replace every `h.repo.` with `h.uc.`.

- [ ] **Step 7: Rename the handler test mock**

In `test/handler/profile_url_handler_test.go`, rename `MockProfileUrlRepository` → `MockProfileUrlUsecase`. Signatures unchanged.

- [ ] **Step 8: Wire the router**

```go
	profileUrlUC := usecase.NewProfileUrlUsecase(profileUrlRepo)
```

```go
	profileUrl := handler.NewProfileUrlHandler(profileUrlUC)
```

- [ ] **Step 9: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

- [ ] **Step 10: Commit**

```bash
git add internal/domain/profile_url.go internal/usecase/profile_url_usecase.go \
        internal/repository/profile_url_repository.go \
        internal/delivery/handler/profile_url_handler.go \
        internal/delivery/route/setup_router.go \
        test/usecase/profile_url_usecase_test.go test/handler/profile_url_handler_test.go
git commit -m "feat(usecase): add profile url usecase, thin repo"
```

---

### Task 6: Skill and Language usecases (no orchestration)

These two have no merge. Their `UpdateInput` fields are plain values, not pointers, so `Update` already receives complete data and the repository has nothing to fetch. Their usecases forward. Do not invent merge semantics for them — that would be a behavior change nobody asked for.

Repository interfaces for these two are **unchanged**.

**Files:**
- Modify: `internal/domain/skill.go`, `internal/domain/language.go`
- Create: `internal/usecase/skill_usecase.go`, `internal/usecase/language_usecase.go`
- Modify: `internal/delivery/handler/skill_handler.go`, `internal/delivery/handler/language_handler.go`
- Modify: `test/handler/skill_handler_test.go`, `test/handler/language_handler_test.go`

**Interfaces:**
- Produces: `domain.SkillUsecase`, `domain.LanguageUsecase`, `usecase.NewSkillUsecase`, `usecase.NewLanguageUsecase`

- [ ] **Step 1: Add the usecase interfaces**

In `internal/domain/skill.go`, add directly above `SkillRepository` (leave `SkillRepository` exactly as it is):

```go
// SkillUsecase is the port the HTTP layer depends on. Skill updates carry
// complete values, so there is no merge step and this forwards to storage.
type SkillUsecase interface {
	CreateSkill(ctx context.Context, s SkillInput) (Skill, error)
	GetSkills(ctx context.Context, profileID string) (Skill, error)
	UpdateSkill(ctx context.Context, id string, s SkillUpdateInput) error
	DeleteSkill(ctx context.Context, id string) error
}
```

In `internal/domain/language.go`, add directly above `LanguageRepository` (leave `LanguageRepository` as it is):

```go
// LanguageUsecase is the port the HTTP layer depends on. Language updates carry
// complete values, so there is no merge step and this forwards to storage.
type LanguageUsecase interface {
	CreateLanguage(ctx context.Context, l LanguageInput) (Language, error)
	GetLanguages(ctx context.Context, profileID string) ([]Language, error)
	UpdateLanguage(ctx context.Context, id string, l LanguageUpdateInput) error
	DeleteLanguage(ctx context.Context, id string) error
}
```

- [ ] **Step 2: Write the usecases**

Create `internal/usecase/skill_usecase.go`:

```go
package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type skillUsecase struct {
	repo domain.SkillRepository
}

func NewSkillUsecase(repo domain.SkillRepository) domain.SkillUsecase {
	return &skillUsecase{repo: repo}
}

func (u *skillUsecase) CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error) {
	return u.repo.CreateSkill(ctx, s)
}

func (u *skillUsecase) GetSkills(ctx context.Context, profileID string) (domain.Skill, error) {
	return u.repo.GetSkills(ctx, profileID)
}

func (u *skillUsecase) UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error {
	return u.repo.UpdateSkill(ctx, id, s)
}

func (u *skillUsecase) DeleteSkill(ctx context.Context, id string) error {
	return u.repo.DeleteSkill(ctx, id)
}
```

Create `internal/usecase/language_usecase.go`:

```go
package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type languageUsecase struct {
	repo domain.LanguageRepository
}

func NewLanguageUsecase(repo domain.LanguageRepository) domain.LanguageUsecase {
	return &languageUsecase{repo: repo}
}

func (u *languageUsecase) CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error) {
	return u.repo.CreateLanguage(ctx, l)
}

func (u *languageUsecase) GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error) {
	return u.repo.GetLanguages(ctx, profileID)
}

func (u *languageUsecase) UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error {
	return u.repo.UpdateLanguage(ctx, id, l)
}

func (u *languageUsecase) DeleteLanguage(ctx context.Context, id string) error {
	return u.repo.DeleteLanguage(ctx, id)
}
```

No usecase tests for these two: there is no logic to test, and testing that a method forwards its argument is the exact waste that got the previous usecase layer deleted in `644dbc0`. The handler tests already cover these paths.

- [ ] **Step 3: Point the handlers at the usecases**

In `internal/delivery/handler/skill_handler.go`:

```go
type SkillHandler struct {
	uc domain.SkillUsecase
}

func NewSkillHandler(u domain.SkillUsecase) *SkillHandler {
	return &SkillHandler{uc: u}
}
```

In `internal/delivery/handler/language_handler.go`:

```go
type LanguageHandler struct {
	uc domain.LanguageUsecase
}

func NewLanguageHandler(u domain.LanguageUsecase) *LanguageHandler {
	return &LanguageHandler{uc: u}
}
```

Replace every `h.repo.` with `h.uc.` in both files.

- [ ] **Step 4: Rename the handler test mocks**

In `test/handler/skill_handler_test.go`, rename `MockSkillRepository` → `MockSkillUsecase`.
In `test/handler/language_handler_test.go`, rename `MockLanguageRepository` → `MockLanguageUsecase`.
Signatures unchanged in both.

- [ ] **Step 5: Wire the router**

```go
	skillUC := usecase.NewSkillUsecase(skillRepo)
	languageUC := usecase.NewLanguageUsecase(languageRepo)
```

```go
	skill := handler.NewSkillHandler(skillUC)
	language := handler.NewLanguageHandler(languageUC)
```

- [ ] **Step 6: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add internal/domain/skill.go internal/domain/language.go \
        internal/usecase/skill_usecase.go internal/usecase/language_usecase.go \
        internal/delivery/handler/skill_handler.go internal/delivery/handler/language_handler.go \
        internal/delivery/route/setup_router.go \
        test/handler/skill_handler_test.go test/handler/language_handler_test.go
git commit -m "feat(usecase): add skill and language usecases"
```

---

### Task 7: Resume usecase

Read-only, single method. Forwards.

**Files:**
- Modify: `internal/domain/resume.go`
- Create: `internal/usecase/resume_usecase.go`
- Modify: `internal/delivery/handler/resume_handler.go`
- Modify: `test/handler/resume_handler_test.go`

**Interfaces:**
- Produces: `domain.ResumeUsecase`, `usecase.NewResumeUsecase(domain.ResumeRepository) domain.ResumeUsecase`

- [ ] **Step 1: Add the usecase interface**

In `internal/domain/resume.go`, add directly above `ResumeRepository` (leave `ResumeRepository` as it is):

```go
// ResumeUsecase is the port the HTTP layer depends on.
type ResumeUsecase interface {
	GetResume(ctx context.Context, id string) (*Resume, error)
}
```

- [ ] **Step 2: Write the usecase**

Create `internal/usecase/resume_usecase.go`:

```go
package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type resumeUsecase struct {
	repo domain.ResumeRepository
}

func NewResumeUsecase(repo domain.ResumeRepository) domain.ResumeUsecase {
	return &resumeUsecase{repo: repo}
}

func (u *resumeUsecase) GetResume(ctx context.Context, id string) (*domain.Resume, error) {
	return u.repo.GetResume(ctx, id)
}
```

- [ ] **Step 3: Point the handler at the usecase**

In `internal/delivery/handler/resume_handler.go`:

```go
type ResumeHandler struct {
	uc domain.ResumeUsecase
}

func NewResumeHandler(u domain.ResumeUsecase) *ResumeHandler {
	return &ResumeHandler{uc: u}
}
```

Replace every `h.repo.` with `h.uc.`.

- [ ] **Step 4: Rename the handler test mock**

In `test/handler/resume_handler_test.go`, rename `MockResumeRepository` → `MockResumeUsecase`. Signature unchanged.

- [ ] **Step 5: Wire the router**

```go
	resumeUC := usecase.NewResumeUsecase(resumeRepo)
```

```go
	resume := handler.NewResumeHandler(resumeUC)
```

- [ ] **Step 6: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

- [ ] **Step 7: Commit**

```bash
git add internal/domain/resume.go internal/usecase/resume_usecase.go \
        internal/delivery/handler/resume_handler.go \
        internal/delivery/route/setup_router.go \
        test/handler/resume_handler_test.go
git commit -m "feat(usecase): add resume usecase"
```

---

### Task 8: Contact form usecase (the real win)

`emailContactRepository` currently takes `domain.ProfileRepository` as a constructor dependency and calls `profileRepo.GetProfile` inside `SendEmail`. That is cross-repository orchestration living in an adapter — a usecase wearing a repository's name. This task separates them: the repository becomes a pure SMTP port that is told where to send, and the usecase does the profile lookup.

**Files:**
- Modify: `internal/domain/email_contact_form.go`
- Create: `internal/usecase/contact_form_usecase.go`
- Modify: `internal/repository/email_contact_repository.go`
- Modify: `internal/delivery/handler/contact_form_handler.go`
- Create: `test/usecase/contact_form_usecase_test.go`
- Modify: `test/handler/contact_form_handler_test.go`

**Interfaces:**
- Consumes: `domain.ProfileRepository.GetProfile` (from Task 4 — signature unchanged by that task)
- Produces: `domain.ContactFormUsecase`, `usecase.NewContactFormUsecase(domain.EmailContactRepository, domain.ProfileRepository) domain.ContactFormUsecase`
- Changes: `domain.EmailContactRepository.SendEmail` gains a `to string` parameter

- [ ] **Step 1: Write the failing usecase test**

Create `test/usecase/contact_form_usecase_test.go`:

```go
package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmailContactRepository struct {
	mock.Mock
}

func (m *MockEmailContactRepository) SendEmail(ctx context.Context, form domain.EmailContactFormInput, to string) error {
	args := m.Called(ctx, form, to)
	return args.Error(0)
}

func TestContactFormUsecaseSendEmail(t *testing.T) {
	form := domain.EmailContactFormInput{
		ProfileID: "1",
		Name:      "Ada",
		Email:     "ada@example.com",
		Subject:   "Hello",
		Message:   "Hi there",
	}

	t.Run("resolves the recipient from the profile then sends", func(t *testing.T) {
		profileRepo := new(MockProfileRepository)
		mailer := new(MockEmailContactRepository)
		uc := usecase.NewContactFormUsecase(mailer, profileRepo)

		profileRepo.On("GetProfile", mock.Anything, "1").
			Return(&domain.Profile{ID: "1", Email: "owner@example.com"}, nil)
		mailer.On("SendEmail", mock.Anything, form, "owner@example.com").Return(nil)

		assert.NoError(t, uc.SendEmail(context.Background(), form))
		mailer.AssertExpectations(t)
	})

	t.Run("does not send when the profile lookup fails", func(t *testing.T) {
		profileRepo := new(MockProfileRepository)
		mailer := new(MockEmailContactRepository)
		uc := usecase.NewContactFormUsecase(mailer, profileRepo)

		profileRepo.On("GetProfile", mock.Anything, "1").Return(nil, errors.New("db down"))

		assert.Error(t, uc.SendEmail(context.Background(), form))
		mailer.AssertNotCalled(t, "SendEmail", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("does not send when the profile is missing", func(t *testing.T) {
		profileRepo := new(MockProfileRepository)
		mailer := new(MockEmailContactRepository)
		uc := usecase.NewContactFormUsecase(mailer, profileRepo)

		profileRepo.On("GetProfile", mock.Anything, "1").Return(nil, nil)

		assert.Error(t, uc.SendEmail(context.Background(), form))
		mailer.AssertNotCalled(t, "SendEmail", mock.Anything, mock.Anything, mock.Anything)
	})
}
```

`MockProfileRepository` is reused from Task 4's file in the same package. Do not redeclare it.

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./test/usecase/ -count=1 -run TestContactFormUsecase`
Expected: FAIL — `undefined: usecase.NewContactFormUsecase`.

- [ ] **Step 3: Change the domain interfaces**

In `internal/domain/email_contact_form.go`, replace the `EmailContactRepository` block:

```go
// ContactFormUsecase is the port the HTTP layer depends on. It resolves the
// recipient from the profile, then hands off to the mail transport.
type ContactFormUsecase interface {
	SendEmail(ctx context.Context, form EmailContactFormInput) error
}

// EmailContactRepository is the mail transport port. It is told where to send;
// it does not look recipients up.
type EmailContactRepository interface {
	SendEmail(ctx context.Context, form EmailContactFormInput, to string) error
}
```

- [ ] **Step 4: Write the usecase**

Create `internal/usecase/contact_form_usecase.go`:

```go
package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type contactFormUsecase struct {
	mailer      domain.EmailContactRepository
	profileRepo domain.ProfileRepository
}

func NewContactFormUsecase(mailer domain.EmailContactRepository, profileRepo domain.ProfileRepository) domain.ContactFormUsecase {
	return &contactFormUsecase{mailer: mailer, profileRepo: profileRepo}
}

func (u *contactFormUsecase) SendEmail(ctx context.Context, form domain.EmailContactFormInput) error {
	profile, err := u.profileRepo.GetProfile(ctx, form.ProfileID)
	if err != nil {
		return fmt.Errorf("failed to get profile: %w", err)
	}
	if profile == nil {
		return fmt.Errorf("profile %s not found", form.ProfileID)
	}

	return u.mailer.SendEmail(ctx, form, profile.Email)
}
```

- [ ] **Step 5: Make the repository a pure transport**

In `internal/repository/email_contact_repository.go`, change the struct, constructor, and `SendEmail`. Leave `emailTemplate` and the entire HTML body untouched.

```go
type emailContactRepository struct {
	dialer *gomail.Dialer
	from   string
}

func NewEmailContactRepository(mail *config.Mail, cfg *config.Config) domain.EmailContactRepository {
	return &emailContactRepository{dialer: mail.Dialer, from: cfg.GmailUser}
}

func (r *emailContactRepository) SendEmail(ctx context.Context, form domain.EmailContactFormInput, to string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", fmt.Sprintf("Portfolio <%s>", r.from))
	m.SetHeader("To", to)
	m.SetHeader("Reply-To", form.Email)
	m.SetHeader("Subject", fmt.Sprintf("[Portfolio] %s - %s", form.Name, form.Subject))

	body := emailTemplate(form)
	m.SetBody("text/html", body)

	if err := r.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
```

Remove the now-unused `profileRepo` field. The `domain` import is still needed for the parameter types.

- [ ] **Step 6: Point the handler at the usecase**

In `internal/delivery/handler/contact_form_handler.go`:

```go
type ContactFormHandler struct {
	uc domain.ContactFormUsecase
}

func NewContactFormHandler(u domain.ContactFormUsecase) *ContactFormHandler {
	return &ContactFormHandler{uc: u}
}
```

Replace every `h.repo.` with `h.uc.`. The handler calls `SendEmail(ctx, input)` with two arguments — that matches `ContactFormUsecase`, so no call site changes.

- [ ] **Step 7: Rename the handler test mock**

In `test/handler/contact_form_handler_test.go`, rename `MockEmailContactRepository` → `MockContactFormUsecase`. Its `SendEmail` keeps the **two-argument** form (`ctx, form`) because it mocks the usecase port, not the transport port. Do not add the `to` parameter here.

- [ ] **Step 8: Wire the router**

In `internal/delivery/route/setup_router.go`, change the repository construction to drop `profileRepo`:

```go
	contactFormRepo := repository.NewEmailContactRepository(mail, cfg)
```

and add the usecase:

```go
	contactFormUC := usecase.NewContactFormUsecase(contactFormRepo, profileRepo)
```

```go
	contactForm := handler.NewContactFormHandler(contactFormUC)
```

- [ ] **Step 9: Run the full suite**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

- [ ] **Step 10: Commit**

```bash
git add internal/domain/email_contact_form.go internal/usecase/contact_form_usecase.go \
        internal/repository/email_contact_repository.go \
        internal/delivery/handler/contact_form_handler.go \
        internal/delivery/route/setup_router.go \
        test/usecase/contact_form_usecase_test.go test/handler/contact_form_handler_test.go
git commit -m "refactor(contact): split mail transport from lookup"
```

---

### Task 9: Update CLAUDE.md

`CLAUDE.md` currently says "There is no usecase layer" and documents a two-layer flow. That is now false. This is the same staleness that made the docs wrong before — do not skip it.

**Files:**
- Modify: `CLAUDE.md`

- [ ] **Step 1: Update the directory tree**

In the Architecture section, add the usecase line to the `internal/` tree, between `db/` and `repository/`:

```
├── usecase/               # Orchestration: fetch -> Merge -> write
├── repository/            # Driven adapters: storage only, never see UpdateInput
```

- [ ] **Step 2: Replace the "no usecase layer" paragraph**

Delete this paragraph:

> There is no usecase layer. These are CRUD endpoints with no orchestration to own, so handlers depend on the domain repository interfaces directly. Add one only when a request needs logic that spans more than one repository.

Replace it with:

```markdown
Handlers depend on `domain.XUsecase`. Usecases own the read-modify-write cycle:
they fetch the current entity, call its `Merge` method, and hand the merged
entity to the repository. Repositories are storage only and never see an
`UpdateInput`.

Skill, language, and resume usecases forward without orchestration — their
updates carry complete values, so there is nothing to merge. That is deliberate.
```

- [ ] **Step 3: Update the Data Flow section**

Replace the numbered list with:

```markdown
1. **HTTP Request** → `handler` binds the DTO
2. **Handler** → `dto.ToXInput()` converts to a domain type
3. **Handler** → calls the `domain.XUsecase` port
4. **Usecase** → for updates: fetch current entity, `current.Merge(in)`, write the result
5. **Repository** → maps domain ↔ `db` models and calls `db.Queries` (sqlc)
6. **Handler** → `dto.ToXDTO()` converts the result back for the response
```

- [ ] **Step 4: Update the Key Patterns partial-updates bullet**

Replace the "Partial updates" bullet with:

```markdown
- **Partial updates**: `UpdateInput` types use pointer fields, where nil means "leave
  this alone". Merging is domain logic: each entity has a `Merge(UpdateInput)` method
  built on `pkg/ptr.Or()`. Orchestration is usecase logic: the usecase fetches the
  current entity, merges, and writes. Repositories take a fully-merged entity.
  Fields that must not change (`ProfileID`) are deliberately absent from `Merge`.
```

- [ ] **Step 5: Update the Testing section**

Add to the test package list, after the `test/domain/` bullet:

```markdown
- `test/usecase/` — mocks the repository interfaces; asserts that updates fetch-then-merge and that a failed fetch never writes.
```

- [ ] **Step 6: Update the test commands**

In Build & Development Commands, add to the per-package list:

```bash
go test ./test/usecase/...
```

- [ ] **Step 7: Verify and commit**

Run: `go build ./... && go vet ./... && go test ./... -count=1`
Expected: PASS.

```bash
git add CLAUDE.md
git commit -m "docs: describe the usecase layer"
```

---

## Verification

After Task 9, confirm the layering actually holds:

```bash
# No handler should import the repository package.
grep -rn 'internal/repository' internal/delivery/handler/ && echo "LEAK: handler imports repository" || echo "OK"

# No repository should mention UpdateInput.
grep -rn 'UpdateInput' internal/repository/ && echo "LEAK: repository sees UpdateInput" || echo "OK"

# Only the router wires concrete types together.
grep -rln 'repository.New' internal/ | grep -v setup_router.go && echo "LEAK: wiring outside router" || echo "OK"
```

All three should print `OK`. The second check has one expected caveat: `grep` exits non-zero when it finds nothing, which is why each line uses `&& ... || echo OK`.

Then confirm the merge tests from `test/domain/merge_test.go` still pass unchanged — that file must not need edits, because `Merge` did not move.
