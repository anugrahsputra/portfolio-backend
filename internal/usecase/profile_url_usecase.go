package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type profileUrlUsecase struct {
	repo domain.ProfileUrlRepository
}

func NewProfileUrlUsecase(repo domain.ProfileUrlRepository) domain.ProfileUrlUsecase {
	return &profileUrlUsecase{repo: repo}
}

func (u *profileUrlUsecase) CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error) {
	return u.repo.CreateProfileUrl(ctx, pu)
}

func (u *profileUrlUsecase) GetProfileUrlByID(ctx context.Context, id string) (domain.ProfileUrl, error) {
	return u.repo.GetProfileUrlByID(ctx, id)
}

func (u *profileUrlUsecase) GetProfileUrl(ctx context.Context, profileID string) ([]domain.ProfileUrl, error) {
	return u.repo.GetProfileUrl(ctx, profileID)
}

func (u *profileUrlUsecase) UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error {
	current, err := u.repo.GetProfileUrlByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to load profile url for update: %w", err)
	}

	return u.repo.UpdateProfileUrl(ctx, current.Merge(pu))
}

func (u *profileUrlUsecase) DeleteProfileUrl(ctx context.Context, id string) error {
	return u.repo.DeleteProfileUrl(ctx, id)
}
