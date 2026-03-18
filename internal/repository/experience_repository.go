package repository

import (
	"context"
	"time"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
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
		return domain.Experience{}, err
	}

	param := db.CreateExperienceParams{
		ProfileID:   profileIDStr,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		StartDate:   pgtype.Date{Time: ex.StartDate, Valid: true},
		EndDate:     pgtype.Date{Time: ex.EndDate, Valid: !ex.EndDate.IsZero()},
	}

	experience, err := r.db.CreateExperience(ctx, param)
	if err != nil {
		return domain.Experience{}, err
	}

	result := mapper.ToExperienceDomain(experience)
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
		exToDomain := mapper.ToExperienceDomain(experience)
		result = append(result, exToDomain)
	}

	return result, nil
}

func (r *experienceRepository) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Experience{}, err
	}

	var ed pgtype.Date
	if ex.EndDate != nil {
		ed = pgtype.Date{Time: *ex.EndDate, Valid: !ex.EndDate.IsZero()}
	}

	param := db.UpdateExperienceParams{
		ID:          idStr,
		Company:     ptr.Or(ex.Company, ""),
		Position:    ptr.Or(ex.Position, ""),
		Description: ptr.Or(ex.Description, []string{}),
		StartDate:   pgtype.Date{Time: ptr.Or(ex.StartDate, time.Time{}), Valid: ex.StartDate != nil},
		EndDate:     ed,
	}

	experience, err := r.db.UpdateExperience(ctx, param)
	if err != nil {
		return domain.Experience{}, err
	}

	result := mapper.ToExperienceDomain(experience)
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
