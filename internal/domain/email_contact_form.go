package domain

import "context"

type EmailContactFormInput struct {
	ProfileID string
	Name      string
	Email     string
	Subject   string
	Message   string
}

type EmailContactRepository interface {
	SendEmail(ctx context.Context, form EmailContactFormInput) error
}
