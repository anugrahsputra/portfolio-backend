package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
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
