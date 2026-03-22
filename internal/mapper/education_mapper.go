package mapper

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToEducationDomain(e db.Education) domain.Education {
	var gd *time.Time
	if e.GraduationDate.Valid {
		gd = &e.GraduationDate.Time
	}

	return domain.Education{
		ID:             e.ID.String(),
		ProfileID:      e.ProfileID.String(),
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      e.StartDate.Time,
		GraduationDate: gd,
		IsPresent:      e.IsPresent,
	}
}
