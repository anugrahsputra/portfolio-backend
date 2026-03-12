package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type LanguageResp struct {
	ID          string `json:"id"`
	ProfileID   string `json:"profile_id"`
	Language    string `json:"language"`
	Proficiency string `json:"proficiency"`
}

type LanguageReq struct {
	ProfileID   string `json:"profile_id"`
	Language    string `json:"language"`
	Proficiency string `json:"proficiency"`
}

type LanguageUpdateReq struct {
	Language    string `json:"language"`
	Proficiency string `json:"proficiency"`
}

func ToLanguageDTO(l *domain.Language) LanguageResp {
	return LanguageResp{
		ID:          l.ID,
		ProfileID:   l.ProfileID,
		Language:    l.Language,
		Proficiency: l.Proficiency,
	}
}

func ToLanguageInput(l *LanguageReq) domain.LanguageInput {
	return domain.LanguageInput{
		ProfileID:   l.ProfileID,
		Language:    l.Language,
		Proficiency: l.Proficiency,
	}
}

func ToLanguageUpdateInput(l *LanguageUpdateReq) domain.LanguageUpdateInput {
	return domain.LanguageUpdateInput{
		Language:    l.Language,
		Proficiency: l.Proficiency,
	}
}
