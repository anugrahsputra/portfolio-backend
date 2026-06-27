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

func (r *projectRepository) UpdateProject(ctx context.Context, id string, pr domain.ProjectUpdateInput) (domain.Project, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Project{}, fmt.Errorf("invalid project id: %w", err)
	}

	currentDB, err := r.db.GetProjectByID(ctx, idStr)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to fetch existing project for update: %w", err)
	}
	current := r.toDomain(currentDB)

	var ed pgtype.Date
	if pr.EndDate != nil {
		ed = pgtype.Date{Time: *pr.EndDate, Valid: !pr.EndDate.IsZero()}
	} else if current.EndDate != nil {
		ed = pgtype.Date{Time: *current.EndDate, Valid: !current.EndDate.IsZero()}
	}

	title := current.Title
	if pr.Title != nil {
		title = *pr.Title
	}
	description := current.Description
	if pr.Description != nil {
		description = *pr.Description
	}
	techStacks := current.TechStacks
	if pr.TechStacks != nil {
		techStacks = *pr.TechStacks
	}
	liveDemoUrl := current.LiveDemoUrl
	if pr.LiveDemoUrl != nil {
		liveDemoUrl = *pr.LiveDemoUrl
	}
	githubRepoUrl := current.GithubRepoUrl
	if pr.GithubRepoUrl != nil {
		githubRepoUrl = *pr.GithubRepoUrl
	}
	isLive := current.IsLive
	if pr.IsLive != nil {
		isLive = *pr.IsLive
	}
	isNda := current.IsNda
	if pr.IsNda != nil {
		isNda = *pr.IsNda
	}
	isFeatured := current.IsFeatured
	if pr.IsFeatured != nil {
		isFeatured = *pr.IsFeatured
	}
	imageUrl := current.ImageUrl
	if pr.ImageUrl != nil {
		imageUrl = *pr.ImageUrl
	}
	company := current.Company
	if pr.Company != nil {
		company = *pr.Company
	}
	startDate := current.StartDate
	if pr.StartDate != nil {
		startDate = *pr.StartDate
	}
	isPresent := current.IsPresent
	if pr.IsPresent != nil {
		isPresent = *pr.IsPresent
	}
	location := current.Location
	if pr.Location != nil {
		location = *pr.Location
	}

	param := db.UpdateProjectParams{
		ID:            idStr,
		Title:         title,
		Description:   description,
		TechStacks:    techStacks,
		LiveDemoUrl:   liveDemoUrl,
		GithubRepoUrl: githubRepoUrl,
		IsLive:        isLive,
		IsNda:         isNda,
		IsFeatured:    isFeatured,
		ImageUrl:      imageUrl,
		Company:       company,
		StartDate:     pgtype.Date{Time: startDate, Valid: pr.StartDate != nil},
		EndDate:       ed,
		IsPresent:     isPresent,
		Location:      location,
	}

	project, err := r.db.UpdateProject(ctx, param)
	if err != nil {
		return domain.Project{}, fmt.Errorf("failed to update project: %w", err)
	}

	result := r.toDomain(project)
	return result, nil
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
