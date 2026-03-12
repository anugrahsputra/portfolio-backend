package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type EducationUsecase interface {
	CreateEducation(ctx context.Context, e domain.EducationInput) error
	GetEducations(ctx context.Context, profileID string) ([]domain.Education, error)
	UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error
	DeleteEducation(ctx context.Context, id string) error
}

type educationUsecase struct {
	repo domain.EducationRepository
}

func NewEducationUsecase(r domain.EducationRepository) EducationUsecase {
	return &educationUsecase{repo: r}
}

func (u *educationUsecase) CreateEducation(ctx context.Context, e domain.EducationInput) error {
	if err := u.repo.CreateEducation(ctx, e); err != nil {
		return err
	}

	return nil
}

func (u *educationUsecase) GetEducations(ctx context.Context, profileID string) ([]domain.Education, error) {
	result, err := u.repo.GetEducations(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *educationUsecase) UpdateEducation(ctx context.Context, id string, e domain.EducationUpdateInput) error {
	if err := u.repo.UpdateEducation(ctx, id, e); err != nil {
		return err
	}

	return nil
}

func (u *educationUsecase) DeleteEducation(ctx context.Context, id string) error {
	if err := u.repo.DeleteEducation(ctx, id); err != nil {
		return err
	}

	return nil
}
