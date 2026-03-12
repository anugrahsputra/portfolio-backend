package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type ProfileUsecase interface {
	CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error)
	GetProfile(ctx context.Context, id string) (*domain.Profile, error)
	UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error
	DeleteProfile(ctx context.Context, id string) error
}

type profileUsecase struct {
	repo domain.ProfileRepository
}

func NewProfileUsecase(r domain.ProfileRepository) ProfileUsecase {
	return &profileUsecase{repo: r}
}

func (u *profileUsecase) CreateProfile(ctx context.Context, p domain.ProfileInput) (*domain.Profile, error) {
	result, err := u.repo.CreateProfile(ctx, p)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *profileUsecase) GetProfile(ctx context.Context, id string) (*domain.Profile, error) {
	result, err := u.repo.GetProfile(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (u *profileUsecase) UpdateProfile(ctx context.Context, id string, p domain.ProfileUpdateInput) error {
	if err := u.repo.UpdateProfile(ctx, id, p); err != nil {
		return err
	}
	return nil
}

func (u *profileUsecase) DeleteProfile(ctx context.Context, id string) error {
	if err := u.repo.DeleteProfile(ctx, id); err != nil {
		return err
	}
	return nil
}
