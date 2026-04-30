package mapper

import (
	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
	"github.com/anugrahsputra/portfolio-backend/pkg/parser"
)

// ToProfileDomain maps db.GetProfileRow (with URLs) to domain.Profile
func ToProfileDomain(p db.GetProfileRow) (domain.Profile, error) {
	var urls []domain.ProfileUrl
	id := p.ID.String()

	urls, err := parser.JsonSliceParser[domain.ProfileUrl](p.Urls, id)
	if err != nil {
		return domain.Profile{}, err
	}

	return domain.Profile{
		ID:      p.ID.String(),
		Name:    p.Name,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
		Url:     urls,
	}, nil
}

// ToProfileDomainFromDB maps db.Profile to domain.Profile
func ToProfileDomainFromDB(p db.Profile) domain.Profile {
	return domain.Profile{
		ID:      p.ID.String(),
		Name:    p.Name,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
	}
}
