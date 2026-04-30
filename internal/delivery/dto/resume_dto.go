package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type ResumeResp struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Title       string           `json:"title"`
	About       string           `json:"about"`
	Address     string           `json:"address"`
	Email       string           `json:"email"`
	Phone       string           `json:"phone"`
	Url         []ProfileUrlResp `json:"urls"`
	Skills      []SkillResp      `json:"skills"`
	Languages   []LanguageResp   `json:"languages"`
	Experiences []ExperienceResp `json:"experiences"`
	Educations  []EducationResp  `json:"educations"`
	Projects    []ProjectResp    `json:"projects"`
}

func ToResumeDTO(r *domain.Resume) ResumeResp {
	urls := make([]ProfileUrlResp, 0, len(r.Url))
	for _, url := range r.Url {
		item := ToProfileUrlDTO(&url)
		urls = append(urls, item)
	}

	skills := make([]SkillResp, 0, len(r.Skills))
	for _, skill := range r.Skills {
		item := ToSkillDTO(&skill)
		skills = append(skills, item)
	}

	languages := make([]LanguageResp, 0, len(r.Languages))
	for _, language := range r.Languages {
		item := ToLanguageDTO(&language)
		languages = append(languages, item)
	}

	experiences := make([]ExperienceResp, 0, len(r.Experiences))
	for _, experience := range r.Experiences {
		item := ToExperienceDTO(&experience)
		experiences = append(experiences, item)
	}

	educations := make([]EducationResp, 0, len(r.Educations))
	for _, education := range r.Educations {
		item := ToEducationDTO(&education)
		educations = append(educations, item)
	}

	projects := make([]ProjectResp, 0, len(r.Projects))
	for _, project := range r.Projects {
		item := ToProjectDTO(&project)
		projects = append(projects, item)
	}

	return ResumeResp{
		ID:          r.ID,
		Name:        r.Name,
		Title:       r.Title,
		About:       r.About,
		Address:     r.Address,
		Email:       r.Email,
		Phone:       r.Phone,
		Url:         urls,
		Skills:      skills,
		Languages:   languages,
		Experiences: experiences,
		Educations:  educations,
		Projects:    projects,
	}
}
