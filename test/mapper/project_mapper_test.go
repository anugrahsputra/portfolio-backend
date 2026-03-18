package mapper_test

import (
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToProjectDomain(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()

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
		Period:        "2023-2024",
		Location:      "Remote",
	}

	domainProject := mapper.ToProjectDomain(dbProject)

	assert.Equal(t, id.String(), domainProject.ID)
	assert.Equal(t, profileID.String(), domainProject.ProfileID)
	assert.Equal(t, dbProject.Title, domainProject.Title)
	assert.Equal(t, dbProject.TechStacks, domainProject.TechStacks)
	assert.Equal(t, dbProject.IsLive, domainProject.IsLive)
}
