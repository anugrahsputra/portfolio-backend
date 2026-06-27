package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type ContactFormReq struct {
	ProfileID string `json:"profile_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Subject   string `json:"subject" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

func ToContactFormInput(cf *ContactFormReq) domain.EmailContactFormInput {
	return domain.EmailContactFormInput{
		ProfileID: cf.ProfileID,
		Name:      cf.Name,
		Email:     cf.Email,
		Subject:   cf.Subject,
		Message:   cf.Message,
	}
}
