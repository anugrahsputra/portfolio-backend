package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type contactFormUsecase struct {
	mailer      domain.EmailContactRepository
	profileRepo domain.ProfileRepository
}

func NewContactFormUsecase(mailer domain.EmailContactRepository, profileRepo domain.ProfileRepository) domain.ContactFormUsecase {
	return &contactFormUsecase{mailer: mailer, profileRepo: profileRepo}
}

func (u *contactFormUsecase) SendEmail(ctx context.Context, form domain.EmailContactFormInput) error {
	profile, err := u.profileRepo.GetProfile(ctx, form.ProfileID)
	if err != nil {
		return fmt.Errorf("failed to get profile: %w", err)
	}
	if profile == nil {
		return fmt.Errorf("profile %s not found", form.ProfileID)
	}

	return u.mailer.SendEmail(ctx, form, profile.Email)
}
