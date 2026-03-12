package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToExperienceDomain(ex db.Experience) domain.Experience {
	return domain.Experience{
		ID:          ex.ID.String(),
		ProfileID:   ex.ProfileID.String(),
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		StartDate:   ex.StartDate,
		EndDate:     ex.EndDate,
	}
}
