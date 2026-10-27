package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
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
		IsComingSoon:  pr.IsComingSoon,
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

	result := r.toDomain(project)
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
		item := r.toDomain(project)
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

	result := r.toDomain(project)
	return result, nil
}

func (r *projectRepository) UpdateProject(ctx context.Context, pr domain.Project) (domain.Project, error) {
	idStr, err := uuid.Parse(pr.ID)
	if err != nil {
		return domain.Project{}, fmt.Errorf("invalid project id: %w", err)
	}

	var ed pgtype.Date
	if pr.EndDate != nil {
		ed = pgtype.Date{Time: *pr.EndDate, Valid: !pr.EndDate.IsZero()}
	}

	param := db.UpdateProjectParams{
		ID:            idStr,
		Title:         pr.Title,
		Description:   pr.Description,
		TechStacks:    pr.TechStacks,
		LiveDemoUrl:   pr.LiveDemoUrl,
		GithubRepoUrl: pr.GithubRepoUrl,
		IsLive:        pr.IsLive,
		IsNda:         pr.IsNda,
		IsComingSoon:  pr.IsComingSoon,
		IsFeatured:    pr.IsFeatured,
		ImageUrl:      pr.ImageUrl,
		Company:       pr.Company,
		StartDate:     pgtype.Date{Time: pr.StartDate, Valid: true},
		EndDate:       ed,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}

	project, err := r.db.UpdateProject(ctx, param)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to update project: %w", err)
	}

	return r.toDomain(project), nil
}

func (r *projectRepository) toDomain(pr db.Project) domain.Project {
	var ed *time.Time
	if pr.EndDate.Valid {
		ed = &pr.EndDate.Time
	}

	return domain.Project{
		ID:            pr.ID.String(),
		ProfileID:     pr.ProfileID.String(),
		Title:         pr.Title,
		Description:   pr.Description,
		TechStacks:    pr.TechStacks,
		LiveDemoUrl:   pr.LiveDemoUrl,
		GithubRepoUrl: pr.GithubRepoUrl,
		IsLive:        pr.IsLive,
		IsNda:         pr.IsNda,
		IsComingSoon:  pr.IsComingSoon,
		IsFeatured:    pr.IsFeatured,
		ImageUrl:      pr.ImageUrl,
		Company:       pr.Company,
		StartDate:     pr.StartDate.Time,
		EndDate:       ed,
		IsPresent:     pr.IsPresent,
		Location:      pr.Location,
	}
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
