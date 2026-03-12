package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
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
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handler := NewExperienceHandler(mockUsecase)
		r := gin.Default()
		r.POST("/experiences", handler.CreateExperience)

		input := dto.ExperienceReq{
			ProfileID:    "1",
			Company:      "Company",
			Position:     "Software Engineer",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		expectedExperience := domain.Experience{ID: "1", ProfileID: "1", Company: "Company"}
		mockUsecase.On("CreateExperience", mock.Anything, mock.Anything).Return(expectedExperience, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handler := NewExperienceHandler(mockUsecase)
		r := gin.Default()
		r.POST("/experiences", handler.CreateExperience)

		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBufferString("invalid json"))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handler := NewExperienceHandler(mockUsecase)
		r := gin.Default()
		r.POST("/experiences", handler.CreateExperience)

		input := dto.ExperienceReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/experiences", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockUsecase.On("CreateExperience", mock.Anything, mock.Anything).Return(domain.Experience{}, errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestExperienceHandler_GetExperiences(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handler := NewExperienceHandler(mockUsecase)
		r := gin.Default()
		r.GET("/profiles/:profile_id/experiences", handler.GetExperiences)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/experiences", nil)
		w := httptest.NewRecorder()

		expectedExperiences := []domain.Experience{
			{ID: "1", ProfileID: "1", Company: "Company"},
		}
		mockUsecase.On("GetExperiences", mock.Anything, "1").Return(expectedExperiences, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockExperienceUsecase)
		handler := NewExperienceHandler(mockUsecase)
		r := gin.Default()
		r.GET("/profiles/:profile_id/experiences", handler.GetExperiences)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/experiences", nil)
		w := httptest.NewRecorder()

		mockUsecase.On("GetExperiences", mock.Anything, "1").Return(nil, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
