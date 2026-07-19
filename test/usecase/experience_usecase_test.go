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
