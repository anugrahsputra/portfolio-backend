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

// MockEducationRepository is a mock implementation of domain.EducationRepository
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

func (m *MockEducationRepository) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	args := m.Called(ctx, id, e)
	return args.Error(0)
}

func (m *MockEducationRepository) DeleteEducation(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateEducation(t *testing.T) {
	mockRepo := new(MockEducationRepository)
	uc := usecase.NewEducationUsecase(mockRepo)

	ctx := context.Background()
	input := domain.EducationInput{
		ProfileID: "profile-1",
		School:    "Test School",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateEducation", ctx, input).Return(nil).Once()

		err := uc.CreateEducation(ctx, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("CreateEducation", ctx, input).Return(errors.New("db error")).Once()

		err := uc.CreateEducation(ctx, input)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetEducations(t *testing.T) {
	mockRepo := new(MockEducationRepository)
	uc := usecase.NewEducationUsecase(mockRepo)

	ctx := context.Background()
	profileID := "profile-1"
	expectedEducations := []domain.Education{
		{
			ID:        "1",
			ProfileID: profileID,
			School:    "Test School",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetEducations", ctx, profileID).Return(expectedEducations, nil).Once()

		result, err := uc.GetEducations(ctx, profileID)

		assert.NoError(t, err)
		assert.Equal(t, expectedEducations, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetEducations", ctx, profileID).Return(nil, errors.New("db error")).Once()

		result, err := uc.GetEducations(ctx, profileID)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateEducation(t *testing.T) {
	mockRepo := new(MockEducationRepository)
	uc := usecase.NewEducationUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	school := "Updated School"
	input := domain.EducationUpdateInput{
		School: &school,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateEducation", ctx, id, input).Return(nil).Once()

		err := uc.UpdateEducation(ctx, id, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("UpdateEducation", ctx, id, input).Return(errors.New("db error")).Once()

		err := uc.UpdateEducation(ctx, id, input)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteEducation(t *testing.T) {
	mockRepo := new(MockEducationRepository)
	uc := usecase.NewEducationUsecase(mockRepo)

	ctx := context.Background()
	id := "1"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("DeleteEducation", ctx, id).Return(nil).Once()

		err := uc.DeleteEducation(ctx, id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("DeleteEducation", ctx, id).Return(errors.New("db error")).Once()

		err := uc.DeleteEducation(ctx, id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

