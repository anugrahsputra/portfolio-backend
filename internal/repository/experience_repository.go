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

type experienceRepository struct {
	db *db.Queries
}

func NewExperienceRepository(database *config.Database) domain.ExperienceRepository {
	return &experienceRepository{db: db.New(database.Pool)}
}

func (r *experienceRepository) CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error) {
	profileIDStr, err := uuid.Parse(ex.ProfileID)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("invalid profile id: %w", err)
	}

	var ed pgtype.Date
	if ex.EndDate != nil {
		ed = pgtype.Date{Time: *ex.EndDate, Valid: !ex.EndDate.IsZero()}
	}

	param := db.CreateExperienceParams{
		ProfileID:   profileIDStr,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		Location:    ex.Location,
		StartDate:   pgtype.Date{Time: ex.StartDate, Valid: true},
		EndDate:     ed,
		IsPresent:   ex.IsPresent,
	}

	experience, err := r.db.CreateExperience(ctx, param)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("failed to create experience: %w", err)
	}

	result := r.toDomain(experience)
	return result, nil
}

func (r *experienceRepository) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	profileIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, fmt.Errorf("invalid profile id: %w", err)
	}

	experiences, err := r.db.GetExperiences(ctx, profileIDStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get experiences: %w", err)
	}

	var result []domain.Experience
	for _, experience := range experiences {
		exToDomain := r.toDomain(experience)
		result = append(result, exToDomain)
	}

	return result, nil
}

func (r *experienceRepository) GetExperienceByID(ctx context.Context, id string) (domain.Experience, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("invalid experience id: %w", err)
	}

	experience, err := r.db.GetExperienceByID(ctx, idStr)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("failed to get experience by id: %w", err)
	}

	result := r.toDomain(experience)
	return result, nil
}

func (r *experienceRepository) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("invalid experience id: %w", err)
	}

	currentDB, err := r.db.GetExperienceByID(ctx, idStr)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("failed to fetch existing experience for update: %w", err)
	}
	current := r.toDomain(currentDB)

	var ed pgtype.Date
	if ex.EndDate != nil {
		ed = pgtype.Date{Time: *ex.EndDate, Valid: !ex.EndDate.IsZero()}
	} else if current.EndDate != nil {
		ed = pgtype.Date{Time: *current.EndDate, Valid: !current.EndDate.IsZero()}
	}

	company := current.Company
	if ex.Company != nil {
		company = *ex.Company
	}
	position := current.Position
	if ex.Position != nil {
		position = *ex.Position
	}
	description := current.Description
	if ex.Description != nil {
		description = *ex.Description
	}
	location := current.Location
	if ex.Location != nil {
		location = *ex.Location
	}
	startDate := current.StartDate
	if ex.StartDate != nil {
		startDate = *ex.StartDate
	}
	isPresent := current.IsPresent
	if ex.IsPresent != nil {
		isPresent = *ex.IsPresent
	}

	param := db.UpdateExperienceParams{
		ID:          idStr,
		Company:     company,
		Position:    position,
		Description: description,
		Location:    location,
		StartDate:   pgtype.Date{Time: startDate, Valid: true},
		EndDate:     ed,
		IsPresent:   isPresent,
	}

	experience, err := r.db.UpdateExperience(ctx, param)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("failed to update experience: %w", err)
	}

	result := r.toDomain(experience)
	return result, nil
}

func (r *experienceRepository) toDomain(ex db.Experience) domain.Experience {
	var ed *time.Time
	if ex.EndDate.Valid {
		ed = &ex.EndDate.Time
	}

	return domain.Experience{
		ID:          ex.ID.String(),
		ProfileID:   ex.ProfileID.String(),
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		Location:    ex.Location,
		StartDate:   ex.StartDate.Time,
		EndDate:     ed,
		IsPresent:   ex.IsPresent,
	}
}

func (r *experienceRepository) DeleteExperience(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid experience id: %w", err)
	}

	if err := r.db.DeleteExperience(ctx, idStr); err != nil {
		return fmt.Errorf("failed to delete experience: %w", err)
	}

	return nil
}
