package domain

import (
	"context"
	"time"
)

type Experience struct {
	ID          string
	ProfileID   string
	Company     string
	Position    string
	Description []string
	Location    string
	StartDate   time.Time
	EndDate     *time.Time
	IsPresent   bool
}

type ExperienceInput struct {
	ProfileID   string
	Company     string
	Position    string
	Description []string
	Location    string
	StartDate   time.Time
	EndDate     *time.Time
	IsPresent   bool
}

type ExperienceUpdateInput struct {
	ProfileID   *string
	Company     *string
	Position    *string
	Description *[]string
	Location    *string
	StartDate   *time.Time
	EndDate     *time.Time
	IsPresent   *bool
}

type ExperienceRepository interface {
	CreateExperience(ctx context.Context, ex ExperienceInput) (Experience, error)
	GetExperiences(ctx context.Context, profileID string) ([]Experience, error)
	GetExperienceByID(ctx context.Context, id string) (Experience, error)
	UpdateExperience(ctx context.Context, id string, ex ExperienceUpdateInput) (Experience, error)
	DeleteExperience(ctx context.Context, id string) error
}

func (e *Experience) SetProfileID(profileID string) {
	e.ProfileID = profileID
}
