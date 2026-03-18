package mapper_test

import (
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/anugrahsputra/portfolio-backend/pkg/ptr"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToProjectDomain(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()
	startDate := time.Now()
	endDate := time.Now().AddDate(1, 0, 0)

	dbProject := db.Project{
		ID:            id,
		ProfileID:     profileID,
		Title:         "Test Project",
		Description:   []string{"Test Description"},
		TechStacks:    []string{"Go", "PostgreSQL"},
		LiveDemoUrl:   "https://live.com",
		GithubRepoUrl: "https://github.com/test",
		IsLive:        true,
		IsNda:         false,
		IsFeatured:    true,
		ImageUrl:      "https://image.com",
		Company:       "Test Company",
		StartDate:     startDate,
		EndDate:       endDate,
		Location:      "Remote",
	}

	domainProject := mapper.ToProjectDomain(dbProject)

	assert.Equal(t, id.String(), domainProject.ID)
	assert.Equal(t, profileID.String(), domainProject.ProfileID)
	assert.Equal(t, dbProject.Title, domainProject.Title)
	assert.Equal(t, dbProject.TechStacks, domainProject.TechStacks)
	assert.Equal(t, dbProject.IsLive, domainProject.IsLive)
	assert.Equal(t, startDate, domainProject.StartDate)
	assert.Equal(t, ptr.To(endDate), domainProject.EndDate)
}
