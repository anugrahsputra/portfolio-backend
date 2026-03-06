package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToEducationDomain(e db.Education) domain.Education {
	return domain.Education{
		ProfileID:      e.ProfileID.String(),
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      e.StartDate,
		GraduationDate: e.GraduationDate,
	}
}
