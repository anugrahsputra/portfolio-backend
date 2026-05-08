package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockEmailContactUsecase is a mock implementation of usecase.EmailContactUsecase
type MockEmailContactUsecase struct {
	mock.Mock
}

func (m *MockEmailContactUsecase) SendEmail(ctx context.Context, form domain.EmailContactFormInput) error {
	args := m.Called(ctx, form)
	return args.Error(0)
}

func TestContactFormHandler_SendMail(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockEmailContactUsecase)
		handlerObj := handler.NewContactFormHandler(mockUsecase)
		app := fiber.New()
		app.Post("/send-email", handlerObj.SendMail)

		input := dto.ContactFormReq{
			ProfileID: "550e8400-e29b-41d4-a716-446655440000",
			Name:      "John Doe",
			Email:     "john@example.com",
			Subject:   "Hello",
			Message:   "Test Message",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/send-email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		mockUsecase.On("SendEmail", mock.Anything, mock.Anything).Return(nil).Once()

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad json", func(t *testing.T) {
		mockUsecase := new(MockEmailContactUsecase)
		handlerObj := handler.NewContactFormHandler(mockUsecase)
		app := fiber.New()
		app.Post("/send-email", handlerObj.SendMail)

		req, _ := http.NewRequest(http.MethodPost, "/send-email", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockEmailContactUsecase)
		handlerObj := handler.NewContactFormHandler(mockUsecase)
		app := fiber.New()
		app.Post("/send-email", handlerObj.SendMail)

		input := dto.ContactFormReq{
			Name: "John Doe",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/send-email", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		mockUsecase.On("SendEmail", mock.Anything, mock.Anything).Return(errors.New("validation failed")).Once()

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		
		var response dto.NoDataResponse
		err = json.NewDecoder(resp.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, response.Status)
		
		mockUsecase.AssertExpectations(t)
	})
}
