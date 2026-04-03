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

// MockProfileUrlRepository is a mock implementation of domain.ProfileUrlRepository
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
	return args.Get(0).([]domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlRepository) UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error {
	args := m.Called(ctx, id, pu)
	return args.Error(0)
}

func (m *MockProfileUrlRepository) DeleteProfileUrl(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateProfileUrl(t *testing.T) {
	mockRepo := new(MockProfileUrlRepository)
	uc := usecase.NewProfileUrlUsecase(mockRepo)

	ctx := context.Background()
	input := domain.ProfileUrlInput{
		ProfileID: "profile-1",
		Label:     "LinkedIn",
		Url:       "https://linkedin.com/in/test",
	}
	expectedProfileUrl := &domain.ProfileUrl{
		ID:        "1",
		ProfileID: "profile-1",
		Label:     "LinkedIn",
		Url:       "https://linkedin.com/in/test",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateProfileUrl", ctx, input).Return(expectedProfileUrl, nil).Once()

		result, err := uc.CreateProfileUrl(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, expectedProfileUrl, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("CreateProfileUrl", ctx, input).Return(nil, errors.New("db error")).Once()

		result, err := uc.CreateProfileUrl(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProfileUrl(t *testing.T) {
	mockRepo := new(MockProfileUrlRepository)
	uc := usecase.NewProfileUrlUsecase(mockRepo)

	ctx := context.Background()
	profileID := "profile-1"
	expectedProfileUrls := []domain.ProfileUrl{
		{ID: "1", ProfileID: profileID, Label: "LinkedIn", Url: "https://linkedin.com/in/test"},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetProfileUrl", ctx, profileID).Return(expectedProfileUrls, nil).Once()

		result, err := uc.GetProfileUrl(ctx, profileID)

		assert.NoError(t, err)
		assert.Equal(t, expectedProfileUrls, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetProfileUrl", ctx, profileID).Return([]domain.ProfileUrl(nil), errors.New("not found")).Once()

		result, err := uc.GetProfileUrl(ctx, profileID)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProfileUrlByID(t *testing.T) {
	mockRepo := new(MockProfileUrlRepository)
	uc := usecase.NewProfileUrlUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	expectedProfileUrl := domain.ProfileUrl{
		ID:        id,
		ProfileID: "profile-1",
		Label:     "LinkedIn",
		Url:       "https://linkedin.com/in/test",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetProfileUrlByID", ctx, id).Return(expectedProfileUrl, nil).Once()

		result, err := uc.GetProfileUrlByID(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, expectedProfileUrl, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetProfileUrlByID", ctx, id).Return(domain.ProfileUrl{}, errors.New("not found")).Once()

		result, err := uc.GetProfileUrlByID(ctx, id)

		assert.Error(t, err)
		assert.Equal(t, domain.ProfileUrl{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateProfileUrl(t *testing.T) {
	mockRepo := new(MockProfileUrlRepository)
	uc := usecase.NewProfileUrlUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	label := "GitHub"
	input := domain.ProfileUrlUpdateInput{
		Label: &label,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateProfileUrl", ctx, id, input).Return(nil).Once()

		err := uc.UpdateProfileUrl(ctx, id, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("UpdateProfileUrl", ctx, id, input).Return(errors.New("db error")).Once()

		err := uc.UpdateProfileUrl(ctx, id, input)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteProfileUrl(t *testing.T) {
	mockRepo := new(MockProfileUrlRepository)
	uc := usecase.NewProfileUrlUsecase(mockRepo)

	ctx := context.Background()
	id := "1"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("DeleteProfileUrl", ctx, id).Return(nil).Once()

		err := uc.DeleteProfileUrl(ctx, id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("DeleteProfileUrl", ctx, id).Return(errors.New("db error")).Once()

		err := uc.DeleteProfileUrl(ctx, id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
