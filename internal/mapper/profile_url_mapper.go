package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

func ToProfileURLDomain(pu db.ProfileUrl) domain.ProfileUrl {
	return domain.ProfileUrl{
		ProfileID: pu.ID.String(),
		Label:     pu.Label,
		Url:       pu.Url,
	}
}
