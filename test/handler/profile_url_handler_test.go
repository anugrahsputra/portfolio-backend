package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/go-chi/chi/v5"
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
		r := chi.NewRouter()
		r.Post("/profile-urls", handlerObj.CreateProfileUrl)

		input := dto.ProfileUrlReq{
			ProfileID: "1",
			Label:     "LinkedIn",
			Url:       "https://linkedin.com/in/johndoe",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expectedProfileUrl := &domain.ProfileUrl{ID: "1", ProfileID: "1", Label: "LinkedIn"}
		mockUsecase.On("CreateProfileUrl", mock.Anything, mock.Anything).Return(expectedProfileUrl, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		r := chi.NewRouter()
		r.Post("/profile-urls", handlerObj.CreateProfileUrl)

		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		r := chi.NewRouter()
		r.Post("/profile-urls", handlerObj.CreateProfileUrl)

		input := dto.ProfileUrlReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUsecase.On("CreateProfileUrl", mock.Anything, mock.Anything).Return(nil, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestProfileUrlHandler_GetProfileUrl(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		r := chi.NewRouter()
		r.Get("/profile-urls/{profile_url_id}", handlerObj.GetProfileUrlByID)

		req, _ := http.NewRequest(http.MethodGet, "/profile-urls/1", nil)
		w := httptest.NewRecorder()

		expectedProfileUrl := domain.ProfileUrl{ID: "1", ProfileID: "1", Label: "LinkedIn"}
		mockUsecase.On("GetProfileUrlByID", mock.Anything, "1").Return(expectedProfileUrl, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockProfileUrlUsecase)
		handlerObj := handler.NewProfileUrlHandler(mockUsecase)
		r := chi.NewRouter()
		r.Get("/profile-urls/{profile_url_id}", handlerObj.GetProfileUrlByID)

		req, _ := http.NewRequest(http.MethodGet, "/profile-urls/1", nil)
		w := httptest.NewRecorder()

		mockUsecase.On("GetProfileUrlByID", mock.Anything, "1").Return(domain.ProfileUrl{}, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
