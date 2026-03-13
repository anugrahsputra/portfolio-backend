package dto

import (
	"time"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type EducationResp struct {
	ID             string    `json:"id"`
	ProfileID      string    `json:"profile_id"`
	School         string    `json:"school"`
	Degree         string    `json:"degree"`
	FieldOfStudy   string    `json:"field_of_study"`
	Gpa            float64   `json:"gpa"`
	StartDate      time.Time `json:"start_date"`
	GraduationDate time.Time `json:"graduation_date"`
}

type EducationReq struct {
	ProfileID      string  `json:"profile_id"`
	School         string  `json:"school"`
	Degree         string  `json:"degree"`
	FieldOfStudy   string  `json:"field_of_study"`
	Gpa            float64 `json:"gpa"`
	StartDate      string  `json:"start_date"`
	GraduationDate string  `json:"graduation_date"`
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
	}
}

func ToEducationInput(e *EducationReq) domain.EducationInput {
	sd, _ := time.Parse("2006-01-02", e.StartDate)
	gd, _ := time.Parse("2006-01-02", e.GraduationDate)
	return domain.EducationInput{
		ProfileID:      e.ProfileID,
		School:         e.School,
		Degree:         e.Degree,
		FieldOfStudy:   e.FieldOfStudy,
		Gpa:            e.Gpa,
		StartDate:      sd,
		GraduationDate: gd,
	}
}

func ToEducationUpdateInput(e *EducationReq) domain.EducationUpdateInput {
	sd, _ := time.Parse("2006-01-02", e.StartDate)
	gd, _ := time.Parse("2006-01-02", e.GraduationDate)
	return domain.EducationUpdateInput{
		ProfileID:      &e.ProfileID,
		School:         &e.School,
		Degree:         &e.Degree,
		FieldOfStudy:   &e.FieldOfStudy,
		Gpa:            &e.Gpa,
		StartDate:      &sd,
		GraduationDate: &gd,
	}
}
