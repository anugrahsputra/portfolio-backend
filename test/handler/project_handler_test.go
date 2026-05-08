package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/delivery/dto"
	"github.com/anugrahsputra/portfolio-backend/internal/delivery/handler"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProjectUsecase struct {
	mock.Mock
}

func (m *MockProjectUsecase) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	args := m.Called(ctx, pr)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectUsecase) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	args := m.Called(ctx, profileID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Project), args.Error(1)
}

func (m *MockProjectUsecase) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	args := m.Called(ctx, id, pr)
	return args.Get(0).(domain.Project), args.Error(1)
}

func (m *MockProjectUsecase) DeleteProject(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestProjectHandler(t *testing.T) {
	mockUsecase := new(MockProjectUsecase)
	h := handler.NewProjectHandler(mockUsecase)

	t.Run("CreateProject - Success", func(t *testing.T) {
		app := fiber.New()
		app.Post("/api/v1/projects", h.CreateProject)

		input := dto.ProjectReq{Title: "New Project", ProfileID: "1", Description: []string{"Desc"}}
		body, _ := json.Marshal(input)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		expected := domain.Project{ID: "1", Title: "New Project", ProfileID: "1", Description: []string{"Desc"}}
		mockUsecase.On("CreateProject", mock.Anything, mock.Anything).Return(expected, nil).Once()

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("GetProjects - Success", func(t *testing.T) {
		app := fiber.New()
		app.Get("/api/v1/profiles/:profile_id/projects", h.GetProjects)

		req, _ := http.NewRequest(http.MethodGet, "/api/v1/profiles/1/projects", nil)

		expected := []domain.Project{{ID: "1", Title: "P1", ProfileID: "1"}}
		mockUsecase.On("GetProjects", mock.Anything, "1").Return(expected, nil).Once()

		resp, err := app.Test(req)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockUsecase.AssertExpectations(t)
	})
}
