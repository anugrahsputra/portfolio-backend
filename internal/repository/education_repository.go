package repository

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/config"
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
)

type educationRepository struct {
	db *db.Queries
}

func NewEducationRepository(database *config.Database) domain.EducationRepository {
	return &educationRepository{db: db.New(database.Pool)}
}

func (r *educationRepository) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	profileID, err := uuid.Parse(e.ProfileID)
	if err != nil {
		return err
	}
	param := db.CreateEducationParams{
		ProfileID:      profileID,
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      e.StartDate,
		GraduationDate: e.GraduationDate,
	}

	if _, err := r.db.CreateEducation(ctx, param); err != nil {
		return err
	}

	return nil
}

func (r *educationRepository) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	profileIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, err
	}

	educations, err := r.db.GetEducations(ctx, profileIDStr)
	if err != nil {
		return nil, err
	}

	var result []domain.Education
	for _, education := range educations {
		ed := mapper.ToEducationDomain(education)
		result = append(result, ed)
	}

	return result, nil
}

func (r *educationRepository) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	param := db.UpdateEducationParams{
		ID:             idStr,
		School:         *e.School,
		Degree:         *e.Degree,
		FieldOfStudy:   *e.FieldOfStudy,
		Gpa:            *e.Gpa,
		StartDate:      *e.StartDate,
		GraduationDate: *e.GraduationDate,
	}

	if _, err := r.db.UpdateEducation(ctx, param); err != nil {
		return err
	}

	return nil
}

func (r *educationRepository) DeleteEducation(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := r.db.DeleteEducation(ctx, idStr); err != nil {
		return err
	}

	return nil
}
