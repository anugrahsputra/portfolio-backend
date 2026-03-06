package repository

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
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
		return domain.Experience{}, err
	}
	param := db.CreateExperienceParams{
		ProfileID:   profileIDStr,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		StartDate:   ex.StartDate,
		EndDate:     ex.EndDate,
	}

	experience, err := r.db.CreateExperience(ctx, param)
	if err != nil {
		return domain.Experience{}, err
	}

	result := mapper.ToExperienceMapper(experience)
	return result, nil
}

func (r *experienceRepository) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	profileIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, err
	}

	experiences, err := r.db.GetExperiences(ctx, profileIDStr)
	if err != nil {
		return nil, err
	}

	var result []domain.Experience
	for _, experience := range experiences {
		exToDomain := mapper.ToExperienceMapper(experience)
		result = append(result, exToDomain)
	}

	return result, nil
}

func (r *experienceRepository) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Experience{}, err
	}

	param := db.UpdateExperienceParams{
		ID:          idStr,
		Company:     *ex.Company,
		Position:    *ex.Position,
		Description: *ex.Description,
		StartDate:   *ex.StartDate,
		EndDate:     *ex.EndDate,
	}

	experience, err := r.db.UpdateExperience(ctx, param)
	if err != nil {
		return domain.Experience{}, err
	}

	result := mapper.ToExperienceMapper(experience)
	return result, nil
}

func (r *experienceRepository) DeleteExperience(ctx context.Context, id string) error {

	idStr, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := r.db.DeleteExperience(ctx, idStr); err != nil {
		return err
	}

	return nil
}
