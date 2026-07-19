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

type MockEmailContactRepository struct {
	mock.Mock
}

func (m *MockEmailContactRepository) SendEmail(ctx context.Context, form domain.EmailContactFormInput, to string) error {
	args := m.Called(ctx, form, to)
	return args.Error(0)
}

func TestContactFormUsecaseSendEmail(t *testing.T) {
	form := domain.EmailContactFormInput{
		ProfileID: "1",
		Name:      "Ada",
		Email:     "ada@example.com",
		Subject:   "Hello",
		Message:   "Hi there",
	}

	t.Run("resolves the recipient from the profile then sends", func(t *testing.T) {
		profileRepo := new(MockProfileRepository)
		mailer := new(MockEmailContactRepository)
		uc := usecase.NewContactFormUsecase(mailer, profileRepo)

		profileRepo.On("GetProfile", mock.Anything, "1").
			Return(&domain.Profile{ID: "1", Email: "owner@example.com"}, nil)
		mailer.On("SendEmail", mock.Anything, form, "owner@example.com").Return(nil)

		assert.NoError(t, uc.SendEmail(context.Background(), form))
		mailer.AssertExpectations(t)
	})

	t.Run("does not send when the profile lookup fails", func(t *testing.T) {
		profileRepo := new(MockProfileRepository)
		mailer := new(MockEmailContactRepository)
		uc := usecase.NewContactFormUsecase(mailer, profileRepo)

		profileRepo.On("GetProfile", mock.Anything, "1").Return(nil, errors.New("db down"))

		assert.Error(t, uc.SendEmail(context.Background(), form))
		mailer.AssertNotCalled(t, "SendEmail", mock.Anything, mock.Anything, mock.Anything)
	})

	t.Run("does not send when the profile is missing", func(t *testing.T) {
		profileRepo := new(MockProfileRepository)
		mailer := new(MockEmailContactRepository)
		uc := usecase.NewContactFormUsecase(mailer, profileRepo)

		profileRepo.On("GetProfile", mock.Anything, "1").Return(nil, nil)

		assert.Error(t, uc.SendEmail(context.Background(), form))
		mailer.AssertNotCalled(t, "SendEmail", mock.Anything, mock.Anything, mock.Anything)
	})
}
