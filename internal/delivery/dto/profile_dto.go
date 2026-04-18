package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type ProfileResp struct {
	ID      string           `json:"id"`
	Name    string           `json:"name"`
	Title   string           `json:"title"`
	About   string           `json:"about"`
	Address string           `json:"address"`
	Email   string           `json:"email"`
	Phone   string           `json:"phone"`
	Url     []ProfileUrlResp `json:"url"`
}

type ProfilePublicResp struct {
	ID      string           `json:"id"`
	Name    string           `json:"name"`
	Title   string           `json:"title"`
	About   string           `json:"about"`
	Address string           `json:"address"`
	Email   string           `json:"email"`
	Url     []ProfileUrlResp `json:"url"`
}

type ProfileReq struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	About   string `json:"about"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func ToProfileDTO(p *domain.Profile) ProfileResp {
	urls := make([]ProfileUrlResp, 0, len(p.Url))
	for _, url := range p.Url {
		item := ToProfileUrlDTO(&url)
		urls = append(urls, item)
	}

	return ProfileResp{
		ID:      p.ID,
		Name:    p.Name,
		Title:   p.Title,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Phone:   p.Phone,
		Url:     urls,
	}
}

func ToProfilePublicDTO(p *domain.Profile) ProfilePublicResp {
	urls := make([]ProfileUrlResp, 0, len(p.Url))
	for _, url := range p.Url {
		item := ToProfileUrlDTO(&url)
		urls = append(urls, item)
	}

	return ProfilePublicResp{
		ID:      p.ID,
		Name:    p.Name,
		Title:   p.Title,
		About:   p.About,
		Address: p.Address,
		Email:   p.Email,
		Url:     urls,
	}
}
