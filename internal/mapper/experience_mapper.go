package mapper

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToExperienceDomain(ex db.Experience) domain.Experience {
	ed := time.Time{}
	if ex.EndDate.Valid {
		ed = ex.EndDate.Time
	}

	return domain.Experience{
		ID:          ex.ID.String(),
		ProfileID:   ex.ProfileID.String(),
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		StartDate:   ex.StartDate.Time,
		EndDate:     ed,
	}
}
