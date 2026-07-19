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
