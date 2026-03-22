package domain

import (
	"context"
	"time"
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

type EducationRepository interface {
	CreateEducation(ctx context.Context, e EducationInput) error
	GetEducations(ctx context.Context, profileID string) ([]Education, error)
	GetEducationByID(ctx context.Context, id string) (Education, error)
	UpdateEducation(ctx context.Context, id string, e EducationUpdateInput) error
	DeleteEducation(ctx context.Context, id string) error
}
