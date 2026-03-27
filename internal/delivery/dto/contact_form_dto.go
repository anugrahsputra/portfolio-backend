package dto

import "github.com/anugrahsputra/portfolio-backend/internal/domain"

type ContactFormReq struct {
	ProfileID string `json:"profile_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
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
