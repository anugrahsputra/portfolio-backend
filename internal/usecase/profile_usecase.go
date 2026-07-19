package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type profileUsecase struct {
	repo domain.ProfileRepository
}

func NewProfileUsecase(repo domain.ProfileRepository) domain.ProfileUsecase {
	return &profileUsecase{repo: repo}
}

func (u *profileUsecase) CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error) {
	return u.repo.CreateProfile(ctx, p)
}

func (u *profileUsecase) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	return u.repo.GetProfiles(ctx)
}

func (u *profileUsecase) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	return u.repo.GetProfile(ctx, id)
}

func (u *profileUsecase) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {
	current, err := u.repo.GetProfile(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to load profile for update: %w", err)
	}
	if current == nil {
		return fmt.Errorf("profile %s not found", id)
	}

	return u.repo.UpdateProfile(ctx, current.Merge(p))
}

func (u *profileUsecase) DeleteProfile(ctx context.Context, id string) error {
	return u.repo.DeleteProfile(ctx, id)
}
