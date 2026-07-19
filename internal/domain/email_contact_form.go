package domain

import "context"

type EmailContactFormInput struct {
	ProfileID string
	Name      string
	Email     string
	Subject   string
	Message   string
}

type ContactFormUsecase interface {
	SendEmail(ctx context.Context, form EmailContactFormInput) error
}

type EmailContactRepository interface {
	SendEmail(ctx context.Context, form EmailContactFormInput, to string) error
}
