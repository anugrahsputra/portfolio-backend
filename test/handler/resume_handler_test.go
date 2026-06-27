package handler_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockResumeRepository struct {
	mock.Mock
}

func (m *MockResumeRepository) GetResume(ctx context.Context, id string) (*domain.Resume, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Resume), args.Error(1)
}

func TestResumeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockResumeRepository)
	h := handler.NewResumeHandler(mockRepo)

	t.Run("GetResume - Success", func(t *testing.T) {
		r := gin.New()
		r.GET("/api/v1/resume/:profile_id", h.GetResume)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/resume/valid-id", nil)
		w := httptest.NewRecorder()

		expected := &domain.Resume{
			ID:    "valid-id",
			Name:  "John Doe",
			Title: "Software Engineer",
		}
		mockRepo.On("GetResume", mock.Anything, "valid-id").Return(expected, nil).Once()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetResume - Not Found", func(t *testing.T) {
		r := gin.New()
		r.GET("/api/v1/resume/:profile_id", h.GetResume)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/resume/invalid-id", nil)
		w := httptest.NewRecorder()

		mockRepo.On("GetResume", mock.Anything, "invalid-id").Return((*domain.Resume)(nil), errors.New("resume not found")).Once()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
		mockRepo.AssertExpectations(t)
	})
}
