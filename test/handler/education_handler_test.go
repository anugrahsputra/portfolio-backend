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

type MockEducationUsecase struct {
	mock.Mock
}

func (m *MockEducationUsecase) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *MockEducationUsecase) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Education), args.Error(1)
}

func (m *MockEducationUsecase) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	args := m.Called(ctx, id, e)
	return args.Error(0)
}

func (m *MockEducationUsecase) DeleteEducation(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestEducationHandler_CreateEducation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockEducationUsecase)
		handlerObj := handler.NewEducationHandler(mockUsecase)
		r := gin.Default()
		r.POST("/educations", handlerObj.CreateEducation)

		input := dto.EducationReq{
			ProfileID:    "1",
			School:       "University",
			Degree:       "Bachelor",
			FieldOfStudy: "Computer Science",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/educations", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockUsecase.On("CreateEducation", mock.Anything, mock.Anything).Return(nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request - invalid json", func(t *testing.T) {
		mockUsecase := new(MockEducationUsecase)
		handlerObj := handler.NewEducationHandler(mockUsecase)
		r := gin.Default()
		r.POST("/educations", handlerObj.CreateEducation)

		req, _ := http.NewRequest(http.MethodPost, "/educations", bytes.NewBufferString("invalid json"))
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("usecase error", func(t *testing.T) {
		mockUsecase := new(MockEducationUsecase)
		handlerObj := handler.NewEducationHandler(mockUsecase)
		r := gin.Default()
		r.POST("/educations", handlerObj.CreateEducation)

		input := dto.EducationReq{
			ProfileID: "1",
		}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/educations", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		mockUsecase.On("CreateEducation", mock.Anything, mock.Anything).Return(errors.New("internal error"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestEducationHandler_GetEducation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		mockUsecase := new(MockEducationUsecase)
		handlerObj := handler.NewEducationHandler(mockUsecase)
		r := gin.Default()
		r.GET("/profiles/:profile_id/educations", handlerObj.GetEducation)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/educations", nil)
		w := httptest.NewRecorder()

		expectedEducations := []domain.Education{
			{ID: "1", ProfileID: "1", School: "University"},
		}
		mockUsecase.On("GetEducations", mock.Anything, "1").Return(expectedEducations, nil)

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		mockUsecase := new(MockEducationUsecase)
		handlerObj := handler.NewEducationHandler(mockUsecase)
		r := gin.Default()
		r.GET("/profiles/:profile_id/educations", handlerObj.GetEducation)

		req, _ := http.NewRequest(http.MethodGet, "/profiles/1/educations", nil)
		w := httptest.NewRecorder()

		mockUsecase.On("GetEducations", mock.Anything, "1").Return(nil, errors.New("not found"))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
