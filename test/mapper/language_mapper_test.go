package mapper_test

import (
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToLanguageDomain(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()

	dbLanguage := &db.Language{
		ID:          id,
		ProfileID:   profileID,
		Language:    "English",
		Proficiency: db.ProficiencyLevelNative,
	}

	domainLanguage := mapper.ToLanguageDomain(dbLanguage)

	assert.Equal(t, id.String(), domainLanguage.ID)
	assert.Equal(t, profileID.String(), domainLanguage.ProfileID)
	assert.Equal(t, "English", domainLanguage.Language)
	assert.Equal(t, "native", domainLanguage.Proficiency)
}
