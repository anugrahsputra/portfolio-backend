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
		r := chi.NewRouter()
		r.Post("/languages", handlerObj.CreateLanguage)

		input := dto.LanguageReq{
			ProfileID:   "1",
			Language:    "English",
			Proficiency: "Native",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expectedLanguage := domain.Language{ID: "1", ProfileID: "1", Language: "English"}
		mockUsecase.On("CreateLanguage", mock.Anything, mock.Anything).Return(expectedLanguage, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		r := chi.NewRouter()
		r.Post("/languages", handlerObj.CreateLanguage)

		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		r := chi.NewRouter()
		r.Post("/languages", handlerObj.CreateLanguage)

		input := dto.LanguageReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockUsecase.On("CreateLanguage", mock.Anything, mock.Anything).Return(domain.Language{}, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestLanguageHandler_GetLanguages(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		r := chi.NewRouter()
		r.Get("/profiles/{profile_id}/languages", handlerObj.GetLanguages)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/languages", nil)
		w := httptest.NewRecorder()

		expectedLanguages := []domain.Language{
			{ID: "1", ProfileID: "1", Language: "English"},
		}
		mockUsecase.On("GetLanguages", mock.Anything, "1").Return(expectedLanguages, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockLanguageUsecase)
		handlerObj := handler.NewLanguageHandler(mockUsecase)
		r := chi.NewRouter()
		r.Get("/profiles/{profile_id}/languages", handlerObj.GetLanguages)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/languages", nil)
		w := httptest.NewRecorder()

		mockUsecase.On("GetLanguages", mock.Anything, "1").Return(nil, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
