package domain

import (
	"context"
	"time"

	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
)

type Education struct {
	ID             string
	ProfileID      string
	School         string
	Degree         string
	FieldOfStudy   string
	Gpa            float64
	StartDate      time.Time
	GraduationDate *time.Time
	IsPresent      bool
}

type EducationInput struct {
	ProfileID      string
	School         string
	Degree         string
	FieldOfStudy   string
	Gpa            float64
	StartDate      time.Time
	GraduationDate *time.Time
	IsPresent      bool
}

type EducationUpdateInput struct {
	ProfileID      *string
	School         *string
	Degree         *string
	FieldOfStudy   *string
	Gpa            *float64
	StartDate      *time.Time
	GraduationDate *time.Time
	IsPresent      *bool
}

// Merge applies the non-nil fields of in on top of e and returns the result.
// ProfileID is not mergeable: an education cannot change owner.
func (e Education) Merge(in EducationUpdateInput) Education {
	e.School = ptr.Or(in.School, e.School)
	e.Degree = ptr.Or(in.Degree, e.Degree)
	e.FieldOfStudy = ptr.Or(in.FieldOfStudy, e.FieldOfStudy)
	e.Gpa = ptr.Or(in.Gpa, e.Gpa)
	e.StartDate = ptr.Or(in.StartDate, e.StartDate)
	e.IsPresent = ptr.Or(in.IsPresent, e.IsPresent)
	if in.GraduationDate != nil {
		e.GraduationDate = in.GraduationDate
	}
	return e
}

type EducationUsecase interface {
	CreateEducation(ctx context.Context, e EducationInput) error
	GetEducations(ctx context.Context, profileID string) ([]Education, error)
	GetEducationByID(ctx context.Context, id string) (Education, error)
	UpdateEducation(ctx context.Context, id string, e EducationUpdateInput) error
	DeleteEducation(ctx context.Context, id string) error
}

type EducationRepository interface {
	CreateEducation(ctx context.Context, e EducationInput) error
	GetEducations(ctx context.Context, profileID string) ([]Education, error)
	GetEducationByID(ctx context.Context, id string) (Education, error)
	UpdateEducation(ctx context.Context, e Education) error
	DeleteEducation(ctx context.Context, id string) error
}


