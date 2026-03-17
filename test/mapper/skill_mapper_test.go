package mapper_test

import (
	"testing"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/mapper"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToSkillDomain(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()

	dbSkill := &db.Skill{
		ID:           id,
		ProfileID:    profileID,
		Tools:        []string{"Git", "Docker"},
		Technologies: []string{"Go", "PostgreSQL"},
		HardSkills:   []string{"Programming"},
		SoftSkills:   []string{"Communication"},
	}

	domainSkill := mapper.ToSkillDomain(dbSkill)

	assert.Equal(t, id.String(), domainSkill.ID)
	assert.Equal(t, profileID.String(), domainSkill.ProfileID)
	assert.Equal(t, []string{"Git", "Docker"}, domainSkill.Tools)
	assert.Equal(t, []string{"Go", "PostgreSQL"}, domainSkill.Technologies)
	assert.Equal(t, []string{"Programming"}, domainSkill.HardSkills)
	assert.Equal(t, []string{"Communication"}, domainSkill.SoftSkills)
}
