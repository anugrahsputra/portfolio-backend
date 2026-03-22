package dto

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type EducationResp struct {
	ID             string     `json:"id"`
	ProfileID      string     `json:"profile_id"`
	School         string     `json:"school"`
	Degree         string     `json:"degree"`
	FieldOfStudy   string     `json:"field_of_study"`
	Gpa            float64    `json:"gpa"`
	StartDate      time.Time  `json:"start_date"`
	GraduationDate *time.Time `json:"graduation_date"`
	IsPresent      bool       `json:"is_present"`
}

type EducationReq struct {
	ProfileID      string  `json:"profile_id"`
	School         string  `json:"school"`
	Degree         string  `json:"degree"`
	FieldOfStudy   string  `json:"field_of_study"`
	Gpa            float64 `json:"gpa"`
	StartDate      string  `json:"start_date"`
	GraduationDate string  `json:"graduation_date"`
	IsPresent      bool    `json:"is_present"`
}

type EducationUpdateReq struct {
	ProfileID      *string  `json:"profile_id"`
	School         *string  `json:"school"`
	Degree         *string  `json:"degree"`
	FieldOfStudy   *string  `json:"field_of_study"`
	Gpa            *float64 `json:"gpa"`
	StartDate      *string  `json:"start_date"`
	GraduationDate *string  `json:"graduation_date"`
	IsPresent      *bool    `json:"is_present"`
}

func ToEducationDTO(e *domain.Education) EducationResp {
	return EducationResp{
		ID:             e.ID,
		ProfileID:      e.ProfileID,
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      e.StartDate,
		GraduationDate: e.GraduationDate,
		IsPresent:      e.IsPresent,
	}
}

func ToEducationInput(e *EducationReq) domain.EducationInput {
	sd, _ := time.Parse("2006-01-02", e.StartDate)
	var gd *time.Time
	if e.GraduationDate != "" {
		parsed, err := time.Parse("2006-01-02", e.GraduationDate)
		if err == nil {
			gd = &parsed
		}
	}
	return domain.EducationInput{
		ProfileID:      e.ProfileID,
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      sd,
		GraduationDate: gd,
		IsPresent:      e.IsPresent,
	}
}

func ToEducationUpdateInput(e *EducationUpdateReq) domain.EducationUpdateInput {
	var sd *time.Time
	if e.StartDate != nil {
		parsed, err := time.Parse("2006-01-02", *e.StartDate)
		if err == nil {
			sd = &parsed
		}
	}

	var gd *time.Time
	if e.GraduationDate != nil {
		parsed, err := time.Parse("2006-01-02", *e.GraduationDate)
		if err == nil {
			gd = &parsed
		}
	}

	return domain.EducationUpdateInput{
		ProfileID:      e.ProfileID,
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      sd,
		GraduationDate: gd,
		IsPresent:      e.IsPresent,
	}
}
