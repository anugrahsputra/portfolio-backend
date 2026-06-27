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

type MockExperienceRepository struct {
	mock.Mock
}

func (m *MockExperienceRepository) CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error) {
	args := m.Called(ctx, ex)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) GetExperienceByID(ctx context.Context, id string) (domain.Experience, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	args := m.Called(ctx, id, ex)
	return args.Get(0).(domain.Experience), args.Error(1)
}

func (m *MockExperienceRepository) DeleteExperience(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestExperienceHandler_CreateExperience(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockExperienceRepository)
		handlerObj := handler.NewExperienceHandler(mockRepo)
		r := gin.New()
		r.POST("/experiences", handlerObj.CreateExperience)

		input := dto.ExperienceReq{
			ProfileID:    "1",
			Company:      "Company",
			Position:     "Software Engineer",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		expectedExperience := domain.Experience{ID: "1", ProfileID: "1", Company: "Company"}
		mockRepo.On("CreateExperience", mock.Anything, mock.Anything).Return(expectedExperience, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockRepo := new(MockExperienceRepository)
		handlerObj := handler.NewExperienceHandler(mockRepo)
		r := gin.New()
		r.POST("/experiences", handlerObj.CreateExperience)

		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockExperienceRepository)
		handlerObj := handler.NewExperienceHandler(mockRepo)
		r := gin.New()
		r.POST("/experiences", handlerObj.CreateExperience)

		input := dto.ExperienceReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockRepo.On("CreateExperience", mock.Anything, mock.Anything).Return(domain.Experience{}, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestExperienceHandler_GetExperiences(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockExperienceRepository)
		handlerObj := handler.NewExperienceHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/experiences", handlerObj.GetExperiences)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/experiences", nil)
		w := httptest.NewRecorder()

		expectedExperiences := []domain.Experience{
			{ID: "1", ProfileID: "1", Company: "Company"},
		}
		mockRepo.On("GetExperiences", mock.Anything, "1").Return(expectedExperiences, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockExperienceRepository)
		handlerObj := handler.NewExperienceHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/experiences", handlerObj.GetExperiences)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/experiences", nil)
		w := httptest.NewRecorder()

		mockRepo.On("GetExperiences", mock.Anything, "1").Return(nil, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestExperienceHandler_UpdateExperience(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockExperienceRepository)
		handlerObj := handler.NewExperienceHandler(mockRepo)
		r := gin.New()
		r.PUT("/experiences/:experience_id", handlerObj.UpdateExperience)

		company := "Updated Company"
		input := dto.ExperienceUpdateReq{Company: &company}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPut, "/experiences/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		expectedExperience := domain.Experience{ID: "1", ProfileID: "1", Company: "Updated Company"}
		mockRepo.On("UpdateExperience", mock.Anything, "1", mock.Anything).Return(expectedExperience, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestExperienceHandler_DeleteExperience(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockExperienceRepository)
		handlerObj := handler.NewExperienceHandler(mockRepo)
		r := gin.New()
		r.DELETE("/experiences/:experience_id", handlerObj.DeleteExperience)

		req, _ := http.NewRequest(http.MethodDelete, "/experiences/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("DeleteExperience", mock.Anything, "1").Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}
