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

// MockProfileUsecase is a mock implementation of usecase.ProfileUsecase
type MockProfileUsecase struct {
	mock.Mock
}

func (m *MockProfileUsecase) CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error) {
	args := m.Called(ctx, p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileUsecase) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileUsecase) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {
	args := m.Called(ctx, id, p)
	return args.Error(0)
}

func (m *MockProfileUsecase) DeleteProfile(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProfileHandler_CreateProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockProfileUsecase)
		handlerObj := handler.NewProfileHandler(mockUsecase)
		app := fiber.New()
		app.Post("/profiles", handlerObj.CreateProfile)

		input := dto.ProfileReq{
			Name: "John Doe",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/profiles", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		expectedProfile := &domain.Profile{ID: "1", Name: "John Doe"}
		mockUsecase.On("CreateProfile", mock.Anything, mock.Anything).Return(expectedProfile, nil)

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		mockUsecase := new(MockProfileUsecase)
		handlerObj := handler.NewProfileHandler(mockUsecase)
		app := fiber.New()
		app.Post("/profiles", handlerObj.CreateProfile)

		req, _ := http.NewRequest(http.MethodPost, "/profiles", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestProfileHandler_GetProfile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockProfileUsecase)
		handlerObj := handler.NewProfileHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profiles/:id", handlerObj.GetProfile)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1", nil)

		expectedProfile := &domain.Profile{ID: "1", Name: "John Doe"}
		mockUsecase.On("GetProfile", mock.Anything, "1").Return(expectedProfile, nil)

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockProfileUsecase)
		handlerObj := handler.NewProfileHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profiles/:id", handlerObj.GetProfile)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1", nil)

		mockUsecase.On("GetProfile", mock.Anything, "1").Return(nil, errors.New("not found"))

		resp, _ := app.Test(req)

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}


