package mapper

import (
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToExperienceDomain(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()
	startDate := time.Now()
	endDate := time.Now().Add(time.Hour * 24 * 365)

	dbExperience := db.Experience{
		ID:          id,
		ProfileID:   profileID,
		Company:     "Tech Corp",
		Position:    "Software Engineer",
		Description: "Developing cool stuff",
		StartDate:   startDate,
		EndDate:     endDate,
	}

	domainExperience := ToExperienceDomain(dbExperience)

	assert.Equal(t, id.String(), domainExperience.ID)
	assert.Equal(t, profileID.String(), domainExperience.ProfileID)
	assert.Equal(t, "Tech Corp", domainExperience.Company)
	assert.Equal(t, "Software Engineer", domainExperience.Position)
	assert.Equal(t, "Developing cool stuff", domainExperience.Description)
	assert.Equal(t, startDate, domainExperience.StartDate)
	assert.Equal(t, endDate, domainExperience.EndDate)
}
