package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type projectUsecase struct {
	repo domain.ProjectRepository
}

func NewProjectUsecase(repo domain.ProjectRepository) domain.ProjectUsecase {
	return &projectUsecase{repo: repo}
}

func (u *projectUsecase) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	return u.repo.CreateProject(ctx, pr)
}

func (u *projectUsecase) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	return u.repo.GetProjects(ctx, profileID)
}

func (u *projectUsecase) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	current, err := u.repo.GetProjectByID(ctx, id)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to load project for update: %w", err)
	}

	return u.repo.UpdateProject(ctx, current.Merge(pr))
}

func (u *projectUsecase) DeleteProject(ctx context.Context, id string) error {
	return u.repo.DeleteProject(ctx, id)
}
