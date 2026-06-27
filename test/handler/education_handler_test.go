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

type MockEducationRepository struct {
	mock.Mock
}

func (m *MockEducationRepository) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *MockEducationRepository) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Education), args.Error(1)
}

func (m *MockEducationRepository) GetEducationByID(ctx context.Context, id string) (domain.Education, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Education), args.Error(1)
}

func (m *MockEducationRepository) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	args := m.Called(ctx, id, e)
	return args.Error(0)
}

func (m *MockEducationRepository) DeleteEducation(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestEducationHandler_CreateEducation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockEducationRepository)
		handlerObj := handler.NewEducationHandler(mockRepo)
		r := gin.New()
		r.POST("/educations", handlerObj.CreateEducation)

		input := dto.EducationReq{
			ProfileID:    "1",
			School:       "University",
			Degree:       "Bachelor",
			FieldOfStudy: "Computer Science",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/educations", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockRepo.On("CreateEducation", mock.Anything, mock.Anything).Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockRepo := new(MockEducationRepository)
		handlerObj := handler.NewEducationHandler(mockRepo)
		r := gin.New()
		r.POST("/educations", handlerObj.CreateEducation)

		req, _ := http.NewRequest(http.MethodPost, "/educations", bytes.NewBufferString("invalid json"))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("repository error", func(t *testing.T) {
		mockRepo := new(MockEducationRepository)
		handlerObj := handler.NewEducationHandler(mockRepo)
		r := gin.New()
		r.POST("/educations", handlerObj.CreateEducation)

		input := dto.EducationReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/educations", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		mockRepo.On("CreateEducation", mock.Anything, mock.Anything).Return(errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestEducationHandler_GetEducation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockEducationRepository)
		handlerObj := handler.NewEducationHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/educations", handlerObj.GetEducation)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/educations", nil)
		w := httptest.NewRecorder()

		expectedEducations := []domain.Education{
			{ID: "1", ProfileID: "1", School: "University"},
		}
		mockRepo.On("GetEducations", mock.Anything, "1").Return(expectedEducations, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockRepo := new(MockEducationRepository)
		handlerObj := handler.NewEducationHandler(mockRepo)
		r := gin.New()
		r.GET("/profiles/:profile_id/educations", handlerObj.GetEducation)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/educations", nil)
		w := httptest.NewRecorder()

		mockRepo.On("GetEducations", mock.Anything, "1").Return(nil, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestEducationHandler_UpdateEducation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockEducationRepository)
		handlerObj := handler.NewEducationHandler(mockRepo)
		r := gin.New()
		r.PUT("/educations/:education_id", handlerObj.UpdateEducation)

		school := "Updated School"
		input := dto.EducationUpdateReq{School: &school}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPut, "/educations/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockRepo.On("UpdateEducation", mock.Anything, "1", mock.Anything).Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

func TestEducationHandler_DeleteEducation(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("success", func(t *testing.T) {
		mockRepo := new(MockEducationRepository)
		handlerObj := handler.NewEducationHandler(mockRepo)
		r := gin.New()
		r.DELETE("/educations/:education_id", handlerObj.DeleteEducation)

		req, _ := http.NewRequest(http.MethodDelete, "/educations/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("DeleteEducation", mock.Anything, "1").Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}
