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
