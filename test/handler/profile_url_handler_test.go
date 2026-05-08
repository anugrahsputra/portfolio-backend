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

type MockProfileUrlUsecase struct {
	mock.Mock
}

func (m *MockProfileUrlUsecase) CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error) {
	args := m.Called(ctx, pu)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlUsecase) GetProfileUrl(ctx context.Context, profileID string) ([]domain.ProfileUrl, error) {
	args := m.Called(ctx, profileID)
	return args.Get(0).([]domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlUsecase) GetProfileUrlByID(ctx context.Context, id string) (domain.ProfileUrl, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlUsecase) UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error {
	args := m.Called(ctx, id, pu)
	return args.Error(0)
}

func (m *MockProfileUrlUsecase) DeleteProfileUrl(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProfileUrlHandler_CreateProfileUrl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		app := fiber.New()
		app.Post("/profile-urls", handlerObj.CreateProfileUrl)

		input := dto.ProfileUrlReq{
			ProfileID: "1",
			Label:     "LinkedIn",
			Url:       "https://linkedin.com/in/johndoe",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		expectedProfileUrl := &domain.ProfileUrl{ID: "1", ProfileID: "1", Label: "LinkedIn"}
		mockUsecase.On("CreateProfileUrl", mock.Anything, mock.Anything).Return(expectedProfileUrl, nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		app := fiber.New()
		app.Post("/profile-urls", handlerObj.CreateProfileUrl)

		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		app := fiber.New()
		app.Post("/profile-urls", handlerObj.CreateProfileUrl)

		input := dto.ProfileUrlReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		mockUsecase.On("CreateProfileUrl", mock.Anything, mock.Anything).Return(nil, errors.New("internal error"))

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}

func TestProfileUrlHandler_GetProfileUrl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profile-urls/:profile_url_id", handlerObj.GetProfileUrlByID)

		req, _ := http.NewRequest(http.MethodGet, "/profile-urls/1", nil)

		expectedProfileUrl := domain.ProfileUrl{ID: "1", ProfileID: "1", Label: "LinkedIn"}
		mockUsecase.On("GetProfileUrlByID", mock.Anything, "1").Return(expectedProfileUrl, nil)

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		app := fiber.New()
		app.Get("/profile-urls/:profile_url_id", handlerObj.GetProfileUrlByID)

		req, _ := http.NewRequest(http.MethodGet, "/profile-urls/1", nil)

		mockUsecase.On("GetProfileUrlByID", mock.Anything, "1").Return(domain.ProfileUrl{}, errors.New("not found"))

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
	})
}
