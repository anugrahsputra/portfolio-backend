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

type MockExperienceUsecase struct {
	mock.Mock
}

func (m *MockExperienceUsecase) CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error) {
	args := m.Called(ctx, ex)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceUsecase) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Experience), args.Error(1)
}

func (m *MockExperienceUsecase) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	args := m.Called(ctx, id, ex)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceUsecase) DeleteExperience(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestExperienceHandler_CreateExperience(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handlerObj := handler.NewExperienceHandler(mockUsecase)
		app := fiber.New()
		app.Post("/experiences", handlerObj.CreateExperience)

		input := dto.ExperienceReq{
			ProfileID:    "1",
			Company:      "Company",
			Position:     "Software Engineer",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		expectedExperience := domain.Experience{ID: "1", ProfileID: "1", Company: "Company"}
		mockUsecase.On("CreateExperience", mock.Anything, mock.Anything).Return(expectedExperience, nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handlerObj := handler.NewExperienceHandler(mockUsecase)
		app := fiber.New()
		app.Post("/experiences", handlerObj.CreateExperience)

		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handlerObj := handler.NewExperienceHandler(mockUsecase)
		app := fiber.New()
		app.Post("/experiences", handlerObj.CreateExperience)

		input := dto.ExperienceReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		mockUsecase.On("CreateExperience", mock.Anything, mock.Anything).Return(domain.Experience{}, errors.New("internal error"))

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestExperienceHandler_GetExperiences(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handlerObj := handler.NewExperienceHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profiles/:profile_id/experiences", handlerObj.GetExperiences)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/experiences", nil)

		expectedExperiences := []domain.Experience{
			{ID: "1", ProfileID: "1", Company: "Company"},
		}
		mockUsecase.On("GetExperiences", mock.Anything, "1").Return(expectedExperiences, nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handlerObj := handler.NewExperienceHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profiles/:profile_id/experiences", handlerObj.GetExperiences)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/experiences", nil)

		mockUsecase.On("GetExperiences", mock.Anything, "1").Return(nil, errors.New("not found"))

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
