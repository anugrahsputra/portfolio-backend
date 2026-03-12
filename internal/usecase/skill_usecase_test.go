package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSkillRepository is a mock implementation of domain.SkillRepository
type MockSkillRepository struct {
	mock.Mock
}

func (m *MockSkillRepository) CreateSkill(ctx context.Context, s domain.SkillInput) (domain.Skill, error) {
	args := m.Called(ctx, s)
	return args.Get(0).(domain.Skill), args.Error(1)
}

func (m *MockSkillRepository) GetSkills(ctx context.Context, profileID string) (domain.Skill, error) {
	args := m.Called(ctx, profileID)
	return args.Get(0).(domain.Skill), args.Error(1)
}

func (m *MockSkillRepository) UpdateSkill(ctx context.Context, id string, s domain.SkillUpdateInput) error {
	args := m.Called(ctx, id, s)
	return args.Error(0)
}

func (m *MockSkillRepository) DeleteSkill(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestCreateSkill(t *testing.T) {
	mockRepo := new(MockSkillRepository)
	uc := NewSkillUsecase(mockRepo)

	ctx := context.Background()
	input := domain.SkillInput{
		ProfileID: "profile-1",
		Tools:     []string{"Go", "Docker"},
	}
	expectedSkill := domain.Skill{
		ID:        "1",
		ProfileID: "profile-1",
		Tools:     []string{"Go", "Docker"},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("CreateSkill", ctx, input).Return(expectedSkill, nil).Once()

		result, err := uc.CreateSkill(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, expectedSkill, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("CreateSkill", ctx, input).Return(domain.Skill{}, errors.New("db error")).Once()

		result, err := uc.CreateSkill(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, domain.Skill{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSkills(t *testing.T) {
	mockRepo := new(MockSkillRepository)
	uc := NewSkillUsecase(mockRepo)

	ctx := context.Background()
	profileID := "profile-1"
	expectedSkill := domain.Skill{
		ID:        "1",
		ProfileID: profileID,
		Tools:     []string{"Go", "Docker"},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetSkills", ctx, profileID).Return(expectedSkill, nil).Once()

		result, err := uc.GetSkills(ctx, profileID)

		assert.NoError(t, err)
		assert.Equal(t, expectedSkill, result)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("GetSkills", ctx, profileID).Return(domain.Skill{}, errors.New("db error")).Once()

		result, err := uc.GetSkills(ctx, profileID)

		assert.Error(t, err)
		assert.Equal(t, domain.Skill{}, result)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateSkill(t *testing.T) {
	mockRepo := new(MockSkillRepository)
	uc := NewSkillUsecase(mockRepo)

	ctx := context.Background()
	id := "1"
	input := domain.SkillUpdateInput{
		Tools: []string{"Go", "Kubernetes"},
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("UpdateSkill", ctx, id, input).Return(nil).Once()

		err := uc.UpdateSkill(ctx, id, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("UpdateSkill", ctx, id, input).Return(errors.New("db error")).Once()

		err := uc.UpdateSkill(ctx, id, input)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteSkill(t *testing.T) {
	mockRepo := new(MockSkillRepository)
	uc := NewSkillUsecase(mockRepo)

	ctx := context.Background()
	id := "1"

	t.Run("success", func(t *testing.T) {
		mockRepo.On("DeleteSkill", ctx, id).Return(nil).Once()

		err := uc.DeleteSkill(ctx, id)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo.On("DeleteSkill", ctx, id).Return(errors.New("db error")).Once()

		err := uc.DeleteSkill(ctx, id)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
