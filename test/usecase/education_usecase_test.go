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

func (m *MockEducationRepository) UpdateEducation(ctx context.Context, e domain.Education) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *MockEducationRepository) DeleteEducation(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestEducationUsecaseUpdate(t *testing.T) {
	start := time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)
	current := domain.Education{ID: "e1", ProfileID: "owner", School: "MIT", Gpa: 3.9, StartDate: start}

	t.Run("fetches current, merges, writes the merged entity", func(t *testing.T) {
		repo := new(MockEducationRepository)
		uc := usecase.NewEducationUsecase(repo)

		expected := current
		expected.School = "Stanford"

		repo.On("GetEducationByID", mock.Anything, "e1").Return(current, nil)
		repo.On("UpdateEducation", mock.Anything, expected).Return(nil)

		err := uc.UpdateEducation(context.Background(), "e1", domain.EducationUpdateInput{
			School: strPtr("Stanford"),
		})

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("does not write when the fetch fails", func(t *testing.T) {
		repo := new(MockEducationRepository)
		uc := usecase.NewEducationUsecase(repo)

		repo.On("GetEducationByID", mock.Anything, "missing").
			Return(domain.Education{}, errors.New("not found"))

		err := uc.UpdateEducation(context.Background(), "missing", domain.EducationUpdateInput{
			School: strPtr("Stanford"),
		})

		assert.Error(t, err)
		repo.AssertNotCalled(t, "UpdateEducation", mock.Anything, mock.Anything)
	})
}
