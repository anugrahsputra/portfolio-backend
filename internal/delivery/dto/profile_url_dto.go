package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type ProfileUrlResp struct {
	ID        string `json:"id"`
	ProfileID string `json:"profile_id"`
	Label     string `json:"label"`
	Url       string `json:"url"`
}

type ProfileUrlReq struct {
	ProfileID string `json:"profile_id"`
	Label     string `json:"label"`
	Url       string `json:"url"`
}

func ToProfileUrlDTO(pu *domain.ProfileUrl) ProfileUrlResp {
	return ProfileUrlResp{
		ID:        pu.ID,
		ProfileID: pu.ProfileID,
		Label:     pu.Label,
		Url:       pu.Url,
	}
}
