package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
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

type MockProjectRepository struct {
	mock.Mock
}

func (m *MockProjectRepository) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	args := m.Called(ctx, pr)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectRepository) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Project), args.Error(1)
}

func (m *MockProjectRepository) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	args := m.Called(ctx, id, pr)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectRepository) DeleteProject(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProjectHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockRepo := new(MockProjectRepository)
	h := handler.NewProjectHandler(mockRepo)

	t.Run("CreateProject - Success", func(t *testing.T) {
		r := gin.New()
		r.POST("/api/v1/projects", h.CreateProject)

		input := dto.ProjectReq{Title: "New Project", ProfileID: "1", Description: []string{"Desc"}}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		expected := domain.Project{ID: "1", Title: "New Project", ProfileID: "1", Description: []string{"Desc"}}
		mockRepo.On("CreateProject", mock.Anything, mock.Anything).Return(expected, nil).Once()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("GetProjects - Success", func(t *testing.T) {
		r := gin.New()
		r.GET("/api/v1/profiles/:profile_id/projects", h.GetProjects)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/profiles/1/projects", nil)
		w := httptest.NewRecorder()

		expected := []domain.Project{{ID: "1", Title: "P1", ProfileID: "1"}}
		mockRepo.On("GetProjects", mock.Anything, "1").Return(expected, nil).Once()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("UpdateProject - Success", func(t *testing.T) {
		r := gin.New()
		r.PUT("/api/v1/projects/:project_id", h.UpdateProject)

		title := "Updated"
		input := dto.ProjectUpdateReq{Title: &title}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPut, "/api/v1/projects/1", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		expected := domain.Project{ID: "1", Title: "Updated", ProfileID: "1"}
		mockRepo.On("UpdateProject", mock.Anything, "1", mock.Anything).Return(expected, nil).Once()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("DeleteProject - Success", func(t *testing.T) {
		r := gin.New()
		r.DELETE("/api/v1/projects/:project_id", h.DeleteProject)

		req, _ := http.NewRequest(http.MethodDelete, "/api/v1/projects/1", nil)
		w := httptest.NewRecorder()

		mockRepo.On("DeleteProject", mock.Anything, "1").Return(nil).Once()

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})
}

