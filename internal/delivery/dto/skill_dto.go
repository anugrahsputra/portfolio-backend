package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type SkillResp struct {
	ID           string   `json:"id"`
	ProfileID    string   `json:"profile_id"`
	Tools        []string `json:"tools"`
	Technologies []string `json:"technologies"`
	HardSkills   []string `json:"hard_skills"`
	SoftSkills   []string `json:"soft_skills"`
}

type SkillReq struct {
	ProfileID    string   `json:"profile_id"`
	Tools        []string `json:"tools"`
	Technologies []string `json:"technologies"`
	HardSkills   []string `json:"hard_skills"`
	SoftSkills   []string `json:"soft_skills"`
}

type SkillUpdateReq struct {
	Tools        []string `json:"tools"`
	Technologies []string `json:"technologies"`
	HardSkills   []string `json:"hard_skills"`
	SoftSkills   []string `json:"soft_skills"`
}

func ToSkillDTO(s *domain.Skill) SkillResp {
	return SkillResp{
		ID:           s.ID,
		ProfileID:    s.ProfileID,
		Tools:        s.Tools,
		Technologies: s.Technologies,
		HardSkills:   s.HardSkills,
		SoftSkills:   s.SoftSkills,
	}
}
