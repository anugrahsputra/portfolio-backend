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

type MockProfileUrlRepository struct {
	mock.Mock
}

func (m *MockProfileUrlRepository) CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error) {
	args := m.Called(ctx, pu)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlRepository) GetProfileUrl(ctx context.Context, profileID string) ([]domain.ProfileUrl, error) {
	args := m.Called(ctx, profileID)
	return args.Get(0).([]domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlRepository) GetProfileUrlByID(ctx context.Context, id string) (domain.ProfileUrl, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.ProfileUrl), args.Error(1)
}

func (m *MockProfileUrlRepository) UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error {
	args := m.Called(ctx, id, pu)
	return args.Error(0)
}

func (m *MockProfileUrlRepository) DeleteProfileUrl(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProfileUrlHandler_CreateProfileUrl(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.POST("/profile-urls", handlerObj.CreateProfileUrl)

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
		mockRepo.On("CreateProfileUrl", mock.Anything, mock.Anything).Return(expectedProfileUrl, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.POST("/profile-urls", handlerObj.CreateProfileUrl)

		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.POST("/profile-urls", handlerObj.CreateProfileUrl)

		input := dto.ProfileUrlReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/profile-urls", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockRepo.On("CreateProfileUrl", mock.Anything, mock.Anything).Return(nil, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestProfileUrlHandler_GetProfileUrlByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.GET("/profile-urls/:profile_url_id", handlerObj.GetProfileUrlByID)

		req, _ := http.NewRequest(http.MethodGet, "/profile-urls/1", nil)
		w := httptest.NewRecorder()

		expectedProfileUrl := domain.ProfileUrl{ID: "1", ProfileID: "1", Label: "LinkedIn"}
		mockRepo.On("GetProfileUrlByID", mock.Anything, "1").Return(expectedProfileUrl, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.GET("/profile-urls/:profile_url_id", handlerObj.GetProfileUrlByID)

		req, _ := http.NewRequest(http.MethodGet, "/profile-urls/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("GetProfileUrlByID", mock.Anything, "1").Return(domain.ProfileUrl{}, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestProfileUrlHandler_GetProfileURL(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/urls", handlerObj.GetProfileURL)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/urls", nil)
		w := httptest.NewRecorder()

		expected := []domain.ProfileUrl{{ID: "1", ProfileID: "1", Url: "https://example.com"}}
		mockRepo.On("GetProfileUrl", mock.Anything, "1").Return(expected, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestProfileUrlHandler_UpdateProfileUrl(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.PUT("/urls/:profile_url_id", handlerObj.UpdateProfileUrl)

		url := "https://updated.com"
		input := dto.ProfileUrlReq{ProfileID: "1", Label: "Updated", Url: url}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPut, "/urls/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockRepo.On("UpdateProfileUrl", mock.Anything, "1", mock.Anything).Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestProfileUrlHandler_DeleteProfileUrl(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileUrlRepository)
		handlerObj := handler.NewProfileUrlHandler(mockRepo)
		r := gin.New()
		r.DELETE("/urls/:profile_url_id", handlerObj.DeleteProfileUrl)

		req, _ := http.NewRequest(http.MethodDelete, "/urls/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("DeleteProfileUrl", mock.Anything, "1").Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

