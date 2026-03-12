package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToSkillDomain(db *db.Skill) domain.Skill {
	return domain.Skill{
		ID:           db.ID.String(),
		ProfileID:    db.ProfileID.String(),
		Tools:        db.Tools,
		Technologies: db.Technologies,
		HardSkills:   db.HardSkills,
		SoftSkills:   db.SoftSkills,
	}
}
