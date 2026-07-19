package domain_test

import (
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/stretchr/testify/assert"
)

func strPtr(s string) *string { return &s }

func TestProjectMerge(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	current := domain.Project{
		ID:        "p1",
		ProfileID: "owner",
		Title:     "old title",
		Location:  "Jakarta",
		StartDate: start,
		EndDate:   &end,
		IsLive:    true,
	}

	t.Run("empty input keeps every field", func(t *testing.T) {
		got := current.Merge(domain.ProjectUpdateInput{})
		assert.Equal(t, current, got)
	})

	t.Run("non-nil fields overwrite", func(t *testing.T) {
		newEnd := time.Date(2025, 2, 2, 0, 0, 0, 0, time.UTC)
		got := current.Merge(domain.ProjectUpdateInput{
			Title:   strPtr("new title"),
			EndDate: &newEnd,
		})
		assert.Equal(t, "new title", got.Title)
		assert.Equal(t, &newEnd, got.EndDate)
		assert.Equal(t, "Jakarta", got.Location, "untouched field must survive")
	})

	t.Run("start date survives a partial update", func(t *testing.T) {
		// Regression: start_date is NOT NULL, so a partial update that omits it
		// must still carry the existing value through.
		got := current.Merge(domain.ProjectUpdateInput{Title: strPtr("x")})
		assert.Equal(t, start, got.StartDate)
	})

	t.Run("false is a real value, not absence", func(t *testing.T) {
		isLive := false
		got := current.Merge(domain.ProjectUpdateInput{IsLive: &isLive})
		assert.False(t, got.IsLive)
	})

	t.Run("owner is not mergeable", func(t *testing.T) {
		got := current.Merge(domain.ProjectUpdateInput{Title: strPtr("x")})
		assert.Equal(t, "owner", got.ProfileID)
	})
}

func TestProfileUrlMerge(t *testing.T) {
	current := domain.ProfileUrl{ID: "u1", ProfileID: "owner", Label: "GitHub", Url: "https://gh"}

	// Regression: this used to dereference nil pointers and panic.
	t.Run("empty input keeps every field", func(t *testing.T) {
		got := current.Merge(domain.ProfileUrlUpdateInput{})
		assert.Equal(t, current, got)
	})

	t.Run("partial input merges only what is set", func(t *testing.T) {
		got := current.Merge(domain.ProfileUrlUpdateInput{Label: strPtr("Portfolio")})
		assert.Equal(t, "Portfolio", got.Label)
		assert.Equal(t, "https://gh", got.Url)
	})
}

func TestProfileMerge(t *testing.T) {
	current := domain.Profile{ID: "1", Name: "Ada", Title: "Engineer", Email: "a@b.c"}

	got := current.Merge(domain.ProfileUpdateInput{Name: strPtr("Grace")})
	assert.Equal(t, "Grace", got.Name)
	assert.Equal(t, "Engineer", got.Title)
	assert.Equal(t, "a@b.c", got.Email)
}

func TestEducationMerge(t *testing.T) {
	start := time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)
	grad := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	current := domain.Education{School: "MIT", Gpa: 3.9, StartDate: start, GraduationDate: &grad}

	got := current.Merge(domain.EducationUpdateInput{School: strPtr("Stanford")})
	assert.Equal(t, "Stanford", got.School)
	assert.Equal(t, 3.9, got.Gpa)
	assert.Equal(t, start, got.StartDate)
	assert.Equal(t, &grad, got.GraduationDate)
}

func TestExperienceMerge(t *testing.T) {
	start := time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC)
	current := domain.Experience{
		Company:     "Acme",
		Position:    "Dev",
		Description: []string{"a", "b"},
		StartDate:   start,
		IsPresent:   true,
	}

	desc := []string{"c"}
	got := current.Merge(domain.ExperienceUpdateInput{Description: &desc})
	assert.Equal(t, []string{"c"}, got.Description)
	assert.Equal(t, "Acme", got.Company)
	assert.Equal(t, start, got.StartDate)
	assert.True(t, got.IsPresent)
}
