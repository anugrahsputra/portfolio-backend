package mapper_test

import (
	"testing"
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
		Description: []string{"Developing cool stuff", "Managing team"},
		Location:    "Jakarta, Indonesia",
		StartDate:   pgtype.Date{Time: startDate, Valid: true},
		EndDate:     pgtype.Date{Time: endDate, Valid: true},
		IsPresent:   false,
	}

	domainExperience := mapper.ToExperienceDomain(dbExperience)

	assert.Equal(t, id.String(), domainExperience.ID)
	assert.Equal(t, profileID.String(), domainExperience.ProfileID)
	assert.Equal(t, "Tech Corp", domainExperience.Company)
	assert.Equal(t, "Software Engineer", domainExperience.Position)
	assert.Equal(t, []string{"Developing cool stuff", "Managing team"}, domainExperience.Description)
	assert.Equal(t, "Jakarta, Indonesia", domainExperience.Location)
	assert.Equal(t, startDate, domainExperience.StartDate)
	assert.Equal(t, &endDate, domainExperience.EndDate)
	assert.Equal(t, false, domainExperience.IsPresent)
}
