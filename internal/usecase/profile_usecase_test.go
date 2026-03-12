package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockProfileRepository is a mock implementation of domain.ProfileRepository
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

func (m *MockProfileRepository) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {
	args := m.Called(ctx, id, p)
	return args.Error(0)
}

func (m *MockProfileRepository) DeleteProfile(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateProfile(t *testing.T) {
	mockRepo := new(MockProfileRepository)
	uc := NewProfileUsecase(mockRepo)

	ctx := context.Background()
	input := domain.ProfileInput{
		Name: "John Doe",
	}
	expectedProfile := &domain.Profile{
		ID:   "1",
		Name: "John Doe",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateProfile", ctx, input).Return(expectedProfile, nil).Once()

		result, err := uc.CreateProfile(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, expectedProfile, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("CreateProfile", ctx, input).Return(nil, errors.New("db error")).Once()

		result, err := uc.CreateProfile(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetProfile(t *testing.T) {
	mockRepo := new(MockProfileRepository)
	uc := NewProfileUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	expectedProfile := &domain.Profile{
		ID:   id,
		Name: "John Doe",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetProfile", ctx, id).Return(expectedProfile, nil).Once()

		result, err := uc.GetProfile(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, expectedProfile, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetProfile", ctx, id).Return(nil, errors.New("not found")).Once()

		result, err := uc.GetProfile(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateProfile(t *testing.T) {
	mockRepo := new(MockProfileRepository)
	uc := NewProfileUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	name := "Updated Name"
	input := domain.ProfileUpdateInput{
		Name: &name,
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateProfile", ctx, id, input).Return(nil).Once()

		err := uc.UpdateProfile(ctx, id, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("UpdateProfile", ctx, id, input).Return(errors.New("db error")).Once()

		err := uc.UpdateProfile(ctx, id, input)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteProfile(t *testing.T) {
	mockRepo := new(MockProfileRepository)
	uc := NewProfileUsecase(mockRepo)

	ctx := context.Background()
	id := "1"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("DeleteProfile", ctx, id).Return(nil).Once()

		err := uc.DeleteProfile(ctx, id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("DeleteProfile", ctx, id).Return(errors.New("db error")).Once()

		err := uc.DeleteProfile(ctx, id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
