package dto

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type ExperienceResp struct {
	ID          string    `json:"id"`
	ProfileID   string    `json:"profile_id"`
	Company     string    `json:"company"`
	Position    string    `json:"position"`
	Description []string  `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type ExperienceReq struct {
	ProfileID   string   `json:"profile_id"`
	Company     string   `json:"company"`
	Position    string   `json:"position"`
	Description []string `json:"description"`
	StartDate   string   `json:"start_date"`
	EndDate     string   `json:"end_date"`
}

type ExperienceUpdateReq struct {
	ProfileID   *string   `json:"profile_id"`
	Company     *string   `json:"company"`
	Position    *string   `json:"position"`
	Description *[]string `json:"description"`
	StartDate   *string   `json:"start_date"`
	EndDate     *string   `json:"end_date"`
}

func ToExperienceDTO(ex *domain.Experience) ExperienceResp {
	return ExperienceResp{
		ID:          ex.ID,
		ProfileID:   ex.ProfileID,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		StartDate:   ex.StartDate,
		EndDate:     ex.EndDate,
	}
}

func ToExperienceInput(ex *ExperienceReq) domain.ExperienceInput {
	sd, _ := time.Parse("2006-01-02", ex.StartDate)
	ed, _ := time.Parse("2006-01-02", ex.EndDate)
	return domain.ExperienceInput{
		ProfileID:   ex.ProfileID,
		Company:     ex.Company,
		Position:    ex.Position,
		Description: ex.Description,
		StartDate:   sd,
		EndDate:     ed,
	}
}

func ToExperienceUpdateInput(ex *ExperienceReq) domain.ExperienceUpdateInput {
	sd, _ := time.Parse("2006-01-02", ex.StartDate)
	ed, _ := time.Parse("2006-01-02", ex.EndDate)
	return domain.ExperienceUpdateInput{
		ProfileID:   &ex.ProfileID,
		Company:     &ex.Company,
		Position:    &ex.Position,
		Description: &ex.Description,
		StartDate:   &sd,
		EndDate:     &ed,
	}
}
