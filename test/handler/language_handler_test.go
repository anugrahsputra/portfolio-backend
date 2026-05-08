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

type MockLanguageUsecase struct {
	mock.Mock
}

func (m *MockLanguageUsecase) CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error) {
	args := m.Called(ctx, l)
	return args.Get(0).(domain.Language), args.Error(1)
}

func (m *MockLanguageUsecase) GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Language), args.Error(1)
}

func (m *MockLanguageUsecase) UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error {
	args := m.Called(ctx, id, l)
	return args.Error(0)
}

func (m *MockLanguageUsecase) DeleteLanguage(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestLanguageHandler_CreateLanguage(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		app := fiber.New()
		app.Post("/languages", handlerObj.CreateLanguage)

		input := dto.LanguageReq{
			ProfileID:   "1",
			Language:    "English",
			Proficiency: "Native",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		expectedLanguage := domain.Language{ID: "1", ProfileID: "1", Language: "English"}
		mockUsecase.On("CreateLanguage", mock.Anything, mock.Anything).Return(expectedLanguage, nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		app := fiber.New()
		app.Post("/languages", handlerObj.CreateLanguage)

		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		app := fiber.New()
		app.Post("/languages", handlerObj.CreateLanguage)

		input := dto.LanguageReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		mockUsecase.On("CreateLanguage", mock.Anything, mock.Anything).Return(domain.Language{}, errors.New("internal error"))

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}

func TestLanguageHandler_GetLanguages(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profiles/:profile_id/languages", handlerObj.GetLanguages)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/languages", nil)

		expectedLanguages := []domain.Language{
			{ID: "1", ProfileID: "1", Language: "English"},
		}
		mockUsecase.On("GetLanguages", mock.Anything, "1").Return(expectedLanguages, nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profiles/:profile_id/languages", handlerObj.GetLanguages)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/languages", nil)

		mockUsecase.On("GetLanguages", mock.Anything, "1").Return(nil, errors.New("not found"))

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	})
}
