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
		StartDate:      pgtype.Date{Time: e.StartDate, Valid: true},
		GraduationDate: pgtype.Date{Time: e.GraduationDate, Valid: !e.GraduationDate.IsZero()},
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

	result := make([]domain.Education, 0, len(educations))
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

	var gd pgtype.Date
	if e.GraduationDate != nil {
		gd = pgtype.Date{Time: *e.GraduationDate, Valid: !e.GraduationDate.IsZero()}
	}

	param := db.UpdateEducationParams{
		ID:             idStr,
		School:         ptr.Or(e.School, ""),
		Degree:         ptr.Or(e.Degree, ""),
		FieldOfStudy:   ptr.Or(e.FieldOfStudy, ""),
		Gpa:            ptr.Or(e.Gpa, 0),
		StartDate:      pgtype.Date{Time: ptr.Or(e.StartDate, time.Time{}), Valid: e.StartDate != nil},
		GraduationDate: gd,
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
