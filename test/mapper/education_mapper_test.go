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

func TestToEducationDomain(t *testing.T) {
	id := uuid.New()
	profileID := uuid.New()
	startDate := time.Now()
	graduationDate := time.Now().Add(time.Hour * 24 * 365 * 4)

	dbEducation := db.Education{
		ID:             id,
		ProfileID:      profileID,
		School:         "University of Technology",
		Degree:         "Bachelor of Science",
		FieldOfStudy:   "Computer Science",
		Gpa:            3.8,
		StartDate:      pgtype.Date{Time: startDate, Valid: true},
		GraduationDate: pgtype.Date{Time: graduationDate, Valid: true},
		IsPresent:      false,
	}

	domainEducation := mapper.ToEducationDomain(dbEducation)

	assert.Equal(t, id.String(), domainEducation.ID)
	assert.Equal(t, profileID.String(), domainEducation.ProfileID)
	assert.Equal(t, "University of Technology", domainEducation.School)
	assert.Equal(t, "Bachelor of Science", domainEducation.Degree)
	assert.Equal(t, "Computer Science", domainEducation.FieldOfStudy)
	assert.Equal(t, 3.8, domainEducation.Gpa)
	assert.Equal(t, startDate, domainEducation.StartDate)
	assert.Equal(t, &graduationDate, domainEducation.GraduationDate)
	assert.Equal(t, false, domainEducation.IsPresent)
}
