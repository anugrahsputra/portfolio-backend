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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockLanguageRepository struct {
	mock.Mock
}

func (m *MockLanguageRepository) CreateLanguage(ctx context.Context, l domain.LanguageInput) (domain.Language, error) {
	args := m.Called(ctx, l)
	return args.Get(0).(domain.Language), args.Error(1)
}

func (m *MockLanguageRepository) GetLanguages(ctx context.Context, profileID string) ([]domain.Language, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Language), args.Error(1)
}

func (m *MockLanguageRepository) UpdateLanguage(ctx context.Context, id string, l domain.LanguageUpdateInput) error {
	args := m.Called(ctx, id, l)
	return args.Error(0)
}

func (m *MockLanguageRepository) DeleteLanguage(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestLanguageHandler_CreateLanguage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockLanguageRepository)
		handlerObj := handler.NewLanguageHandler(mockRepo)
		r := gin.New()
		r.POST("/languages", handlerObj.CreateLanguage)

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
		mockRepo.On("CreateLanguage", mock.Anything, mock.Anything).Return(expectedLanguage, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockRepo := new(MockLanguageRepository)
		handlerObj := handler.NewLanguageHandler(mockRepo)
		r := gin.New()
		r.POST("/languages", handlerObj.CreateLanguage)

		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockLanguageRepository)
		handlerObj := handler.NewLanguageHandler(mockRepo)
		r := gin.New()
		r.POST("/languages", handlerObj.CreateLanguage)

		input := dto.LanguageReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/languages", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockRepo.On("CreateLanguage", mock.Anything, mock.Anything).Return(domain.Language{}, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestLanguageHandler_GetLanguages(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockLanguageRepository)
		handlerObj := handler.NewLanguageHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/languages", handlerObj.GetLanguages)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/languages", nil)
		w := httptest.NewRecorder()

		expectedLanguages := []domain.Language{
			{ID: "1", ProfileID: "1", Language: "English"},
		}
		mockRepo.On("GetLanguages", mock.Anything, "1").Return(expectedLanguages, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockLanguageRepository)
		handlerObj := handler.NewLanguageHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/languages", handlerObj.GetLanguages)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/languages", nil)
		w := httptest.NewRecorder()

		mockRepo.On("GetLanguages", mock.Anything, "1").Return(nil, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestLanguageHandler_UpdateLanguage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockLanguageRepository)
		handlerObj := handler.NewLanguageHandler(mockRepo)
		r := gin.New()
		r.PUT("/languages/:language_id", handlerObj.UpdateLanguage)

		language := "Updated Language"
		input := dto.LanguageUpdateReq{Language: language}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPut, "/languages/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockRepo.On("UpdateLanguage", mock.Anything, "1", mock.Anything).Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestLanguageHandler_DeleteLanguage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockLanguageRepository)
		handlerObj := handler.NewLanguageHandler(mockRepo)
		r := gin.New()
		r.DELETE("/languages/:language_id", handlerObj.DeleteLanguage)

		req, _ := http.NewRequest(http.MethodDelete, "/languages/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("DeleteLanguage", mock.Anything, "1").Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}
