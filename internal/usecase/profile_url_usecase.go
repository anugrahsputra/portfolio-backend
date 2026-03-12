package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type ProfileUrlUsecase interface {
	CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error)
	GetProfileUrl(ctx context.Context, id string) (domain.ProfileUrl, error)
	UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error
	DeleteProfileUrl(ctx context.Context, id string) error
}

type profileUrlUsecase struct {
	repo domain.ProfileUrlRepository
}

func NewProfileUrlUsecase(r domain.ProfileUrlRepository) ProfileUrlUsecase {
	return &profileUrlUsecase{repo: r}
}

func (u *profileUrlUsecase) CreateProfileUrl(ctx context.Context, pu domain.ProfileUrlInput) (*domain.ProfileUrl, error) {
	result, err := u.repo.CreateProfileUrl(ctx, pu)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *profileUrlUsecase) GetProfileUrl(ctx context.Context, id string) (domain.ProfileUrl, error) {
	result, err := u.repo.GetProfileUrl(ctx, id)
	if err != nil {
		return domain.ProfileUrl{}, err
	}

	return result, nil
}

func (u *profileUrlUsecase) UpdateProfileUrl(ctx context.Context, id string, pu domain.ProfileUrlUpdateInput) error {
	if err := u.repo.UpdateProfileUrl(ctx, id, pu); err != nil {
		return err
	}
	return nil
}

func (u *profileUrlUsecase) DeleteProfileUrl(ctx context.Context, id string) error {
	if err := u.repo.DeleteProfileUrl(ctx, id); err != nil {
		return err
	}
	return nil
}
