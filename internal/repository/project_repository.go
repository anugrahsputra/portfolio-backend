package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type projectRepository struct {
	db *db.Queries
}

func NewProjectRepository(database *config.Database) domain.ProjectRepository {
	return &projectRepository{db: db.New(database.Pool)}
}

func (r *projectRepository) CreateProject(ctx context.Context, pr domain.ProjectInput) (domain.Project, error) {
	profileIDStr, err := uuid.Parse(pr.ProfileID)
	if err != nil {
		return domain.Project{}, fmt.Errorf("invalid profile id: %w", err)
	}

	var ed pgtype.Date
	if pr.EndDate != nil {
		ed = pgtype.Date{Time: *pr.EndDate, Valid: !pr.EndDate.IsZero()}
	}

	param := db.CreateProjectParams{
		ProfileID:     profileIDStr,
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
		StartDate:     pgtype.Date{Time: pr.StartDate, Valid: true},
		EndDate:       ed,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}

	project, err := r.db.CreateProject(ctx, param)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to create project: %w", err)
	}

	result := mapper.ToProjectDomain(project)
	return result, nil
}

func (r *projectRepository) GetProjects(ctx context.Context, profileID string) ([]domain.Project, error) {
	projectIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, fmt.Errorf("invalid profile id: %w", err)
	}

	projects, err := r.db.GetProjects(ctx, projectIDStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects: %w", err)
	}

	result := make([]domain.Project, 0, len(projects))
	for _, project := range projects {
		item := mapper.ToProjectDomain(project)
		result = append(result, item)
	}

	return result, nil
}

func (r *projectRepository) GetProjectByID(ctx context.Context, id string) (domain.Project, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Project{}, fmt.Errorf("invalid project id: %w", err)
	}

	project, err := r.db.GetProjectByID(ctx, idStr)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to get project by id: %w", err)
	}

	result := mapper.ToProjectDomain(project)
	return result, nil
}

func (r *projectRepository) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Project{}, fmt.Errorf("invalid project id: %w", err)
	}

	currentDB, err := r.db.GetProjectByID(ctx, idStr)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to fetch existing project for update: %w", err)
	}
	current := mapper.ToProjectDomain(currentDB)

	var ed pgtype.Date
	if pr.EndDate != nil {
		ed = pgtype.Date{Time: *pr.EndDate, Valid: !pr.EndDate.IsZero()}
	} else if current.EndDate != nil {
		ed = pgtype.Date{Time: *current.EndDate, Valid: !current.EndDate.IsZero()}
	}

	param := db.UpdateProjectParams{
		ID:            idStr,
		Title:         ptr.Or(pr.Title, current.Title),
		Description:   ptr.Or(pr.Description, current.Description),
		TechStacks:    ptr.Or(pr.TechStacks, current.TechStacks),
		LiveDemoUrl:   ptr.Or(pr.LiveDemoUrl, current.LiveDemoUrl),
		GithubRepoUrl: ptr.Or(pr.GithubRepoUrl, current.GithubRepoUrl),
		IsLive:        ptr.Or(pr.IsLive, current.IsLive),
		IsNda:         ptr.Or(pr.IsNda, current.IsNda),
		IsFeatured:    ptr.Or(pr.IsFeatured, current.IsFeatured),
		ImageUrl:      ptr.Or(pr.ImageUrl, current.ImageUrl),
		Company:       ptr.Or(pr.Company, current.Company),
		StartDate:     pgtype.Date{Time: ptr.Or(pr.StartDate, time.Time{}), Valid: pr.StartDate != nil},
		EndDate:       ed,
		IsPresent:     ptr.Or(pr.IsPresent, current.IsPresent),
		Location:      ptr.Or(pr.Location, current.Location),
	}

	project, err := r.db.UpdateProject(ctx, param)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to update project: %w", err)
	}

	result := mapper.ToProjectDomain(project)
	return result, nil
}

func (r *projectRepository) DeleteProject(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid project id: %w", err)
	}

	if err := r.db.DeleteProject(ctx, idStr); err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}

	return nil
}
