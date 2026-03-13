package mapper

import (
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToProfileDomain(t *testing.T) {
	id := uuid.New()
	
	t.Run("with URLs", func(t *testing.T) {
		// Mock data as it might come from pgx (already parsed) or as a slice of maps
		urlsData := []map[string]interface{}{
			{
				"ID":    "url-id",
				"Label": "GitHub",
				"Url":   "https://github.com/user",
			},
		}
		dbProfile := db.GetProfileRow{
			ID:      id,
			Name:    "John Doe",
			About:   "Developer",
			Address: "123 Street",
			Email:   "john@example.com",
			Phone:   "123456789",
			Urls:    urlsData,
		}

		domainProfile, err := ToProfileDomain(dbProfile)
		assert.NoError(t, err)
		assert.Equal(t, id.String(), domainProfile.ID)
		assert.Equal(t, "John Doe", domainProfile.Name)
		assert.Len(t, domainProfile.Url, 1)
		assert.Equal(t, "GitHub", domainProfile.Url[0].Label)
		// Verify that the mapper correctly populated the missing ProfileID
		assert.Equal(t, id.String(), domainProfile.Url[0].ProfileID)
	})

	t.Run("without URLs", func(t *testing.T) {
		dbProfile := db.GetProfileRow{
			ID:      id,
			Name:    "John Doe",
			About:   "Developer",
			Address: "123 Street",
			Email:   "john@example.com",
			Phone:   "123456789",
			Urls:    nil,
		}

		domainProfile, err := ToProfileDomain(dbProfile)
		assert.NoError(t, err)
		assert.Equal(t, id.String(), domainProfile.ID)
		assert.Empty(t, domainProfile.Url)
	})
}

func TestToProfileDomainFromDB(t *testing.T) {
	id := uuid.New()
	dbProfile := db.Profile{
		ID:      id,
		Name:    "John Doe",
		About:   "Developer",
		Address: "123 Street",
		Email:   "john@example.com",
		Phone:   "123456789",
	}

	domainProfile := ToProfileDomainFromDB(dbProfile)

	assert.Equal(t, id.String(), domainProfile.ID)
	assert.Equal(t, "John Doe", domainProfile.Name)
}
