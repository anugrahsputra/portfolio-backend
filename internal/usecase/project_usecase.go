package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type projectUsecase struct {
	repo domain.ProjectRepository
}

type ProjectUsecase interface {
	CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error)
	GetProjects(ctx context.Context, profileID string) ([]domain.Project, error)
	UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error)
	DeleteProject(ctx context.Context, id string) error
}

func NewProjectUsecase(r domain.ProjectRepository) ProjectUsecase {
	return &projectUsecase{repo: r}
}

func (r *projectUsecase) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	result, err := r.repo.CreateProject(ctx, pr)
	if err != nil {
		return domain.Project{}, nil
	}

	return result, nil
}

func (r *projectUsecase) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	result, err := r.repo.GetProjects(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *projectUsecase) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	result, err := r.repo.UpdateProject(ctx, id, pr)
	if err != nil {
		return domain.Project{}, err
	}

	return result, nil
}

func (r *projectUsecase) DeleteProject(ctx context.Context, id string) error {
	if err := r.repo.DeleteProject(ctx, id); err != nil {
		return err
	}

	return nil
}
