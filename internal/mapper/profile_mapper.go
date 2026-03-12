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
		var urlsData []byte
		switch v := p.Urls.(type) {
		case []byte:
			urlsData = v
		case string:
			urlsData = []byte(v)
		}

		if len(urlsData) > 0 {
			if err := json.Unmarshal(urlsData, &urls); err != nil {
				return domain.Profile{}, err
			}

			// Manually populate ProfileID for each URL as it's not included in the JSON from DB
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
