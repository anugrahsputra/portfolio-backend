package mapper_test

import (
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToProfileURLDomain(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()

	dbProfileUrl := db.ProfileUrl{
		ID:        id,
		ProfileID: profileID,
		Label:     "LinkedIn",
		Url:       "https://linkedin.com/in/user",
	}

	domainProfileUrl := mapper.ToProfileURLDomain(dbProfileUrl)

	assert.Equal(t, id.String(), domainProfileUrl.ID)
	assert.Equal(t, profileID.String(), domainProfileUrl.ProfileID)
	assert.Equal(t, "LinkedIn", domainProfileUrl.Label)
	assert.Equal(t, "https://linkedin.com/in/user", domainProfileUrl.Url)
}
