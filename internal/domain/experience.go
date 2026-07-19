package domain

import (
	"context"
	"time"

	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
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

// Merge applies the non-nil fields of in on top of ex and returns the result.
// ProfileID is not mergeable: an experience cannot change owner.
func (ex Experience) Merge(in ExperienceUpdateInput) Experience {
	ex.Company = ptr.Or(in.Company, ex.Company)
	ex.Position = ptr.Or(in.Position, ex.Position)
	ex.Description = ptr.Or(in.Description, ex.Description)
	ex.Location = ptr.Or(in.Location, ex.Location)
	ex.StartDate = ptr.Or(in.StartDate, ex.StartDate)
	ex.IsPresent = ptr.Or(in.IsPresent, ex.IsPresent)
	if in.EndDate != nil {
		ex.EndDate = in.EndDate
	}
	return ex
}

type ExperienceRepository interface {
	CreateExperience(ctx context.Context, ex ExperienceInput) (Experience, error)
	GetExperiences(ctx context.Context, profileID string) ([]Experience, error)
	GetExperienceByID(ctx context.Context, id string) (Experience, error)
	UpdateExperience(ctx context.Context, id string, ex ExperienceUpdateInput) (Experience, error)
	DeleteExperience(ctx context.Context, id string) error
}


