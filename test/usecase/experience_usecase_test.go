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

// MockExperienceRepository is a mock implementation of domain.ExperienceRepository
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

func (m *MockExperienceRepository) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	args := m.Called(ctx, id, ex)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) DeleteExperience(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateExperience(t *testing.T) {
	mockRepo := new(MockExperienceRepository)
	uc := usecase.NewExperienceUsecase(mockRepo)

	ctx := context.Background()
	input := domain.ExperienceInput{
		ProfileID: "profile-1",
		Company:   "Test Company",
	}
	expectedExperience := domain.Experience{
		ID:        "1",
		ProfileID: "profile-1",
		Company:   "Test Company",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateExperience", ctx, input).Return(expectedExperience, nil).Once()

		result, err := uc.CreateExperience(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, expectedExperience, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("CreateExperience", ctx, input).Return(domain.Experience{}, errors.New("db error")).Once()

		result, err := uc.CreateExperience(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, domain.Experience{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetExperiences(t *testing.T) {
	mockRepo := new(MockExperienceRepository)
	uc := usecase.NewExperienceUsecase(mockRepo)

	ctx := context.Background()
	profileID := "profile-1"
	expectedExperiences := []domain.Experience{
		{
			ID:        "1",
			ProfileID: profileID,
			Company:   "Test Company",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetExperiences", ctx, profileID).Return(expectedExperiences, nil).Once()

		result, err := uc.GetExperiences(ctx, profileID)

		assert.NoError(t, err)
		assert.Equal(t, expectedExperiences, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetExperiences", ctx, profileID).Return(nil, errors.New("db error")).Once()

		result, err := uc.GetExperiences(ctx, profileID)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateExperience(t *testing.T) {
	mockRepo := new(MockExperienceRepository)
	uc := usecase.NewExperienceUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	company := "Updated Company"
	input := domain.ExperienceUpdateInput{
		Company: &company,
	}
	expectedExperience := domain.Experience{
		ID:      id,
		Company: company,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateExperience", ctx, id, input).Return(expectedExperience, nil).Once()

		result, err := uc.UpdateExperience(ctx, id, input)

		assert.NoError(t, err)
		assert.Equal(t, expectedExperience, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("UpdateExperience", ctx, id, input).Return(domain.Experience{}, errors.New("db error")).Once()

		result, err := uc.UpdateExperience(ctx, id, input)

		assert.Error(t, err)
		assert.Equal(t, domain.Experience{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteExperience(t *testing.T) {
	mockRepo := new(MockExperienceRepository)
	uc := usecase.NewExperienceUsecase(mockRepo)

	ctx := context.Background()
	id := "1"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("DeleteExperience", ctx, id).Return(nil).Once()

		err := uc.DeleteExperience(ctx, id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("DeleteExperience", ctx, id).Return(errors.New("db error")).Once()

		err := uc.DeleteExperience(ctx, id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
