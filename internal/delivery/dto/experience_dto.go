package dto

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type ExperienceResp struct {
	ID          string     `json:"id"`
	ProfileID   string     `json:"profile_id"`
	Company     string     `json:"company"`
	Position    string     `json:"position"`
	Description []string   `json:"description"`
	Location    string     `json:"location"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	IsPresent   bool       `json:"is_present"`
}

type ExperienceReq struct {
	ProfileID   string   `json:"profile_id"`
	Company     string   `json:"company"`
	Position    string   `json:"position"`
	Description []string `json:"description"`
	Location    string   `json:"location"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
	IsPresent   bool     `json:"is_present"`
}

type ExperienceUpdateReq struct {
	ProfileID   *string   `json:"profile_id"`
	Company     *string   `json:"company"`
	Position    *string   `json:"position"`
	Description *[]string `json:"description"`
	Location    *string   `json:"location"`
	StartDate   *string   `json:"start_date"`
	EndDate     *string   `json:"end_date"`
	IsPresent   *bool     `json:"is_present"`
}

func ToExperienceDTO(ex *domain.Experience) ExperienceResp {
	return ExperienceResp{
		ID:          ex.ID,
		ProfileID:   ex.ProfileID,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		Location:    ex.Location,
		StartDate:   ex.StartDate,
		EndDate:     ex.EndDate,
		IsPresent:   ex.IsPresent,
	}
}

func ToExperienceInput(ex *ExperienceReq) domain.ExperienceInput {
	sd, _ := time.Parse("2006-01-02", ex.StartDate)
	var ed *time.Time
	if ex.EndDate != "" {
		parsed, err := time.Parse("2006-01-02", ex.EndDate)
		if err == nil {
			ed = &parsed
		}
	}
	return domain.ExperienceInput{
		ProfileID:   ex.ProfileID,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		Location:    ex.Location,
		StartDate:   sd,
		EndDate:     ed,
		IsPresent:   ex.IsPresent,
	}
}

func ToExperienceUpdateInput(ex *ExperienceUpdateReq) domain.ExperienceUpdateInput {
	var sd *time.Time
	if ex.StartDate != nil {
		parsed, err := time.Parse("2006-01-02", *ex.StartDate)
		if err == nil {
			sd = &parsed
		}
	}

	var ed *time.Time
	if ex.EndDate != nil {
		parsed, err := time.Parse("2006-01-02", *ex.EndDate)
		if err == nil {
			ed = &parsed
		}
	}

	return domain.ExperienceUpdateInput{
		ProfileID:   ex.ProfileID,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		Location:    ex.Location,
		StartDate:   sd,
		EndDate:     ed,
		IsPresent:   ex.IsPresent,
	}
}
