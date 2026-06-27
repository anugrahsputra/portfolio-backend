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

type MockProfileRepository struct {
	mock.Mock
}

func (m *MockProfileRepository) CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error) {
	args := m.Called(ctx, p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Profile), args.Error(1)
}

func (m *MockProfileRepository) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {
	args := m.Called(ctx, id, p)
	return args.Error(0)
}

func (m *MockProfileRepository) DeleteProfile(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProfileHandler_CreateProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileRepository)
		handlerObj := handler.NewProfileHandler(mockRepo)
		r := gin.New()
		r.POST("/profiles", handlerObj.CreateProfile)

		input := dto.ProfileReq{
			Name: "John Doe",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/profiles", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		expectedProfile := &domain.Profile{ID: "1", Name: "John Doe"}
		mockRepo.On("CreateProfile", mock.Anything, mock.Anything).Return(expectedProfile, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		mockRepo := new(MockProfileRepository)
		handlerObj := handler.NewProfileHandler(mockRepo)
		r := gin.New()
		r.POST("/profiles", handlerObj.CreateProfile)

		req, _ := http.NewRequest(http.MethodPost, "/profiles", bytes.NewBufferString("invalid json"))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestProfileHandler_GetProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileRepository)
		handlerObj := handler.NewProfileHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:id", handlerObj.GetProfile)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1", nil)
		w := httptest.NewRecorder()

		expectedProfile := &domain.Profile{ID: "1", Name: "John Doe"}
		mockRepo.On("GetProfile", mock.Anything, "1").Return(expectedProfile, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockProfileRepository)
		handlerObj := handler.NewProfileHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:id", handlerObj.GetProfile)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("GetProfile", mock.Anything, "1").Return(nil, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestProfileHandler_GetProfiles(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileRepository)
		handlerObj := handler.NewProfileHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles", handlerObj.GetProfiles)

		req, _ := http.NewRequest(http.MethodGet, "/profiles", nil)
		w := httptest.NewRecorder()

		expected := []domain.Profile{{ID: "1", Name: "John Doe"}}
		mockRepo.On("GetProfiles", mock.Anything).Return(expected, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestProfileHandler_UpdateProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileRepository)
		handlerObj := handler.NewProfileHandler(mockRepo)
		r := gin.New()
		r.PUT("/profiles/:id", handlerObj.UpdateProfile)

		name := "Updated Name"
		input := dto.ProfileUpdateReq{Name: &name}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPut, "/profiles/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockRepo.On("UpdateProfile", mock.Anything, "1", mock.Anything).Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestProfileHandler_DeleteProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockProfileRepository)
		handlerObj := handler.NewProfileHandler(mockRepo)
		r := gin.New()
		r.DELETE("/profiles/:id", handlerObj.DeleteProfile)

		req, _ := http.NewRequest(http.MethodDelete, "/profiles/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("DeleteProfile", mock.Anything, "1").Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}
