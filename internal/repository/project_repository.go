package repository

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
)

type projectRepository struct {
	db *db.Queries
}

func NewProjectRepository(database *config.Database) domain.ProjectRepository {
	return &projectRepository{db: db.New(database.Pool)}
}

func (r *projectRepository) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	projectIDStr, err := uuid.Parse(pr.ProfileID)
	if err != nil {
		return domain.Project{}, err
	}

	param := db.CreateProjectParams{
		ProfileID:     projectIDStr,
		Title:         pr.Title,
		Description:   pr.Description,
		TechStacks:    pr.TechStacks,
		LiveDemoUrl:   pr.LiveDemoUrl,
		GithubRepoUrl: pr.GithubRepoUrl,
		IsLive:        pr.IsLive,
		IsNda:         pr.IsNda,
		IsFeatured:    pr.IsFeatured,
		ImageUrl:      pr.ImageUrl,
		Company:       pr.Company,
		Period:        pr.Period,
		Location:      pr.Location,
	}

	project, err := r.db.CreateProject(ctx, param)
	if err != nil {
		return domain.Project{}, err
	}

	result := mapper.ToProjectDomain(project)
	return result, nil
}

func (r *projectRepository) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	projectIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, err
	}

	projects, err := r.db.GetProjects(ctx, projectIDStr)
	if err != nil {
		return nil, err
	}

	result := make([]domain.Project, 0, len(projects))
	for _, project := range projects {
		item := mapper.ToProjectDomain(project)
		result = append(result, item)
	}

	return result, nil
}

func (r *projectRepository) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Project{}, err
	}

	param := db.UpdateProjectParams{
		ID:            idStr,
		Title:         *pr.Title,
		Description:   *pr.Description,
		TechStacks:    *pr.TechStacks,
		LiveDemoUrl:   *pr.LiveDemoUrl,
		GithubRepoUrl: *pr.GithubRepoUrl,
		IsLive:        *pr.IsLive,
		IsNda:         *pr.IsNda,
		IsFeatured:    *pr.IsFeatured,
		ImageUrl:      *pr.ImageUrl,
		Company:       *pr.Company,
		Period:        *pr.Period,
		Location:      *pr.Location,
	}

	project, err := r.db.UpdateProject(ctx, param)
	if err != nil {
		return domain.Project{}, err
	}

	result := mapper.ToProjectDomain(project)
	return result, nil
}

func (r *projectRepository) DeleteProject(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := r.db.DeleteProject(ctx, idStr); err != nil {
		return err
	}

	return nil
}
