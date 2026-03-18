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

func (m *MockProjectRepository) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	args := m.Called(ctx, id, pr)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectRepository) DeleteProject(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProjectUsecase(t *testing.T) {
	mockRepo := new(MockProjectRepository)
	uc := usecase.NewProjectUsecase(mockRepo)
	ctx := context.Background()

	t.Run("CreateProject - Success", func(t *testing.T) {
		input := domain.ProjectInput{Title: "New Project"}
		expected := domain.Project{ID: "1", Title: "New Project"}

		mockRepo.On("CreateProject", ctx, input).Return(expected, nil).Once()

		result, err := uc.CreateProject(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetProjects - Success", func(t *testing.T) {
		profileID := "profile-1"
		expected := []domain.Project{{ID: "1", Title: "P1"}}

		mockRepo.On("GetProjects", ctx, profileID).Return(expected, nil).Once()

		result, err := uc.GetProjects(ctx, profileID)

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, expected, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateProject - Success", func(t *testing.T) {
		id := "1"
		title := "Updated"
		input := domain.ProjectUpdateInput{Title: &title}
		expected := domain.Project{ID: "1", Title: "Updated"}

		mockRepo.On("UpdateProject", ctx, id, input).Return(expected, nil).Once()

		result, err := uc.UpdateProject(ctx, id, input)

		assert.NoError(t, err)
		assert.Equal(t, expected, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteProject - Error", func(t *testing.T) {
		id := "1"
		mockRepo.On("DeleteProject", ctx, id).Return(errors.New("db error")).Once()

		err := uc.DeleteProject(ctx, id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
