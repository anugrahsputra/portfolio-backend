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
