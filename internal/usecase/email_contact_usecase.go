package usecase

import (
	"context"
	"errors"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type EmailContactUsecase interface {
	SendEmail(ctx context.Context, form domain.EmailContactFormInput) error
}

type emailContactUsecase struct {
	repo domain.EmailContactRepository
}

func NewEmailContactUsecase(r domain.EmailContactRepository) EmailContactUsecase {
	return &emailContactUsecase{repo: r}
}

func (u *emailContactUsecase) SendEmail(ctx context.Context, form domain.EmailContactFormInput) error {
	if err := validate(form); err != nil {
		return err
	}

	return u.repo.SendEmail(ctx, form)
}

func validate(form domain.EmailContactFormInput) error {
	if form.ProfileID == "" {
		return errors.New("profile_id is required")
	}
	if form.Name == "" {
		return errors.New("name is required")
	}

	if form.Email == "" {
		return errors.New("email is required")
	}

	if form.Subject == "" {
		return errors.New("subject is required")
	}
	if form.Message == "" {
		return errors.New("message is required")
	}

	return nil
}
