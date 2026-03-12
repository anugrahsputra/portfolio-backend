package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockLanguageRepository is a mock implementation of domain.LanguageRepository
type MockLanguageRepository struct {
	mock.Mock
}

func (m *MockLanguageRepository) CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error) {
	args := m.Called(ctx, l)
	return args.Get(0).(domain.Language), args.Error(1)
}

func (m *MockLanguageRepository) GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Language), args.Error(1)
}

func (m *MockLanguageRepository) UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error {
	args := m.Called(ctx, id, l)
	return args.Error(0)
}

func (m *MockLanguageRepository) DeleteLanguage(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateLanguage(t *testing.T) {
	mockRepo := new(MockLanguageRepository)
	uc := NewLanguageUsecase(mockRepo)

	ctx := context.Background()
	input := domain.LanguageInput{
		ProfileID:   "profile-1",
		Language:    "English",
		Proficiency: "Native",
	}
	expectedLanguage := domain.Language{
		ID:          "1",
		ProfileID:   "profile-1",
		Language:    "English",
		Proficiency: "Native",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateLanguage", ctx, input).Return(expectedLanguage, nil).Once()

		result, err := uc.CreateLanguage(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, expectedLanguage, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("CreateLanguage", ctx, input).Return(domain.Language{}, errors.New("db error")).Once()

		result, err := uc.CreateLanguage(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, domain.Language{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetLanguages(t *testing.T) {
	mockRepo := new(MockLanguageRepository)
	uc := NewLanguageUsecase(mockRepo)

	ctx := context.Background()
	profileID := "profile-1"
	expectedLanguages := []domain.Language{
		{
			ID:          "1",
			ProfileID:   profileID,
			Language:    "English",
			Proficiency: "Native",
		},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetLanguages", ctx, profileID).Return(expectedLanguages, nil).Once()

		result, err := uc.GetLanguages(ctx, profileID)

		assert.NoError(t, err)
		assert.Equal(t, expectedLanguages, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetLanguages", ctx, profileID).Return(nil, errors.New("db error")).Once()

		result, err := uc.GetLanguages(ctx, profileID)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateLanguage(t *testing.T) {
	mockRepo := new(MockLanguageRepository)
	uc := NewLanguageUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	input := domain.LanguageUpdateInput{
		Language:    "English",
		Proficiency: "Fluent",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateLanguage", ctx, id, input).Return(nil).Once()

		err := uc.UpdateLanguage(ctx, id, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("UpdateLanguage", ctx, id, input).Return(errors.New("db error")).Once()

		err := uc.UpdateLanguage(ctx, id, input)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteLanguage(t *testing.T) {
	mockRepo := new(MockLanguageRepository)
	uc := NewLanguageUsecase(mockRepo)

	ctx := context.Background()
	id := "1"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("DeleteLanguage", ctx, id).Return(nil).Once()

		err := uc.DeleteLanguage(ctx, id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("DeleteLanguage", ctx, id).Return(errors.New("db error")).Once()

		err := uc.DeleteLanguage(ctx, id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
