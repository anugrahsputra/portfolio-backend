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

// MockEmailContactRepository is a mock implementation of domain.EmailContactRepository
type MockEmailContactRepository struct {
	mock.Mock
}

func (m *MockEmailContactRepository) SendEmail(ctx context.Context, form domain.EmailContactFormInput) error {
	args := m.Called(ctx, form)
	return args.Error(0)
}

func TestSendEmail(t *testing.T) {
	mockRepo := new(MockEmailContactRepository)
	uc := usecase.NewEmailContactUsecase(mockRepo)

	ctx := context.Background()
	input := domain.EmailContactFormInput{
		ProfileID: "550e8400-e29b-41d4-a716-446655440000",
		Name:      "John Doe",
		Email:     "john@example.com",
		Subject:   "Test Subject",
		Message:   "Test Message",
	}

	t.Run("success", func(t *testing.T) {
		mockRepo.On("SendEmail", ctx, input).Return(nil).Once()

		err := uc.SendEmail(ctx, input)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("validation error - missing profile_id", func(t *testing.T) {
		invalidInput := input
		invalidInput.ProfileID = ""

		err := uc.SendEmail(ctx, invalidInput)

		assert.Error(t, err)
		assert.Equal(t, "profile_id is required", err.Error())
	})

	t.Run("validation error - missing name", func(t *testing.T) {
		invalidInput := input
		invalidInput.Name = ""

		err := uc.SendEmail(ctx, invalidInput)

		assert.Error(t, err)
		assert.Equal(t, "name is required", err.Error())
	})

	t.Run("validation error - missing email", func(t *testing.T) {
		invalidInput := input
		invalidInput.Email = ""

		err := uc.SendEmail(ctx, invalidInput)

		assert.Error(t, err)
		assert.Equal(t, "email is required", err.Error())
	})

	t.Run("validation error - missing subject", func(t *testing.T) {
		invalidInput := input
		invalidInput.Subject = ""

		err := uc.SendEmail(ctx, invalidInput)

		assert.Error(t, err)
		assert.Equal(t, "subject is required", err.Error())
	})

	t.Run("validation error - missing message", func(t *testing.T) {
		invalidInput := input
		invalidInput.Message = ""

		err := uc.SendEmail(ctx, invalidInput)

		assert.Error(t, err)
		assert.Equal(t, "message is required", err.Error())
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo.On("SendEmail", ctx, input).Return(errors.New("smtp error")).Once()

		err := uc.SendEmail(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, "smtp error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}
