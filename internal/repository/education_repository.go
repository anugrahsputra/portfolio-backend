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

type educationRepository struct {
	db *db.Queries
}

func NewEducationRepository(database *config.Database) domain.EducationRepository {
	return &educationRepository{db: db.New(database.Pool)}
}

func (r *educationRepository) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	profileID, err := uuid.Parse(e.ProfileID)
	if err != nil {
		return fmt.Errorf("invalid profile id: %w", err)
	}

	var gd pgtype.Date
	if e.GraduationDate != nil {
		gd = pgtype.Date{Time: *e.GraduationDate, Valid: !e.GraduationDate.IsZero()}
	}

	param := db.CreateEducationParams{
		ProfileID:      profileID,
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      pgtype.Date{Time: e.StartDate, Valid: true},
		GraduationDate: gd,
		IsPresent:      e.IsPresent,
	}

	if _, err := r.db.CreateEducation(ctx, param); err != nil {
		return fmt.Errorf("failed to create education: %w", err)
	}

	return nil
}

func (r *educationRepository) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	profileIDStr, err := uuid.Parse(profileID)
	if err != nil {
		return nil, fmt.Errorf("invalid profile id: %w", err)
	}

	educations, err := r.db.GetEducations(ctx, profileIDStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get educations: %w", err)
	}

	result := make([]domain.Education, 0, len(educations))
	for _, education := range educations {
		ed := r.toDomain(education)
		result = append(result, ed)
	}

	return result, nil
}

func (r *educationRepository) GetEducationByID(ctx context.Context, id string) (domain.Education, error) {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return domain.Education{}, fmt.Errorf("invalid education id: %w", err)
	}

	education, err := r.db.GetEducationByID(ctx, idStr)
	if err != nil {
		return domain.Education{}, fmt.Errorf("failed to get education by id: %w", err)
	}

	result := r.toDomain(education)

	return result, nil
}

func (r *educationRepository) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid education id: %w", err)
	}

	currentDB, err := r.db.GetEducationByID(ctx, idStr)
	if err != nil {
		return fmt.Errorf("failed to fetch existing education for update: %w", err)
	}

	merged := r.toDomain(currentDB).Merge(e)

	var gd pgtype.Date
	if merged.GraduationDate != nil {
		gd = pgtype.Date{Time: *merged.GraduationDate, Valid: !merged.GraduationDate.IsZero()}
	}

	param := db.UpdateEducationParams{
		ID:             idStr,
		School:         merged.School,
		Degree:         merged.Degree,
		FieldOfStudy:   merged.FieldOfStudy,
		Gpa:            merged.Gpa,
		StartDate:      pgtype.Date{Time: merged.StartDate, Valid: true},
		GraduationDate: gd,
		IsPresent:      merged.IsPresent,
	}

	if _, err := r.db.UpdateEducation(ctx, param); err != nil {
		return fmt.Errorf("failed to update education: %w", err)
	}

	return nil
}

func (r *educationRepository) toDomain(e db.Education) domain.Education {
	var gd *time.Time
	if e.GraduationDate.Valid {
		gd = &e.GraduationDate.Time
	}

	return domain.Education{
		ID:             e.ID.String(),
		ProfileID:      e.ProfileID.String(),
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      e.StartDate.Time,
		GraduationDate: gd,
		IsPresent:      e.IsPresent,
	}
}

func (r *educationRepository) DeleteEducation(ctx context.Context, id string) error {
	idStr, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid education id: %w", err)
	}

	if err := r.db.DeleteEducation(ctx, idStr); err != nil {
		return fmt.Errorf("failed to delete education: %w", err)
	}

	return nil
}
