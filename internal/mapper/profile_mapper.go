package mapper

import (
	"encoding/json"

	"github.com/anugrahsputra/portfolio-backend/internal/db"
	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

// ToProfileDomain maps db.GetProfileRow (with URLs) to domain.Profile
func ToProfileDomain(p db.GetProfileRow) (domain.Profile, error) {
	var urls []domain.ProfileUrl

	if p.Urls != nil {
		var data []byte
		switch v := p.Urls.(type) {
		case []byte:
			data = v
		case string:
			data = []byte(v)
		default:
			data, _ = json.Marshal(v)
		}

		if len(data) > 0 {
			if err := json.Unmarshal(data, &urls); err != nil {
				return domain.Profile{}, err
			}

			for i := range urls {
				urls[i].ProfileID = p.ID.String()
			}
		}
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
