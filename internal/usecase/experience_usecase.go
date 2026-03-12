package usecase

import (
	"context"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type ExperienceUsecase interface {
	CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error)
	GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error)
	UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error)
	DeleteExperience(ctx context.Context, id string) error
}

type experienceUsecase struct {
	repo domain.ExperienceRepository
}

func NewExperienceUsecase(r domain.ExperienceRepository) ExperienceUsecase {
	return &experienceUsecase{repo: r}
}

func (u *experienceUsecase) CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error) {
	result, err := u.repo.CreateExperience(ctx, ex)
	if err != nil {
		return domain.Experience{}, err
	}

	return result, nil
}

func (u *experienceUsecase) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	result, err := u.repo.GetExperiences(ctx, profileID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *experienceUsecase) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	result, err := u.repo.UpdateExperience(ctx, id, ex)
	if err != nil {
		return domain.Experience{}, err
	}

	return result, nil
}

func (u *experienceUsecase) DeleteExperience(ctx context.Context, id string) error {
	if err := u.repo.DeleteExperience(ctx, id); err != nil {
		return err
	}

	return nil
}
