package usecase

import (
	"context"
	"fmt"

	"github.com/anugrahsputra/portfolio-backend/internal/domain"
)

type experienceUsecase struct {
	repo domain.ExperienceRepository
}

func NewExperienceUsecase(repo domain.ExperienceRepository) domain.ExperienceUsecase {
	return &experienceUsecase{repo: repo}
}

func (u *experienceUsecase) CreateExperience(ctx context.Context, ex domain.ExperienceInput) (domain.Experience, error) {
	return u.repo.CreateExperience(ctx, ex)
}

func (u *experienceUsecase) GetExperiences(ctx context.Context, profileID string) ([]domain.Experience, error) {
	return u.repo.GetExperiences(ctx, profileID)
}

func (u *experienceUsecase) GetExperienceByID(ctx context.Context, id string) (domain.Experience, error) {
	return u.repo.GetExperienceByID(ctx, id)
}

func (u *experienceUsecase) UpdateExperience(ctx context.Context, id string, ex domain.ExperienceUpdateInput) (domain.Experience, error) {
	current, err := u.repo.GetExperienceByID(ctx, id)
	if err != nil {
		return domain.Experience{}, fmt.Errorf("failed to load experience for update: %w", err)
	}

	return u.repo.UpdateExperience(ctx, current.Merge(ex))
}

func (u *experienceUsecase) DeleteExperience(ctx context.Context, id string) error {
	return u.repo.DeleteExperience(ctx, id)
}
