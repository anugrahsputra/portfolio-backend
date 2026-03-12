package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToLanguageDomain(db *db.Language) domain.Language {
	return domain.Language{
		ID:          db.ID.String(),
		ProfileID:   db.ProfileID.String(),
		Language:    db.Language,
		Proficiency: string(db.Proficiency),
	}
}
